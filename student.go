package main


type students struct {
	id int
	name string
	class string
	subject map[string]string
}


type Course struct{
	subject map[string]string
}


var StudentsList =map[int]students{
	//1: {
	//	id: 1,
	//	name: "jim",
	//	class: "js1",
	//
	//},

}


func(s students) TakeCourse(c Course) map[int]students  {

	s.subject = c.subject
	StudentsList[s.id] = s

	return StudentsList
}