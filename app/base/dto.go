package base

type (
	PagingDto struct {
		PageSize int `query:"page""`
		PageNum  int `query:"per_page"`
	}

	// this struct for join query
	// could change format key and value
	QueryOption struct {
		Preloads *[]string `json:"preloads"`
	}
)
