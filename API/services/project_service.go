package services

import (
	"llio-api/models/DAOs"
	"llio-api/models/DTOs"
	"llio-api/models/enums"
	"llio-api/repositories"
	"strconv"

	"github.com/jinzhu/copier"
)

func VerifyProjectJSON(projectDTO *DTOs.ProjectDTO) []DTOs.FieldErrorDTO {
	var errors []DTOs.FieldErrorDTO

	if !projectDTO.CreatedAt.IsZero() {
		// Vérifier que CreatedAt est avant EndAt
		if !projectDTO.EndAt.IsZero() && projectDTO.CreatedAt.After(projectDTO.EndAt) {
			errors = append(errors, DTOs.FieldErrorDTO{
				Field:   "endAt",
				Message: "La date de fin doit être après la date de création",
			})
		}

		// Vérifier que CreatedAt est avant UpdatedAt
		if !projectDTO.UpdatedAt.IsZero() && projectDTO.CreatedAt.After(projectDTO.UpdatedAt) {
			errors = append(errors, DTOs.FieldErrorDTO{
				Field:   "updatedAt",
				Message: "La date de mise à jour doit être après la date de création",
			})
		}
	}

	return errors
}

func CreateProject(projectDTO *DTOs.ProjectDTO) (*DTOs.ProjectDTO, error) {

	project := &DAOs.Project{}
	projectDTO.Description = ""
	err := copier.Copy(project, projectDTO)
	if err != nil {
		return nil, err
	}

	projectDAOAdded, err := repositories.CreateProject(project)
	if err != nil {
		return nil, err
	}

	projectDTOResponse := &DTOs.ProjectDTO{}
	err = copier.Copy(projectDTOResponse, projectDAOAdded)
	return projectDTOResponse, err
}

func GetProjects() ([]map[string]any, error) {
	projects, err := repositories.GetProjects()
	if err != nil {
		return nil, err
	}

	var result []map[string]any
	for _, project := range projects {
		projectDTO := &DTOs.ProjectDTO{}
		err = copier.Copy(projectDTO, project)
		formattedProject := formatProjectDTO(projectDTO)
		result = append(result, formattedProject)
	}

	return result, err
}

func GetProjectsByManagerId(id int) ([]map[string]any, error) {
	projects, err := repositories.GetProjectsByManagerId(id)
	if err != nil {
		return nil, err
	}

	var result []map[string]any
	for _, project := range projects {
		projectDTO := &DTOs.ProjectDTO{}
		err = copier.Copy(projectDTO, project)
		formattedProject := formatProjectDTO(projectDTO)
		result = append(result, formattedProject)
	}

	return result, err
}

func GetProjectById(id string) (*DTOs.ProjectDTO, error) {
	project, err := repositories.GetProjectById(id)
	if err != nil {
		return nil, err
	}

	projectDTO := &DTOs.ProjectDTO{}
	err = copier.Copy(projectDTO, project)

	return projectDTO, err
}

func UpdateProject(projectDTO *DTOs.ProjectDTO) (*DTOs.ProjectDTO, error) {

	projectDAO := &DAOs.Project{}
	err := copier.Copy(projectDAO, projectDTO)
	if err != nil {
		return nil, err
	}

	projectDAOUpdated, err := repositories.UpdateProject(projectDAO)
	if err != nil {
		return nil, err
	}

	projectDTOResponse := &DTOs.ProjectDTO{}
	err = copier.Copy(projectDTOResponse, projectDAOUpdated)
	return projectDTOResponse, err
}

func formatProjectDTO(projectDTO *DTOs.ProjectDTO) map[string]any {
	manager, _ := GetUserById(strconv.Itoa(projectDTO.ManagerId))
	lead := manager.FirstName + " " + manager.LastName

	employeesMap := make(map[string][]map[string]any)

	// Process activities to build employees and categories
	for _, activity := range projectDTO.Activities {
		user, _ := GetUserById(strconv.Itoa(activity.UserId))
		userName := user.FirstName + " " + user.LastName

		// Get category name
		category, _ := GetCategoryById(strconv.Itoa(activity.CategoryId))

		// Add or update employee's categories
		found := false
		for _, cat := range employeesMap[userName] {
			if cat["name"] == category.Name {
				cat["timeSpent"] = cat["timeSpent"].(float64) + activity.EndDate.Sub(activity.StartDate).Hours()
				found = true
				break
			}
		}

		if !found {
			employeesMap[userName] = append(employeesMap[userName], map[string]any{
				"name":          category.Name,
				"timeSpent":     activity.EndDate.Sub(activity.StartDate).Hours(),
				"timeEstimated": 0, // You'll need to add this to your category model
			})
		}
	}

	// Convert employees map to array
	var employees []map[string]any
	for name, categories := range employeesMap {
		employees = append(employees, map[string]any{
			"name":       name,
			"categories": categories,
		})
	}

	// Calculate total time spent
	var totalTimeSpent float64
	for _, empl := range employees {
		for _, cat := range empl["categories"].([]map[string]any) {
			totalTimeSpent += cat["timeSpent"].(float64)
		}
	}

	// Build the final response
	return map[string]any{
		"id":             projectDTO.Id,
		"name":           projectDTO.Name,
		"description":    projectDTO.Description,
		"lead":           lead,
		"coLeads":        []string{}, // Pas de fetch de coLeads pour l'instant
		"employees":      employees,
		"totalTimeSpent": totalTimeSpent,
		"isArchived":     projectDTO.Status == enums.ProjectStatus(enums.Finish),
	}
}
