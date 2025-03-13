package database

import (
	"errors"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type OfficialCoinRate struct {
	GUID      uuid.UUID `gorm:"primaryKey" json:"guid"`
	AssetName string    `json:"asset_name"`
	BaseAsset string    `json:"price_usdt"`
	Price     string    `json:"price"`
	Timestamp uint64
}

type OfficialCoinRateView interface {
	QueryOfficialCoinRateByAsset(string) ([]*OfficialCoinRate, error)
}

type OfficialCoinRateDB interface {
	OfficialCoinRateView

	StoreOfficialCoinRate([]OfficialCoinRate) error
}

type officialCoinRateDB struct {
	gorm *gorm.DB
}

func (o *officialCoinRateDB) QueryOfficialCoinRateByAsset(s string) ([]*OfficialCoinRate, error) {
	var rateList []*OfficialCoinRate
	err := o.gorm.Table("official_coin_rate").Find(&rateList).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return rateList, nil
}

func (o *officialCoinRateDB) StoreOfficialCoinRate(rateList []OfficialCoinRate) error {
	result := o.gorm.Table("official_coin_rate").CreateInBatches(&rateList, len(rateList))
	return result.Error
}

func NewOfficialCoinRateDB(db *gorm.DB) OfficialCoinRateDB {
	return &officialCoinRateDB{gorm: db}
}
