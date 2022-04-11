package controllers

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"oauth-2-youtube/config"
)

func GoogleLogin(w http.ResponseWriter, r *http.Request) {
	googleConfig := config.SetupConfig()

	url := googleConfig.AuthCodeURL("somestate")

	http.Redirect(w, r, url, http.StatusSeeOther)
}

func GoogleCallback(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL.Query())
	state := r.URL.Query()["state"][0]
	if state != "somestate" {
		fmt.Fprintln(w, "state does not match")
		return
	}

	code := r.URL.Query()["code"][0]
	googleConfig := config.SetupConfig()

	token, err := googleConfig.Exchange(context.Background(), code)
	if err != nil {
		fmt.Fprintln(w, "token exchange failed")
		return
	}

	res, err := http.Get("https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + token.AccessToken)
	if err != nil {
		fmt.Fprintln(w, "could not retrieve user info from token")
		return
	}

	userData, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Fprintln(w, "could not unmarshal response body")
		return
	}

	fmt.Fprintln(w, string(userData))
}

func GoogleValidityCallback(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL.Query())
	// state := r.URL.Query()["state"][0]
	// if state != "somestate" {
	// 	fmt.Fprintln(w, "state does not match")
	// 	return
	// }

	code := r.URL.Query()["code"][0]
	googleConfig := config.SetupConfig()

	token, err := googleConfig.Exchange(context.Background(), code)
	if err != nil {
		fmt.Fprintln(w, "token exchange failed")
		return
	}

	if !token.Valid() {
		fmt.Fprintln(w, "invalid token")
		return
	}

	// res, err := http.Get("https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + token.AccessToken)
	// if err != nil {
	// 	fmt.Fprintln(w, "could not retrieve user info from token")
	// 	return
	// }

	// userData, err := ioutil.ReadAll(res.Body)
	// if err != nil {
	// 	fmt.Fprintln(w, "could not unmarshal response body")
	// 	return
	// }

	fmt.Fprintln(w, token.Expiry.Unix())
}

// Expected response:
// {
// 	"id": "115054411203375952094",
// 	"email": "sajirt@gmail.com",
// 	"verified_email": true,
// 	"name": "sajir mohamed",
// 	"given_name": "sajir",
// 	"family_name": "mohamed",
// 	"picture": "https://lh3.googleusercontent.com/a/AATXAJzUsfQlawes6HsBwKBTY_a_gu0HilfgK3_t6RY-2A=s96-c",
// 	"locale": "en-GB"
// }
