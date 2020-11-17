package req

// Create represents incoming req body via HTTP
type Create struct {
	Name        string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
}
