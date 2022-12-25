package lead

import (
	database "go-fiber-crm-basic/databes"

	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"
)

type Lead struct {
	gorm.Model
	Name string `json:"name"`
	Comapany string `json:"company"`
	Email string `json:"email"`
	Phone string `json:"phone"`
}

func GetLeads(c *fiber.Ctx){
	db := database.DBConn
	var leads []Lead
	db.Find(&leads)
	c.JSON(leads)
}
func GetLead(c *fiber.Ctx){
	id := c.Params("id")
	db := database.DBConn
	var lead Lead
	db.Find(&lead,id)
	c.JSON(lead)
}
func NewLead(c *fiber.Ctx){
	db := database.DBConn
	lead := new(Lead)
	if err := c.BodyParser(lead); err!= nil{
		c.Status(503).Send(err)
		return
	}
	db.Create(&lead)
	c.JSON(lead)
}
func DeleteLead(c *fiber.Ctx){
	id := c.Params("id")
	db := database.DBConn
	
	var lead Lead
	db.First(&lead,id)
	if lead.Name == ""{
		c.Status(500).Send("no lead found with id")
	}
	db.Delete(&lead)
	c.Send("lead successfully deleted")
}