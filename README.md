# BiTE — Bitcoin High School

**Bitcoin High School** ([bitcoinhighschool.com](https://bitcoinhighschool.com)) is an online STEM education platform that incentivises learning with real Bitcoin rewards. Students sign up, attempt timed quizzes across categories like Game Theory, Computer Science, Cryptography, AI, and Robotics, and earn Satoshi (sats) based on their scores. Rewards are paid out instantly via the Bitcoin Lightning Network to a linked [Blink](https://www.blink.sv/) wallet.

<!-- ### Key Features

- **STEM Quizzes** — Timed, category-based quizzes with multiple-choice questions covering mathematics, AI, robotics, cryptography, and more.
- **AI-Powered Previews** — Gemini-generated explanations for each question to help students understand correct answers.
- **Bitcoin Rewards** — A learn-to-earn model where quiz performance is rewarded with Satoshi, delivered instantly via the Lightning Network.
- **Chess Puzzles** — Interactive chess puzzles with a Glicko-2 rating system for competitive skill tracking.
- **Category Certificates** — Downloadable PDF certificates awarded upon completing all quizzes in a category.
- **Leaderboard & Rankings** — Global student rankings based on quiz performance.
- **Admin Dashboard** — Full quiz, question, and user management for administrators. -->

---

## Prerequisites

| Tool           | Version | Installation                                                                                                                                     |
| -------------- | ------- | ------------------------------------------------------------------------------------------------------------------------------------------------ |
| **Node.js**    | ≥ 20    | [nodejs.org](https://nodejs.org/)                                                                                                                |
| **pnpm**       | ≥ 8     | [pnpm.io/installation](https://pnpm.io/installation)                                                                                             |
| **Go**         | 1.25    | [go.dev/doc/install](https://go.dev/doc/install)                                                                                                 |
| **PostgreSQL** | ≥ 17    | [postgresql.org/download](https://www.postgresql.org/download/)                                                                                  |
| **Make**       | —       | Pre-installed on macOS/Linux. Windows: [Run Makefile on Windows](https://medium.com/@samsorrahman/how-to-run-a-makefile-in-windows-b4d115d7c516) |

---

## 1 — Clone the Repository

```sh
git clone https://github.com/Tibz-Dankan/BiTE.git
cd BiTE
```

---

## 2 — Client (Frontend)

The client is a **React 19 + TypeScript** app built with **Vite** and **TailwindCSS 4**.

### 2.1 Install dependencies

```sh
cd client
pnpm install
```

### 2.2 Set up environment variables

Create a `.env` file inside the `client/` directory:

```sh
# client/.env

VITE_PUBLIC_POSTHOG_KEY=<your-posthog-project-api-key>
VITE_PUBLIC_POSTHOG_HOST=https://us.i.posthog.com
```

| Variable                   | Type   | Description                |
| -------------------------- | ------ | -------------------------- |
| `VITE_PUBLIC_POSTHOG_KEY`  | string | PostHog project API key    |
| `VITE_PUBLIC_POSTHOG_HOST` | string | PostHog ingestion host URL |

### 2.3 Start the development server

```sh
pnpm dev
```

The client will be available at **http://localhost:5173**.

### 2.4 Build for production (optional)

```sh
pnpm build
pnpm preview
```

---

## 3 — Server (Backend)

The server is a **Go 1.25** application using **Fiber**, **GORM**, and **PostgreSQL**.

### 3.1 Install dependencies

```sh
cd server
make install
```

### 3.2 Set up environment variables

Create a `.env` file inside the `server/` directory:

```sh
# server/.env

# Database
BiTE_DEV_DSN="host=localhost user=postgres password=<db_password> dbname=<db_name> port=5432 sslmode=disable"
BiTE_PROD_DSN="<production_dsn>"

# JWT
JWT_SECRET="<your_jwt_secret>"

# Admin
ADMIN_EMAIL=<admin_email>

# Mailjet
MJ_SENDER_MAIL=<mailjet_sender_email>
MJ_APIKEY_PUBLIC=<mailjet_public_key>
MJ_APIKEY_PRIVATE=<mailjet_private_key>

# AWS S3
S3_ACCESS_KEY=<aws_s3_secret_access_key>
S3_ACCESS_KEY_ID=<aws_s3_access_key_id>
S3_BUCKET_NAME=<s3_bucket_name>
AWS_REGION=<aws_region>

# Blink (Bitcoin Lightning)
BLINK_API_URL=https://api.blink.sv/graphql
BLINK_API_KEY=<blink_api_key>
BLINK_WALLET_ID=<blink_wallet_id>

# Resend
RESEND_API_KEY=<resend_api_key>
RESEND_SENDER_EMAIL=<resend_sender_email>

# PostHog
POSTHOG_API_KEY=<posthog_api_key>
POSTHOG_HOST=https://us.i.posthog.com
POSTHOG_BITE_PERSONAL_KEY=<posthog_personal_key>
```

| Variable                    | Type   | Description                                |
| --------------------------- | ------ | ------------------------------------------ |
| `BiTE_DEV_DSN`              | string | DSN for the local/dev PostgreSQL database  |
| `BiTE_PROD_DSN`             | string | DSN for the production PostgreSQL database |
| `JWT_SECRET`                | string | Secret key used to sign JWT tokens         |
| `ADMIN_EMAIL`               | string | Comma-separated admin email addresses      |
| `MJ_SENDER_MAIL`            | string | Mailjet sender email address               |
| `MJ_APIKEY_PUBLIC`          | string | Mailjet public API key                     |
| `MJ_APIKEY_PRIVATE`         | string | Mailjet private API key                    |
| `S3_ACCESS_KEY`             | string | AWS S3 secret access key                   |
| `S3_ACCESS_KEY_ID`          | string | AWS S3 access key ID                       |
| `S3_BUCKET_NAME`            | string | AWS S3 bucket name                         |
| `AWS_REGION`                | string | AWS region (e.g. `eu-central-1`)           |
| `BLINK_API_URL`             | string | Blink Bitcoin Lightning API URL            |
| `BLINK_API_KEY`             | string | Blink API key                              |
| `BLINK_WALLET_ID`           | string | Blink wallet ID                            |
| `RESEND_API_KEY`            | string | Resend email API key                       |
| `RESEND_SENDER_EMAIL`       | string | Resend sender email address                |
| `POSTHOG_API_KEY`           | string | PostHog project API key                    |
| `POSTHOG_HOST`              | string | PostHog ingestion host URL                 |
| `POSTHOG_BITE_PERSONAL_KEY` | string | PostHog personal API key                   |

### 3.3 Set up the database

Ensure PostgreSQL is running, then create a database matching the name in your `BiTE_DEV_DSN`:

```sh
sudo -u postgres psql
```

```sql
CREATE DATABASE bite_db;
```

> **Note:** The server auto-migrates tables on startup via GORM, so no manual migrations are needed.

### 3.4 Start the development server

```sh
make run
```

The server will be available at **http://localhost:5000**.

---

## Project Structure

```
BiTE/
├── client/          # React + Vite + TypeScript frontend
├── server/          # Go (Fiber + GORM) backend
├── scripts/         # Utility scripts
├── Dockerfile       # Production Docker image (server)
├── docker-compose.yaml
└── README.md
```

---

## Useful Commands

| Command        | Location  | Description                      |
| -------------- | --------- | -------------------------------- |
| `pnpm dev`     | `client/` | Start client dev server          |
| `pnpm build`   | `client/` | Build client for production      |
| `pnpm preview` | `client/` | Preview production build         |
| `make install` | `server/` | Install Go dependencies          |
| `make run`     | `server/` | Start server in development mode |
