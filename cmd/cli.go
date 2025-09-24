// Package cmd contient les commandes CLI de l'application
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// rootCmd représente la commande de base quand appelée sans sous-commandes
var rootCmd = &cobra.Command{
	Use:   "loganizer",
	Short: "LogAnalyzer - Outil d'analyse de logs distribuée",
	Long: `LogAnalyzer est un outil en ligne de commande pour analyser
des fichiers de logs de manière distribuée et parallèle.

Il permet de traiter plusieurs logs simultanément avec gestion
des erreurs et export des résultats au format JSON.`,
}

// Execute ajoute toutes les commandes enfant à la commande root
func Execute() error {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Erreur lors de l'exécution de '%s'\n", err)
		return err
	}
	return nil
}

func init() {
	rootCmd.Flags().BoolP("version", "v", false, "Afficher la version")
}
