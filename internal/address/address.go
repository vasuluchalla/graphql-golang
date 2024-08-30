package address

import (
	"log"

	database "github.com/vasuluchalla/graphql-golang/internal/database"
)

type Address struct {
	ID    string
	Line1 string
	Line2 string
	City  string
	State string
	Zip   int
}

func (address *Address) Create() (int64, error) {
	statement, err := database.Db.Prepare("INSERT INTO Address(Line1, Line2, City, State, Zip)) VALUES(?,?,?,?,?)")
	print(statement)
	if err != nil {
		log.Fatal(err)
	}
	defer statement.Close()

	res, err := statement.Exec(address.Line1, address.Line2, address.City, address.State, address.Zip)
	if err != nil {
		log.Printf("Error fetching last insert ID: %v", err)
		return 0, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		log.Printf("Error fetching last insert ID: %v", err)
		return 0, err
	}
	return id, nil
}
