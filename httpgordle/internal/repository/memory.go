package repository

import (
	"fmt"
	"learngo-pockets/httpgordle/internal/session"
)

type GameRepository struct {
	storage map[session.GameID]session.Game
}

func New() *GameRepository {
	return &GameRepository{
		storage: make(map[session.GameID]session.Game),
	}
}

func (gr *GameRepository) Add(game session.Game) error {
	_, ok := gr.storage[game.ID]
	if ok {
		return fmt.Errorf("%w (%s)", ErrConflictingID, game.ID)
	}

	gr.storage[game.ID] = game

	return nil
}

func (gr *GameRepository) Find(id session.GameID) (session.Game, error) {
	game, ok := gr.storage[id]
	if !ok {
		return session.Game{}, fmt.Errorf("can't find game %s: %w", id, ErrNotFound)
	}

	return game, nil
}

func (gr *GameRepository) Update(game session.Game) error {
	_, ok := gr.storage[game.ID]
	if !ok {
		return fmt.Errorf("can't find game %s: %w", game.ID, ErrNotFound)
	}

	gr.storage[game.ID] = game

	return nil
}
