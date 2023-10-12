package functionsfolder

import (
	"github.com/gocql/gocql"
	"github.com/gofiber/fiber/v2"
)



func GetSrData(c *fiber.Ctx, session *gocql.Session) error {
	query := "Select no,description,Type,status,reporter from serviceRequest"
	var toDoData []serviceRequest
	var inProgressData []serviceRequest
	var doneData []serviceRequest
	var rejecedData []serviceRequest
	var acceptedData []serviceRequest
	scanner := session.Query(query).Iter().Scanner()
	for scanner.Next() {
		var no int64
		var description string
		var Type string
		var status string
		var reporter string
		err := scanner.Scan(&no, &description, &Type, &status, &reporter)
		if err != nil {
			return c.SendStatus(fiber.StatusInternalServerError)
		}
		data1 := serviceRequest{No: no, Description: description, Type: Type, Status: status, Reporter: reporter}
		if data1.Status == "ToDo" {
			toDoData = append(toDoData, data1)
		} else if data1.Status == "InProgress" {
			inProgressData = append(inProgressData, data1)
		} else if data1.Status == "Done" {
			doneData = append(doneData, data1)
		} else if data1.Status == "Accepted" {
			acceptedData = append(acceptedData, data1)
		} else {
			rejecedData = append(rejecedData, data1)
		}
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"toDo":       toDoData,
		"inProgress": inProgressData,
		"done":       doneData,
		"rejected":   rejecedData,
		"accepted":   acceptedData,
	})
}

func CreateNewSr(c *fiber.Ctx, session *gocql.Session) error {
	p := new(serviceRequest)
	if err := c.BodyParser(p); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"Error": "Details are missing"})
	}

	query := session.Query(
		"INSERT INTO serviceRequest (no,assignee,description,reporter,status,type) VALUES (?, ?, ?, ?, ?, ?)",
		p.No,
		p.Assignee,
		p.Description,
		p.Reporter,
		p.Status,
		p.Type,
	)
	if err := query.Exec(); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"Err": "Could not create SR"})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"Message": "SR Successfully created"})

}
