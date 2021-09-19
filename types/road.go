package types

type Road struct {
	ID               int64  `gorm:"AUTO_INCREMENT" json:"id,omitmenty"`
	SenderName       string `gorm:"sendername" json:"sender_name,omitmenty"`
	ReceiverName     string `gorm:"receivername" json:"receiver_name,omitempty"`
	SenderLocation   string `gorm:"senderlocation" json:"sender_location,omitempty"`
	ReceiverLocation string `gorm:"receiverlocation" json:"receiver_location,omitempty"`
	Tracker          string `gorm:"tracker" json:"tracker,omitempty"`
	Bill             int    `gorm:"bill" json:"bill,omitempty"`
}
