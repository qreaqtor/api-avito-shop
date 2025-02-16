package schema

import (
	"github.com/uptrace/bun"
)

type MerchSchema struct {
	bun.BaseModel `bun:"table:merch"`

	Type  string `bun:"merch_type,pk"`
	Price uint   `bun:"price,notnull"`
}
