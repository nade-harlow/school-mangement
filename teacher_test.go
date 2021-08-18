package main

import (
	"reflect"
	"testing"
)

func TestTeachers_Grade(t *testing.T)  {
	grade:= []struct {
		input1 Teachers
		input2 students
		expect students
	}{
		{
			input1: Teachers{
				"Miss Ann",
			},
			input2: students{
				id: 1,
				name: "jim",
				class: "js1",
				subject: map[string]string{
					"Math": "A",
					"English": "B",
				},
			},
			expect: students{id: 1 ,name: "jim", class: "js1", subject: map[string]string{"Math":"A","English":"B"}},
		},
	}
	for _, v := range grade{
		f := v.input1.Grade(&v.input2)
		if !reflect.DeepEqual(f,v.expect){
			t.Fatalf("Expecting: %v, Got %v", v.expect, f)
		}
	}
}
