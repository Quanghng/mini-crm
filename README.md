⚔️ Voici un **README professionnel** pour ton projet Go Mini CRM :

---

# Mini CRM (Go CLI)

## 📌 Description

Mini CRM est une petite application **en ligne de commande** écrite en Go.
Elle permet de **gérer une liste de contacts** avec des opérations simples :

* Ajouter un contact
* Lister tous les contacts
* Supprimer un contact
* Mettre à jour un contact

Deux modes d’utilisation :

1. **Mode interactif (menu)** → Lancement classique du programme, l’utilisateur navigue dans un menu CLI.
2. **Mode flags (exécution directe)** → Utilisation d’options en ligne de commande (`--addContact`, `--nom`, `--email`) pour automatiser certaines actions.

---

## 📂 Structure du projet

```
TP/
│── go.mod
├── cmd/  
|   └── main.go             # Point d'entrée : menu CLI + gestion des flags
│
└── internal/
    ├── domain/
    │   └── contact.go      # Définition du modèle Contact
    │
    └── handler/
        └── contact_handler.go      # Logique métier (CRUD sur les contacts)
```

* **`internal/domain/contact.go`** : Définit la structure de données `Contact`.
* **`internal/handler/contact_handler.go`** : Contient la logique pour ajouter, supprimer, lister et mettre à jour les contacts. Les contacts sont stockés en mémoire dans une `map[int]Contact`.
* **`main.go`** : Interface utilisateur. Gère :

  * le menu CLI interactif
  * l’exécution via flags (`flag` package)

---

## ⚙️ Prérequis

* **Go 1.21+** (testé sur Go 1.22)
* OS supportés : Linux, macOS, Windows

Vérifie que Go est installé :

```bash
go version
```

---

## 🚀 Installation & Exécution

### 1. Cloner le projet

```bash
git clone https://github.com/Quanghng/mini-crm
cd Mini CRM
```

### 2. Initialiser les dépendances

```bash
go mod tidy
```

### 3. Lancer en mode interactif (menu)

```bash
go run .
```

Exemple :

```
----- Mini CRM -----
1. Ajouter un contact
2. Lister tous les contacts
3. Supprimer un contact
4. Mettre à jour un contact
5. Quitter l'application
Sélectionnez votre option:
```

### 4. Lancer en mode flags (ajout rapide d’un contact)

```bash
go run . --addContact --nom="Alice" --email="alice@mail.com"
```

---

## 🔮 Améliorations futures possibles

* Persistance des données (sauvegarde des contacts dans un fichier JSON/SQLite).
* Ajout d’un flag `--listContacts` pour lister les contacts sans menu.
* Export/Import des contacts.
* Ajout de tests unitaires.



