package user

import "google.golang.org/api/oauth2/v2"

func getTokenInfo(idToken string) (*oauth2.Tokeninfo, error) {
oauth2Service,err := oauth2.New(&http.Client{})
if err != nil {
    return nil, err
}
tokenInfoCall := oauth2Service.Tokeninfo()
tokenInfoCall.IdToken(idToken)
return tokenInfoCall.Do()
}