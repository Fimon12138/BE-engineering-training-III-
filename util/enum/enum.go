package enum

const (
	V1 = "/v1"

	EMAILADDRESSREGEXP = "\\w+([-+.]\\w+)*@\\w+([-.]\\w+)*\\.\\w+([-.]\\w+)*"
	LENOFTELEPHONE     = 11

	TABLENAME_ACCOUNT        = "accounts"
	TABLENAME_USER           = "users"
	TABLENAME_MERCHANT       = "merchants"
	TABLENAME_ZJPAY          = "zjpays"
	TABLENAME_ORDER          = "orders"
	TABLENAME_TICKET         = "tickets"
	TABLENAME_COMMENT        = "comments"
	TABLENAME_RECOMMENDATION = "recommendations"
	TABLENAME_FAVORITE       = "favorites"

	YES = "yes"
	NO  = "no"

	ACCOUNT_TYPE_USER     = "user"
	ACCOUNT_TYPE_MERCHANT = "merchant"

	ORDER_STATUS_PROCESSING = "processing"
	ORDER_STATUS_FINISHED   = "finished"

	ORDERBY_COMMON_CREATETIME = "create_time"
	ORDERBY_COMMON_UPDATETIME = "update_time"

	ORDERBY_TICKET_SUBSCRIBECOUNT = "subscribe_count"

	ORDER_DESC              = "desc"
	ORDER_ASC               = "asc"
	TICKET_TYPE_CONCERT     = "concert"
	TICKET_TYPE_SPORT       = "sport"
	TICKET_TYPE_EXHIBITION_ = "exhibition"
	TICKET_TYPE_OPERA       = "opera"

)

func GetAccountType() []string {
	return []string{ACCOUNT_TYPE_MERCHANT, ACCOUNT_TYPE_USER}
}

func GetBoolType() []string {
	return []string{YES, NO}
}

func GetOrderValues() []string {
	return []string{ORDER_ASC, ORDER_DESC}
}

func GetOrderStatusValues() []string {
	return []string{ORDER_STATUS_PROCESSING, ORDER_STATUS_FINISHED}
}

func GetTicketOrderByValues() []string {
	return []string{ORDERBY_COMMON_UPDATETIME, ORDERBY_TICKET_SUBSCRIBECOUNT}
}

func GetTicketTypeValues() []string {
	return []string{TICKET_TYPE_CONCERT, TICKET_TYPE_SPORT, TICKET_TYPE_EXHIBITION_, TICKET_TYPE_OPERA}
}

