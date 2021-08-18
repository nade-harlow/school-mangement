package main

import (
	"reflect"
	"testing"
)

func TestPrincipal_Admit(t *testing.T)  {
	admit:= []struct {
	input1 Principal
	input2 applicant
	expect map[int]students
	}{
		{
			input1: Principal{name: "tim"},
			input2: applicant{
				name: "jim",
				age: 12,
			},
			expect: map[int]students{1: {id: 1 ,name: "jim", class: "js1"}},
		},
		{
			//input1: Principal{name: "tim"},
			input2: applicant{
				name: "timmy",
				age: 14,
			},
			expect: map[int]students{1: {id: 1 ,name: "jim", class: "js1"}, 2: {id: 2 ,name: "timmy", class: "js2"}},
		},

	}
	for _, v := range admit {
		f := v.input1.Admit(v.input2)
		if !reflect.DeepEqual(f,v.expect){
			t.Fatalf("Expecting: %v, Got %v", v.expect, f)
		}
	}
}

func TestPrincipal_Expel(t *testing.T)  {
	expel:= []struct {
		input1 Principal
		input2 students
		expect map[int]students
	}{
		{
			input1: Principal{name: "tim"},
			input2: students{
				id: 1,
				name: "jim",
				class: "js1",

			},
			expect: map[int]students{2:{id: 2 ,name: "timmy", class:"js2" }},
		},
		}

	for _, v := range expel {
		f := v.input1.Expel(v.input2)
		if !reflect.DeepEqual(f,v.expect){
			t.Fatalf("Expecting: %v, Got %v", v.expect, f)
		}
	}
}
