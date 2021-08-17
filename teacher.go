package main

type Teachers struct{
	name string
	//subject string
}

func(t Teachers) Grade(s *students) students {
	t.name= "tim"
	//c := Course{
	//	subject: map[string]string{},
	//}
	s.subject = map[string]string{
		"English": "A",
		"Math": "B",
		"Chemistry": "A",
	}
	return *s

}

func (c *Course) cc(res ...string)  {
	c.subject = map[string]string{
		"math":res[0],
		"english":res[1],
	}
}


