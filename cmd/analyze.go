package cmd

import (
	"fmt"

	"tp3-loganizer/internal"

	"github.com/spf13/cobra"
)

var (
	configPath   string
	outputPath   string
	statusFilter string
)

var analyzeCmd = &cobra.Command{
	Use:   "analyze",
	Short: "Analyser des fichiers de logs en parallèle",
	Long: `La commande analyze traite plusieurs fichiers de logs en parallèle
selon une configuration JSON et génère un rapport des résultats.

Exemple d'utilisation:
  loganizer analyze -c config.json
  loganizer analyze -c config.json -o report.json
  loganizer analyze -c config.json -o report.json --status FAILED`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if configPath == "" {
			return fmt.Errorf("le drapeau --config (-c) est requis")
		}

		return internal.ProcessAnalysis(configPath, outputPath, statusFilter)
	},
}

func init() {
	rootCmd.AddCommand(analyzeCmd)

	analyzeCmd.Flags().StringVarP(&configPath, "config", "c", "", "Chemin vers le fichier de configuration JSON (requis)")
	analyzeCmd.Flags().StringVarP(&outputPath, "output", "o", "", "Chemin vers le fichier de sortie JSON (optionnel)")
	analyzeCmd.Flags().StringVar(&statusFilter, "status", "", "Filtrer les résultats par statut (OK, FAILED)")

	analyzeCmd.MarkFlagRequired("config")
}
