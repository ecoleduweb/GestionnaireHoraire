package services

import (
	"llio-api/models/DAOs"
	"llio-api/models/DTOs"
	"llio-api/models/enums"
	"llio-api/repositories"

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
	err := copier.Copy(project, projectDTO)
	if err != nil {
		return nil, err
	}

	projectDAOAdded, err := repositories.CreateProject(project)
	if err != nil {
		return nil, err
	}

	_, err = repositories.CreateCategory(&DAOs.Category{
		Name:        "Par défaut",
		Description: "Catégorie par défaut",
		ProjectId:   projectDAOAdded.Id,
		CreatedAt:   projectDAOAdded.CreatedAt,
		UpdatedAt:   projectDAOAdded.UpdatedAt,
		Activities:  []DAOs.Activity{},
		UserId:      projectDAOAdded.ManagerId,
	})
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

	result, err := formatProjects(projects, err)
	if err != nil {
		return nil, err
	}
	return result, err
}

func GetProjectsList() ([]*DTOs.ProjectDTO, error) {
	projects, err := repositories.GetProjectsList()
	if err != nil {
		return nil, err
	}

	if len(projects) == 0 {
		return make([]*DTOs.ProjectDTO, 0), nil
	}

	projectsDTO := make([]*DTOs.ProjectDTO, len(projects))
	for i := range projectsDTO {
		projectsDTO[i] = &DTOs.ProjectDTO{}
	}

	if err := copier.Copy(&projectsDTO, &projects); err != nil {
		return nil, err
	}

	return projectsDTO, nil
}

func GetProjectsByManagerId(id int) ([]map[string]any, error) {
	projects, err := repositories.GetProjectsByManagerId(id)
	if err != nil {
		return nil, err
	}

	result, err := formatProjects(projects, err)
	if err != nil {
		return nil, err
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

func formatProjects(projects []*DAOs.Project, err error) ([]map[string]any, error) {
	users, err := repositories.GetAllUsers()
	if err != nil {
		return nil, err
	}
	userMap := make(map[int]*DAOs.User)
	for _, user := range users {
		userMap[user.Id] = user
	}

	categories, err := repositories.GetCategories()
	if err != nil {
		return nil, err
	}
	categoryMap := make(map[int]*DAOs.Category)
	for _, cat := range categories {
		categoryMap[cat.Id] = cat
	}

	var result []map[string]any
	for _, project := range projects {
		tempActivities, err := repositories.GetProjectActivities(project.Id)
		if err != nil {
			return nil, err
		}

		activities := make([]DAOs.Activity, len(tempActivities))
		for i, result := range tempActivities {
			activities[i] = DAOs.Activity{
				UserId:     result.UserID,
				CategoryId: result.CategoryID,
				ProjectId:  result.ProjectID,
				TimeSpent:  result.TimeSpent,
			}
		}

		formattedProject := formatProjectWithActivities(project, activities, userMap, categoryMap)
		result = append(result, formattedProject)
	}

	return result, nil
}

func formatProjectWithActivities(project *DAOs.Project, activities []DAOs.Activity, userMap map[int]*DAOs.User, categoryMap map[int]*DAOs.Category) map[string]any {
	employeesMap := make(map[int]map[string]any)

	for _, activity := range activities {
		user, exists := userMap[activity.UserId]
		if !exists {
			continue
		}

		category, exists := categoryMap[activity.CategoryId]
		if !exists {
			continue
		}

		if _, ok := employeesMap[user.Id]; !ok {
			employeesMap[user.Id] = map[string]any{
				"name":       user.FirstName + " " + user.LastName,
				"categories": make([]map[string]any, 0),
			}
		}

		categories := employeesMap[user.Id]["categories"].([]map[string]any)
		found := false
		for i, cat := range categories {
			if cat["name"] == category.Name {
				categories[i]["timeSpent"] = activity.TimeSpent
				found = true
				break
			}
		}

		if !found {
			categories = append(categories, map[string]any{
				"name":          category.Name,
				"timeSpent":     activity.TimeSpent,
				"timeEstimated": 0,
			})
			employeesMap[user.Id]["categories"] = categories
		}
	}

	// Convert to array
	var employees []map[string]any
	for _, emp := range employeesMap {
		employees = append(employees, emp)
	}

	// Calculate total time
	var totalTimeSpent float64
	for _, emp := range employees {
		for _, cat := range emp["categories"].([]map[string]any) {
			totalTimeSpent += cat["timeSpent"].(float64)
		}
	}

	// Get manager info
	manager, exists := userMap[project.ManagerId]
	lead := ""
	if exists {
		lead = manager.FirstName + " " + manager.LastName
	}

	var totalTimeEstimated float64 = 0 // À ajouter aux modèles

	return map[string]any{
		"id":                 project.Id,
		"name":               project.Name,
		"description":        project.Description,
		"lead":               lead,
		"coLeads":            []string{},
		"employees":          employees,
		"totalTimeEstimated": totalTimeEstimated,
		"totalTimeRemaining": totalTimeEstimated - totalTimeSpent,
		"totalTimeSpent":     totalTimeSpent,
		"isArchived":         project.Status == enums.ProjectStatus(enums.Finish),
	}
}
