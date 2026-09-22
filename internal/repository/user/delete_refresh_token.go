package user

import (
	"context"
	"errors"
)

func (r *userRepository) DeleteRefreshTokenByUserID(ctx context.Context, userID int64) error {
	result, err := r.db.ExecContext(ctx, "DELETE FROM refresh_tokens WHERE user_id = ?", userID)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return errors.New("Nothing to delete!")
	}

	return nil
}
