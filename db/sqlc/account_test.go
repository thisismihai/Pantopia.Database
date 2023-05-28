package db

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/mbaxamb3/pantopia/util"
	"github.com/stretchr/testify/require"
)

func CreateRandomAccount(t *testing.T) Account {
	arg := CreateAccountParams{
		Owner:          util.RandomOwner(),
		CompanyName:    util.RandomCompanyName(),
		Industry:       util.RandomIndustry(),
		TargetIndustry: util.RandomIndustry(),
		TargetPosition: util.RandomPosition(),
		TargetSize:     util.RandomTargetSize(),
	}

	account, err := testQueries.CreateAccount(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, account)

	require.Equal(t, arg.Owner, account.Owner)
	require.Equal(t, arg.CompanyName, account.CompanyName)
	require.Equal(t, arg.Industry, account.Industry)
	require.Equal(t, arg.TargetIndustry, account.TargetIndustry)
	require.Equal(t, arg.TargetPosition, account.TargetPosition)
	require.Equal(t, arg.TargetSize, account.TargetSize)

	require.NotZero(t, account.ID)
	require.NotZero(t, account.CreatedAt)

	return account
}

func TestCreateAccount(t *testing.T) {
	CreateRandomAccount(t)
}

func TestGetAccoun(t *testing.T) {
	//create user account
	account1 := CreateRandomAccount(t)
	account2, err := testQueries.GetAccount(context.Background(), account1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, account2)

	require.Equal(t, account1.ID, account2.ID)
	require.Equal(t, account1.Owner, account2.Owner)
	require.Equal(t, account1.Industry, account2.Industry)
	require.Equal(t, account1.TargetIndustry, account2.TargetIndustry)

	require.WithinDuration(t, account1.CreatedAt, account2.CreatedAt, time.Second)
}

func TestUpdateAccount(t *testing.T) {
	account1 := CreateRandomAccount(t)

	arg := UpdateAccountParams{
		ID:          account1.ID,
		CompanyName: util.RandomCompanyName(),
	}
	account2, err := testQueries.UpdateAccount(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, account2)

	require.Equal(t, account1.ID, account2.ID)
	require.Equal(t, account1.Owner, account2.Owner)
	require.Equal(t, account1.Industry, account2.Industry)
	require.Equal(t, account1.TargetIndustry, account2.TargetIndustry)

	require.WithinDuration(t, account1.CreatedAt, account2.CreatedAt, time.Second)

}

func TestDeleteAccount(t *testing.T) {
	account1 := CreateRandomAccount(t)
	err := testQueries.DeleteAccount(context.Background(), account1.ID)

	require.NoError(t, err)

	account2, err := testQueries.GetAccount(context.Background(), account1.ID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, account2)
}
