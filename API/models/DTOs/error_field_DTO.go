package DTOs

//Modèle de retour pour les erreurs de validation des champs
type FieldErrorDTO struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}
