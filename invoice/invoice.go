package invoice

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Invoice struct {
	ID                   primitive.ObjectID    `bson:"_id,omitempty"`
	UserID               primitive.ObjectID    `bson:"u,omitempty"`
	CreatedAt            time.Time             `bson:"ca,omitempty"`
	Amount               uint32                `bson:"a,omitempty"`
	Status               status                `bson:"s,omitempty"`
	Kind                 kind                  `bson:"k,omitempty"`
	DebitRobokassa       *debitRobokassa       `bson:"dr,omitempty"`
	CreditCompanyPremium *creditCompanyPremium `bson:"cc,omitempty"`
	CreditDataPremium    *creditDataPremium    `bson:"cd,omitempty"`
}

type debitRobokassa struct {
	InvoiceID uint64 `bson:"i,omitempty"`
}

type creditCompanyPremium struct {
	CompanyID   primitive.ObjectID `bson:"c,omitempty"`
	MonthAmount uint32             `bson:"m,omitempty"`
}

type creditDataPremium struct {
	MonthAmount uint32 `bson:"m,omitempty"`
}

type kind uint8

const (
	_ kind = iota
	kind_debitRobokassa
	kind_creditCompanyPremium
	kind_debitManual
	kind_creditDataPremium
)

type status uint8

const (
	_ status = iota
	status_pending
	status_success
	status_fail
)
