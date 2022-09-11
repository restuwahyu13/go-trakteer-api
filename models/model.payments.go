package models

import "time"

type Payments struct {
	Id            uint      `db:"id"`
	TransactionId string    `db:"transaction_id"`
	ExternalId    string    `db:"external_id"`
	Status        string    `db:"status"`
	Amount        uint64    `db:"amount"`
	IsPaid        bool      `db:"is_paid"`
	Categorie     string    `db:"categorie"`
	GoalId        uint      `db:"goal_id"`
	SupporterId   uint      `db:"supporter_id"`
	CreatedAt     time.Time `db:"created_at"`
	Supporter     Supporter
}
