package models

import (
	"encoding/json"
	"time"
)

type PaymentLogs struct {
	Id            int             `json:"id" db:"id"`
	TransactionId string          `json:"transaction_id" db:"transaction_id"`
	ExternalId    string          `json:"external_id" db:"external_id"`
	Categorie     string          `json:"categorie" db:"categorie"`
	ResPayload    json.RawMessage `json:"res_payload" db:"res_payload"`
	ResourceId    uint            `json:"resource_id" db:"resource_id"`
	ResourceType  string          `json:"resource_type" db:"resource_type"`
	CreatedAt     time.Time       `json:"created_at" db:"created_at"`
}
