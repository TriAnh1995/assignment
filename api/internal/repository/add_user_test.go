package repository

import (
	"assignment/internal/model"
	"assignment/internal/repository/testdata"
	"context"
	"database/sql"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/require"
)

func TestImpl_AddUserToDatabase(t *testing.T) {

	testCases := []struct {
		Name        string
		User        model.User
		ExpectedErr error
		WantDBFail  bool
	}{
		{
			Name:        "Success",
			User:        model.User{Name: "NewUser", Email: "new-user-email@example.com"},
			ExpectedErr: nil,
			WantDBFail:  false,
		},
		{
			Name:        "Error",
			User:        model.User{Name: "AlreadyExistUser", Email: "already-exist-email@example.com"},
			ExpectedErr: UserAlreadyExist,
			WantDBFail:  false,
		},
		{
			Name:        "Internal Server Error",
			User:        model.User{Name: "InternalErrorUser", Email: "internal-error-email@example.com"},
			ExpectedErr: InternalErrorAddUser,
			WantDBFail:  true,
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

				err := repo.AddUser(ctx, tc.User)

				if err != nil {
					require.EqualError(t, err, tc.ExpectedErr.Error())
				} else {
					require.Equal(t, err, tc.ExpectedErr)
				}
			})
		})
	}

}
