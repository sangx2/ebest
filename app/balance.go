package app

import (
	"fmt"
	"github.com/sangx2/ebest-sdk/ebest"
	"github.com/sangx2/ebest-sdk/res"
	"github.com/sangx2/ebest/model"
	"github.com/sangx2/ebest/store"
	"github.com/sangx2/ebest/utils"
	log "github.com/sangx2/golog"
)

func (es *EBestServer) InitBalance() error {
	// 1. 계좌 목록 조회
	for _, account := range es.GetAccounts() {
		req := model.NewQueryRequest(ebest.T0424, false, res.T0424InBlock{Accno: account.Number})
		if err := es.requestServer.Request(ebest.T0424, req); err != nil {
			return fmt.Errorf("InitBalances: %v", err)
		} else {
			if resp := <-req.RespChan; resp.Error != nil {
				return fmt.Errorf("InitBalances: %v", resp.Error)
			} else {
				for _, outBlock := range resp.OutBlocks[1].([]res.T0424OutBlock1) {
					// 2. 데이터 조회
					queryBalance := <-es.store.Balance().Get(account.Number, outBlock.Expcode, utils.GetDateString())
					if queryBalance.Err != nil {
						log.Error("InitBalances", log.Err(queryBalance.Err))
						continue
					}

					var balance *model.Balance
					if queryBalance.Data != nil {
						balance = queryBalance.Data.(*model.Balance)
						balance.T0424OutBlock1 = outBlock
					} else {
						balance = model.NewBalance(outBlock)
					}

					es.Balances[account.Number] = append(es.Balances[account.Number], balance)
				}
			}
		}
	}
	log.Info("balance 초기화 완료")

	return nil
}

func (es *EBestServer) FinalizeBalance() error {
	for acntNo, balances := range es.Balances {
		var queryBalance store.Result

		for _, balance := range balances {
			queryBalance = <-es.store.Balance().Save(acntNo, balance)
			if queryBalance.Err != nil {
				return fmt.Errorf("FinalizeBalances: %v", queryBalance.Err)
			}
		}
	}
	log.Info("balance 정보 저장 완료")

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
