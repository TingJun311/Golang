package StackArray

type StackArray struct {
	Stack []interface{}
}

func (sa *StackArray) Push(content interface{}) {
	head := *sa
	head = append(head, content)
}

func (sa *StackArray) Pop() {
	if 
}

func (sa *StackArray) Length() (length int){
	for _, i := range sa {
		length++
	}
	return length
}
