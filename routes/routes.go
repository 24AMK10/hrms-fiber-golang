package routes

import (
	"fmt"
	"net/http"
	"encoding/json"
	"hrms.co/example/json_structs"
	// "io"
	// "strings"
	// "log"
	"hrms.co/example/mongo_ops"
	"github.com/golang-jwt/jwt/v5"
)



func Home(w http.ResponseWriter, req *http.Request){

	if req.Method == http.MethodGet {
		jsonData := json_structs.Res{
			Status:true, 
			Message:"nil"}
		js, _ := json.Marshal(jsonData)
		w.Header().Set("content-type", "application/json")
		w.Write(js)
		return
	}
	http.Error(w, "Invalid Method.", http.StatusNotFound)
	return
	
}

func RegisterAdmin(w http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodPost {
		req.Header.Set("Content-Type", "application/json")

		var admin *json_structs.Admin
		err := json.NewDecoder(req.Body).Decode(&admin)
		if err != nil {
			http.Error(w, "Invalid Request", http.StatusBadRequest)
			panic(err)
		}
		fmt.Println("admin >>>", admin)
		response := mongo_ops.RegisterAdmin(admin)
		
		jsonData := json_structs.Res{
			Status:response.Status, 
			Message:response.Message,
			RegistrationId: response.RequestId,
		}
		js, _ := json.Marshal(jsonData)
		w.Header().Set("content-type", "application/json")
		w.Write(js)
		return

	}
	http.Error(w, "Invalid Method.", http.StatusNotFound)
	return
}	

func RegisterEmp(w http.ResponseWriter, req *http.Request)  {
	if req.Method == http.MethodPost{
		var employeeStruct *json_structs.EmployeeDetails
		decoder := json.NewDecoder(req.Body)
		if err := decoder.Decode(&employeeStruct); err != nil {
			fmt.Println(err)
			http.Error(w, "Invalid Request", http.StatusBadRequest)
			return
		}

		response := mongo_ops.RegisterEmployee(employeeStruct)
		jsonData := json_structs.Res{
			Status:response.Status, 
			Message:response.Message,
			RegistrationId: response.RequestId,
		}
		js, _ := json.Marshal(jsonData)
		w.Header().Set("content-type", "application/json")
		w.Write(js)
		return
	}
	http.Error(w, "Invalid Method !", http.StatusNotFound)
	return
}

func GetAdminsList(w http.ResponseWriter, req *http.Request){
	if req.Method == http.MethodPost{
		return
	}
	return
}

func GetEmployeeList(w http.ResponseWriter, req *http.Request){
	data := mongo_ops.GetAllEmployee()
	fmt.Println(data)
	if data != nil {
		w.Header().Set("content-type", "application/json")
		js, _ := json.Marshal(data)
		w.Write(js)
		return 
	}
	http.Error(w, "Internal Error / zero Employees", http.StatusInternalServerError )
}

func SignIn(w http.ResponseWriter, req *http.Request){
	if req.Method == http.MethodPost{
		var reqJs json_structs.SignInDetails 
		decoder := json.NewDecoder(req.Body)
		err := decoder.Decode(&reqJs)
		if err != nil {
			http.Error(w, "Bad Request credentials", http.StatusBadRequest)
			return
		}
		res, msg := mongo_ops.UserSignIn(reqJs.Email, reqJs.Password)
		if res {
			response := json_structs.Res{Status: true, Message : msg}
			js, _ := json.Marshal(response)
			
			w.Write(js)
		} 
	}
	http.Error(w, "Invalid Method", http.StatusNotFound)
	return
}