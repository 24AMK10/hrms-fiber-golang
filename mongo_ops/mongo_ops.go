package mongo_ops

import (
	"fmt"
	// "encoding/json"
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"

	// "go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"hrms.co/example/json_structs"
	// "reflect"
)

const (
	MongoURI     = "mongodb://localhost:27017"
	databaseName = "hrms"
	employee_col = "employee"
	admin_col    = "admin"
)

var client *mongo.Client
var employeeDB *mongo.Collection
var adminDB *mongo.Collection
var AllEmployees []json_structs.EmployeeDetails

func ConnectMongo(MongoURI string) error {
	clientOptions := options.Client().ApplyURI(MongoURI)

	var err error
	client, err = mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		fmt.Println("cant connect right now ::: >>>", err)
		return err
	}
	adminDB = client.Database(databaseName).Collection(admin_col)
	employeeDB = client.Database(databaseName).Collection(employee_col)
	return nil
}

func DisconnectMongo() error {
	if err := client.Disconnect(context.Background()); err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func MongoInsertOne() {
	return
}
func RegisterAdmin(admin *json_structs.Admin) json_structs.RegistrationRes {

	var res json_structs.Admin
	var response json_structs.RegistrationRes
	adminJs := bson.D{{Key: "org_email", Value: admin.OrgEmail}}
	if err := adminDB.FindOne(context.TODO(), adminJs).Decode(&res); err != nil {
		if err == mongo.ErrNoDocuments {
			fmt.Println(" there is no data found ")
			admin.ApprovalStatus = "pending"
			admin.CreatedAt = time.Now()
			js := bson.D{
				{Key: "name", Value: admin.Name},
				{Key: "job_role", Value: admin.JobRole},
				{Key: "date_of_birth", Value: admin.DateOfBirth},
				{Key: "blood_group", Value: admin.Team},
				{Key: "team", Value: admin.Team},
				{Key: "manager", Value: admin.Manager},
				{Key: "hr", Value: admin.Hr},
				{Key: "org_email", Value: admin.OrgEmail},
				{Key: "password", Value: admin.Password},
				{Key: "personal_email", Value: admin.PersonalEmail},
				{Key: "created_at", Value: admin.CreatedAt},
				{Key: "approval_status", Value: admin.ApprovalStatus},
				{Key: "approved_by", Value: ""},
			}
			if err != nil {
				fmt.Println("there was an error")
			}
			res, err := adminDB.InsertOne(context.TODO(), js)
			if err != nil {
				fmt.Println(err)
			}
			if res != nil {
				fmt.Println("inserted the data", res)
				response.Status = true
				response.RequestId = res.InsertedID.(string)
				response.Message = "Request sent for approval"
			}
		}

	}

	if res.Name == admin.Name {
		response.Status = false
		response.RequestId = "Refer to original request"
		response.Message = "In Progress/ Email already registered"
	}

	return response
}

func RegisterEmployee(employee *json_structs.EmployeeDetails) json_structs.RegistrationRes {
	filter := bson.D{{Key: "org_email", Value: employee.OrgEmail}}
	var employeeDetails json_structs.EmployeeDetails
	var regRes json_structs.RegistrationRes

	if err := employeeDB.FindOne(context.TODO(), filter).Decode(&employeeDetails); err != nil {
		if err == mongo.ErrNoDocuments {
			entryTime := time.Now()
			js := bson.D{
				{Key: "name", Value: employee.Name},
				{Key: "job_role", Value: employee.JobRole},
				{Key: "date_of_birth", Value: employee.DateOfBirth},
				{Key: "blood_group", Value: employee.BloodGroup},
				{Key: "team", Value: employee.Team},
				{Key: "manager", Value: employee.Manager},
				{Key: "hr", Value: employee.Hr},
				{Key: "joining_date", Value: employee.JoiningDate},
				{Key: "org_email", Value: employee.OrgEmail},
				{Key: "password", Value: employee.Password},
				{Key: "personal_email", Value: employee.PersonalEmail},
				{Key: "created_at", Value: entryTime},
			}
			fmt.Println(employeeDetails)
			fmt.Println(js)
			res, err := employeeDB.InsertOne(context.TODO(), js)
			if err != nil {
				fmt.Println(time.Now(), ":: employee insertion failed", err)
			}
			if res != nil {
				regRes.Status = true

				regRes.Message = "Request sent for approval"
				return regRes
			}

		}
	}
	regRes.Status = false
	regRes.Message = "Email used / In Progress"
	regRes.RequestId = "NA"
	return regRes

}

func GetActiveAdmins() {
	filter := bson.D{
		{Key: "approval_status", Value: "completed"},
	}
	var response json_structs.Res
	res, err := adminDB.Find(context.TODO(), filter)
	if err != nil {
		response.Status = false
		response.Message = "There was error getting the data"
	}
	if res != nil {
		fmt.Println(res)
	}

}

func GetAllEmployee() interface{} {
	// filterId := bson.D{{Key: "_id", Value:false},}
	filter := bson.M{}
	findOptions := options.Find().SetProjection(bson.M{"created_at": 0})

	data, err := employeeDB.Find(context.TODO(), filter, findOptions)
	if err != nil {
		log.Println(time.Now(), err)
		return nil
	}

	err = data.All(context.TODO(), &AllEmployees)
	if err != nil {
		fmt.Println(err)
	}
	return AllEmployees
}

func UserSignIn(email, password string) (bool, string) {
	filter := bson.D{
		{Key: "org_email", Value: email},
	}
	var employee json_structs.EmployeeDetails
	err := employeeDB.FindOne(context.TODO(), filter).Decode(&employee)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return false, "Please Register"
		}
	}
	if employee.OrgEmail == email && employee.Password == password {
		return true, "Success"
	}
	return false, "Invalid Credentials"
}
