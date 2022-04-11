package config

import (
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

func SetupConfig() *oauth2.Config {
	conf := &oauth2.Config{
		ClientID:     "138856099997-hs2p78c5o8snvvkislp54u32501fgkto.apps.googleusercontent.com",
		ClientSecret: "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",
		RedirectURL:  "http://localhost:3080/google/callback",
		Scopes: []string{
			"https://www.googleapis.com/auth/userinfo.email",
			"https://www.googleapis.com/auth/userinfo.profile",
		},
		Endpoint: google.Endpoint,
	}

	return conf
}
