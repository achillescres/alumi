package valueobject

type UserType string

const (
	UserTypeMenti  UserType = "menti"
	UserTypeMentor UserType = "mentor"
)

var userTypes = []string{
	string(UserTypeMenti),
	string(UserTypeMentor),
}

func (ut UserType) Validate() bool {
	for _, u := range userTypes {
		if u == string(ut) {
			return true
		}
	}
	return false
}

func (UserType) Values() []string {
	return userTypes
}
