package ports

type UserServices interface {
	// EcnrypPassword generates a hashed password from the input password string using bcrypt.
	// 
	// Parameters:
	//   - password: the password string to be hashed.
	//   - string: the hashed password.
	//   - error: an error if any occurs during the hashing process.
	EcnrypPassword(password string) (string, error)
	// ComparePassword compares a hashed password with a plaintext password.
	// 
	// Parameters:
	//   - hashedPassword: the hashed password to compare.
	//   - password: the plaintext password to compare.
	// 
	// Return type: error.
	ComparePassword(hashedPassword, password string) error
}