package data

type Account struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

var AccountData = []*Account{
	{
		Id:       1,
		Name:     "account1",
		Email:    "acc1@gmail.com",
		Password: "123456",
	},
	{
		Id:       2,
		Name:     "account2",
		Email:    "acc2@gmail.com",
		Password: "123456",
	},
	{
		Id:       3,
		Name:     "account3",
		Email:    "acc3@gmail.com",
		Password: "123456",
	},
}
