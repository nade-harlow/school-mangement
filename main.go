package main

import (
	"fmt"
)

func main()  {
	p:= Principal{name: "james"}
	a:= applicant{name: "jimmy",
		age: 12}
	a1:= applicant{name: "lyn",
		age: 12}
	a2:= applicant{name: "sike",
		age: 11}
	a3:= applicant{name: "toliver",
		age: 15}

	fmt.Println(p.Admit(a))
	//fmt.Println(c.TakeCourse)
	fmt.Println(p.Admit(a1))
	fmt.Println(p.Admit(a2))
	fmt.Println(p.Admit(a3))
	fmt.Println("............................")
	//c := Course{
	//	subject: map[string]string{
	//		"Eng": "A",
	//		"Math": "B",
	//	},
	//}
	s := students{
		name: "John",
		id: 2,
	}

	//c.TakeCourse(s)
	fmt.Println(StudentsList)

	t := Teachers{
		name: "Tim",
	}
	fmt.Println(t.Grade(&s))

	//s := students{id: 1,name: "jimmy"}
	//fmt.Println(p.Expel(s))

}
