package utils

import (
	"encoding/base64"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func Reverse[T any](s []T) []T {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return s
}

func CheckParse(c *gin.Context, model interface{}) interface{} {

	if err := c.ShouldBindJSON(&model); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		c.Abort()
		return nil
	}

	return model
}

type PaginationInfo struct {
	NextCursor string `json:"nextCursor"`
	PrevCursor string `json:"prevCursor"`
}

type Cursor map[string]interface{}

func CreateCursor(id string, createdAt time.Time, pointNext bool) Cursor {
	return Cursor{
		"id":        id,
		"createdAt": createdAt,
		"pointNext": pointNext,
	}
}

func GeneratePager(next Cursor, prev Cursor) PaginationInfo {
	return PaginationInfo{
		NextCursor: encodeCursor(next),
		PrevCursor: encodeCursor(prev),
	}
}

func encodeCursor(cursor Cursor) string {
	if len(cursor) == 0 {
		return ""
	}
	serialize, err := json.Marshal(cursor)
	if err != nil {
		return ""
	}
	return base64.StdEncoding.EncodeToString(serialize)
}

func DecodeCursor(encoded string) (Cursor, error) {
	decode, err := base64.StdEncoding.DecodeString(encoded)
	if err != nil {
		return nil, err
	}
	var cursor Cursor

	err = json.Unmarshal(decode, &cursor)
	if err != nil {
		return nil, err
	}
	return cursor, nil
}

func GetPaginationOperator(pointsNext bool, sortOrder string) (string, string) {
	if pointsNext && sortOrder == "asc" {
		return ">", ""
	}
	if pointsNext && sortOrder == "desc" {
		return "<", ""
	}
	if !pointsNext && sortOrder == "asc" {
		return "<", "desc"
	}
	if !pointsNext && sortOrder == "desc" {
		return ">", "asc"
	}

	return "", ""
}
