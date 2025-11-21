package note

import (
	"errors"
	"fmt"
	"time"
)

type Note struct {
	title     string
	content   string
	createdAt time.Time
}

//this become a method by adding
//(somethinginside) that passes cvalues from the note struct
// the word before the name of the struct is a nickname
// that

func (note Note) Display() {
	fmt.Printf("Your note titled %v has the following content:\n\n%v\n\n", note.title, note.content)
}

func NewNote(title, content string) (*Note, error) {
	if title == "" || content == "" {
		return &Note{}, errors.New("invalid input")
	}

	return &Note{
		title:     title,
		content:   content,
		createdAt: time.Now(),
	}, nil
}
