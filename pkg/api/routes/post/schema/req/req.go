package req

// Create represents incoming req body via HTTP
type Create struct {
	Title string  `json:"title" binding:"required"`
	Body  string  `json:"body" binding:"required"`
	Tags  []int64 `json:"tags"`
}
