package marshal

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"
)

type Good struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (g Good) MarshalJSON() ([]byte, error) {
	return Marshal(g)
}

func TestMarshal(t *testing.T) {
	good := Good{123, "jock", time.Now(), time.Now()}
	bytes, _ := json.Marshal(good)
	// {"id":123,"name":"jock","created_at":"2024-03-14 17:55:03","updated_at":"2024-03-14 17:55:03"}
	fmt.Printf("%s\n", bytes)
}
