package valueobject

type Email string

func (e Email) Validate() error {
	return nil
}