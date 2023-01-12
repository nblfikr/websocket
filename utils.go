package main

import (
	"github.com/google/uuid"
)

func uuuid() string {
	return uuid.New().String()
}
