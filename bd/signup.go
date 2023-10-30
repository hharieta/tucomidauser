package bd

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/hharieta/tucomidauser/models"
	"github.com/hharieta/tucomidauser/tools"
)

func SignUp(sig models.SignUp) error {
	fmt.Println("Start Register")

	err := DbConnect()

	if err != nil {
		return err
	}

	defer Db.Close()

	sqlStatement := "INSERT INTO users (User_Email, User_UUID, User_DateAdd) VALUES (?, ?, ?)"

	fmt.Println(sqlStatement)

	_, err = Db.Exec(sqlStatement, sig.UserEmail, sig.UserUUID, tools.DateMySQL())

	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	fmt.Println("Signup > Successful execution")
	return nil
}
