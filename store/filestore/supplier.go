package filestore

import (
	"github.com/sangx2/ebest/store"
	"os"
)

var PATHS = []string{
	assetsPath,
	balancesPath,
	fngPath,
}

const (
	assetsPath   = "\\assets"
	balancesPath = "\\balances"
	fngPath      = "\\FNGs"

	FileType = ".json"
)

type FileSupplierStores struct {
	asset   store.AssetStore
	balance store.BalanceStore
	FNG     store.FNGStore
}

type FileSupplier struct {
	basePath string

	fileInfo os.FileInfo

	stores FileSupplierStores
}

func NewFileSupplier(basePath string) *FileSupplier {
	supplier := &FileSupplier{
		basePath: basePath,
	}

	if e := supplier.Open(); e != nil {
		return nil
	}

	supplier.stores.asset = NewFileAssetStore(supplier)
	supplier.stores.balance = NewFileBalanceStore(supplier)
	supplier.stores.FNG = NewFileFNGStore(supplier)

	return supplier
}

func (fs *FileSupplier) Open() error {
	if fileInfo, e := os.Stat(fs.basePath); os.IsNotExist(e) {
		e := os.MkdirAll(fs.basePath, os.ModeDir)
		if e != nil {
			return e
		}

		fs.fileInfo = fileInfo
	}

	for _, path := range PATHS {
		if _, e := os.Stat(fs.basePath + path); os.IsNotExist(e) {
			e := os.MkdirAll(fs.basePath+path, os.ModeDir)
			if e != nil {
				return e
			}
		}
	}

	return nil
}

func (fs *FileSupplier) Close() error {
	return nil
}

func (fs *FileSupplier) GetBasePath() string {
	return fs.basePath
}

func (fs *FileSupplier) Asset() store.AssetStore {
	return fs.stores.asset
}

func (fs *FileSupplier) Balance() store.BalanceStore {
	return fs.stores.balance
}

func (fs *FileSupplier) FNG() store.FNGStore {
	return fs.stores.FNG
}
