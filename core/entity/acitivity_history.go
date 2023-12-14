package entity

type ActivityHistoryType string

const (
	ActivityHistorySendMoney    ActivityHistoryType = "send_money"
	ActivityHistoryReceiveMoney ActivityHistoryType = "receive_money"
	ActivityHistoryTypeLogin    ActivityHistoryType = "login"
	ActivityHistoryTypeLogout   ActivityHistoryType = "logout"
)

type ActivityHistory struct {
	UserID       string
	ActivityType ActivityHistoryType
	Description  string
}
