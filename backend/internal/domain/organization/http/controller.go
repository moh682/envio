package organization_http

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
	"github.com/moh682/envio/backend/internal/domain/organization"
	"github.com/supertokens/supertokens-golang/recipe/session"
)

type Controller interface {
	GetOrganizationByUserId() http.HandlerFunc
	CreateOrganization() http.HandlerFunc
}

type httpController struct {
	organizationService organization.Service
}

type CreateOrganizationRequest struct {
	Name               string `json:"name"`
	InvoiceNumberStart int32  `json:"invoiceNumberStart"`
}

// CreateOrganization implements Controller.
func (h *httpController) CreateOrganization() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		sessionContainer := session.GetSessionFromRequestContext(r.Context())
		userID := sessionContainer.GetUserID()

		userUUID := uuid.MustParse(userID)

		defer r.Body.Close()

		var reqBody CreateOrganizationRequest

		err := json.NewDecoder(r.Body).Decode(&reqBody)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		organization, err := h.organizationService.CreateOrganization(r.Context(), userUUID, reqBody.Name, reqBody.InvoiceNumberStart)
		if err != nil {
			w.WriteHeader(500)
			return
		}

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(organization)
	}
}

// GetOrganizationByUserId implements Controller.
func (h *httpController) GetOrganizationByUserId() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		sessionContainer := session.GetSessionFromRequestContext(r.Context())
		userID := sessionContainer.GetUserID()

		userUUID := uuid.MustParse(userID)

		organization, err := h.organizationService.GetOrganizationByUserId(r.Context(), userUUID)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(err)
			return
		}

		if organization == nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(organization)
	}
}

func NewHttpController(organizationService organization.Service) Controller {
	return &httpController{organizationService}
}
