package partials

import (
	"time"

	"github.com/jackc/pgx/v5/pgtype"
)

func ConvertDate(dateString string) string {
	date, err := time.Parse("{2006-01-02 15:04:05.999999 -0700 MST finite true}", dateString)
	if err != nil {
		return "-"
	}
	return date.Format("2 Jan, 2006")
}

func ParseId(id pgtype.UUID) string {
	value, err := id.Value()
	if err != nil {
		return "Invalid Id"
	}
	return value.(string)
}
