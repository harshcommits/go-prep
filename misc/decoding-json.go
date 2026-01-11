package misc

import (
	"encoding/json"
	"fmt"
	"log"
)

func DecodedJsonFinal() {
	jsonValue := `[
        {
            "species": "pigeon",
            "description": "likes to perch on rocks"
        },
        {
            "species":"eagle",
            "description":"bird of prey"
        }
    ]`

	// Option 1: Output raw, then decode
	DecodeJSONResponse(jsonValue)

	// Option 2: Using generics (1.18+)
	var birds []Bird
	if err := DecodeJSON[[]Bird](jsonValue, &birds); err != nil {
		log.Fatal(err)
	}
	fmt.Println("\nDecoded with generics:", birds)
}

// Generic decoder function (Go 1.18+)
func DecodeJSON[T any](content string, v *T) error {
	return json.Unmarshal([]byte(content), v)
}

func DecodeJSONResponse(content string) {
	// First, output the raw response
	var rawData interface{}
	err := json.Unmarshal([]byte(content), &rawData)
	if err != nil {
		log.Fatal("Error decoding raw data: ", err)
	}
	fmt.Println("Raw Response:", rawData)

	// Then decode to typed struct
	var birds []Bird
	json.Unmarshal([]byte(content), &birds)

	for _, bird := range birds {
		fmt.Println(bird.Species)
	}
}

type Bird struct {
	Species     string `json:"species"`
	Description string `json:"description"`
}
