package app

import (
	"fmt"
	"github.com/sangx2/ebest-sdk/ebest"
	"github.com/sangx2/ebest-sdk/res"
	"github.com/sangx2/ebest/model"
	log "github.com/sangx2/golog"
)

// InitAssets 자산 정보 초기화
// - 계좌 정보 초기화(InitAccounts)가 요구됨
func (es *EBestServer) InitAssets() error {
	// 1. 계좌 목록 조회
	for _, account := range es.GetAccounts() {
		pwd, ok := es.config.AccountSettings.Accounts[account.Number]
		if !ok {
			log.Warn("InitAssets: 설정 파일에 계좌번호 없음", log.String("계좌번호", account.Number))
			continue
		}
		cspaq12200Req := model.NewQueryRequest(ebest.CSPAQ12200, false,
			res.CSPAQ12200InBlock1{RecCnt: "1", AcntNo: account.Number, Pwd: pwd})
		if err := es.requestServer.Request(ebest.CSPAQ12200, cspaq12200Req); err != nil {
			return err
		} else {
			if resp := <-cspaq12200Req.RespChan; resp.Error != nil {
				return fmt.Errorf("InitAssets: %v", resp.Error)
			} else {
				es.Assets[account.Number] = model.NewAsset(account.Number, resp.OutBlocks[1].(res.CSPAQ12200OutBlock2))
			}
		}
	}
	log.Info("asset 초기화 완료")

	return nil
}

func (es *EBestServer) UpdateAssets() {
	for _, account := range es.GetAccounts() {
		pwd, ok := es.config.AccountSettings.Accounts[account.Number]
		if !ok {
			log.Warn("UpdateAssets: 설정 파일에 계좌번호 없음", log.String("account", account.Number))
			continue
		}
		cspaq12200Req := model.NewQueryRequest(ebest.CSPAQ12200, false,
			res.CSPAQ12200InBlock1{RecCnt: "1", AcntNo: account.Number, Pwd: pwd})
		if err := es.requestServer.Request(ebest.CSPAQ12200, cspaq12200Req); err != nil {
			log.Error("UpdateAssets", log.Err(err))
		} else {
			if resp := <-cspaq12200Req.RespChan; resp.Error != nil {
				log.Error("UpdateAssets", log.Err(resp.Error))
			} else {
				es.assetsMutex.Lock()
				asset := es.Assets[account.Number]
				asset.CSPAQ12200OutBlock2 = resp.OutBlocks[1].(res.CSPAQ12200OutBlock2)
				es.Assets[account.Number] = asset
				es.assetsMutex.Unlock()
			}
		}
	}
}

func (es *EBestServer) GetAssets() map[string]*model.Asset {
	es.assetsMutex.RLock()
	assets := es.Assets
	es.assetsMutex.RUnlock()

	return assets
}
