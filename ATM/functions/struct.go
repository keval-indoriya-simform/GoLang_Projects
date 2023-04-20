package functions

// This is a struct for storing user details
type User struct {
	CardNum  int
	UserName string
	PIN      int
	Balance  float64
	logs     []Log
}

// This is a struct for storing log for user's transaction
type Log struct {
	TransactionId     int
	DateTime          string
	TypeOfTransaction string
	Amount            float64
	OpeningBalance    float64
	ClosingBalance    float64
}

// This methods are for sorting the struct slice by transactionId
type ById []Log

func (a ById) Len() int           { return len(a) }
func (a ById) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ById) Less(i, j int) bool { return a[i].TransactionId > a[j].TransactionId }
