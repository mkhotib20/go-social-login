package socialauth

/** DirectVerify
 * Easy way to verify token
 * You can directly use this in one shot
 * I recommend to use this way if you implement omni-channel authentication
 */
func DirectVerify(token, channel string) (dt *DataClaim, err error) {
	switch channel {
	case "facebook":
		err = verifyFacebookToken(token, dt)
		return
	case "google":
		err = verifyGoogleToken(token, dt)
		return
	}
	return
}
