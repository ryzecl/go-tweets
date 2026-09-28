# go-tweets 🐦

A production-ready Twitter/X backend REST API clone built with **Go (Golang)**, **Gin Framework**, and **MySQL**, adhering to **Clean Architecture** principles.

---

## 🚀 Tech Stack & Highlights

- **Language:** Go (1.25+)
- **HTTP Framework:** [Gin Web Framework](https://github.com/gin-gonic/gin)
- **Database:** MySQL 8.0 (Containerized via Docker Compose)
- **Database Migrations:** [dbmate](https://github.com/amacneil/dbmate)
- **Input Validation:** [go-playground/validator/v10](https://github.com/go-playground/validator)
- **Authentication:** JWT (JSON Web Tokens) with Token Rotation (`golang-jwt/jwt/v5`)
- **Password Hashing:** `golang.org/x/crypto/bcrypt`
- **Architecture Pattern:** Clean Architecture (Transport, DTO, Service, Repository, Entity Model)
- **Query Optimization:** Two-Step Eager Loading with In-Memory Hash Map to eliminate the N+1 Query Problem

---

## 🏛️ Arsitektur & Struktur Folder

Proyek ini mematuhi prinsip *Separation of Concerns* (Pemisahan Tanggung Jawab) dengan membagi aplikasi ke dalam layer-layer yang decoupled dan mudah di-test:

```text
go-tweets/
├── cmd/
│   └── main.go                 # Entry point aplikasi & Dependency Injection Wiring
├── db/
│   └── migrations/             # Plain SQL migration files (dbmate)
├── docs/
│   └── learn/                  # Dokumentasi pembelajaran & bedah arsitektur
├── internal/
│   ├── config/                 # Konfigurasi aplikasi & parsing .env
│   ├── dto/                    # Request & Response Data Transfer Objects
│   ├── handler/                # HTTP Transport Layer (Gin Controllers)
│   │   ├── comment/            # Handler modul komentar
│   │   ├── post/               # Handler modul tweet/post
│   │   └── user/               # Handler modul autentikasi & user
│   ├── middleware/             # Gin Middlewares (Auth JWT, CORS, dll.)
│   ├── model/                  # Entitas database & plain structs
│   ├── repository/             # Data Access Layer (Plain SQL Query)
│   │   ├── comment/
│   │   ├── post/
│   │   └── user/
│   └── service/                # Business Logic Layer (Use Cases)
│       ├── comment/
│       ├── post/
│       └── user/
├── pkg/
│   └── internalsql/            # Shared MySQL connection pool
├── docker-compose.yml          # Container MySQL 8.0
├── go.mod
└── README.md
```

---

## 📊 Database Schema (ERD)

```mermaid
erDiagram
    users ||--o{ posts : "creates"
    users ||--o{ post_likes : "likes"
    users ||--o{ comments : "writes"
    users ||--o{ comment_likes : "likes"
    posts ||--o{ post_likes : "receives"
    posts ||--o{ comments : "has"
    comments ||--o{ comment_likes : "receives"

    users {
        bigint id PK
        varchar email UK
        varchar username UK
        varchar password
        datetime created_at
        datetime updated_at
    }

    posts {
        bigint id PK
        bigint user_id FK
        varchar title
        text content
        datetime created_at
        datetime updated_at
        datetime deleted_at "Soft Delete"
    }

    post_likes {
        bigint id PK
        bigint post_id FK
        bigint user_id FK
        datetime created_at
    }

    comments {
        bigint id PK
        bigint post_id FK
        bigint user_id FK
        text content
        datetime created_at
        datetime updated_at
        datetime deleted_at "Soft Delete"
    }

    comment_likes {
        bigint id PK
        bigint comment_id FK
        bigint user_id FK
        datetime created_at
    }
```

---

## 🛠️ Getting Started

### 1. Prasyarat
Pastikan environment lokal kamu memiliki:
- [Go](https://go.dev/dl/) (>= 1.25)
- [Docker Desktop](https://www.docker.com/products/docker-desktop/)
- [dbmate](https://github.com/amacneil/dbmate) *(Opsional, untuk CLI migrasi database)*

### 2. Konfigurasi Environment
Salin file template `.env.example` menjadi `.env`:

```bash
cp .env.example .env
```

Sesuaikan konfigurasi kredensial database dan JWT:
```env
DATABASE_URL="mysql://root:password@127.0.0.1:3306/go_tweets"
PORT=8080
APP_ENV=development
SECRET_JWT=your_super_secret_jwt_key_here
```

### 3. Jalankan Database (Docker)
Nyalakan container MySQL 8.0 di background:

```bash
docker compose up -d
```

### 4. Eksekusi Migrasi Database
Jalankan file migrasi tabel menggunakan `dbmate`:

```bash
dbmate up
```

### 5. Jalankan Server API
Jalankan aplikasi Go:

```bash
go run cmd/main.go
```

Server aktif dan siap melayani request di `http://127.0.0.1:8080`.

---

## 📡 Dokumentasi Endpoint API

### 🏥 Health Check

#### `GET /check`
Mengecek status ketersediaan server.
- **Auth:** Publik
- **Response (200 OK):**
  ```json
  {
    "mesage": "App is running"
  }
  ```

---

### 🔐 Modul Autentikasi (`/auth`)

#### 1. `POST /auth/register`
Mendaftarkan akun pengguna baru dengan verifikasi keunikan email/username dan kecocokan konfirmasi password.
- **Auth:** Publik
- **Request Body:**
  ```json
  {
    "email": "user@example.com",
    "username": "user123",
    "password": "secretpassword",
    "password_confirm": "secretpassword"
  }
  ```
- **Response (201 Created):**
  ```json
  {
    "id": 1
  }
  ```

#### 2. `POST /auth/login`
Masuk menggunakan kredensial email & password untuk memperoleh *Access Token* (masa aktif singkat) dan *Refresh Token* (masa aktif panjang).
- **Auth:** Publik
- **Request Body:**
  ```json
  {
    "email": "user@example.com",
    "password": "secretpassword"
  }
  ```
- **Response (200 OK):**
  ```json
  {
    "access_token": "eyJhbGciOiJIUzI1NiIsIn...",
    "refresh_token": "eyJhbGciOiJIUzI1NiIsIn..."
  }
  ```

#### 3. `POST /auth/refresh`
Memperbarui token akses yang kedaluwarsa dengan mekanisme *Token Rotation*.
- **Auth:** Protected *(Bearer Token menggunakan Refresh Token)*
- **Headers:**
  ```http
  Authorization: Bearer <refresh_token>
  ```
- **Response (200 OK):**
  ```json
  {
    "access_token": "eyJhbGciOiJIUzI1NiIsIn...",
    "refresh_token": "eyJhbGciOiJIUzI1NiIsIn..."
  }
  ```

---

### 📝 Modul Tweets / Postingan (`/tweets`)

#### 1. `GET /tweets/` *(Feed Publik Berpaginasi)*
Mengambil seluruh postingan aktif yang diurutkan dari yang paling baru (`ORDER BY created_at DESC`), lengkap dengan data author, total like, dan daftar komentar.
- **Auth:** Publik
- **Query Parameters:**
  - `page` (opsional, default: `1`): Nomor halaman saat ini.
  - `limit` (opsional, default: `10`): Jumlah tweet per halaman.
- **Request Example:**
  ```http
  GET /tweets/?page=1&limit=5
  ```
- **Response (200 OK):**
  ```json
  {
    "total_page": 4,
    "current_page": 1,
    "limit": 5,
    "data": [
      {
        "id": 10,
        "username": "ferry",
        "title": "Membangun API Go dengan Clean Architecture",
        "content": "Pola Eager Loading dan in-memory grouping sangat ampuh mengatasi N+1 query problem.",
        "like_count": 8,
        "comments": [
          {
            "id": 1,
            "username": "reviewer1",
            "content": "Setuju! Performa Go sangat terasa stabil.",
            "like_count": 3,
            "created_at": "2026-09-29 04:00:00 +0700 WIB",
            "updated_at": "2026-09-29 04:00:00 +0700 WIB"
          }
        ],
        "created_at": "2026-09-29 03:30:00 +0700 WIB",
        "updated_at": "2026-09-29 03:30:00 +0700 WIB"
      }
    ]
  }
  ```

#### 2. `GET /tweets/:post_id/detail`
Mendapatkan detail 1 tweet tertentu beserta seluruh komentarnya yang diurutkan berdasarkan jumlah like tertinggi (`ORDER BY like_count DESC`).
- **Auth:** Publik
- **Response (200 OK):**
  ```json
  {
    "id": 10,
    "username": "ferry",
    "title": "Membangun API Go dengan Clean Architecture",
    "content": "Pola Eager Loading dan in-memory grouping sangat ampuh mengatasi N+1 query problem.",
    "like_count": 8,
    "comments": [
      {
        "id": 1,
        "username": "reviewer1",
        "content": "Setuju! Performa Go sangat terasa stabil.",
        "like_count": 3,
        "created_at": "2026-09-29 04:00:00 +0700 WIB",
        "updated_at": "2026-09-29 04:00:00 +0700 WIB"
      }
    ],
    "created_at": "2026-09-29 03:30:00 +0700 WIB",
    "updated_at": "2026-09-29 03:30:00 +0700 WIB"
  }
  ```

#### 3. `POST /tweets/`
Membuat tweet baru untuk akun yang sedang login.
- **Auth:** Protected *(Bearer Access Token)*
- **Headers:** `Authorization: Bearer <access_token>`
- **Request Body:**
  ```json
  {
    "title": "Halo Komunitas Go!",
    "content": "Ini adalah tweet pertama saya melalui API go-tweets."
  }
  ```
- **Response (201 Created):**
  ```json
  {
    "id": 10
  }
  ```

#### 4. `PUT /tweets/:post_id/update`
Memperbarui judul dan isi konten tweet milik pengguna (dilengkapi validasi hak kepemilikan / ownership check).
- **Auth:** Protected *(Bearer Access Token)*
- **Headers:** `Authorization: Bearer <access_token>`
- **Request Body:**
  ```json
  {
    "title": "Judul Baru yang Diperbaiki",
    "content": "Isi tweet baru setelah melalui proses pengeditan."
  }
  ```
- **Response (200 OK):**
  ```json
  {
    "id": 10
  }
  ```

#### 5. `DELETE /tweets/:post_id/delete`
Menghapus tweet secara aman (*Soft Delete* dengan mengisi timestamp `deleted_at`).
- **Auth:** Protected *(Bearer Access Token)*
- **Headers:** `Authorization: Bearer <access_token>`
- **Response (200 OK):**
  ```json
  {
    "message": "Post deleted successfully"
  }
  ```

#### 6. `POST /tweets/action` *(Like / Unlike Toggle)*
Menyukai atau membatalkan suka pada sebuah tweet secara otomatis (*Toggle*).
- **Auth:** Protected *(Bearer Access Token)*
- **Headers:** `Authorization: Bearer <access_token>`
- **Request Body:**
  ```json
  {
    "post_id": 10
  }
  ```
- **Response (200 OK):**
  ```json
  {
    "message": "succesfully liked or unliked post"
  }
  ```

---

### 💬 Modul Komentar (`/comment`)

#### 1. `POST /comment/`
Menambahkan komentar baru ke tweet tertentu.
- **Auth:** Protected *(Bearer Access Token)*
- **Headers:** `Authorization: Bearer <access_token>`
- **Request Body:**
  ```json
  {
    "post_id": 10,
    "content": "Komentar yang sangat bermanfaat!"
  }
  ```
- **Response (200 OK):**
  ```json
  {
    "message": "comment created successfully"
  }
  ```

#### 2. `POST /comment/action` *(Like / Unlike Toggle Komentar)*
Menyukai atau membatalkan suka pada sebuah komentar secara otomatis (*Toggle*).
- **Auth:** Protected *(Bearer Access Token)*
- **Headers:** `Authorization: Bearer <access_token>`
- **Request Body:**
  ```json
  {
    "comment_id": 1
  }
  ```
- **Response (200 OK):**
  ```json
  {
    "message": "succesfully liked or unliked comment"
  }
  ```

---

## 📚 Catatan Belajar & Deep Dive Arsitektur (`docs/learn/`)

Seluruh riwayat pembuatan fitur, bedah sintaks Go, tips debugging, serta analogi perbandingan dengan Laravel (PHP) dan Next.js (TypeScript) didokumentasikan secara terstruktur:

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
- [11 - Fitur Like & Unlike Komentar (Toggle Action)](docs/learn/penjelasan_11_fitur_like_dan_unlike_komentar_toggle.md)
- [12 - Fitur Detail Tweet & Pola Eager Loading Komentar](docs/learn/penjelasan_12_fitur_detail_tweet_dan_eager_loading_komentar.md)
- [13 - Fitur Get All Tweets & Pagination (Offset-Based)](docs/learn/penjelasan_13_fitur_get_all_tweet_dan_pagination.md)

---

## 📄 License
This project is open-source and available under the [MIT License](LICENSE).
