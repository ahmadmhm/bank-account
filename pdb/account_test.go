package bankAccount

import (
	"context"
    "testing"
    "time"

	"github.com/stretchr/testify/require"
	"bankAccount/util"
)

func CreateRandomAccount(t *testing.T) Account {
	arg := CreateAccountParams {
		Owner : util.RandomName(),
		Balance:  util.RandomBalance(),
		Currency: util.RandomCurrency(),
	}

	account, err := testQueries.CreateAccount(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, account)
	require.Equal(t, arg.Owner, account.Owner)
	require.Equal(t, arg.Balance, account.Balance)
	require.Equal(t, arg.Currency, account.Currency)
	require.NotZero(t, account.ID)
	require.NotZero(t, account.CreatedAt)

	return account
}

func TestCreateAccount(t *testing.T) {
	CreateRandomAccount(t)
}


func TestGetAccount(t *testing.T) {
	account := CreateRandomAccount(t)
	gotAccount, err := testQueries.GetAccount(context.Background(), account.ID)

	require.NoError(t, err)
	require.NotEmpty(t, gotAccount)
	require.Equal(t, account.ID, gotAccount.ID)
	require.Equal(t, account.Owner, gotAccount.Owner)
	require.Equal(t, account.Balance, gotAccount.Balance)
	require.Equal(t, account.Currency, gotAccount.Currency)
	require.Equal(t, account.CreatedAt, gotAccount.CreatedAt)
	require.WithinDuration(t, account.CreatedAt, gotAccount.CreatedAt, time.Second)
}

func TestUpdateAccount(t *testing.T) {
	account := CreateRandomAccount(t)
	arg := UpdateAccountParams {
		ID : account.ID,
		Owner:  account.Owner,
		Balance:  util.RandomBalance(),
		Currency:  account.Currency,
	}
	err := testQueries.UpdateAccount(context.Background(), arg)
	gotAccount, getErr := testQueries.GetAccount(context.Background(), account.ID)

	require.NoError(t, err)
	require.NoError(t, getErr)
	require.NotEmpty(t, gotAccount)
	require.Equal(t, account.ID, gotAccount.ID)
	require.Equal(t, account.Owner, gotAccount.Owner)
	require.Equal(t, arg.Balance, gotAccount.Balance)
	require.Equal(t, account.Currency, gotAccount.Currency)
	require.Equal(t, account.CreatedAt, gotAccount.CreatedAt)
	require.WithinDuration(t, account.CreatedAt, gotAccount.CreatedAt, time.Second)
}


func TestDeleteAccount(t *testing.T) {
	account := CreateRandomAccount(t)

	err := testQueries.DeleteAccount(context.Background(), account.ID)
	gotAccount, getErr := testQueries.GetAccount(context.Background(), account.ID)

	require.NoError(t, err)
	require.ErrorContains(t, getErr, "no rows in result set")
	require.Empty(t, gotAccount)
}

func TestListAccounts(t *testing.T) {
	var lastAccount Account
	for i := 0; i < 10; i++ {
		lastAccount = CreateRandomAccount(t)
	}
	arg := ListAccountsParams {
		Owner:  lastAccount.Owner,
		Limit:  5,
		Offset: 0,
	} 

	accounts, err := testQueries.ListAccounts(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, accounts)
	require.Len(t, accounts, 1)
}