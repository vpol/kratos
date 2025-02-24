package code

import (
	"context"

	"github.com/gofrs/uuid"
)

type (
	RecoveryCodePersister interface {
		CreateRecoveryCode(ctx context.Context, dto *CreateRecoveryCodeParams) (*RecoveryCode, error)
		UseRecoveryCode(ctx context.Context, fID uuid.UUID, code string) (*RecoveryCode, error)
		DeleteRecoveryCodesOfFlow(ctx context.Context, fID uuid.UUID) error
	}

	RecoveryCodePersistenceProvider interface {
		RecoveryCodePersister() RecoveryCodePersister
	}
)
