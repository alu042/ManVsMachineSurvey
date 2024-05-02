# Helseveileder
> Bachelor-project ~ Spring 2024

## 📋 Table of Contents

1. 📝 [Info](#info)
2. ⚙️ [Tech Stack](#tech-stack)
3. 🔋 [Features](#features)
4. 🤸 [Quick Start](#quick-start)
5. 🔗 [Links](#links)
6. 📄 [License](#license)

## 📝 Info

This repository contains our Bachelor project for Spring 2024. This project is a collaborative effort with medical students and researchers from UiB (University of Bergen). 

As a part of this project we needed to scrape questions and answers from [Studenterspør.no](https://studenterspor.no/). This involved developing a webscraper to extract the necessary data. You can explore the source code [here](https://github.com/SindreKjelsrud/helseveileder_webscraper).

## ⚙️ Tech Stack

- Go
- Svelte
- TypeScript
- Tailwind CSS
- PostgreSQL

## 🔋 Features

👉 **Accessibility:** Ensure access & usability for all users.

👉 **Bug Reporting:** Simplify reporting for prompt issue resolution.

👉 **Evaluation of Three Custom Criteria:** Enable customized item evaluation for informed decision-making.

## 🤸 Quick Start

Follow these steps to set up the project locally on your machine.

**Prerequisites**

Make sure you have the following installed on your machine:

- Docker
- Go

**Cloning the repository**

```bash
git clone git@github.com:haraldnilsen/helseveileder.git 
cd helseveileder
```

**Installation**

```bash
cd frontend && npm i
```

**Running the project locally**

```bash
cd backend
docker compose -f docker-compose.dev.yml up -d
```

**Deploying the project**

Change .env info in both `frontend/` & `/backend`, then:

```bash
cd backend
docker compose -f docker-compose.deploy.yml up -d
```

## 🔗 Links

See also links below for related info about project:

- [Helseveileder-webscraper](https://github.com/SindreKjelsrud/helseveileder_webscraper)

## 📄 Licence

This project is licensed under the [MIT License](./LICENSE).
