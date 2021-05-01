package repositories

import "go-course/demo/app/shared/domain/shared"

type PageableRepository interface {
	FindPageable(limit, page int64, query interface{}) (*shared.Pageable, error)
}
