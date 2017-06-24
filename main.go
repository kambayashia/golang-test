package main

import (
	"github.com/dghubble/go-twitter/twitter"
	"flag"
	"github.com/dghubble/oauth1"
	"fmt"
	"net/http"
)

func init() {

}
func main() {
	consumerKey := flag.String("consumer-key", "", "Twitter application consumer key")
	consumerSecret := flag.String("consumer-secret", "", "Twitter application consumer secret")
	// accessToken := flag.String("access-token", "", "Twitter application access token")
	// accessSecret := flag.String("access-secret", "", "Twitter application access secret")
	var _ = flag.String("access-token", "", "Twitter application access token")
	var _ = flag.String("access-secret", "", "Twitter application access secret")
	flag.Parse()

	var requestSecret string

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		config := oauthConfig(*consumerKey, *consumerSecret)

		var requestToken string
		requestToken, requestSecret, _ = config.RequestToken()

		url, _ := config.AuthorizationURL(requestToken)
		http.Redirect(w, r, url.String(), http.StatusFound)
	})
	http.HandleFunc("/auth", func(w http.ResponseWriter, r *http.Request) {
		requestToken, verifier, _ := oauth1.ParseAuthorizationCallback(r)
		config := oauthConfig(*consumerKey, *consumerSecret)
		accessToken, accessSecret, _ := config.AccessToken(requestToken, requestSecret, verifier)
		fmt.Fprintf(w, "accessToken: %v accessSecret:%v", accessToken, accessSecret)
	})

	err := http.ListenAndServe(":3000", nil)
	fmt.Errorf("error: %v", err)
}

func oauthConfig(consumerKey, consumerSecret string) oauth1.Config{
	return oauth1.Config{
		ConsumerKey: consumerKey,
		ConsumerSecret: consumerSecret,
		CallbackURL: "http://localhost:3000/auth",
		Endpoint: oauth1.Endpoint{
			AuthorizeURL: "https://api.twitter.com/oauth/authorize",
			AccessTokenURL: "https://api.twitter.com/oauth/access_token",
			RequestTokenURL: "https://api.twitter.com/oauth/request_token",
		},
	}
}

func twitterAccess(consumerKey, consumerSecret, accessToken, accessSecret string) {
	config := oauthConfig(consumerKey, consumerSecret)
	token := oauth1.NewToken(accessToken, accessSecret)
	// OAuth1 http.Client will automatically authorize Requests
	httpClient := config.Client(oauth1.NoContext, token)

	// Twitter client
	client := twitter.NewClient(httpClient)

	user, _, err := client.Accounts.VerifyCredentials(&twitter.AccountVerifyParams{
		SkipStatus: twitter.Bool(false),
		IncludeEmail: twitter.Bool(true),
	})
	fmt.Errorf("%v", err)
	fmt.Printf("User's ACCOUNT:\n%+v\n", user)
}
