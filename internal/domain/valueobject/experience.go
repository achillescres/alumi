package valueobject

import "time"

type Experience struct {
	Description string
	Place       string
	Start, End  time.Time
	Who         string
}
