package valueobject

type MatchStatus string

const (
	MatchStatusPending  MatchStatus = "pending"
	MatchStatusAccepted MatchStatus = "accepted"
	MatchStatusDeclined MatchStatus = "declined"
)

var matchStatuses = []string{
	string(MatchStatusPending),
	string(MatchStatusAccepted),
	string(MatchStatusDeclined),
}

func (ms MatchStatus) Validate() bool {
	for _, m := range matchStatuses {
		if m == string(ms) {
			return true
		}
	}
	return false
}

func (ms MatchStatus) Values() []string {
	return matchStatuses
}
