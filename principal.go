package main

import "fmt"

type Principal struct {
	name string

}

func (p Principal) Admit(a applicant) map[int]students {
	//a.Applicants("frank", 12)
	if a.age <= 12{
		StudentsList[len(StudentsList)+1]= students{
			id: len(StudentsList)+1,
			name: a.name,
			class: "js1",

		}
	}
	if a.age > 12 && a.age <=15{
		StudentsList[len(StudentsList)+1]= students{
			id: len(StudentsList)+1,
			name: a.name,
			class: "js2",
		}
	}else{
		fmt.Println("You are too old for to attend our school")
	}

	return StudentsList
}

func(p Principal) Expel(s students)map[int]students  {
	delete(StudentsList, s.id)
	return StudentsList
}

