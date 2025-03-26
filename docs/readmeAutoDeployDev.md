# Workflow de Déploiement LLIO - Documentation

Ce document explique le fonctionnement du workflow GitHub Actions qui déploie automatiquement l'application GestionnaireHoraire.

## Structure du workflow

```yaml
name: Deployer LLIO
on:
  push:
    branches:
      - main
```

- Nom du workflow et déclencheur (s'active sur push vers main)

```yaml
jobs:
  deploy:
    runs-on: ubuntu-latest
    steps:
      - name: Deploy Full Application
        uses: appleboy/ssh-action@v0.1.2
        with:
          host: ${{secrets.DEV_SSH_HOST}}
          key: ${{secrets.DEV_SSH_KEY}}
          username: ${{ secrets.SSH_USERNAME}}
          debug: true
```

- Un job unique exécuté sur un runner Ubuntu
- Utilise l'action SSH pour se connecter au serveur distant (environnement de dev)
- Paramètres de connexion stockés dans les secrets GitHub

## Explication du script de déploiement

### Configuration initiale

```bash
set -e # Exit immediately if a command exits with a non-zero status.
```

- Active le mode strict: arrête le script si une commande échoue

### Arrêt des services

```bash
sudo systemctl stop llio-api.service || echo "API service not running"
sudo systemctl stop llio-frontend.service || echo "Frontend service not running"
```

- Arrête les services existants
- Le || est utile pour que le script continue même si les services ne sont pas actifs

### Rafraîchissement du code source

```bash
cd /root/ProjetLLIO2025
rm -rf GestionnaireHoraire
git clone https://github.com/ecoleduweb/GestionnaireHoraire.git
```

- Supprime l'ancienne version du code
- Clone la dernière version depuis GitHub

### Configuration et build de l'API

```bash
cd /root/ProjetLLIO2025/GestionnaireHoraire/API
cat > .env << EOF
DB_USER=${{ secrets.DB_USER }}
DB_PASSWORD=${{ secrets.DB_PASSWORD }}
# ... autres variables d'environnement
EOF

/usr/local/go/bin/go mod tidy
```

- Crée le fichier de configuration .env avec les secrets GitHub
- Télécharge et organise les dépendances Go

### Configuration et build du frontend

```bash
cd /root/ProjetLLIO2025/GestionnaireHoraire/LLIO2025
cat > .env << EOF
VITE_API_BASE_URL=${{ secrets.VITE_API_BASE_URL }}
EOF

npm install
npm run build
```

- Crée le fichier .env pour le frontend
- Installe les dépendances et compile l'application

### Démarrage et vérification

```bash
sudo systemctl start llio-api || echo "Failed to start API service"
sudo systemctl start llio-frontend || echo "Failed to start Frontend service"

sudo systemctl status llio-api --no-pager
sudo systemctl status llio-frontend --no-pager
```

- Démarre les services
- Affiche leur statut pour vérification si tout fonctionne correctement

## Points techniques importants

1. **Mode strict (`set -e`)**: Arrête l'exécution en cas d'erreur critique.

2. **Opérateur OR (`||`)**: Utilisé pour la gestion d'erreur, permet de continuer le script même si une commande échoue.

3. **Here documents**: La syntaxe `cat > file << EOF...EOF` crée des fichiers multi-lignes pour les configurations.

4. **Stratégie de déploiement**: Suppression complète et clone frais plutôt que mise à jour, pour éviter les conflits.

5. **Variables d'environnement**: Toutes les valeurs sensibles sont injectées depuis les secrets GitHub.

## Prérequis

1. **Secrets GitHub**: Configurer toutes les variables référencées dans le script. À noter que si vous ajoutez des variables dans le .env, ne pas oublier d'ajouter leurs contenus dans les secrets GitHub, sinon rien ne sera ajouté au .env

2. **Configuration du serveur**:
   - Services /etc/systemd/system `llio-api.service` et `llio-frontend.service` configurés
   - Permissions sudo pour l'utilisateur SSH
   - Go, Node.js, npm et Git installés

## Résolution des problèmes

1. **Services ne démarrant pas**:

   - Vérifier les logs: `journalctl -u llio-api` ou `journalctl -u llio-frontend`
   - Confirmer que les fichiers .env sont correctement générés

2. **Échec lors du clonage**:

   - Vérifier l'accès au repository et les permissions
   - Vérifier que la connection SSH s'effectue correctement et sans erreur

3. **Échec d'installation des dépendances**:

   - Vérifier les versions de Node/npm sur le serveur
   - Regarder si le serveur a assez de stockage ET DE MÉMOIRE

4. **Problèmes avec les variables d'environnement**:
   - Vérifier que tous les secrets sont configurés dans GitHub
   - S'assurer que les noms des variables correspondent à ceux attendus dans les secrets ou dans le .env

## Maintenance

Pour modifier ce workflow:

1. Éditer le fichier `.github/workflows/deploy.yml` dans le repository
2. Tester les modifications sur une branche séparée avant de les fusionner dans main
3. Si vous ajoutez de nouvelles variables d'environnement, n'oubliez pas de les configurer dans les secrets GitHub
