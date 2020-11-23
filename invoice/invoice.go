package invoice

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

// Robokassa log item:
// ID
// UserID
// Amount
// Op_debit
// Status
// CreatedAt
// RKInvoiceID
//
// Internal invoice log item:
// ID
// UserID
// Amount
// Op_credit
// Status
// CreatedAt
// CompanyID
// MonthAmount
type Invoice struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	UserID      primitive.ObjectID `bson:"u,omitempty"`
	CompanyID   primitive.ObjectID `bson:"c,omitempty"`
	RKInvoiceID uint64             `bson:"r,omitempty"`
	MonthAmount uint32             `bson:"ma,omitempty"`
	CreatedAt   time.Time          `bson:"ca,omitempty"`
	Amount      uint32             `bson:"a,omitempty"`
	Op          op                 `bson:"o,omitempty"`
	Status      status             `bson:"s,omitempty"`
}

type op uint8

const (
	_ op = iota
	Op_debit
	Op_credit
)

type status uint8

const (
	_ status = iota
	Status_pending
	Status_success
	Status_fail
)
