package user

// Dummy validateUser function for demonstration
func ValidateUnencodedUser(username, password string) (bool, error) {
	// In a real world application, you should check the username and password against your database or another form of storage.
	// Here we are simply checking against hard coded values for demonstration purpose.
	if username == "testuser" && password == "testpass" {
		// If the username and the password is correct return true, nil
		return true, nil
	}

	return false, nil
}
