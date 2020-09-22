package orm

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOpen(t *testing.T) {
	//testutil.Ok(t, CreateORM("gorm.db"))
	var orm = new(ORM)

	assert.Nil(t, orm.Open("doyo.db"))
}
