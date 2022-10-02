package models

import "time"

type Payments struct {
	Id            int        `json:"id" db:"id"`
	TransactionId string     `json:"transaction_id" db:"transaction_id"`
	ExternalId    string     `json:"external_id" db:"external_id"`
	Status        string     `json:"status" db:"status"`
	Amount        uint64     `json:"amount" db:"amount"`
	IsPaid        bool       `json:"is_paid" db:"is_paid"`
	Categorie     string     `json:"categorie" db:"categorie"`
	GoalId        uint       `json:"goal_id" db:"goal_id"`
	SupporterId   uint       `json:"supporter_id" db:"supporter_id"`
	CreatedAt     time.Time  `json:"created_at" db:"created_at"`
	Supporter     *Supporter `json:"supporter,omitempty"`
}
