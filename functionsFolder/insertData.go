package functionsfolder

import (
	"fmt"
	"github.com/gocql/gocql"
    "github.com/gofiber/fiber/v2"
    "strconv"
)



func InsertData(session *gocql.Session,c *fiber.Ctx) string{
    p := new(TableData);
    if err := c.BodyParser(p); err != nil{
        return err.Error();
    }
    age,err := strconv.ParseInt(p.Age,10,64);
    if err != nil{
        fmt.Println("Invalid data");
        return err.Error();
    }
    query := session.Query(
        "INSERT INTO my_table (id, name, age) VALUES (?, ?, ?)",
         gocql.TimeUUID(),
         p.Name,       
         age,            
    )
	var response string;
    if err := query.Exec(); err != nil {
		response = err.Error();
        fmt.Println("Error inserting data into Cassandra:", err)
    } else {
		response = "Data inserted successfully."
        fmt.Println("Data inserted successfully.")
    }
   return response;
}