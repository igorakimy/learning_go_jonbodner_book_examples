package stub

type Person struct {
	Name string
	Age  string
}

type User struct {
	Name string
	Age  int
}

type Pet struct {
	Name  string
	Owner User
}

type Entities interface {
	GetUser(id string) (User, error)
	GetPets(userID string) ([]Pet, error)
	GetChildren(userID string) ([]Person, error)
	GetFriends(userID string) ([]Person, error)
	SaveUser(user User) error
}

type Logic struct {
	Entities Entities
}

func (l Logic) GetPetNames(userId string) ([]string, error) {
	pets, err := l.Entities.GetPets(userId)
	if err != nil {
		return nil, err
	}
	out := make([]string, 0, len(pets))
	for _, p := range pets {
		out = append(out, p.Name)
	}
	return out, nil
}
