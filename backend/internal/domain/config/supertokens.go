package config

import (
	"github.com/supertokens/supertokens-golang/recipe/emailpassword"
	"github.com/supertokens/supertokens-golang/recipe/session"
	"github.com/supertokens/supertokens-golang/supertokens"
)

func main() {
	apiBasePath := "/auth"
	websiteBasePath := "/auth"
	err := supertokens.Init(supertokens.TypeInput{
		Supertokens: &supertokens.ConnectionInfo{
			// We use try.supertokens for demo purposes.
			// At the end of the tutorial we will show you how to create
			// your own SuperTokens core instance and then update your config.
			ConnectionURI: "https://try.supertokens.io",
			// APIKey: <YOUR_API_KEY>
		},
		AppInfo: supertokens.AppInfo{
			AppName:         "be",
			APIDomain:       "http://localhost:3567",
			WebsiteDomain:   "http://localhost:3000",
			APIBasePath:     &apiBasePath,
			WebsiteBasePath: &websiteBasePath,
		},
		RecipeList: []supertokens.Recipe{
			emailpassword.Init(nil),
			session.Init(nil),
		},
	})

	if err != nil {
		panic(err.Error())
	}
}
