package fsm

type State int

const (
	CREATED  State = 0
	APPROVED State = 1
	ERROR    State = 2
)

func (state State) Name() string {
	return []string{"CREATED", "APPROVED", "ERROR"}[state]
}
