package manager

type Pagination struct {
	Page int `json:"page"`
	PageSize int `json:"page_size"`
	PageCount int `json:"page_count"`
	Next int `json:"next"`
	Prev int `json:"prev"`
	RowsCount int64 `json:"rows_count"`
}

func (p *Pagination) GetOffset() int {
	return (p.GetPage() - 1) * p.GetPageSize()
}

func (p *Pagination) GetPageSize() int {
	if p.PageSize == 0 {
		p.PageSize = 10
	}
	return p.PageSize
}

func (p *Pagination) GetPage() int {
	if p.Page == 0 {
		p.Page = 1
	}
	return p.Page
}

func (p *Pagination) GetPrevPage() int {
	if p.Page == 1 {
		p.Prev = 1
	} else if p.Page >= p.PageCount {
		p.Prev = p.PageCount
	} else {
		p.Prev = p.Page - 1
	}
	return p.Prev
}

func (p *Pagination) GetNextPage() int {
	if p.Page >= p.PageCount {
		p.Next = p.PageCount
	} else {
		p.Next = p.Page + 1
	}
	return p.Next
}

