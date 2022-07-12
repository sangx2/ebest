package filestore

import (
	"fmt"
	"github.com/sangx2/ebest/model"
	"github.com/sangx2/ebest/store"
	"os"
)

type FileAssetStore struct {
	FileStore
}

func NewFileAssetStore(store FileStore) *FileAssetStore {
	return &FileAssetStore{
		FileStore: store,
	}
}

func (f *FileAssetStore) Get(account, date string) store.Channel {
	return store.Do(func(result *store.Result) {
		path := f.GetBasePath() + assetsPath + account + "\\" + date + FileType

		if _, e := os.Stat(path); e == nil {
			if file, e := os.Open(path); e != nil {
				result.Err = fmt.Errorf("Get: %v", e)
			} else {
				result.Data = model.AssetFromJson(file)
			}
		} else if !os.IsNotExist(e) {
			result.Err = fmt.Errorf("Get: %v", e)
		}
	})
}

func (f *FileAssetStore) Save(asset *model.Asset) store.Channel {
	return store.Do(func(result *store.Result) {
		path := f.GetBasePath() + assetsPath + "\\" + asset.AccountNum
		if _, e := os.Stat(path); os.IsNotExist(e) {
			e := os.MkdirAll(path, os.ModeDir)
			if e != nil {
				result.Err = fmt.Errorf("Create: %v", e)
				return
			}
		}

		file, e := os.OpenFile(path+"\\"+asset.Date+FileType, os.O_CREATE|os.O_WRONLY, 0644)
		if e != nil {
			result.Err = fmt.Errorf("Save: %v", e)
			return
		}
		defer file.Close()

		_, e = file.Write(asset.ToJson())
		if e != nil {
			result.Err = fmt.Errorf("Save: %v", e)
			return
		}
	})
}
