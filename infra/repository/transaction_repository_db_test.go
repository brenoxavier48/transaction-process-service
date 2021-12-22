package repository

import (
	"os"
	"testing"

	"github.com/brenoxavier48/imersaofc-gateway/domain/entity"
	"github.com/brenoxavier48/imersaofc-gateway/infra/repository/fixture"
	"github.com/stretchr/testify/assert"
)

func TestTransactionRepositoryDB_Insert(t *testing.T) {
	migrationDir := os.DirFS("fixture/sql")
	db := fixture.Up(migrationDir)
	defer fixture.Down(db, migrationDir)

	repository := NewTransactionRepositoryDB(db)
	err := repository.Insert("1", "1", 900, entity.ACCEPTED, "")
	assert.Nil(t, err)
}
