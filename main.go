package main

import (
	"fmt"
	"net/http"
	"hrms.co/example/routes"
	"hrms.co/example/mongo_ops"
	// "time"
)

func main()  {
	res := mongo_ops.ConnectMongo(mongo_ops.MongoURI)
	if res != nil {
		fmt.Println("Connect to mongo Failed",res)
	}
	fmt.Println("Started thed connection with mongo", res)
	
	http.HandleFunc("/home", routes.Home)
	http.HandleFunc("/admin_registration", routes.RegisterAdmin)
	http.HandleFunc("/employee_registration", routes.RegisterEmp)
	http.HandleFunc("/get_employees", routes.GetEmployeeList)

	http.ListenAndServe(":8081", nil)
}