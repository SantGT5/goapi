package api

type CoinBalanceParams struct {
	Username string
}

type CoinBalanceResponse struct {
	// Success COde, Usually 200
	Code int

	// Account Balance
	Balance int64
}

type Error struct {
	// Error code
	Code int

	// Error Message
	Message string
}
