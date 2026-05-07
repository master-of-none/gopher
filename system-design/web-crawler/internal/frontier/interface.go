package frontier

type Frontier interface {
	Push(url string)
	Pop() string
}
