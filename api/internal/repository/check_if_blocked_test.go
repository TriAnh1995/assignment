package repository

import (
	"assignment/internal/repository/testdata"
	"context"
	"database/sql"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/require"
)

func TestRepoImplement_CheckIfBlocked(t *testing.T) {
	testCases := []struct {
		Name          string
		Emails        []string
		ExpectedExist bool
		ExpectedErr   error
		WantDBFail    bool
	}{
		{
			Name:          "Success",
			Emails:        []string{"new_requester@example.com", "new_target@example.com"},
			ExpectedErr:   nil,
			ExpectedExist: false,
			WantDBFail:    false,
		},
		{
			Name:          "Blocked Target",
			Emails:        []string{"requester2@example.com", "target2@example.com"},
			ExpectedErr:   nil,
			ExpectedExist: true,
			WantDBFail:    false,
		},
		{
			Name:          "Non Blocked Target",
			Emails:        []string{"requester1@example.com", "target1@example.com"},
			ExpectedExist: false,
			ExpectedErr:   nil,
			WantDBFail:    false,
		},
		{
			Name:          "Internal Server Error",
			Emails:        []string{"error_requester@example.com", "error_target@example.com"},
			ExpectedErr:   InternalErrorCheckIfBlocked,
			ExpectedExist: false,
			WantDBFail:    true,
		},
	}
	ctx := context.Background()
	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			testdata.TestDatabase(t, func(tx *sql.Tx) {
				testdata.LoadTestSQLFile(t, tx, "testdata/testdata_for_subscription.sql")
				repo := New(tx)
				if tc.WantDBFail {
					dbMock, _, _ := sqlmock.New()
					repo = New(dbMock)
				}

				exists, err := repo.CheckIfBlocked(ctx, tc.Emails)

				if err != nil {
					require.EqualError(t, err, tc.ExpectedErr.Error())
				} else {
					require.Equal(t, tc.ExpectedExist, exists)
				}
			})
		})
	}
}
