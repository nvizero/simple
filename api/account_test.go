package api

import (
	db "simple/db/sqlc"
	"simple/mockdb"
	"simple/util"
	"testing"

	"github.com/golang/mock/gomock"
)

func TestGetAccountAPI(t *testing.T) {
	account := rendomAccount()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	store := mockdb.NewMockStore(ctrl)

	store.EXPECT().
		GetAccount(gomock.Any(), gomock.Eq(account.ID)).
		Times(1).
		Return(account, nil)

}

func rendomAccount() db.Account {
	return db.Account{
		ID:       util.RandomInt(1, 1000),
		Owner:    util.RandomOwner(),
		Balance:  util.RandomMoney(),
		Currency: util.RandomCurrency(),
	}
}
