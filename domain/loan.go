package domain

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Loan Status Enum
type LoanStatus string

const (
	Statusapproved LoanStatus = "approved"
	Statusrejected LoanStatus = "rejected"
)

// Loan Category Enum
type LoanCategory string

const (
	CategoryPersonalLoan LoanCategory = "Personal Loans"
	CategoryAutoLoan     LoanCategory = "Auto Loans"
	CategoryMortgageLoan LoanCategory = "Mortgage Loans"
)

// Loan struct
type Loan struct {
	ID          primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Amount      float64            `json:"amount,omitempty" bson:"amount,omitempty"`
	Category    LoanCategory       `json:"category,omitempty" bson:"category,omitempty"`
	Description string             `json:"description,omitempty" bson:"description,omitempty"`
	CreatedAt   time.Time          `json:"created_at,omitempty" bson:"created_at,omitempty"`
	UpdatedAt   time.Time          `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
	Status      LoanStatus         `json:"status,omitempty" bson:"status,omitempty"`
}

// CreateLoan struct for input
type CreateLoan struct {
	Amount      float64      `json:"amount,omitempty" bson:"amount,omitempty"`
	Category    LoanCategory `json:"category,omitempty" bson:"category,omitempty"`
	Description string       `json:"description,omitempty" bson:"description,omitempty"`
	Status      LoanStatus   `json:"status,omitempty" bson:"status,omitempty"`
}

// ResponseLoan struct for output
type ResponseLoan struct {
	ID          primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Amount      float64            `json:"amount,omitempty" bson:"amount,omitempty"`
	Category    LoanCategory       `json:"category,omitempty" bson:"category,omitempty"`
	Description string             `json:"description,omitempty" bson:"description,omitempty"`
	Status      LoanStatus         `json:"status,omitempty" bson:"status,omitempty"`
}
