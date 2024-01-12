package types

type Vote struct {
	Voter    User
	Category Category
	Target   User
}
