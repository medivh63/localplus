package models

import "encoding/json"

type Question struct {
	ID      string          `json:"id"`
	Content string          `json:"content"`
	Images  string          `json:"images"`
	Options json.RawMessage `json:"options"`
}
