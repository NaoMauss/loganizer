package cmd

import (
	"encoding/json"
	"fmt"
	"os"

	"tp3-loganizer/internal/config"

	"github.com/spf13/cobra"
)

var (
	logID      string
	logPath    string
	logType    string
	configFile string
)

var addLogCmd = &cobra.Command{
	Use:   "add-log",
	Short: "Ajouter une configuration de log au fichier config.json",
	Long: `La commande add-log permet d'ajouter manuellement une nouvelle
configuration de log à un fichier config.json existant.

Exemple d'utilisation:
  loganizer add-log --id web-server-3 --path /var/log/nginx/error.log --type nginx-error --file config.json`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if logID == "" || logPath == "" || logType == "" || configFile == "" {
			return fmt.Errorf("tous les drapeaux sont requis: --id, --path, --type, --file")
		}

		return addLogToConfig()
	},
}

func addLogToConfig() error {
	var configs []config.LogConfig

	if _, err := os.Stat(configFile); err == nil {
		existingConfigs, err := config.LoadConfig(configFile)
		if err != nil {
			return fmt.Errorf("erreur lors de la lecture du fichier de configuration: %w", err)
		}
		configs = existingConfigs
	}

	for _, cfg := range configs {
		if cfg.ID == logID {
			return fmt.Errorf("un log avec l'ID '%s' existe déjà", logID)
		}
	}

	newConfig := config.LogConfig{
		ID:   logID,
		Path: logPath,
		Type: logType,
	}
	configs = append(configs, newConfig)

	file, err := os.Create(configFile)
	if err != nil {
		return fmt.Errorf("impossible de créer le fichier de configuration: %w", err)
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	if err := encoder.Encode(configs); err != nil {
		return fmt.Errorf("erreur lors de la sauvegarde: %w", err)
	}

	fmt.Printf("Configuration ajoutée avec succès:\n")
	fmt.Printf("  ID: %s\n", logID)
	fmt.Printf("  Path: %s\n", logPath)
	fmt.Printf("  Type: %s\n", logType)
	fmt.Printf("  Fichier: %s\n", configFile)

	return nil
}

func init() {
	rootCmd.AddCommand(addLogCmd)

	// Drapeaux pour la commande add-log
	addLogCmd.Flags().StringVar(&logID, "id", "", "Identifiant unique du log (requis)")
	addLogCmd.Flags().StringVar(&logPath, "path", "", "Chemin vers le fichier de log (requis)")
	addLogCmd.Flags().StringVar(&logType, "type", "", "Type du log (requis)")
	addLogCmd.Flags().StringVar(&configFile, "file", "", "Chemin vers le fichier config.json (requis)")

	// Marquer tous les flags comme requis
	addLogCmd.MarkFlagRequired("id")
	addLogCmd.MarkFlagRequired("path")
	addLogCmd.MarkFlagRequired("type")
	addLogCmd.MarkFlagRequired("file")
}
