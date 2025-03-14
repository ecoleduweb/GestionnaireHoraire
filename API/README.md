# API LLIO

Une API développée en Go avec le framework Gin-Gonic.

---

## Structure du projet

```plaintext
API/
├── controllers/   # Logique des routes
├── database/      # Connexion et configuration de la base de données
├── main.go        # Point d'entrée de l'application
├── middleware/    # Middlewares
├── models/        # Modèles de données
│   ├── DAOs/      # Objets d'accès aux données
│   └── DTOs/      # Objets de transfert de données
├── repositories/  # Gestion des accès aux données
├── routes/        # Définition des routes
├── services/      # Logique métier
├── tests/         # Tests unitaires et d'intégration
└── useful/        # Fonctions utilitaires
```
# Golang
1. Installation de la derniere version stable de go sur golang.org
    (go version go1.23.5 windows/amd64)
   
2. Instalation des modules GO
    go mod init llio-api

# Gin
- Installation des modules gin-gonic
   go get -u github.com/gin-gonic/gin

# Base de données
1. Installation des modules de l'ORM pour l'utilisation d'une BD MariaDB
   go get -u gorm.io/gorm
   go get -u gorm.io/driver/mysql

2. Créer une base de données SQL

# Variables d'environnement
1. Installation du module qui nous permet d'utiliser les variables d'environnement
   go get github.com/joho/godotenv

2. Structure
   Voici la structure de fichier .env pour que l'API puisse fonctionner:
   ```plaintext
    DB_USER=user
    DB_PASSWORD=password
    DB_HOST=127.0.0.1
    DB_PORT=3306
    DB_NAME_RUN=nom_bd
    DB_NAME_TEST=nom_bd_test
    ENV=DEV ou TEST ou PROD
    AZUREAD_CLIENT_ID=client_id_azure
    AZUREAD_CLIENT_SECRET=client_secret_azuread
    AZUREAD_CALLBACK_URL=http://localhost:8080/auth/azureadv2/callback
    SESSION_SECRET=secret
    SESSION_MAX_AGE=86400
    PORT = port_de_l'application
    FRONT_ADDRESS = http://localhost:xxxx
   ```

# Migrations et mise en place de la base de données
1. Créer une base de données vide nommée llio et une autre nommée llio_test
2. À la racine du projet API, exécuter la commande `go run main.go migrate up`. Cette commande permet d'exécuter les nouvelles migrataions en se basant sur la table schema_migrations.

## Pour créer une nouvelle migration
1. Exécuter la commande  `go run main.go migrate create un_nom_significatif_pour_ta_migration`. Cela va ajouter deux fichiers dans le dossier database/migrations.
2. Complète les fichiers de migration. L'ia peut t'aider à le faire, mais assures-toi de bien valider ce qui est produit! Le ficher up applique le changement, le fichier down revient en arrière.

Test ensuite ta migration. Si elle échoue, la table schema_migration aura ton numéro de migration et la valeur 1 à isDirty qui montre que ta migration a échouée.

## Pour annuler une migration
La commande  `go run main.go migrate down` annule la dernière migration qui a été exécutée.

## Si jamais ma migration a tout cassé?
Le champ isDirty dans schema_migration sera à 1. Le changer pour 0 dans la base de données et re rouler la migration!

# Test
## Installation de la librairie de test pour du golang : testify
Installation de la librairie de base:
go get -u github.com/stretchr/testify

Pour l'utilisation graphique des tests avec l'extension Testing, il faut déactiver le cache des tests antérieurs (mesure de précaution)
Étapes :
1. Ouvrez les Paramètres de VS Code (via Ctrl+, ou Cmd+, sur macOS).

2. Recherchez go.testFlags.

3. Ajoutez -count=1 à la liste des flags de test:
`"go.testFlags": ["-count=1"]`


Pour excécution des tests en ligne de commande:
`go test -v -count=1 ./...`

# Démarrage du serveur
go run main.go
