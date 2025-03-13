package database

import (
	"errors"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type MarketPrice struct {
	GUID      uuid.UUID `gorm:"primaryKey" json:"guid"`
	AssetName string    `json:"asset_name"`
	PriceUsdt string    `json:"price_usdt"`
	Volume    string    `json:"volume"`
	Rate      string    `json:"rate"`
	Timestamp uint64
}

type MarketPriceView interface {
	QueryMarketPriceByAsset(string) ([]*MarketPrice, error)
}

type MarketPriceDB interface {
	MarketPriceView

	StoreMarketPrice([]MarketPrice) error
}

type marketPriceDB struct {
	gorm *gorm.DB
}

func (m *marketPriceDB) QueryMarketPriceByAsset(s string) ([]*MarketPrice, error) {
	var marketPriceList []*MarketPrice
	err := m.gorm.Table("market_price").Find(&marketPriceList).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return marketPriceList, nil
}

func (m *marketPriceDB) StoreMarketPrice(priceList []MarketPrice) error {
	result := m.gorm.Table("market_price").CreateInBatches(&priceList, len(priceList))
	return result.Error
}

func NewMarketPriceDB(db *gorm.DB) MarketPriceDB {
	return &marketPriceDB{gorm: db}
}
