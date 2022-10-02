package models

import (
	"encoding/json"
	"time"
)

type PaymentLogs struct {
	Id            int             `json:"id,omitempty" db:"id"`
	TransactionId string          `json:"transaction_id,omitempty" db:"transaction_id"`
	ExternalId    string          `json:"external_id,omitempty" db:"external_id"`
	Categorie     string          `json:"categorie,omitempty" db:"categorie"`
	ResPayload    json.RawMessage `json:"res_payload,omitempty" db:"res_payload"`
	ResourceId    uint            `json:"resource_id,omitempty" db:"resource_id"`
	ResourceType  string          `json:"resource_type,omitempty" db:"resource_type"`
	CreatedAt     time.Time       `json:"created_at,omitempty" db:"created_at"`
}
