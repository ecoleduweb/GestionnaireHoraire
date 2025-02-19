package DTOs

//Modèle de retour pour les erreurs de validation des champs
type FieldError struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}
