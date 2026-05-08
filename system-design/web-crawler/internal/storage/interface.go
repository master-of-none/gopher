package storage

type VisitedStore interface {
	Seen(url string) bool
	MarkInProgress(url string) bool
	RemoveInProgress(url string) error
}
