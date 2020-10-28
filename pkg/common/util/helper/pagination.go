package helper

const (
	paginationDefaultLimit  = 10
	paginationMaxLimit      = 20
	paginationDefaultPageNo = 1
)

// PaginationReq holds pagination http fields and tags
type PaginationReq struct {
	PageNo   int
	PageSize int
}

// Pagination data
type Pagination struct {
	Limit  int
	Offset int
}

// Transform checks and converts http pagination into database pagination model
func (p *PaginationReq) Transform() Pagination {
	if p.PageSize < 1 {
		p.PageSize = paginationDefaultLimit
	}

	if p.PageSize > paginationMaxLimit {
		p.PageSize = paginationMaxLimit
	}

	if p.PageNo < 1 {
		p.PageNo = paginationDefaultPageNo
	}

	return Pagination{Limit: p.PageSize, Offset: (p.PageNo - 1) * p.PageSize}
}
