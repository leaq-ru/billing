package invoice

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

// Robokassa log item:
// ID
// UserID
// Amount
// Op_income
// Status
//
// Internal invoice log item:
// ID
// UserID
// Amount
// Op_outcome
// Status
// CompanyID
// CompanyPremiumDeadline
type Invoice struct {
	ID                     primitive.ObjectID `bson:"_id,omitempty"`
	UserID                 primitive.ObjectID `bson:"u,omitempty"`
	CompanyID              primitive.ObjectID `bson:"c,omitempty"`
	CompanyPremiumDeadline time.Time          `bson:"cp,omitempty"`
	Amount                 uint32             `bson:"a,omitempty"`
	Op                     op                 `bson:"o,omitempty"`
	Status                 status             `bson:"s,omitempty"`
}

type op uint8

const (
	_ op = iota
	Op_income
	Op_outcome
)

type status uint8

const (
	_ status = iota
	Status_pending
	Status_success
	Status_fail
)
