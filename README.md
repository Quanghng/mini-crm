# Mini CRM

Mini CRM est une application CLI (ligne de commande) simple pour gérer des contacts. Elle permet d’ajouter, lister, mettre à jour et supprimer des contacts. L’application supporte trois types de stockage :

* **GORM/SQLite** : persistance dans une base SQLite.
* **JSON** : persistance dans un fichier JSON.
* **Memory** : stockage éphémère en RAM, utile pour les tests.

---

## Fonctionnalités

* **CRUD complet** : Ajouter, Lister, Mettre à jour, Supprimer des contacts.
* **Interface CLI** : Commandes claires et standardisées via [Cobra](https://github.com/spf13/cobra).
* **Stockage configurable** : Changez le backend (GORM, JSON, Memory) via `config.yaml` sans recompiler.
* **Persistance des données** : Support SQLite et JSON pour conserver les données entre les exécutions.

---

## Structure du projet

```
mini-crm/
├─ cmd/                   # Commandes Cobra
│  ├─ root.go
│  ├─ add.go
│  ├─ list.go
│  ├─ update.go
│  └─ delete.go
├─ internal/
│  ├─ app/                # Interface CLI interactive
│  │  └─ app.go
│  ├─ config/             # Gestion de la configuration
│  │  └─ config.go
│  ├─ database/           # Connexion SQLite/GORM
│  │  └─ database.go
│  ├─ models/             # Structures de données
│  │  └─ contact.go
│  ├─ repository/         # GORM et JSON stores
│  │  └─ contact_repository.go
│  ├─ storage/            # Interface Storer + MemoryStore
|  |  ├─ json_store.go
│  |  ├─ storage.go
│  |  └─ memory.go
|  └─ utils/ 
|     └─ utils.go           
├─ data/                  # Dossier pour DB ou JSON
│  └─ contacts.db / contacts.json
├─ config.yaml            # Configuration externe
└─ go.mod
```

---

## Installation

### Prérequis

* Go >= 1.21 installé ([Télécharger Go](https://go.dev/dl/))
* Git installé ([Télécharger Git](https://git-scm.com/downloads))

### Cloner le projet

```bash
git clone https://github.com/ton-utilisateur/mini-crm.git
cd mini-crm
```

---

## Configuration

Le fichier `config.yaml` permet de choisir le type de stockage et le chemin des fichiers :

```yaml
storage:
  type: gorm   # valeurs possibles: gorm | json | memory

database:
  gorm:
    name: contacts.db
    dsn: ./data/contacts.db
  json:
    name: contacts.json
    dsn: ./data/contacts.json

app:
  environment: development
```

* `storage.type` : sélection du backend.
* `database.gorm.dsn` : chemin du fichier SQLite.
* `database.json.dsn` : chemin du fichier JSON.
* `memory` : ne nécessite pas de fichier, données perdues à la fermeture.

---

## Build & Run

### Windows

1. Ouvrir `cmd` ou PowerShell.
2. Compiler :

```powershell
go build -o minicrm.exe ./cmd
```

3. Exécuter :

```powershell
.\minicrm.exe
```

### Mac / Linux

1. Ouvrir le terminal.
2. Compiler :

```bash
go build -o minicrm ./cmd
```

3. Exécuter :

```bash
./minicrm
```

---

## CLI

### Commande interactive

Si aucun flag n’est passé, le programme lance le menu interactif :

```
Welcome to Mini CRM!
--- Main Menu ---
1. Add a contact
2. List contacts
3. Update a contact
4. Delete a contact
5. Exit
```

### Commandes Cobra

* **Ajouter un contact**

```bash
./minicrm add --name "John Doe" --email "john@example.com"
```

* **Lister les contacts**

```bash
./minicrm list
```

* **Mettre à jour un contact**

```bash
./minicrm update --id 1 --name "John Smith"
```

* **Supprimer un contact**

```bash
./minicrm delete --id 1
```

---

## Tests / Mémoire

Pour tester rapidement sans persistance, configurez :

```yaml
storage:
  type: memory
```

Toutes les modifications sont perdues à la fermeture de l’application.

---

## Notes

* Le backend **JSON** crée ou met à jour `contacts.json` dans le dossier `data/`.
* Le backend **GORM** crée ou met à jour `contacts.db` SQLite.
* Tous les changements de backend se font via `config.yaml`, aucune recompilation nécessaire.
