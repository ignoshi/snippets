package snippets

import (
	"errors"
	"time"

	"github.com/ignoshi/core/db"
	"github.com/ignoshi/core/tags"
	"gopkg.in/mgo.v2/bson"
)

// Snippet represents a code snipper model
type Snippet struct {
	ID        bson.ObjectId `json:"id,omitempty" bson:"_id"`
	Title     string        `json:"title" bson:"title"`
	Body      string        `json:"body" bson:"body"`
	Tags      []tags.Tag    `json:"tags" body:"tags"`
	CreatedAt time.Time     `json:"created_at" bosn:"created_at"`
	UpdatedAt time.Time     `json:"updated_at" bosn:"updated_at"`
}

type SnippetsManager struct{}

// IsValid checks if model fields are valid before save
func (s *Snippet) IsValid() (ok bool, errs []error) {
	ok = true
	if s.Title == "" {
		ok = false
		errs = append(errs, errors.New("Title is required"))
	}
	if s.Body == "" {
		ok = false
		errs = append(errs, errors.New("Body is required"))
	}

	return
}

// Save saves a snippet into database, either update it or create new one
func (s *Snippet) Save() error {
	db := db.GetDB()
	s.UpdatedAt = time.Now()
	var err error
	if s.ID == bson.ObjectId("") {
		s.ID = bson.NewObjectId()
		s.CreatedAt = s.UpdatedAt
		err = db.C("snippets").Insert(s)
	} else {
		err = db.C("snippets").Update(bson.M{"_id": s.ID}, s)
	}
	return err
}

// ListSnippets returns a list of all snippets exists in the db
func (m *SnippetsManager) Find(queryParams bson.M) ([]Snippet, error) {
	items := []Snippet{}
	db := db.GetDB()
	err := db.C("snippets").Find(queryParams).All(&items)
	return items, err
}

func (m *SnippetsManager) FindOne(queryParams bson.M) (Snippet, error) {
	item := Snippet{}
	db := db.GetDB()
	err := db.C("snippets").Find(queryParams).One(&item)
	return item, err
}
