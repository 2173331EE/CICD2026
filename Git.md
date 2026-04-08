
# 🌱 Introduction à Git : branches locales, branches GitHub et échanges de modifications

Git est un outil qui permet de suivre l’évolution d’un projet et de travailler à plusieurs sans écraser le travail des autres.  
Dans ce laboratoire CI/CD, tu vas utiliser Git pour envoyer ton code vers GitHub, ce qui déclenchera automatiquement un pipeline de déploiement.

Ce document t’explique simplement :

- la différence entre les **branches locales** et les **branches GitHub**,
- comment les modifications circulent entre ton ordinateur et GitHub,
- les commandes essentielles pour travailler correctement.

---

# Étape de préconfif config

git config --global user.name "Danny Brabant"
<br>
git config --global user.email "danny.brabant@cegepmontpetit.ca"

# 🧩 1. Branches locales (sur ton ordinateur)

Une **branche locale** est une version du projet qui existe **uniquement sur ta machine**.

Quand tu modifies des fichiers dans VS Code, tu modifies **ta copie locale** du projet.

Pour voir tes branches locales :

```bash
git branch
```

Exemples de branches locales dans ce labo :

- `main` (qui est la branche prod)
- `dev`
- `qa`

👉 **Tu dois toujours vérifier sur quelle branche tu te trouves avant de modifier des fichiers.**

---

# ☁️ 2. Branches distantes (sur GitHub)

GitHub possède aussi ses propres branches, appelées **branches distantes**.

Pour les voir :

```bash
git branch -r
```

Exemples :

- `origin/main`  (qui est la branche prod)
- `origin/dev`
- `origin/qa`

👉 Ce sont ces branches qui déclenchent le pipeline CI/CD dans ton labo.

---

# 🔄 3. Comment les modifications circulent entre local et GitHub

Voici le cycle complet :

```
(1) Tu modifies des fichiers dans VS Code
(2) git add .
(3) git commit -m "message"
(4) git push origin dev
```

Et GitHub reçoit tes changements.

### 🧠 Résumé simple :

| Action | Où ? | Commande |
|--------|------|----------|
| Modifier des fichiers | Local | (VS Code) |
| Enregistrer un snapshot | Local | `git commit` |
| Envoyer sur GitHub | Distant | `git push` |
| Récupérer depuis GitHub | Local | `git pull` |

---

# 🌿 4. Pourquoi utiliser plusieurs branches ?

Dans ce laboratoire CI/CD :

| Branche Git | Namespace k3s | Environnement |
|-------------|---------------|---------------|
| `dev` | `dev` | Développement |
| `qa` | `qa` | Qualité |
| `main` | `prod` | Production |

Chaque branche correspond à un environnement différent dans Kubernetes.

👉 **C’est pour ça que tu dois te positionner sur la bonne branche avant de modifier ton code.**

Exemple :

```bash
git checkout dev
```

---

# 🚀 5. Exemple concret : modifier le message pour DEV

1. Tu vas sur la branche dev :

```bash
git checkout dev
```

2. Tu modifies `server.js`

3. Tu enregistres tes modifications :

```bash
git add .
git commit -m "Changer le message pour DEV"
```

4. Tu envoies sur GitHub :

```bash
git push origin dev
```

👉 GitHub Actions démarre automatiquement  
👉 Docker Hub reçoit une nouvelle image  
👉 k3s déploie dans le namespace `dev`

---

# 🎓 6. Métaphore simple pour comprendre

Imagine :

- **Ton ordinateur = ton cahier personnel**
- **GitHub = le cahier partagé de la classe**
- **Une branche = une version différente du projet**

Quand tu fais :

- `git commit` → tu écris dans ton cahier
- `git push` → tu recopies dans le cahier de la classe
- `git pull` → tu recopies ce que les autres ont ajouté

---

# 🧠 7. Commandes Git essentielles pour ce labo

Voici les commandes que tu vas utiliser le plus souvent :

### Voir les branches
```bash
git branch
```

### Changer de branche
```bash
git checkout dev
```

### Ajouter les fichiers modifiés
```bash
git add .
```

### Faire un commit
```bash
git commit -m "Message clair"
```

### Envoyer sur GitHub
```bash
git push origin dev
```

### Récupérer les changements depuis GitHub
```bash
git pull origin dev
```

### Déclencher un pipeline sans changer de fichier
```bash
git commit --allow-empty -m "Trigger pipeline"
git push origin dev
```

---

# 🧾 Tableau de résumé – Branches Git & échanges local ↔ GitHub

| Concept | Description | Exemple / Commande |
|--------|-------------|--------------------|
| **Branche locale** | Version du projet qui existe sur ton ordinateur. | `git branch` |
| **Branche distante** | Version du projet stockée sur GitHub. | `git branch -r` |
| **Changer de branche** | Permet de travailler dans un environnement spécifique. | `git checkout dev` |
| **Créer une branche** | Crée une nouvelle version du projet. | `git checkout -b qa` |
| **Ajouter des fichiers** | Prépare les fichiers modifiés pour un commit. | `git add .` |
| **Créer un commit** | Enregistre un ensemble de modifications localement. | `git commit -m "Message"` |
| **Envoyer vers GitHub** | Envoie les commits locaux vers la branche distante. | `git push origin dev` |
| **Récupérer depuis GitHub** | Met à jour ta branche locale avec les changements distants. | `git pull origin dev` |
| **Déclencher un pipeline** | Envoie un commit vide pour lancer GitHub Actions. | `git commit --allow-empty -m "Trigger"` |
| **Vérifier l’état du projet** | Montre les fichiers modifiés, ajoutés ou non suivis. | `git status` |
| **Voir l’historique** | Liste les commits de la branche actuelle. | `git log --oneline` |

---

# 🧠 Conclusion

Pour réussir ce laboratoire CI/CD, retiens ceci :

- Tu travailles **localement** sur ton ordinateur.
- Tu envoies tes modifications sur **GitHub** avec `git push`.
- GitHub déclenche automatiquement le pipeline CI/CD.
- Chaque branche correspond à un environnement Kubernetes différent.
- Toujours vérifier ta branche avant de modifier des fichiers.

