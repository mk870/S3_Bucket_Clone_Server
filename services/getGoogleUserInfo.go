package services

import (
	"context"

	"google.golang.org/api/idtoken"
)

func RetrieveUser(tokenString string, audience string) *idtoken.Payload {
	payload, err := idtoken.Validate(context.Background(), tokenString, audience)
	if err != nil {
		panic(err.Error())
	}
	return payload
}
