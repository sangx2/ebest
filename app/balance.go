package app

import (
	"fmt"
	"github.com/sangx2/ebest-sdk/ebest"
	"github.com/sangx2/ebest-sdk/res"
	"github.com/sangx2/ebest/model"
	log "github.com/sangx2/golog"
)

// InitBalance 잔고 정보 초기화
func (es *EBestServer) InitBalance() error {
	for _, account := range es.GetAccounts() {
		req := model.NewQueryRequest(ebest.T0424, false, res.T0424InBlock{Accno: account.Number})
		if err := es.requestServer.Request(ebest.T0424, req); err != nil {
			return fmt.Errorf("InitBalances: %v", err)
		} else {
			if resp := <-req.RespChan; resp.Error != nil {
				return fmt.Errorf("InitBalances: %v", resp.Error)
			} else {
				for _, outBlock := range resp.OutBlocks[1].([]res.T0424OutBlock1) {
					es.Balances[account.Number] = append(es.Balances[account.Number], model.NewBalance(outBlock))
				}
			}
		}
	}
	log.Info("balance 초기화 완료")

	return nil
}

func (es *EBestServer) UpdateBalances(accountNumber string) {
	var balances []*model.Balance

	req := model.NewQueryRequest(ebest.T0424, false, res.T0424InBlock{Accno: accountNumber})
	if e := es.requestServer.Request(ebest.T0424, req); e != nil {
		log.Error("UpdateBalances", log.Err(e))
	} else {
		if resp := <-req.RespChan; resp.Error != nil {
			log.Error("UpdateBalances", log.Err(resp.Error))
		} else {
			for _, outBlock := range resp.OutBlocks[1].([]res.T0424OutBlock1) {
				balances = append(balances, model.NewBalance(outBlock))
			}
			es.balancesMutex.Lock()
			es.Balances[accountNumber] = balances
			es.balancesMutex.Unlock()
		}
	}
}

func (es *EBestServer) GetBalances(accountNumber string) []*model.Balance {
	es.balancesMutex.RLock()
	balances := es.Balances[accountNumber]
	es.balancesMutex.RUnlock()

	return balances
}
