// Error codes
package errors

const (
	INIT_SUCCESS string = "INIT_SUCCESS"
	INIT_FAILED  string = "INIT_FAILED"
)

const (
	CREATE_ORDER_SUCCESS int = 1
	CREATE_ORDER_ERROR   int = 0

	CANCEL_ORDER_SUCCESS int = 1
	CANCEL_ORDER_ERROR   int = 0
)

// telegram
const (
	TELEGRAM_BOT_NOT_SUCCESSFULLY_SENT string = "TELEGRAM_BOT_NOT_SUCCESSFULLY_SENT"
	TELEGRAM_BOT_NOT_INITIALIZED       string = "TELEGRAM_BOT_NOT_INITIALIZED"
	TELEGRAM_BOT_SEND_MESSAGE_ERROR    string = "TELEGRAM_BOT_SEND_MESSAGE_ERROR"
)
