package todo

import (
	"fmt"

	"github.com/google/uuid"
)

type Item struct {
	ID    string `json:"id"`
	Title string `json:"title"`
}

func New(title string) Item {
	return Item{
		ID:    uuid.New().String(),
		Title: title,
	}
}

func (t *Item) IsValid() error {
	if err := uuid.Validate(t.ID); err != nil {
		return fmt.Errorf("invalid todo ID: %w", err)
	}
	return nil
}
