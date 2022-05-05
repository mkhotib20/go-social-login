package socialauth

/** Verify
 * Create new socialauth Instance
 * After this you can call Google(), Facebook(), etc
 * I recommend to use this way if you implement ala-carte authentication
 */
func Verify(token string) *SocialAuth {
	return &SocialAuth{
		token: token,
	}
}

/** Google()
 * Verify google token
 */
func (r *SocialAuth) Google() (dt *DataClaim, err error) {
	err = verifyGoogleToken(r.token, dt)
	return
}

/** Facebook()
 * Verify facebook token
 */
func (r *SocialAuth) Facebook() (dt *DataClaim, err error) {
	err = verifyFacebookToken(r.token, dt)
	return
}
