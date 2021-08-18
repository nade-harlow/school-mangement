package main



type applicant struct {
	name string
	age int
}


func (a applicant) Applicants( name string, age int)  {
	a.name= name
	a.age= age
}