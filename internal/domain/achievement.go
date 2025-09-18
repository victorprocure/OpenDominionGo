package domain

type Achievement struct {
	Name        *string
	Description *string
	Icon        *string
}

func NewAchievement(name, description, icon *string) *Achievement {
	return &Achievement{Name: name, Description: description, Icon: icon}
}
