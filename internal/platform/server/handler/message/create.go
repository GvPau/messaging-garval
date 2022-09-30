package message

// Defines the structure of a message
type createRequest struct {
	ID            string `json:"id" binding:"required"`
	Content       string `json:"content" binding:"required"`
	User_sender   uint   `json:"user_sender" binding:"required"`
	User_receiver uint   `json:"user_receiver" binding:"required"`
	Time_sent     string `json:"time_sent" binding:"required"`
}

// CreateHandler returns an HTTP handler for messages creation
//func CreateHandler(messageRepository mooc.messageRepository)
