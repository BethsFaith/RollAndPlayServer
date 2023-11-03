package Db

type Skill struct {
	id         int
	Name       string
	Icon       string
	CategoryId int
}

func Create(db *Common, skill *Skill) error {
	_, err := db.Create(insertQ+skillsT+skillsP+"values ($1, $2, $3)",
		skill.Name, skill.Icon, skill.CategoryId)
	return err
}

func DeleteById(db *Common, skill *Skill) error {
	_, err := db.Delete(deleteQ+skillsT+"where id = ", skill.id)
	return err
}
