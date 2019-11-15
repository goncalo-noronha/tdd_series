package db

type Appointment struct {
	ID string
	name string
}

type AppointmentRepository struct {
	dao DAO
}

func (repository *AppointmentRepository) FindOneByID(ID string) (Appointment, error) {

	input := map[string]string{
		"ID": ID,
	}

	result, err := repository.dao.Read(input)
	if err != nil {
		return Appointment{}, err
	}

	return result.(Appointment), nil
}
