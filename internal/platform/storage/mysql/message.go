package mysql

const (
	sqlMessageTable = "messages"
)

type sqlCourse struct {
	ID           string `db:"id"`
	Message      string `db:"message"`
	UserSender   string `db:"user_sender"`
	UserReceiver string `db:"user_receiver"`
	TimeSent     string `db:"time_sent"`
}
