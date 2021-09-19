package types

type Sea struct {
	SenderName       string `gorm:"sendername" json:"sender_name"`
	Receiver_Name    string `gorm:"receivername" json:"receiver_name"`
	SenderLocation   string `gorm:"senderlocation" json:"sender_location"`
	ReceiverLocation string `gorm:"receiverlocation" json:"receiver_location"`
	Tracker          string `gorm:"tracker" json:"tracker"`
	Bill             int    `gorm:"bill" json:"bill"`
}
