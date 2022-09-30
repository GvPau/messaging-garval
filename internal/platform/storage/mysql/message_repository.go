package mysql

import "database/sql"

// MessageRepository is a MySQL mooc.MessageRepository implementation
type MessageRepository struct {
	db *sql.DB
}

// func NewMessageRepository initializes a MySQL-based implementation of mooc.MessageRepository.
