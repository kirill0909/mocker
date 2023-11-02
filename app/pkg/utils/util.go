package utils

import (
	"github.com/google/uuid"
	// "math/rand"
	// "time"
	"fmt"
	"log"
)

func GetRandomUUID() string {
	// rand.Seed(time.Now().UnixNano())

	uuid, err := uuid.NewRandom()
	if err != nil {
		log.Printf("utils.GetRandomUUID.NewRandom(). %s", err)
	}

	return fmt.Sprintf("%v", uuid)
}
