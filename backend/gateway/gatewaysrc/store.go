package gatewaysrc

//GatewayStore stores database methods for the gateway.
type GatewayStore interface {
	//User model methods
	GetUserInfo(guser *GoogleUser) (*User, error) //Gets user information from google, results will be encoded on each request so we know which user is making the request
}
