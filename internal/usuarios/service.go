package internal

type Service interface {
	GetAll() ([]Usuario, error)
	Store(nombre, apellido string, email string, edad int, altura float64, activo bool, fecha_creacion string) (Usuario, error)
	Update(id int, nombre, apellido string, email string, edad int, altura float64, activo bool, fecha_creacion string) (Usuario, error)
	UpdateNombre(id int, nombre string) (Usuario, error)
	Delete(id int) error
}

type service struct {
	repository Repository
}

func NewService(repository Repository) Service {
	return &service{repository: repository}
}

func (s *service) GetAll() ([]Usuario, error) {
	usuarios, err := s.repository.GetAll()
	if err != nil {
		return nil, err
	}
	return usuarios, nil
}

func (s *service) Store(nombre, apellido string, email string, edad int, altura float64, activo bool, fecha_creacion string) (Usuario, error) {
	ultimoId, err := s.repository.LastId()
	if err != nil {
		return Usuario{}, err
	}

	usr, err := s.repository.Store(ultimoId+1, nombre, apellido, email, edad, altura, activo, fecha_creacion)
	if err != nil {
		return Usuario{}, err
	}
	return usr, nil
}

func (ser *service) Update(id int, nombre, apellido string, email string, edad int, altura float64, activo bool, fecha_creacion string) (Usuario, error) {
	return ser.repository.Update(id, nombre, apellido, email, edad, altura, activo, fecha_creacion)
}

func (ser *service) UpdateNombre(id int, nombre string) (Usuario, error) {
	return ser.repository.UpdateNombre(id, nombre)
}

func (ser *service) Delete(id int) error {
	return ser.repository.Delete(id)
}
