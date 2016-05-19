package client

// Credentials is the interface for any type that is able to return an id and a secret
type Credentials interface {
	Id() string
	Secret() string
}

// CredentialsProvider is the interface for any type that is able to return Credentials
type CredentialsProvider interface {
	Get() (Credentials, error)
}
