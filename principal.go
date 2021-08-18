package main

type Principal struct {
	name string

}

func (p Principal) Admit(a applicant) map[int]students {
	//a.Applicants("frank", 12)
	if a.age <= 12{
		StudentsList[1]= students{
			id: 1,
			name: a.name,
			class: "js1",

		}
	}
	if a.age > 12 && a.age <=15{
		StudentsList[1]= students{
			id: 1,
			name: a.name,
			class: "js2",
		}
	}

	return StudentsList
}

func(p Principal) Expel(s students)map[int]students  {
	delete(StudentsList, s.id)
	return StudentsList
}

