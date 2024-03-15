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
	PlayTime  time.Time `json:"play_time"`
	CreatedAt time.Time `json:"created_at" datetime:"2006-01-02,omitempty"`
	UpdatedAt time.Time `json:"updated_at" datetime:"2006-01-02 15:04:05,omitempty"`
}

func (g Good) MarshalJSON() ([]byte, error) {
	return Marshal(g)
}

func TestMarshal(t *testing.T) {
	d, _ := time.Parse(time.DateTime, "0000-00-00 00:00:00")
	good := Good{ID: 123, Name: "jock", PlayTime: time.Now(), CreatedAt: time.Now(), UpdatedAt: d}
	bytes, _ := json.Marshal(good)
	// {"id":123,"name":"jock","play_time":"2024-03-15 10:28:38","created_at":"2024-03-15","updated_at":""}
	fmt.Printf("%s\n", bytes)
}
