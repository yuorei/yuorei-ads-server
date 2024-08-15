package domain

import (
	"fmt"

	"github.com/google/uuid"
)

func NewUUID() string {
	uuidObj, err := uuid.NewUUID()
	if err != nil {
		return ""
	}
	return uuidObj.String()
}

func NewAdID() string {
	return fmt.Sprintf("ad_%s", NewUUID())
}

func NewCampaignID() string {
	return fmt.Sprintf("campaign_%s", NewUUID())
}
