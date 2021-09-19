package services

import (
	"context"

	"gorm.io/gorm"
)

type Services interface {
	Create(ctx context.Context, db *gorm.DB) error
	// GetSenderName() (string, error)
	// GetReceiverName() (string, error)
	// GetSenderLocation() (string, error)
	// GetReceiverLocation() (string, error)
	// GetBill() (int, error)
}
