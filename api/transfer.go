package api

type transferRequest struct {
	FromAccountID int64  `json:"from_account_id" binding:"required,min=1"`
	ToAccountID   int64  `json:"to_account_id" binding:"required,min=1`
	Currency      string `json:"currency" binding:"required,oneof=USD EUR"`
}
