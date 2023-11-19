type Category struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Description *string   `json:"description,omitempty"`
	Coursers    []*Course `json:"coursers"`
}
