package Events

type CommandEvent uint

// string mapping
const (
	ArtisanStarting CommandEvent = iota + 1 // single instance of the service will be created.
	CommandFinished
	CommandStarting
	ScheduledTaskFinished
	ScheduledTaskSkipped
	ScheduledTaskStarting
)

func (s CommandEvent) String() string {
	switch s {
	case ArtisanStarting:
		return "ArtisanStarting"
	case CommandFinished:
		return "CommandFinished"
	case CommandStarting:
		return "CommandStarting"
	case ScheduledTaskFinished:
		return "ScheduledTaskFinished"
	case ScheduledTaskSkipped:
		return "ScheduledTaskSkipped"
	case ScheduledTaskStarting:
		return "ScheduledTaskStarting"
	}
	return "unknown"
}

func FromInt(l uint) CommandEvent {
	switch l {
	case 1:
		return ArtisanStarting
	case 2:
		return CommandFinished
	case 3:
		return CommandStarting
	case 4:
		return ScheduledTaskFinished
	case 5:
		return ScheduledTaskSkipped
	case 6:
		return ScheduledTaskStarting
	}
	return ArtisanStarting
}
