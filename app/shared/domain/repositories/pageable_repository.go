package repositories

import "go-course/demo/app/shared/domain/entities"

type PageableRepository interface {
	FindPageable(limit, page int64, query interface{}) (*entities.Pageable, error)
}
