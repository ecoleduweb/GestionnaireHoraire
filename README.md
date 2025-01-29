# API LLIO

Une API développée en Go avec le framework Gin-Gonic.

---

## Structure du projet

```plaintext
llio-api/
├── main.go        # Point d'entrée de l'application
├── controllers/   # Logique des routes
├── models/        # Modèles de données
├── routes/        # Définition des routes
└── services/      # Logique métier
```
# Préalable
1. Installation de la derniere version stable de go sur golang.org (go version go1.23.5 windows/amd64)
   
2. Instalation des modules GO
go mod init llio-api

3. Installation des modules gin-gonic
go get -u github.com/gin-gonic/gin

# Démarrage du serveur
go run main.go