package core

const (
	LoginUrl      = "https://gf2-bbs-api.exiliumgf.com/login/account"
	MemberInfoUrl = "https://gf2-bbs-api.exiliumgf.com/community/member/info"
	ListUrl       = "https://gf2-bbs-api.exiliumgf.com/community/topic/list?sort_type=1&category_id=1&query_type=1&last_tid=0&pub_time=0&reply_time=0&hot_value=0"
	Info          = "https://gf2-bbs.exiliumgf.com/threadInfo?id=%d&hash_flag=1"
	Like          = "https://gf2-bbs-api.exiliumgf.com/community/topic/like/%d?id=%d"
	Share         = "https://gf2-bbs-api.exiliumgf.com/community/topic/share/%d?id=%d"
	ExchangeList  = "https://gf2-bbs-api.exiliumgf.com/community/item/exchange_list"
	Exchange      = "https://gf2-bbs-api.exiliumgf.com/community/item/exchange"
	SignIn        = "https://gf2-bbs-api.exiliumgf.com/community/task/sign_in"
	SignInStatus  = "https://gf2-bbs-api.exiliumgf.com/community/task/get_current_sign_in_status"
)

var Authorization string
var ExchangeAllowed = []int{2, 3, 4, 5}
var Score = 0
