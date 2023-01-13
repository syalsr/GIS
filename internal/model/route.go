package model

type ResponseRoute struct {
	TotalTime float64
	Stops     []string
}

type RequestRoute struct {
	From string
	To   string
}
