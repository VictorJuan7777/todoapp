package token

import "time"

type Maker interface {
	CreateToken(username string, duration time.Duration) (string, error, *Payload)
	VerifyToken(token string) (*Payload, error)
}
