package db

type Appointment struct {
	ID      string `json:"id,omitempty"`
	Patient struct {
		Name       string `json:"name"`
		DocumentId string `json:"document_id"`
	}
	Specialty string `json:"specialty"`
	Date      string `json:"date"`
}

type AppointmentRepository struct {
	Dao DAO
}

func (repository *AppointmentRepository) FindOneByID(ID string) (Appointment, error) {

	input := map[string]interface{}{
		"id": ID,
	}

	result, err := repository.Dao.Read(input)
	if err != nil {
		return Appointment{}, err
	}

	return result.(Appointment), nil
}
