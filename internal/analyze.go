package internal

import (
	"errors"
	"fmt"
	"math/rand"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"

	"tp3-loganizer/internal/analyzer"
	"tp3-loganizer/internal/config"
	"tp3-loganizer/internal/reporter"
)

type AnalyzeResult struct {
	LogConfig config.LogConfig
	Report    reporter.LogReport
}

func AnalyzeLog(logConfig config.LogConfig) reporter.LogReport {
	if _, err := os.Stat(logConfig.Path); err != nil {
		if os.IsNotExist(err) {
			fileErr := analyzer.NewFileNotFoundError(logConfig.Path, err)
			return reporter.LogReport{
				LogID:        logConfig.ID,
				FilePath:     logConfig.Path,
				Status:       "FAILED",
				Message:      "Fichier introuvable.",
				ErrorDetails: fileErr.Error(),
			}
		}
		return reporter.LogReport{
			LogID:        logConfig.ID,
			FilePath:     logConfig.Path,
			Status:       "FAILED",
			Message:      "Fichier inaccessible.",
			ErrorDetails: err.Error(),
		}
	}

	sleepDuration := time.Duration(rand.Intn(151)+50) * time.Millisecond
	time.Sleep(sleepDuration)

	return reporter.LogReport{
		LogID:    logConfig.ID,
		FilePath: logConfig.Path,
		Status:   "OK",
		Message:  "Analyse terminée avec succès.",
	}
}

func AnalyzeLogs(logConfigs []config.LogConfig) []reporter.LogReport {
	var wg sync.WaitGroup
	resultsChan := make(chan reporter.LogReport, len(logConfigs))

	for _, logConfig := range logConfigs {
		wg.Add(1)
		go func(lc config.LogConfig) {
			defer wg.Done()
			result := AnalyzeLog(lc)
			resultsChan <- result
		}(logConfig)
	}

	go func() {
		wg.Wait()
		close(resultsChan)
	}()

	var reports []reporter.LogReport
	for report := range resultsChan {
		reports = append(reports, report)
	}

	return reports
}

func ProcessAnalysis(configPath, outputPath, statusFilter string) error {
	logConfigs, err := config.LoadConfig(configPath)
	if err != nil {
		var parseErr *analyzer.ParseError
		if errors.As(err, &parseErr) {
			return fmt.Errorf("erreur de parsing de la configuration: %w", err)
		}
		return fmt.Errorf("erreur lors du chargement de la configuration: %w", err)
	}

	if len(logConfigs) == 0 {
		return fmt.Errorf("aucune configuration de log trouvée dans le fichier")
	}

	fmt.Printf("Analyse de %d log(s) en cours...\n\n", len(logConfigs))
	reports := AnalyzeLogs(logConfigs)

	if statusFilter != "" {
		reports = reporter.FilterReports(reports, statusFilter)
		fmt.Printf("Filtrage par statut: %s\n", statusFilter)
	}

	displayResults(reports)

	if outputPath != "" {
		timestampedPath := addTimestampToFilename(outputPath)
		if err := reporter.ExportReports(reports, timestampedPath); err != nil {
			return fmt.Errorf("erreur lors de l'export: %w", err)
		}
		fmt.Printf("\nRapport exporté vers: %s\n", timestampedPath)
	}

	return nil
}

func displayResults(reports []reporter.LogReport) {
	successCount := 0
	failedCount := 0

	for _, report := range reports {
		fmt.Printf("ID: %s\n", report.LogID)
		fmt.Printf("Chemin: %s\n", report.FilePath)
		fmt.Printf("Statut: %s\n", report.Status)
		fmt.Printf("Message: %s\n", report.Message)
		if report.ErrorDetails != "" {
			fmt.Printf("Détails de l'erreur: %s\n", report.ErrorDetails)
		}
		fmt.Println("---")

		if report.Status == "OK" {
			successCount++
		} else {
			failedCount++
		}
	}

	fmt.Printf("\nRésumé: %d succès, %d échecs\n", successCount, failedCount)
}

func addTimestampToFilename(filePath string) string {
	now := time.Now()
	timestamp := now.Format("060102")

	dir := filepath.Dir(filePath)
	ext := filepath.Ext(filePath)
	name := strings.TrimSuffix(filepath.Base(filePath), ext)

	return filepath.Join(dir, fmt.Sprintf("%s_%s%s", timestamp, name, ext))
}
