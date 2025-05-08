package bindings

// QCorechainQuery contains QCorechain custom queries.
type QCoreQuery struct {
	Minters *Minters `json:"minters,omitempty"`
}

type Minters struct {
}

type MintersResponse struct {
	AnnualProvisions string `json:"annual_provisions"`
	CurrentEpoch     string `json:"current_epoch"`
	Identifier       string `json:"identifier"`
}
