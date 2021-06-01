package sqlite

import (
	"context"
	"database/sql"

	"github.com/pkg/errors"

	domainmodels "code-cadets-2021/homework_3/calculator/internal/domain/models"
	storagemodels "code-cadets-2021/homework_3/calculator/internal/infrastructure/sqlite/models"
)

// BetRepository provides methods that operate on bets SQLite database.
type BetRepository struct {
	dbExecutor DatabaseExecutor
	betMapper  BetMapper
}

// NewBetRepository creates and returns a new BetRepository.
func NewBetRepository(dbExecutor DatabaseExecutor, betMapper BetMapper) *BetRepository {
	return &BetRepository{
		dbExecutor: dbExecutor,
		betMapper:  betMapper,
	}
}

// InsertBet inserts the provided bet into the database. An error is returned if the operation
// has failed.
func (r *BetRepository) InsertBet(ctx context.Context, bet domainmodels.Bet) error {
	storageBet := r.betMapper.MapDomainBetToStorageBet(bet)
	err := r.queryInsertBet(ctx, storageBet)
	if err != nil {
		return errors.Wrap(err, "bet repository failed to insert a bet with id "+bet.Id)
	}
	return nil
}

func (r *BetRepository) queryInsertBet(ctx context.Context, bet storagemodels.Bet) error {
	insertBetSQL := "INSERT INTO bets(id, selection_id, selection_coefficient, payment) VALUES (?, ?, ?, ?)"
	statement, err := r.dbExecutor.PrepareContext(ctx, insertBetSQL)
	if err != nil {
		return err
	}

	_, err = statement.ExecContext(ctx, bet.Id, bet.SelectionId, bet.SelectionCoefficient, bet.Payment)
	return err
}

// GetBetsBySelectionID fetches bets from the database and returns them.
func (r *BetRepository) GetBetsBySelectionID(ctx context.Context, id string) ([]domainmodels.Bet, bool, error) {
	storageBets, err := r.queryGetBetsBySelectionID(ctx, id)
	if err == sql.ErrNoRows {
		return []domainmodels.Bet{}, false, nil
	}
	if err != nil {
		return []domainmodels.Bet{}, false, errors.Wrap(err, "bet repository failed to get bets with id "+id)
	}

	var bets []domainmodels.Bet

	for _, storageBet := range storageBets {
		bets = append(bets, r.betMapper.MapStorageBetToDomainBet(storageBet))
	}

	return bets, true, nil
}

func (r *BetRepository) queryGetBetsBySelectionID(ctx context.Context, id string) ([]storagemodels.Bet, error) {
	row, err := r.dbExecutor.QueryContext(ctx, "SELECT * FROM bets WHERE selection_id='"+id+"';")
	if err != nil {
		return []storagemodels.Bet{}, err
	}
	defer row.Close()

	var bets []storagemodels.Bet

	for row.Next() {
		var id string
		var selectionId string
		var selectionCoefficient int
		var payment int

		err = row.Scan(&id, &selectionId, &selectionCoefficient, &payment)
		if err != nil {
			return []storagemodels.Bet{}, err
		}

		bets = append(bets, storagemodels.Bet{
			Id:                   id,
			SelectionId:          selectionId,
			SelectionCoefficient: selectionCoefficient,
			Payment:              payment,
		})
	}

	return bets, nil
}
