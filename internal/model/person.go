package model

type Person struct {
	ID          int    `db:"id"`
	Name        string `db:"name"`
	Surname     string `db:"surname"`
	Patronymic  string `db:"patronymic"`
	Age         int    `db:"age"`
	Gender      string `db:"gender"`
	Nationality string `db:"nationality"`
}
