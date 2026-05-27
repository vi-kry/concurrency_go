package compute

const (
	UnknownCommandID = iota
	SetCommandID
	GetCommandID
	DeleteCommandID

	UnknownCommand = "UNKNOWN"
	SetCommand     = "SET"
	GetCommand     = "GET"
	DeleteCommand  = "DELETE"

	setCommandArgumentsNumber     = 2
	getCommandArgumentsNumber     = 1
	deleteCommandArgumentsNumber  = 1
	unknownCommandArgumentsNumber = 0
)

func mapCommandNameToCommandID(name string) int {
	switch name {
	case SetCommand:
		return SetCommandID
	case GetCommand:
		return GetCommandID
	case DeleteCommand:
		return DeleteCommandID
	default:
		return UnknownCommandID
	}
}

func getArgumentsNumberByCommandID(commandID int) int {
	switch commandID {
	case SetCommandID:
		return setCommandArgumentsNumber
	case GetCommandID:
		return getCommandArgumentsNumber
	case DeleteCommandID:
		return deleteCommandArgumentsNumber
	default:
		return unknownCommandArgumentsNumber
	}
}
