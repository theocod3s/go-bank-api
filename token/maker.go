package token

import "time"

type Maker interface {
	// creates and signs a token for a specific username and duration.
	CreateToken(username string, duration time.Duration) (string, error)

	// check if the input token is valid or not
	VerifyToken(token string) (*Payload, error)
}
