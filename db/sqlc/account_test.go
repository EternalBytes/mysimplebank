package db

import (
	"database/sql"
	"testing"
	"time"

	"github.com/eternalbytes/simplebank/util"
	"github.com/stretchr/testify/require"
)

func createRandomAccount(t *testing.T) *Account {
	arg := CreateAccountParams{
		Owner:    util.RandomOwner(),
		Balance:  util.RandomMoney(),
		Currency: util.RandomCurrency(),
	}

	acc, err := testQueries.CreateAccount(ctx, arg)
	require.NoError(t, err)
	require.NotEmpty(t, acc)

	require.Equal(t, arg.Owner, acc.Owner)
	require.Equal(t, arg.Balance, acc.Balance)
	require.Equal(t, arg.Currency, acc.Currency)

	require.NotZero(t, acc.ID)
	require.NotZero(t, acc.CreatedAt)
	return &acc
}

func TestCreateAccount(t *testing.T) {
	createRandomAccount(t)
}

func TestGetAccount(t *testing.T) {
	acc := createRandomAccount(t)
	acc2, err := testQueries.GetAccount(ctx, acc.ID)
	require.NoError(t, err)
	require.NotEmpty(t, acc2)

	require.Equal(t, acc.ID, acc2.ID)
	require.Equal(t, acc.Owner, acc2.Owner)
	require.Equal(t, acc.Balance, acc2.Balance)
	require.Equal(t, acc.Currency, acc2.Currency)
	require.WithinDuration(t, acc.CreatedAt, acc2.CreatedAt, time.Second)
	deleteAccountFunc(t, acc.ID)
}

func deleteAccountFunc(t *testing.T, id int64) {
	err := testQueries.DeleteAccount(ctx, id)
	require.NoError(t, err)

	acc2, err := testQueries.GetAccount(ctx, id)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, acc2)
}

/* func TestListAccounts(t *testing.T) {
	arg2 := ListAccountsParams{
		Limit:  5,
		Offset: 0, // from the first line
	}

	accSlc, err := testQueries.ListAccounts(context.Background(), arg2)
	require.NoError(t, err)
	require.Len(t, accSlc, 5)

	for _, acc := range accSlc {
		require.NotEmpty(t, acc)
	}

	fmt.Println(accSlc)
} */
