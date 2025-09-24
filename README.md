# 🚀 LogAnalyzer - Outil d'Analyse de Logs Distribuée

![Go](https://img.shields.io/badge/go-%2300ADD8.svg?style=for-the-badge&logo=go&logoColor=white)
![CLI](https://img.shields.io/badge/CLI-Tool-brightgreen?style=for-the-badge)
![Concurrency](https://img.shields.io/badge/Concurrent-Processing-orange?style=for-the-badge)

**LogAnalyzer** est un outil en ligne de commande (CLI) développé en Go pour analyser des fichiers de logs de manière distribuée et parallèle. Il permet aux administrateurs système de centraliser l'analyse de multiples logs simultanément avec une gestion robuste des erreurs.

## 📋 Table des matières

- [Fonctionnalités](#-fonctionnalités)
- [Installation](#-installation)
- [Utilisation](#-utilisation)
- [Architecture](#-architecture)
- [Exemples](#-exemples)
- [Fonctionnalités Bonus](#-fonctionnalités-bonus)
- [Équipe de développement](#-équipe-de-développement)
- [Contribution](#-contribution)

## ✨ Fonctionnalités

### Fonctionnalités principales

- **🔄 Traitement concurrent** : Analyse plusieurs logs en parallèle grâce aux goroutines
- **⚡ Interface CLI intuitive** : Commandes simples avec Cobra CLI framework
- **📊 Export JSON** : Génération de rapports détaillés au format JSON
- **🎯 Gestion d'erreurs robuste** : Erreurs personnalisées avec `errors.Is` et `errors.As`
- **📁 Architecture modulaire** : Code organisé en packages logiques

### Fonctionnalités bonus implémentées

- **📅 Horodatage automatique** : Ajout automatique de la date aux fichiers de sortie
- **📂 Création de répertoires** : Création automatique des dossiers d'export s'ils n'existent pas
- **🔍 Filtrage par statut** : Filtrage des résultats par statut (OK/FAILED)
- **➕ Ajout de configurations** : Commande pour ajouter des logs à la configuration

## 🚀 Installation

### Prérequis

- Go 1.21 ou supérieur
- Git

### Étapes d'installation

1. **Cloner le repository**
   ```bash
   git clone https://github.com/votre-username/loganizer.git
   cd loganizer
   ```

2. **Initialiser le module Go**
   ```bash
   go mod tidy
   ```

3. **Compiler l'application**
   ```bash
   go build -o loganizer main.go
   ```

4. **Installer globalement (optionnel)**
   ```bash
   go install
   ```

## 📖 Utilisation

### Commande `analyze`

Analyse des logs selon un fichier de configuration JSON.

```bash
# Analyse basique
./loganizer analyze -c config.json

# Analyse avec export JSON
./loganizer analyze -c config.json -o rapport.json

# Analyse avec filtrage par statut
./loganizer analyze -c config.json -o rapport.json --status FAILED
```

#### Format du fichier de configuration

```json
[
  {
    "id": "web-server-1",
    "path": "/var/log/nginx/access.log",
    "type": "nginx-access"
  },
  {
    "id": "app-backend-2",
    "path": "/var/log/my_app/errors.log",
    "type": "custom-app"
  }
]
```

#### Format du rapport de sortie

```json
[
  {
    "log_id": "web-server-1",
    "file_path": "/var/log/nginx/access.log",
    "status": "OK",
    "message": "Analyse terminée avec succès.",
    "error_details": ""
  },
  {
    "log_id": "invalid-path",
    "file_path": "/non/existent/log.log",
    "status": "FAILED",
    "message": "Fichier introuvable.",
    "error_details": "fichier introuvable ou inaccessible: /non/existent/log.log"
  }
]
```

### Commande `add-log` (Bonus)

Ajouter une nouvelle configuration de log à un fichier existant.

```bash
./loganizer add-log --id web-server-3 --path /var/log/nginx/error.log --type nginx-error --file config.json
```

### Aide et documentation

```bash
# Aide générale
./loganizer --help

# Aide pour une commande spécifique
./loganizer analyze --help
./loganizer add-log --help
```

## 🏗️ Architecture

Le projet suit une architecture modulaire avec séparation des responsabilités :

```
loganizer/
├── cmd/                    # Commandes CLI
│   ├── cli.go             # Commande root
│   ├── analyze.go         # Commande analyze
│   └── add_log.go         # Commande add-log (bonus)
├── internal/              # Packages internes
│   ├── analyze.go         # Logique principale d'analyse
│   ├── config/           # Gestion des configurations
│   │   └── config.go
│   ├── analyzer/         # Erreurs personnalisées
│   │   └── errors.go
│   └── reporter/         # Export et rapport
│       └── reporter.go
├── main.go               # Point d'entrée
├── go.mod               # Module Go
├── test_config.json     # Configuration de test
└── README.md           # Documentation
```

### Packages détaillés

#### `internal/config`
- **Responsabilité** : Lecture et parsing des fichiers de configuration JSON
- **Fonctions principales** :
  - `LoadConfig()` : Charge la configuration depuis un fichier JSON

#### `internal/analyzer`
- **Responsabilité** : Définition des erreurs personnalisées
- **Types d'erreurs** :
  - `FileNotFoundError` : Fichier introuvable ou inaccessible
  - `ParseError` : Erreur de parsing de configuration

#### `internal/reporter`
- **Responsabilité** : Génération et export des rapports
- **Fonctions principales** :
  - `ExportReports()` : Export JSON avec création automatique des répertoires
  - `FilterReports()` : Filtrage des rapports par statut

#### `internal/analyze.go`
- **Responsabilité** : Logique principale d'analyse avec traitement concurrent
- **Fonctions principales** :
  - `AnalyzeLog()` : Analyse d'un log individuel
  - `AnalyzeLogs()` : Traitement concurrent avec goroutines et WaitGroup
  - `ProcessAnalysis()` : Orchestration complète du processus

## 💡 Exemples

### Exemple 1 : Analyse basique

```bash
# Créer un fichier de configuration
echo '[{"id": "test", "path": "main.go", "type": "source"}]' > test.json

# Exécuter l'analyse
./loganizer analyze -c test.json
```

**Sortie :**
```
Analyse de 1 log(s) en cours...

ID: test
Chemin: main.go
Statut: OK
Message: Analyse terminée avec succès.
---

Résumé: 1 succès, 0 échecs
```

### Exemple 2 : Analyse avec export et fichier inexistant

```bash
# Configuration mixte (fichier existant + inexistant)
./loganizer analyze -c test_config.json -o rapports/240924_analyse.json
```

**Sortie :**
```
Analyse de 3 log(s) en cours...

ID: sample-log
Chemin: test_logs/sample.log
Statut: OK
Message: Analyse terminée avec succès.
---
ID: main-go
Chemin: main.go
Statut: OK
Message: Analyse terminée avec succès.
---
ID: invalid-path
Chemin: /non/existent/log.log
Statut: FAILED
Message: Fichier introuvable.
Détails de l'erreur: fichier introuvable ou inaccessible: /non/existent/log.log
---

Résumé: 2 succès, 1 échecs

Rapport exporté vers: rapports/240924_analyse.json
```

### Exemple 3 : Ajout de configuration (Bonus)

```bash
./loganizer add-log --id database-1 --path /var/log/mysql/error.log --type mysql --file config.json
```

## 🎁 Fonctionnalités Bonus

### ✅ Implémentées

1. **Gestion des dossiers d'exportation** : Création automatique des répertoires de sortie
2. **Horodatage des exports JSON** : Format automatique AAMMJJ_nomfichier.json
3. **Commande `add-log`** : Ajout interactif de configurations
4. **Filtrage des résultats** : Option `--status` pour filtrer par statut

### 🔄 Traitement concurrent

L'application utilise les patterns de concurrence Go :

```go
// Utilisation de goroutines et WaitGroup
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
    
    // Collection des résultats via channel
    var reports []reporter.LogReport
    for report := range resultsChan {
        reports = append(reports, report)
    }
    
    return reports
}
```

### 🚨 Gestion d'erreurs personnalisées

```go
// Erreurs typées avec support d'unwrapping
type FileNotFoundError struct {
    FilePath string
    Err      error
}

// Utilisation avec errors.As et errors.Is
var parseErr *analyzer.ParseError
if errors.As(err, &parseErr) {
    return fmt.Errorf("erreur de parsing de la configuration: %w", err)
}
```

## 🧪 Tests

### Tests manuels

```bash
# Test avec configuration valide
./loganizer analyze -c test_config.json

# Test avec fichier inexistant
echo '[{"id": "missing", "path": "/missing.log", "type": "test"}]' > missing.json
./loganizer analyze -c missing.json

# Test d'ajout de configuration
./loganizer add-log --id test-new --path ./test.log --type test --file test_config.json
```

### Validation des fonctionnalités

- ✅ Traitement concurrent avec goroutines
- ✅ Gestion des erreurs personnalisées
- ✅ Export JSON avec structure correcte
- ✅ Interface CLI avec drapeaux
- ✅ Création automatique des répertoires
- ✅ Horodatage des fichiers de sortie
- ✅ Filtrage par statut
- ✅ Ajout de configurations

## 🤝 Contribution

- Nao MAUSSERVEY
- Younes ESSLIMANI

**LogAnalyzer v1.0** - Développé avec ❤️ en Go