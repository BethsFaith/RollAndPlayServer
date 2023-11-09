package model

import "database/sql"

func getDefaultOrValue(defaultValue int, value sql.NullInt64) int {
	if value.Valid {
		return int(value.Int64)
	} else {
		return defaultValue
	}
}
