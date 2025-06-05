package api

const (
	// GameID is the name of the field that stores the game's identifier
	GameID = "id"

	// NewGameRoute is the path to create a new game.
	NewGameRoute = "/games"

	// GetStatusRoute is the path to get the status of a game identified by
	// its id.
	GetStatusRoute = "/games/{" + GameID + "}"

	// GuessRoute is the path to play a guess in a game, identified by its id.
	GuessRoute = "/games/{" + GameID + "}"
)

// GuessRequest is the structureof the message used wen submitting a
// guess.
type GuessRequest struct {
	Guess string `json:"guess"`
}

// GameResponse contains the information about a game.
type GameResponse struct {
	ID           string  `json:"id"`
	AttemptsLeft byte    `json:"attempts_left"`
	Guesses      []Guess `json:"guesses"`
	WordLength   byte    `json:"word_length"`
	Solution     string  `json:"solution,omitempty"`
	Status       string  `json:"status"`
}

// Guess is a pair of a word (submitted by the player) and its feedback
// (provided by Gordle).
type Guess struct {
	Word     string `json:"word"`
	Feedback string `json:"feedback"`
}
