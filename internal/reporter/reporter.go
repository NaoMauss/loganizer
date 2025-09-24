package reporter

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

type LogReport struct {
	LogID        string `json:"log_id"`
	FilePath     string `json:"file_path"`
	Status       string `json:"status"`
	Message      string `json:"message"`
	ErrorDetails string `json:"error_details,omitempty"`
}

func ExportReports(reports []LogReport, outputPath string) error {
	if err := os.MkdirAll(filepath.Dir(outputPath), 0755); err != nil {
		return fmt.Errorf("impossible de créer le répertoire de sortie: %w", err)
	}

	file, err := os.Create(outputPath)
	if err != nil {
		return fmt.Errorf("impossible de créer le fichier de sortie %s: %w", outputPath, err)
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	if err := encoder.Encode(reports); err != nil {
		return fmt.Errorf("erreur lors de l'export JSON: %w", err)
	}

	return nil
}

func FilterReports(reports []LogReport, status string) []LogReport {
	if status == "" {
		return reports
	}

	var filtered []LogReport
	for _, report := range reports {
		if report.Status == status {
			filtered = append(filtered, report)
		}
	}
	return filtered
}
