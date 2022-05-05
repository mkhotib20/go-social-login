package socialauth

import (
	"errors"

	"github.com/mkhotib20/verify-apple-id-token/utils/fetch"
)

const facebookBaseUrl = "https://graph.facebook.com/me?fields=id,name,email,picture&access_token="

type fbData struct {
	url string
}

type fbPicture struct {
	data fbData
}

type FacebookAuthResult struct {
	email   string
	name    string
	picture fbPicture
	error   *string
}

func verifyFacebookToken(token string, sc *DataClaim) error {
	grs := &FacebookAuthResult{}
	err := fetch.Get(facebookBaseUrl + token).Json(&grs)
	if err != nil {
		return err
	}
	if grs.error != nil {
		return errors.New(*grs.error)
	}
	sc.Avatar = grs.picture.data.url
	sc.FullName = grs.name
	sc.Email = grs.email
	return nil
}
