package domain

import (
	"encoding/json"
	"fmt"
	"testing"

	//"go.mongodb.org/mongo-driver/v2/bson"
)

func TestStrategyJSONMArshal(t *testing.T) {
	s := Strategy{
		//ID:          bson.NewObjectID(),
		ID: "12345",
		Name:        "strategy name",
		Description: "strategy description",
	}

	jsonBytes, _ := json.Marshal(s)

	fmt.Println(string(jsonBytes))
}
