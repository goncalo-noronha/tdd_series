package app

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
	Gateway PersistenceGateway
}

func (repository *AppointmentRepository) FindOneByID(ID string) (Appointment, error) {

	input := map[string]interface{}{
		"id": ID,
	}

	result, err := repository.Gateway.Read(input)
	if err != nil {
		return Appointment{}, err
	}

	return result.(Appointment), nil
}
