package filestore

import (
	"github.com/sangx2/ebest/model"
	"github.com/sangx2/ebest/store"
	"io/ioutil"
	"os"
)

type FileFNGStore struct {
	FileStore
}

func NewFileFNGStore(store FileStore) *FileFNGStore {
	return &FileFNGStore{
		FileStore: store,
	}
}

func (f *FileFNGStore) GetAll() store.Channel {
	return store.Do(func(result *store.Result) {
		FNGs := make(map[string]*model.FNG)

		path := f.GetBasePath() + fngPath

		if fileInfos, e := ioutil.ReadDir(path); e != nil {
			result.Err = e
			return
		} else {
			for _, fileInfo := range fileInfos {
				if !fileInfo.IsDir() {
					if file, e := os.Open(path + fileInfo.Name() + FileType); e != nil {
						if os.IsNotExist(e) {
							continue
						} else {
							result.Err = e
						}
					} else {
						FNGs[fileInfo.Name()] = model.FNGFromJson(file)
					}
				}
			}
		}

		result.Data = FNGs
	})
}

func (f *FileFNGStore) SaveAll(FNGs map[string]*model.FNG) store.Channel {
	return store.Do(func(result *store.Result) {
		path := f.GetBasePath() + fngPath

		for _, fng := range FNGs {
			if fng != nil {
				file, e := os.OpenFile(path+"\\"+fng.Code+FileType, os.O_CREATE|os.O_WRONLY, 0644)
				if e != nil {
					result.Err = e
					return
				}

				_, e = file.Write(fng.ToJson())
				if e != nil {
					result.Err = e
					return
				}
			}
		}
	})
}
