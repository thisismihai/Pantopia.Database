package db

import (
	"context"
	"testing"
	"time"

	"github.com/mbaxamb3/pantopia/util"
	"github.com/stretchr/testify/require"
)

func CreateRandomLead(t *testing.T, account Account) Lead {

	arg := CreateLeadParams{
		AccountID:       account.ID,
		Email:           util.RandomEmail(),
		TelephoneNumber: util.RandomTelephoneNumber(),
		TargetEmail:     util.RandomEmail(),
		CompanyName:     util.RandomCompanyName(),
		Conversation:    util.RandomShortText(10),
	}

	lead, err := testQueries.CreateLead(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, lead)

	require.Equal(t, arg.AccountID, lead.AccountID)
	require.Equal(t, arg.Email, lead.Email)
	require.Equal(t, arg.TelephoneNumber, lead.TelephoneNumber)
	require.Equal(t, arg.TargetEmail, lead.TargetEmail)
	require.Equal(t, arg.CompanyName, lead.CompanyName)
	require.Equal(t, arg.Conversation, lead.Conversation)

	return lead
}

func TestCreateLead(t *testing.T) {
	account := CreateRandomAccount(t)
	CreateRandomLead(t, account)
}

func TestGetLead(t *testing.T) {
	//create user account
	lead1 := CreateRandomAccount(t)
	lead2, err := testQueries.GetAccount(context.Background(), lead1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, lead2)

	require.Equal(t, lead1.ID, lead2.ID)

	require.WithinDuration(t, lead1.CreatedAt, lead2.CreatedAt, time.Second)
}

func TestListLeads(t *testing.T) {
	account := CreateRandomAccount(t)
	for i := 0; i < 10; i++ {
		CreateRandomLead(t, account)
	}
	arg := ListLeadsParams{
		AccountID: account.ID,
		Limit:     5,
		Offset:    5,
	}
	leads, err := testQueries.ListLeads(context.Background(), arg)

	require.NoError(t, err)
	require.Len(t, leads, 5)

	for _, lead := range leads {
		require.NotEmpty(t, lead)
		require.Equal(t, arg.AccountID, lead.AccountID)
	}
}
