package mooc

import (
	"context"
	"errors"
	"fmt"

	"github.com/google/uuid"
)

var ErrInvalidMessageID = errors.New("invalid message ID")
var ErrEmptyMessageContent = errors.New("the field Message Content can not be empty")
var ErrEmptyMessageUserSender = errors.New("the field Message User Sender can not be empty")
var ErrEmptyMessageUserReceiver = errors.New("the field Message User Receiver can not be empty")
var ErrEmptyMessageTimeSent = errors.New("the field Message Time Sent can not be empty")

type Message struct {
	id            MessageID
	content       MessageContent
	user_sender   MessageUserSender
	user_receiver MessageUserReceiver
	time_sent     MessageTimeSent
}

// MessageID represents the message unique identifier
type MessageID struct {
	value string
}

// NewMessageID instatiate the VO for MessageID
func NewMessageID(value string) (MessageID, error) {
	v, err := uuid.Parse(value)
	if err != nil {
		return MessageID{}, fmt.Errorf("%w: %s", ErrInvalidMessageID, value)
	}

	return MessageID{
		value: v.String(),
	}, nil
}

// String type converts the MessageID into string
func (id MessageID) String() string {
	return id.value
}

// MessageContent represents the content of the message
type MessageContent struct {
	value string
}

// NewMessageContent instatiate the VO for the MessageContent
func NewMessageContent(value string) (MessageContent, error) {
	if value == "" {
		return MessageContent{}, ErrEmptyMessageContent
	}

	return MessageContent{
		value: value,
	}, nil
}

// String type converts the MessageContent into string
func (content MessageContent) String() string {
	return content.value
}

// MessageUserSender represents the user who has sent the message
type MessageUserSender struct {
	value string
}

// NewMessageUserSender instatiate the Vo for the UserSender
func NewMessageUserSender(value string) (MessageUserSender, error) {
	if value == "" {
		return MessageUserSender{}, ErrEmptyMessageUserSender
	}

	return MessageUserSender{
		value: value,
	}, nil
}

// String type converts the MessageUserSender into string
func (userSender MessageUserSender) String() string {
	return userSender.value
}

// MessageUserReceiver represents the user who we sent the message
type MessageUserReceiver struct {
	value string
}

// NewMessageUserSender instatiate the Vo for the UserReceiver
func NewMessageUserReceiver(value string) (MessageUserReceiver, error) {
	if value == "" {
		return MessageUserReceiver{}, ErrEmptyMessageUserReceiver
	}

	return MessageUserReceiver{
		value: value,
	}, nil
}

// String type converts the MessageUserReceiver into string
func (userReceiver MessageUserReceiver) String() string {
	return userReceiver.value
}

// MessageTimeSent represents the time the message was sent
type MessageTimeSent struct {
	value string
}

// NewMessageTimeSent instatiate the Vo for the TimeSent
func NewMessageTimeSent(value string) (MessageTimeSent, error) {
	if value == "" {
		return MessageTimeSent{}, ErrEmptyMessageTimeSent
	}

	return MessageTimeSent{
		value: value,
	}, nil
}

// String type converts the MessageTimeSent into string
func (timeSent MessageTimeSent) String() string {
	return timeSent.value
}

// MessageRepository defines the expected behaviour from a message storage.
type Messagerepository interface {
	Save(ctx context.Context, message Message) error
}

// NewMessage create a new message
func NewMessage(id, content, user_sender, user_receiver, time_sent string) (Message, error) {
	idVO, err := NewMessageID(id)
	if err != nil {
		return Message{}, err
	}

	contentVO, err := NewMessageContent(content)
	if err != nil {
		return Message{}, err
	}

	user_senderVO, err := NewMessageUserSender(user_sender)
	if err != nil {
		return Message{}, err
	}

	user_receiverVO, err := NewMessageUserReceiver(user_receiver)
	if err != nil {
		return Message{}, err
	}

	time_sentVO, err := NewMessageTimeSent(time_sent)
	if err != nil {
		return Message{}, err
	}

	return Message{
		id:            idVO,
		content:       contentVO,
		user_sender:   user_senderVO,
		user_receiver: user_receiverVO,
		time_sent:     time_sentVO,
	}, nil
}

// ID returns the message unique identifier
func (m Message) ID() MessageID {
	return m.id
}

// Content returns the message content
func (m Message) Content() MessageContent {
	return m.content
}

// User_sender returns the message user sender
func (m Message) User_sender() MessageUserSender {
	return m.user_sender
}

// User_receiver returns the message user receiver
func (m Message) User_receiver() MessageUserReceiver {
	return m.user_receiver
}

// Time_sent returns the message time sent
func (m Message) Time_sent() MessageTimeSent {
	return m.time_sent
}
