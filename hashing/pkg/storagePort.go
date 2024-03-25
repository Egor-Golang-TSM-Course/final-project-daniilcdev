package client

type HashStorage interface {
	Contains(hash string) bool
	Get(payload string) (string, error)
	Add(payload, hash string)
}
