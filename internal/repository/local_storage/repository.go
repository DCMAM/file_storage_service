package local_storage

// Repository represents all the needed depedencies for users
type Repository struct{}

// NewRepository will initiate UserRepository's provider
func NewRepository() Repository {
	return Repository{}
}
