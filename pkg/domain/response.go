package domain

import "github.com/google/uuid"

type Metadata struct {
	Days                int       `json:"days" bson:"days,omitempty"`
	AverageClosingPrice float64   `json:"average_closing_price" bson:"average_closing_price,omitempty"`
	Stock               string    `json:"stock" bson:"stock,omitempty"`
	CreatedDate         string    `json:"created_time" bson:"created_time,omitempty"`
	RequestId           uuid.UUID `json:"request_id" bson:"request_id,omitempty"`
}

type Record struct {
	Metadata Metadata            `json:"metadata" bson:"metadata,omitempty"`
	Daily    []map[string]string `json:"daily" bson:"daily,omitempty"`
}
