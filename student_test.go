package main

import (
	"reflect"
	"testing"
)

func TestStudents_TakeCourse(t *testing.T)  {
	course:= []struct {
		input1 students
		input2 Course
		expect map[int]students
	}{
		{
			input1: students{
				id: 3,
				name: "tim",
				class: "js1",
			},
			input2: Course{
				subject: map[string]string{"Math":"","English":""},
			},
			expect: map[int]students{3: students{id: 3 ,name: "tim", class: "js1", subject: map[string]string{"Math":"","English":""}}, 2:{id: 2, name: "timmy", class: "js2"}},
		},
	}
	for _, v := range course{
		f := v.input1.TakeCourse(v.input2)
		if !reflect.DeepEqual(f,v.expect){
			t.Fatalf("Expecting: %v, Got %v", v.expect, f)
		}
	}
}

