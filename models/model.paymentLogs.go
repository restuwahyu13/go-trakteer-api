package models

import (
	"encoding/json"
	"time"
)

type PaymentLogs struct {
	Id            uint            `db:"id"`
	TransactionId string          `db:"transaction_id"`
	ExternalId    string          `db:"external_id"`
	Categorie     string          `db:"categorie"`
	ResPayload    json.RawMessage `db:"res_payload"`
	ResourceId    uint            `db:"resource_id"`
	ResourceType  string          `db:"resource_type"`
	CreatedAt     time.Time       `db:"created_at"`
}
