package base

import "context"

type (
	IBaseServiceSqlx[T any, ID any] interface {
		Paging(ctx context.Context, condition any, pagingInput PagingDto) ([]T, int64, error)
		GetByID(ctx context.Context, id ID) (*T, error)
	}

	BaseServiceSqlx[T any, ID any] struct {
		Repo IBaseRepositorySqlx[T, ID]
	}
)

func NewBaseService[T any, ID any](repo IBaseRepositorySqlx[T, ID]) *BaseServiceSqlx[T, ID] {
	service := &BaseServiceSqlx[T, ID]{Repo: repo}

	// check compile-time with type assertion
	var _ IBaseServiceSqlx[T, ID] = service
	return service
}

func (service *BaseServiceSqlx[T, ID]) Paging(ctx context.Context, condition any, pagingInput PagingDto) ([]T, int64, error) {
	entities, count, err := service.Repo.Paging(ctx, condition, pagingInput)
	if err != nil {
		return nil, 0, err
	}

	return entities, count, nil
}

func (service *BaseServiceSqlx[T, ID]) GetByID(ctx context.Context, id ID) (*T, error) {
	return service.Repo.GetById(ctx, id)
}
