package controllers

import "go-api/app"
import "github.com/revel/revel"
import "go-api/app/models"

type App struct {
	*revel.Controller
}

func (c App) Index() revel.Result {

	rows, err := app.DB.Query("SELECT username FROM users")
	if err != nil {
		revel.INFO.Println("[pg_error]", err)
	}
	defer rows.Close()

	list := []models.User{}
	for rows.Next() {
		var u models.User
		if err := rows.Scan(&u.Username); err != nil {
			revel.INFO.Println(err)
		}
		list = append(list, u)
	}

	return c.RenderJson(list)
}
