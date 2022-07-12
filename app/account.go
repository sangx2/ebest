package app

import (
	"github.com/sangx2/ebest/model"
	log "github.com/sangx2/golog"
)

func (es *EBestServer) InitAccount() error {
	for _, accountNum := range es.EBest.GetAccountList() {
		name := es.EBest.GetAccountName(accountNum)
		detailName := es.EBest.GetAccountDetailName(accountNum)
		nickName := es.EBest.GetAccountNickName(accountNum)

		account := model.NewAccount(accountNum, name, detailName, nickName)
		es.Accounts = append(es.Accounts, account)
	}
	log.Info("account 초기화 완료")

	return nil
}

func (es *EBestServer) GetAccounts() []*model.Account {
	es.accountsMutex.RLock()
	accounts := es.Accounts
	es.accountsMutex.RUnlock()

	return accounts
}
