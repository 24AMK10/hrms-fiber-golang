package json_structs
import (
	"time"
)

var TheSecret = "87g42@!2hius3##7y0rwhuiefdwy790"

type Res struct {
	Status bool
	Message string
	RegistrationId string
}

type SignInDetails struct {
	Email string `json:"email"`
	Password string `json:"password"`

}

type RegistrationRes struct {
	Status bool
	RequestId string
	Message string
}

// don't give spaces in while declaring the json
type EmployeeDetails struct { // this is the struct for the employee
	Id string 			`json:"_id"  bson:"_id"`
	Name string 		`json:"name" bson:"name"`
	JobRole string 		`json:"job_role" bson:"job_role"`
	DateOfBirth string 	`json:"date_of_birth" bson:"date_of_birth"`
	BloodGroup string 	`json:"blood_group" bson:"blood_group"`
	Team string 		`json:"team" bson:"team"`
	Manager string 		`json:"manager" bson:"manager"`
	Hr string 			`json:"hr" bson:"hr"`
	JoiningDate string 	`json:"joining_date" bson:"joining_date"`
	OrgEmail string 	`json:"org_email" bson:"org_email"`
	Password string 	`json:"password" bson:"password"`
	PersonalEmail string `json:"personal_email" bson:"personal_email"`
	CreatedAt time.Time  `json:"created_at" bson:"created_at"`

}

type Admin struct { // this is for the admin previlages 
	Name string 		`json:"name" bson:"name"`
	JobRole string 		`json:"job_role" bson:"job_role"`
	DateOfBirth string 	`json:"date_of_birth" bson:"date_of_birth"`
	Team string 		`json:"team" bson:"team"`
	Manager string 		`json:"manager" bson:"manager"`
	Hr string 			`json:"hr" bson:"hr"`	
	OrgEmail string 	`json:"org_email" bson:"org_email" bson:"org_email" bson:"org_email"`
	Password string 	`json:"password" bson:"password"`
	PersonalEmail string `json:"personal_email" bson:"personal_email"`
	CreatedAt time.Time  `json:"created_at" bson:"created_at"`
	ApprovalStatus string `json:"approval_status" bson:"approval_status"`
	ApprovedBy string 	`json:"approved_by" bson:"approved_by"`
}
	