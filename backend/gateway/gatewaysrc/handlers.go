package gatewaysrc

import (
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"time"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

var (
	googleOauthConfig = &oauth2.Config{
		RedirectURL:  "https://groups.cahillaw.me/callback",
		ClientID:     os.Getenv("GOOGLE_CLIENT_ID"),
		ClientSecret: os.Getenv("GOOGLE_CLIENT_SECRET"),
		Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email", "https://www.googleapis.com/auth/userinfo.profile", "openid"},
		Endpoint:     google.Endpoint,
	}
)

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func randSeq(n int) string {
	rand.Seed(time.Now().UnixNano())
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

//HandleHome login page might be removed lol
func HandleHome(w http.ResponseWriter, r *http.Request) {
	var htmlIndex = `<html>
		<body>
			<a href="/login">Google Log In</a>
		</body>
	</html>`

	fmt.Fprintf(w, htmlIndex)
}

//HandleLogin is the login controller, handles requests to login through google
//inputs: none
//outputs: Google authorization token
func HandleLogin(w http.ResponseWriter, r *http.Request) {
	state := randSeq(16)
	cookie := http.Cookie{
		Name:  "google_auth_state",
		Value: state,
	}
	http.SetCookie(w, &cookie)

	url := googleOauthConfig.AuthCodeURL(state)
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}

//HandleCallback handles callback after authenticating
func HandleCallback(w http.ResponseWriter, r *http.Request) {
	cookie, _ := r.Cookie("google_auth_state")

	if r.FormValue("state") != cookie.Value {
		fmt.Println("state is not valid")
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	token, err := googleOauthConfig.Exchange(oauth2.NoContext, r.FormValue("code"))
	if err != nil {
		fmt.Printf("could not get token: %s\n", err.Error())
		http.Redirect(w, r, "/test", http.StatusTemporaryRedirect)
		return
	}

	fmt.Println(token.AccessToken)

	resp, _ := http.Get("https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + token.AccessToken)
	fmt.Println(resp)

	frontBase := os.Getenv("FRONT_BASE")

	url := frontBase + "/redirect?access_token=" + token.AccessToken
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}
