package functionsfolder

import (
	"fmt"
	"github.com/gocql/gocql"
)

type sampleTableData struct {
	Name string `json:"name"`
	Age  int64  `json:"age"`
}

func ReadData(session *gocql.Session) []sampleTableData{
	query := "Select id,name,age from my_table"
	fmt.Println("Query is ", query)
	var data []sampleTableData
	scanner := session.Query(query).Iter().Scanner()
	for scanner.Next() {
		var id gocql.UUID
		var name string
		var age int64
		err := scanner.Scan(&id, &name, &age)
		if err != nil {
			fmt.Println(err)
			fmt.Println("Error in sccaning")
			break
		}
		data1 := sampleTableData{Name: name, Age: age}
		data = append(data, data1)
	}
	return data
}
