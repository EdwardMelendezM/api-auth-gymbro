package mysql

import (
	"context"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"

	mockClock "github.com/EdwardMelendezM/info-code-api-shared-v1/clock/mocks"
)

func TestRepositoryAuth_CheckExistenceByUsername(t *testing.T) {
	_, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	t.Run("When check existence with username user", func(t *testing.T) {
		username := "user@mail.com"
		existingUser := 1
		row := sqlmock.NewRows([]string{"total"}).AddRow(existingUser)

		mock.
			ExpectQuery(QueryCheckExistenceByUsername).
			WithArgs(username).
			WillReturnRows(row)

		clock := &mockClock.Clock{}
		r := NewAuthRepository(clock, 60)
		var totalExpected int
		totalExpected = r.CheckExistenceByUsername(context.Background(), username)
		assert.Equal(t, existingUser, totalExpected)
	})
}
