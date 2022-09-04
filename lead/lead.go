package lead

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/gofiber/fiber"
	"github.com/devsylva/go-fiber-crm-basic/database"
)


type Lead struct{
	gorm.Model
	Name         string      `json:"name"`
	Company      string      `json:"company"`
	Email        string 	 `json:"email"`
	Phone        int         `json:"phone"`
}


func GetLeads(c *fiber.ctx){
	db := database.DBConn
	var leads []Lead
	db.Find(&leads)
	c.JSON(leads)
}

func GetLead(c *fiber.ctx){
	id := c.Params("id")
	db := database.DBConn
	var lead Lead 
	db.Find(&lead, id)
	c.JSON(lead)
}

func NewLead(c *fiber.ctx){
	db := database.DBConn
	lead := new(Lead)
	if err := c.BodyParser(lead); err != nil {
		return
	}
	db.Create(&lead)
	c.JSON(lead)
}

func DeleteLead(c *fiber.ctx){
	id := c.Params("id")
	db := database.DBConn

	var lead Lead
	db.First(&lead, id)
	if lead.Name == ""{
		c.Status(500).Send("No lead found wit id")
	}
	db.Delete(&lead)
	c.Send("Lead successfully deleted")
}