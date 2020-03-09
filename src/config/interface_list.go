package config

// Person 人class
type Person struct {
	name string 
	age int
	des string
}

// GetSchoolMenber 外部调用
func GetSchoolMenber()(p []Person){
	person := []Person{
		{
			name : "abd",
			age : 14,
			des : "I am a student.",
		},
		{
			name : "dfe",
			age : 14,
			des : "I am a student.",
		},
		{
			name : "hgf",
			age : 14,
			des : "I am a student.",
		},
	}

	
	return person
}
