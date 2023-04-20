package functions

// This function will return the index where your card number is matched
// if not found it will return -1
func UserIndex(users []User, cardNo int) int {
	for i, user := range users {
		if user.CardNum == cardNo {
			return i
		}
	}
	return -1
}
