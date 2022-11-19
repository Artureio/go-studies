package clients

type Owner struct {
	Name       string
	Age        int
	Phone      string
	CPF        string
	Profession string
}

func CreateNewOwner(Name string, Age int, Phone string, CPF string, Profession string) Owner {
	return Owner{Name, Age, Phone, CPF, Profession}
}
