# Helseveileder
> Bachelor-project ~ Spring 2024

## 📋 Table of Contents

1. 📝 [Info](#info)
2. ⚙️ [Tech Stack](#tech-stack)
3. 🔋 [Features](#features)
4. 🤸 [Quick Start](#quick-start)
5. 🎥 [Demo](#demo)
6. 🔗 [Links](#links)
7. 📄 [License](#license)

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

👉 **Bug Reporting:** Respondents can report a bug if they found one.

👉 **Evaluation of survey:** Respondents can evaluate survey if wanted.

## 🤸 Quick Start

### Server Setup

The server relies on `Docker Compose` for environment management. Ensure that it's installed on your system. 

- **Getting Started:**  
    i. Go into the server directory: `cd backend`  
    ii. Copy `.env.example` to `.env`  
    iii. Build Docker environment: `docker compose -f docker-compose.dev.yml build`  
    iv. Launch the Docker environment; `docker compose -f docker-compose.dev.yml up -d`  
- **API Access:** The API is now accessible at `http://localhost:8080`.
### Client Setup 

The client uses `npm` for package management. Ensure that it's installed before proceeding.

- **Getting Started:**  
    i. Navigate to the client directory: `cd frontend`  
    ii. Copy `.env.example` to `.env`  
    iii. Install dependencies: `npm i`  
    iv. Start dev-server: `npm dev run`  
- **Website Access:** The site is now available at `http://localhost:5173`.

## 🎥 Demo

https://github.com/haraldnilsen/helseveileder/assets/93219711/0b9590a5-736a-41d5-9888-51a1232dc591

## 🔗 Links

See also links below for related info about project:

- [Helseveileder-webscraper](https://github.com/SindreKjelsrud/helseveileder_webscraper)

## 📄 Licence

This project is licensed under the [MIT License](./LICENSE).
