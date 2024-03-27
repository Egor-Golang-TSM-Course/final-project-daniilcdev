package shared

type HashStorage interface {
	Contains(hash string) bool
	Get(payload string) (string, error)
	Add(payload, hash string)
}
