package repository

import (
	"assignment/internal/repository/testdata"
	"context"
	"database/sql"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/require"
)

func TestImpl_CheckUserByEmail(t *testing.T) {

	testCases := []struct {
		Name          string
		WantDBFail    bool
		UserEmail     string
		ExpectedExist bool
		ExpectedErr   error
	}{
		{
			Name:          "Success",
			WantDBFail:    false,
			UserEmail:     "new-user-email@example.com",
			ExpectedExist: false,
			ExpectedErr:   nil,
		},
		{
			Name:          "User Already Exist",
			WantDBFail:    false,
			UserEmail:     "already-exist-email@example.com",
			ExpectedExist: true,
			ExpectedErr:   nil,
		},
		{
			Name:          "Internal Server Error",
			WantDBFail:    true,
			UserEmail:     "internal-error-email@example.com",
			ExpectedExist: false,
			ExpectedErr:   InternalErrorCheckEmail,
		},
	}

	ctx := context.Background()

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {

			testdata.TestDatabase(t, func(tx *sql.Tx) {

				testdata.LoadTestSQLFile(t, tx, "testdata/testdata_for_user_accounts.sql")

				repo := New(tx)
				if tc.WantDBFail {
					dbMock, _, _ := sqlmock.New()
					repo = New(dbMock)
				}

				exists, err := repo.CheckUserByEmail(ctx, tc.UserEmail)

				if err != nil {
					require.EqualError(t, err, tc.ExpectedErr.Error())
				} else {
					require.Equal(t, tc.ExpectedExist, exists)
				}
			})
		})
	}

}
