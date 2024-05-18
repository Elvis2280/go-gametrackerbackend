package types

import "gametracker/utils"

type ResponseDto struct {
	Status     int                  `json:"status"`
	Data       interface{}          `json:"data"`
	Pagination utils.PaginationInfo `json:"pagination"`
}
