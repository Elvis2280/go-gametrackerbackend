package utils

import "gorm.io/gorm"

type Pagination struct {
	TotalPages   int
	ItemsPerPage int
	CurrentPage  int
	SortKey      string
	IsAscending  bool
	Status       string
	//Tags         []string
}

type PaginatedData struct {
	TotalPages  int
	NextPage    int
	PrevPage    int
	CurrentPage int
}

func HandlePagination(db *gorm.DB, pag *Pagination, totalItems int64) (*gorm.DB, *PaginatedData) {
	offset := pag.ItemsPerPage * (pag.CurrentPage - 1)
	db = db.Offset(offset).Limit(pag.ItemsPerPage)

	totalPages := int(totalItems) / pag.ItemsPerPage

	var nextPage int
	if pag.CurrentPage < totalPages {
		nextPage = pag.CurrentPage + 1
	} else {
		nextPage = totalPages
	}
	var prevPage int
	if pag.CurrentPage > 1 {
		prevPage = pag.CurrentPage - 1
	} else {
		prevPage = 1
	}

	//if pag.SortKey != "" {
	//	searchKey := pag.SortKey
	//	db = db.Where("game.name LIKE ?", "%"+searchKey+"%")
	//
	//}

	//if pag.Status != "" {
	//	db = db.Where("status = ?", pag.Status)
	//}

	//if len(pag.Tags) > 0 {
	//	db = db.Table("game").Joins("JOIN game_tags ON game_tags.game_id = game.id").
	//		Joins("JOIN tags ON tags.id = game_tags.tag_id").
	//		Where("tags.name IN (?)", pag.Tags)
	//}

	//if pag.IsAscending {
	//	db = db.Order("game.id ASC")
	//} else {
	//	db = db.Order("game.id DESC")
	//}

	return db, &PaginatedData{
		TotalPages:  totalPages,
		NextPage:    nextPage,
		PrevPage:    prevPage,
		CurrentPage: pag.CurrentPage,
	}
}
