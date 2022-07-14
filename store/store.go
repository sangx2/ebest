package store

import "github.com/sangx2/ebest/model"

type Result struct {
	Data interface{}
	Err  error
}

type Channel chan Result

func Do(f func(result *Result)) Channel {
	storeChannel := make(Channel, 1)
	go func() {
		result := Result{}
		f(&result)
		storeChannel <- result
		close(storeChannel)
	}()
	return storeChannel
}

type Store interface {
	Asset() AssetStore
	Balance() BalanceStore
	FNG() FNGStore
}

type AssetStore interface {
	Get(accountNumber, date string) Channel
	Save(asset *model.Asset) Channel
}

type BalanceStore interface {
	Get(accountNumber, stockCode, date string) Channel
	Save(acntNo string, balance *model.Balance) Channel
}

type FNGStore interface {
	GetAll() Channel
	SaveAll(FNGs map[string]*model.FNG) Channel
}
