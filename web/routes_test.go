package web

import (
	"log"
	"os"
	"rvweb/repository"
	"testing"

	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
)

var r repository.Repository

// test setup
func TestMain(m *testing.M) {

	db, err := sqlx.Connect("sqlite3", ":memory:")
	if err != nil {
		log.Fatalln(err)
	}

	err = repository.Migrate(db.DB, "file://../migrations")
	if err != nil {
		log.Fatalln(err)
	}

	r = repository.Repository{DB: db}

	code := m.Run()
	os.Exit(code)
}

func TestGetPosts(t *testing.T) {

	posts, err := r.GetPosts()
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, 0, len(*posts))
}

func TestCreatePosts(t *testing.T) {

	_, err := r.CreatePost("title", "body")
	if err != nil {
		t.Error(err)
	}

	c, err := r.GetPosts()
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, 1, len(*c))
}
