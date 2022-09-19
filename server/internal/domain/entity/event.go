package entity

type Action string

const (
	CreateEvent = Action("create")
	DeleteEvent = Action("delete")
	UpdateEvent = Action("update")
)

type Message struct {
	ItemID    string `json:"item_id"`
	Action    string `json:"action"`
	Timestamp int64  `json:"timestamp"`
}
