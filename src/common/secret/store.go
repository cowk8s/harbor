package secret

// Store the secrets and provides methods to validate secrets
type Store struct {
	// the key is secret
	// the value is username
	secrets map[string]string
}

// NewStore ...
func NewStore(secrets map[string]string) *Store {
	return &Store{
		secrets: secrets,
	}
}

// IsValid returns whether the secret is valid
func (s *Store) IsValid(secret string) bool {
	return len(s.GetUsername(secret)) != 0
}

// GetUsername return the corresponding the name of the secret
func (s *Store) GetUsername(secret string) string {
	return s.secrets[secret]
}
