package functionsfolder

import (
	"fmt"
	"github.com/gocql/gocql"
)


func ReadData(session *gocql.Session) []TableData {
	query := "Select id,name,age from my_table"
	fmt.Println("Query is ", query)
	var data []TableData
	scanner := session.Query(query).Iter().Scanner()
	for scanner.Next() {
		var id gocql.UUID
		var name string
		var age string
		err := scanner.Scan(&id, &name, &age)
		if err != nil {
			fmt.Println(err)
			fmt.Println("Error in scanning")
			break
		}
		data1 := TableData{Name: name, Age: age}
		data = append(data, data1)
	}
	return data
}
