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
}

type httpController struct {
	organizationService organization.Service
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
