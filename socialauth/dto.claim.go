package socialauth

/*
 ** General data claim as general as posibble
 */
type DataClaim struct {
	FullName string
	Avatar   string
	Email    string
}

type SocialAuth struct {
	token string
}
