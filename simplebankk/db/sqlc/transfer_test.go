package db

import (
	"context"
	"testing"
	"time"

	"github.com/GO/simplebankk/util"
	"github.com/stretchr/testify/require"
)

func createRandomTransfer(t *testing.T, account1, account2 Account) Transfer {

	arg := CreateTransferParams{
		FromAccountID: account1.ID,
		ToAccountID:   account2.ID,
		Amount:        util.RandomMoney(),
	}

	transfer, err := testQueries.CreateTransfer(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, transfer)

	require.Equal(t, arg.FromAccountID, transfer.FromAccountID)
	require.Equal(t, arg.ToAccountID, transfer.ToAccountID)
	require.Equal(t, arg.Amount, transfer.Amount)

	require.NotZero(t, transfer.ID)
	require.NotZero(t, transfer.CreatedAt)

	return transfer

}

func TestCreateTransfer(t *testing.T) {
	account1 := createRandomAccount(t)
	account2 := createRandomAccount(t)
	createRandomTransfer(t, account1, account2)
}

func TestGetTransfer(t *testing.T) {
	account1 := createRandomAccount(t)
	account2 := createRandomAccount(t)
	transfer1 := createRandomTransfer(t, account1, account2)

	transfer2, err := testQueries.GetTransfer(context.Background(), transfer1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, transfer2)

	require.Equal(t, transfer1.ID, transfer2.ID)
	require.Equal(t, transfer1.FromAccountID, transfer2.FromAccountID)
	require.Equal(t, transfer1.ToAccountID, transfer2.ToAccountID)
	require.Equal(t, transfer1.Amount, transfer2.Amount)
	require.WithinDuration(t, transfer1.CreatedAt, transfer2.CreatedAt, time.Second)
}

func TestListTransfers(t *testing.T) {
	// 1. Creează două conturi random în baza de date folosind funcții ajutătoare.
	account1 := createRandomAccount(t)
	account2 := createRandomAccount(t)

	// 2. Creează transferuri între cele două conturi. De cinci ori din contul 1 în contul 2 și de cinci ori invers.
	for i := 0; i < 5; i++ {
		createRandomTransfer(t, account1, account2)
		createRandomTransfer(t, account2, account1)
	}

	// 3. Definește parametrii pentru listarea transferurilor, specificând contul 1 ca fiind atât `FromAccountID` cât și `ToAccountID`, limitând la 5 rezultate și cu un offset de 5.
	arg := ListTransfersParams{
		FromAccountID: account1.ID,
		ToAccountID:   account1.ID,
		Limit:         5,
		Offset:        5,
	}

	// 4. Apelează funcția `ListTransfers` pentru a obține lista de transferuri conform parametrilor specificați.
	transfers, err := testQueries.ListTransfers(context.Background(), arg)

	// 5. Verifică dacă nu a apărut nicio eroare la recuperarea transferurilor.
	require.NoError(t, err)

	// 6. Verifică dacă numărul de transferuri recuperate este exact 5.
	require.Len(t, transfers, 5)

	// 7. Pentru fiecare transfer din lista recuperată:
	for _, transfer := range transfers {
		// a. Verifică dacă transferul nu este gol.
		require.NotEmpty(t, transfer)

		// b. Verifică dacă `FromAccountID` sau `ToAccountID` al transferului este egal cu `account1.ID`.
		require.True(t, transfer.FromAccountID == account1.ID || transfer.ToAccountID == account1.ID)
	}
}
