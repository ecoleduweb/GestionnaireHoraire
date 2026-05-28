package services

import (
	"fmt"
	"io"
	"llio-api/models/enums"
	"llio-api/repositories"
	"time"

	"github.com/xuri/excelize/v2"
)

func GenerateExcel(w io.Writer) error {
	activities, err := repositories.GetAllForExport()
	if err != nil {
		return fmt.Errorf("failed to fetch activities: %w", err)
	}

	loc, err := time.LoadLocation("America/Toronto")
	if err != nil {
		return fmt.Errorf("failed to load timezone: %w", err)
	}

	f := excelize.NewFile()
	sheet := "Sheet1"

	headers := []string{
		"Prénom", "Nom",
		"Projet", "Unique ID", "Statut du projet",
		"Activité", "Catégorie",
		"Date de début", "Date de fin", "Temps passé (h)",
	}

	for i, h := range headers {
		col := string(rune('A' + i))
		f.SetCellValue(sheet, col+"1", h)
	}

	row := 2

	for _, a := range activities {
		status := "inconnu"

		switch a.Project.Status {
		case enums.ProjectStatus(enums.Archived):
			status = "archivé"

		case enums.ProjectStatus(enums.InProgress):
			status = "actif"

		case enums.ProjectStatus(enums.NotStart):
			status = "non-démarré"

		case enums.ProjectStatus(enums.Cancel):
			status = "annulé"
		}

		f.SetCellValue(sheet, fmt.Sprintf("A%d", row), a.User.FirstName)
		f.SetCellValue(sheet, fmt.Sprintf("B%d", row), a.User.LastName)

		f.SetCellValue(sheet, fmt.Sprintf("C%d", row), a.Project.Name)
		f.SetCellValue(sheet, fmt.Sprintf("D%d", row), a.Project.UniqueId)
		f.SetCellValue(sheet, fmt.Sprintf("E%d", row), status)

		f.SetCellValue(sheet, fmt.Sprintf("F%d", row), a.Name)
		f.SetCellValue(sheet, fmt.Sprintf("G%d", row), a.Category.Name)

		f.SetCellValue(sheet, fmt.Sprintf("H%d", row), a.StartDate.In(loc).Format("2006-01-02 15:04"))
		f.SetCellValue(sheet, fmt.Sprintf("I%d", row), a.EndDate.In(loc).Format("2006-01-02 15:04"))
		f.SetCellValue(sheet, fmt.Sprintf("J%d", row), a.TimeSpent)

		row++
	}

	return f.Write(w)
}
