import pymongo
import datetime


client = pymongo.MongoClient("mongodb://localhost:27017")

db = client.hrms

emp = db.employee

admin = db.admin


emp.drop()
admin.drop()

employee_one = {
   
	"name":"ashish",
	"job_role":"director",
	"date_of_birth":"24101996",
	"blood_group":"b-negative",
	"team":"administration",
	"manager":"self",
	"hr":"priyanka",
	"joining_date":"01012024",
	"org_email":"ashish.s.maurya@gmail.com",
	"password":"theguy",
	"personal_email":"ashish@gamil.com",
	"created_at":str(datetime.datetime.now())
}
amdin_one = {
    "name":"ashish",
	"job_role":"director",
	"date_of_birth":"24101996",
	"team":"administration",
	"manager":"self",
	"hr":"priyanka",
	"joining_date":"01012024",
	"org_email":"ashish@akm.com",
	"password":"theguy",
	"personal_email":"ashish@gamil.com",
	"created_at": str(datetime.datetime.now()),
	"approval_status":"completed",
	"approved_by":"ashish",
    "admin-category": "super-admin"
}
emp.insert_one(employee_one)

res = admin.insert_one(amdin_one)

print(res.acknowledged)

