package test

import (
	"encoding/json"
	"fmt"
	"github.com/fire1220/goutils/marshal"
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
	return marshal.Marshal(g)
}

func TestMarshal(t *testing.T) {
	d, _ := time.Parse(time.DateTime, "0000-00-00 00:00:00")
	good := Good{ID: 123, Name: "jock", PlayTime: time.Now(), CreatedAt: time.Now(), UpdatedAt: d}
	bytes, _ := json.Marshal(good)
	// {"id":123,"name":"jock","play_time":"2024-03-14 19:40:19","created_at":"2024-03-14 19:40:19","updated_at":"00:00:00"}
	fmt.Printf("%s\n", bytes)
}
