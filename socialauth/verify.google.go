package socialauth

import (
	"errors"

	"github.com/mkhotib20/verify-apple-id-token/utils/fetch"
)

const googleBaseUrl = "https://oauth2.googleapis.com/tokeninfo?id_token="

type GoogleTokenResult struct {
	email   string
	error   *string
	name    string
	picture string
}

func verifyGoogleToken(token string, sc *DataClaim) error {
	grs := &GoogleTokenResult{}
	err := fetch.Get(googleBaseUrl + token).Json(&grs)
	if err != nil {
		return err
	}
	if grs.error != nil {
		return errors.New(*grs.error)
	}
	sc.Avatar = grs.picture
	sc.FullName = grs.name
	sc.Email = grs.email
	return nil
}
