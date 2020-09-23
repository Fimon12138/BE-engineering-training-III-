package model

const (
	TICKET_FIELD_NAME      = "name"
	MYSQL_CONTANINS_STRING = " LIKE ?"
)

type Pagination struct {
	Offset int
	Size   int
}

type OrderFilter struct {
	Field     string
	Direction string
}
