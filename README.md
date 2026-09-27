# go-tweets 🐦

A robust Twitter/X backend REST API clone built with **Go (Golang)**, **Gin Framework**, and **MySQL**, adhering to **Clean Architecture** principles.

---

## 🚀 Tech Stack

- **Language:** Go (1.25+)
- **HTTP Framework:** [Gin Web Framework](https://github.com/gin-gonic/gin)
- **Database:** MySQL 8.0
- **Database Migrations:** [dbmate](https://github.com/amacneil/dbmate)
- **Validation:** [go-playground/validator/v10](https://github.com/go-playground/validator)
- **Security:** `golang.org/x/crypto/bcrypt`
- **Containerization:** Docker & Docker Compose

---

## 📁 Project Structure

Proyek ini menerapkan pemisahan layer tanggung jawab (*Separation of Concerns*):

```text
go-tweets/
├── cmd/
│   └── main.go                 # Application entry point & dependency wiring
├── db/
│   └── migrations/             # Plain SQL migration files (dbmate)
├── docs/
│   └── learn/                  # Learning logs & deep-dive notes
├── internal/
│   ├── config/                 # Environment & app configurations
│   ├── dto/                    # Request & Response Data Transfer Objects
│   ├── handler/                # HTTP Transport Layer (Gin Controllers)
│   ├── model/                  # Database Entities / Models
│   ├── repository/             # Data Access Layer (Plain SQL queries)
│   └── service/                # Business Logic Layer
├── pkg/
│   └── internalsql/            # Shared database connection pool
├── docker-compose.yml          # MySQL container service
├── go.mod
└── README.md
```

---

## 🛠️ Getting Started

### 1. Prerequisites
Pastikan kamu sudah menginstal:
- [Go](https://go.dev/dl/) (>= 1.25)
- [Docker Desktop](https://www.docker.com/products/docker-desktop/)
- [dbmate](https://github.com/amacneil/dbmate) *(opsional, untuk migrasi database)*

### 2. Clone & Setup Environment
Salin file `.env.example` menjadi `.env`:

```bash
cp .env.example .env
```

Pastikan konfigurasi database di file `.env` sudah sesuai:
```env
DATABASE_URL="mysql://root:password@127.0.0.1:3306/go_tweets"
PORT=8080
APP_ENV=development
SECRET_JWT=your_super_secret_jwt_key_here
```

### 3. Jalankan Database (Docker)
Nyalakan container MySQL:

```bash
docker compose up -d
```

### 4. Jalankan Migrasi Database
Jalankan migrasi menggunakan `dbmate`:

```bash
dbmate up
```

### 5. Jalankan Aplikasi
Jalankan server Go:

```bash
go run cmd/main.go
```

Server akan aktif di `http://127.0.0.1:8080`.

---

## 📡 API Endpoints

### Health Check
- **`GET /check`**  
  Response: `{"mesage": "App is running"}`

### Authentication
- **`POST /auth/register`**  
  Mendaftarkan pengguna baru dengan validasi format email, username, dan konfirmasi password.

  **Request Body:**
  ```json
  {
    "email": "user@example.com",
    "username": "user123",
    "password": "secretpassword",
    "password_confirm": "secretpassword"
  }
  ```

  **Response (201 Created):**
  ```json
  {
    "id": 1
  }
  ```

- **`POST /auth/login`**  
  Masuk ke aplikasi menggunakan email & password untuk mendapatkan Access Token dan Refresh Token.

  **Request Body:**
  ```json
  {
    "email": "user@example.com",
    "password": "secretpassword"
  }
  ```

  **Response (200 OK):**
  ```json
  {
    "access_token": "eyJhbGciOiJIUzI1Ni...",
    "refresh_token": "eyJhbGciOiJIUzI1Ni..."
  }
  ```

- **`POST /auth/refresh`** *(Protected via Refresh Token)*  
  Memperbarui Access Token dan Refresh Token baru (Token Rotation).

  **Headers:**
  ```http
  Authorization: Bearer <refresh_token>
  ```

  **Response (200 OK):**
  ```json
  {
    "access_token": "eyJhbGciOiJIUzI1Ni...",
    "refresh_token": "eyJhbGciOiJIUzI1Ni..."
  }
  ```

### Tweets / Posts

- **`POST /tweets/`** *(Protected via Access Token)*  
  Membuat postingan/tweet baru untuk pengguna yang sedang login.

  **Headers:**
  ```http
  Authorization: Bearer <access_token>
  ```

  **Request Body:**
  ```json
  {
    "title": "Halo Dunia",
    "content": "Ini adalah postingan pertama saya di go-tweets!"
  }
  ```

  **Response (201 Created):**
  ```json
  {
    "id": 1
  }
  ```

- **`PUT /tweets/:post_id/update`** *(Protected via Access Token & Ownership Check)*  
  Memperbarui judul dan isi tweet yang dimiliki pengguna.

  **Headers:**
  ```http
  Authorization: Bearer <access_token>
  ```

  **Request Body:**
  ```json
  {
    "title": "Judul Baru yang Diperbarui",
    "content": "Konten tweet yang sudah diedit."
  }
  ```

  **Response (200 OK):**
  ```json
  {
    "id": 1
  }
  ```

- **`DELETE /tweets/:post_id/delete`** *(Protected via Access Token & Ownership Check)*  
  Menghapus tweet secara lunak (*Soft Delete*) dengan mengisi kolom `deleted_at`.

  **Headers:**
  ```http
  Authorization: Bearer <access_token>
  ```

  **Response (200 OK):**
  ```json
  {
    "message": "Post deleted successfully"
  }
  ```

- **`POST /tweets/action`** *(Protected via Access Token)*  
  Menyukai (*Like*) atau membatalkan suka (*Unlike*) sebuah tweet secara otomatis (*Toggle*).

  **Headers:**
  ```http
  Authorization: Bearer <access_token>
  ```

  **Request Body:**
  ```json
  {
    "post_id": 1
  }
  ```

  **Response (200 OK):**
  ```json
  {
    "message": "succesfully liked or unliked post"
  }
  ```

### Comments

- **`POST /comment/`** *(Protected via Access Token)*  
  Membuat komentar/balasan pada sebuah tweet tertentu.

  **Headers:**
  ```http
  Authorization: Bearer <access_token>
  ```

  **Request Body:**
  ```json
  {
    "post_id": 1,
    "content": "Komentar pertama saya pada tweet ini!"
  }
  ```

  **Response (200 OK):**
  ```json
  {
    "message": "comment created successfully"
  }
  ```

---

## 📚 Catatan Belajar (Learning Docs)

Dokumentasi konsep arsitektur, perbedaan ekosistem (Go vs Laravel vs Next.js), dan cara membaca sintaks Go tersedia di folder `docs/learn/`:
- [01 - Setup Database, Migration, & Gin](docs/learn/penjelasan_01_setup_database_migration.md)
- [02 - Konvensi Pesan Commit Git](docs/learn/penjelasan_02_konvensi_pesan_commit_git.md)
- [03 - Clean Architecture, Fitur Register, & Cara Baca Kodingan Go](docs/learn/penjelasan_03_fitur_register_dan_arsitektur.md)
- [04 - Fitur Login Pengguna, Autentikasi JWT, dan Refresh Token](docs/learn/penjelasan_04_fitur_login_jwt_dan_refresh_token.md)
- [05 - Fitur Refresh Token, Middleware Autentikasi JWT, dan Token Rotation](docs/learn/penjelasan_05_fitur_refresh_token_dan_middleware_auth.md)
- [06 - Fitur Create Tweet / Post & Protected Route Middleware](docs/learn/penjelasan_06_fitur_create_tweet_dan_protected_route.md)
- [07 - Fitur Update Tweet / Postingan & Ownership Authorization](docs/learn/penjelasan_07_fitur_update_tweet_dan_authorization.md)
- [08 - Fitur Soft Delete Tweet & Data Persistence](docs/learn/penjelasan_08_fitur_soft_delete_tweet_dan_data_persistence.md)
- [09 - Fitur Like & Unlike Tweet (Toggle Action)](docs/learn/penjelasan_09_fitur_like_dan_unlike_tweet_toggle_action.md)
- [10 - Modul Komentar & Cross-Repository Dependency Injection](docs/learn/penjelasan_10_modul_komentar_dan_cross_repository_injection.md)
