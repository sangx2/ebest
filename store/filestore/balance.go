package filestore

import (
	"fmt"
	"github.com/sangx2/ebest/model"
	"github.com/sangx2/ebest/store"
	"os"
)

type FileBalanceStore struct {
	FileStore
}

func NewFileBalanceStore(store FileStore) *FileBalanceStore {
	return &FileBalanceStore{
		FileStore: store,
	}
}

func (f *FileBalanceStore) Get(acntNo, stockCode, date string) store.Channel {
	return store.Do(func(result *store.Result) {
		path := f.GetBasePath() + balancesPath + "\\" + acntNo + "\\" + stockCode + "\\" + date + FileType

		if _, e := os.Stat(path); e == nil {
			if file, e := os.Open(path); e != nil {
				result.Err = fmt.Errorf("Get: %v", e)
			} else {
				result.Data = model.BalanceFromJson(file)
			}
		} else if !os.IsNotExist(e) {
			result.Err = fmt.Errorf("Get: %v", e)
		}
	})
}

func (f *FileBalanceStore) Save(acntNo string, balance *model.Balance) store.Channel {
	return store.Do(func(result *store.Result) {
		path := f.GetBasePath() + balancesPath + "\\" + acntNo + "\\" + balance.Expcode
		if _, e := os.Stat(path); os.IsNotExist(e) {
			e := os.MkdirAll(path, os.ModeDir)
			if e != nil {
				result.Err = fmt.Errorf("Create: %v", e)
				return
			}
		}

		file, e := os.OpenFile(path+"\\"+balance.Date+FileType, os.O_CREATE|os.O_WRONLY, 0644)
		if e != nil {
			result.Err = fmt.Errorf("Save: %v", e)
			return
		}
		defer file.Close()

		_, e = file.Write(balance.ToJson())
		if e != nil {
			result.Err = fmt.Errorf("Save: %v", e)
			return
		}
	})
}
