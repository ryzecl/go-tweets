# Catatan Pembelajaran #01: Setup Fondasi, Docker MySQL, Database Migration, dan Server Gin

Dokumentasi ini mencatat rangkuman arsitektur, kode, dan konsep fundamental yang dipelajari pada tahap inisialisasi project **go-tweets** (Twitter/X backend clone).

---

## 1. Ringkasan Fitur yang Dibangun
1. **Docker Compose & MySQL**: Menjalankan MySQL 8.0 di atas Docker Desktop WSL 2 menggunakan *Named Volume* yang aman dan persisten.
2. **Database Migration (`dbmate`)**: Merancang dan mengeksekusi skema relasi database Twitter clone menggunakan DDL Plain SQL.
3. **Environment & Configuration Layer**: Load variabel lingkungan dari `.env` ke dalam Go struct bertipe kuat (*strongly typed*).
4. **Database Connection Pool**: Membuka pool koneksi MySQL menggunakan package standar `database/sql` dan driver MySQL.
5. **HTTP Server (Gin Framework)**: Menginisialisasi router Gin dan mendengarkan request HTTP di port lokal.

---

## 2. Tabel Analogi Konsep (Golang vs Laravel vs Next.js)

Bagi developer dengan latar belakang Laravel atau Next.js/TypeScript, berikut peta analogi arsitekturnya:

| Komponen di Fitur Ini | Padanan di Laravel | Padanan di Next.js / TypeScript | Fungsi Utama |
| :--- | :--- | :--- | :--- |
| `cmd/main.go` | `public/index.php` + `bootstrap/app.php` | `server.ts` / Next.js bootstrap | Titik masuk aplikasi, routing awal, dan start server |
| `internal/config/config.go` | `config/*.php` + `env()` | `process.env` / Zod env schema | Mengumpulkan konfigurasi `.env` ke satu tempat terpusat |
| `pkg/internalsql/mysql.go` | `config/database.php` (DB Facade) | `lib/prisma.ts` / Pool `mysql2` | Mengatur connection pool database |
| `db/migrations/*.sql` | `database/migrations/*.php` | `prisma/migrations/` | DDL SQL untuk membuat & memodifikasi struktur tabel |
| `dbmate` CLI | `php artisan migrate` | `npx prisma migrate dev` | Eksekutor migrasi database |

---

## 3. Alur Lifecycle Bootstrap Aplikasi

```
1. main() dipanggil (cmd/main.go)
       │
       ▼
2. Load Config (internal/config/config.go)
   - Baca file .env menggunakan godotenv
   - Simpan nilai ke dalam struct Config
       │
       ▼
3. Inisialisasi Pool MySQL (pkg/internalsql/mysql.go)
   - Rakit DSN: user:pass@tcp(host:port)/dbname?parseTime=true
   - Buka pool koneksi: sql.Open("mysql", dsn)
       │
       ▼
4. Inisialisasi Gin Engine (gin.Default())
   - Pasang middleware logger & recovery (anti-crash)
       │
       ▼
5. Jalankan HTTP Listener (r.Run("127.0.0.1:8080"))
```

---

## 4. Bedah Kode & Struktur Layer

### A. Skema Database (`db/migrations/`)
Model relasional yang dibuat mencakup:
- **`users`**: Tabel pengguna (`username`, `email`, `password` hash).
- **`posts`**: Tweet / postingan (`user_id`, `title`, `content`, `deleted_at` untuk soft delete).
- **`comments`**: Komentar pada postingan (`user_id`, `post_id`, `content`).
- **`post_likes` & `comment_likes`**: Tabel pivot relasi *many-to-many* untuk fitur like.
- **`refresh_tokens`**: Penyimpanan token refresh untuk rotasi sesi JWT.

### B. Configuration (`internal/config/config.go`)
```go
type Config struct {
    Port           string
    DBHost         string
    DBPort         string
    DBUser         string
    DBPassword     string
    DBName         string
    ...
}
```
* **Mengapa ditampung di struct?**
  Golang adalah bahasa bertipe statis (*statically typed*). Mengumpulkan environment variable ke dalam satu `struct Config` mencegah typo nama variabel di tengah kode aplikasi dan menerapkan prinsip *fail-fast* (langsung gagal saat startup jika `.env` tidak valid).

### C. Database Connection (`pkg/internalsql/mysql.go`)
```go
import (
    "database/sql"
    _ "github.com/go-sql-driver/mysql" // Blank identifier untuk side-effect
)

func ConnectMySQL(cfg *config.Config) (*sql.DB, error) {
    dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true&loc=%s", ...)
    db, err := sql.Open("mysql", dsn)
    ...
}
```
* Objek `*sql.DB` di Go bukanlah koneksi tunggal (*single connection*), melainkan **Connection Pool**. Go secara otomatis mengelola buka-tutup koneksi fisik ke MySQL saat query dijalankan.

---

## 5. Learning Corner: Konsep Fundamental Golang

1. **Aturan Huruf Kapital (Exported vs Unexported)**:
   - Tidak ada keyword `public` atau `private` di Go.
   - Nama identifier yang diawali **Huruf Besar** (contoh: `LoadConfig`, `Config`, `Port`) bersifat **Exported / Public** (dapat diakses package luar).
   - Nama identifier yang diawali **Huruf Kecil** (contoh: `dsn`, `server`) bersifat **Unexported / Private** (hanya bisa diakses di dalam package yang sama).

2. **Side-Effect Import (`_ "github.com/..."`)**:
   - Tanda underscore `_` memberitahu compiler: *"Jangan buang import ini walau saya tidak memanggil fungsinya secara langsung"*. Driver database mendaftarkan dirinya secara otomatis ke `database/sql` lewat fungsi bawaan `init()`.

3. **Pointer (`*Config` vs `Config`)**:
   - `*Config` merujuk ke alamat memori. Mengoper struct berukuran besar lewat pointer jauh lebih efisien daripada menduplikasi nilainya (*pass-by-value*).

4. **Explicit Error Handling (`if err != nil`)**:
   - Di Go, error bukanlah Exception seperti `try-catch` di Laravel/JavaScript, melainkan *return value* kedua yang wajib diperiksa sebelum menggunakan hasil fungsinya.

---

## 6. Catatan Pengalaman & Troubleshooting Penting

### A. Docker Desktop di WSL 2: Named Volume vs Bind Mount
* **Masalah**: Mengarahkan folder database MySQL ke folder Windows host (`C:\...` atau `./data`) menyebabkan crash InnoDB / permission error karena sistem file NTFS tidak mendukung spesifikasi POSIX file locking dan ownership `chown mysql:mysql`.
* **Solusi**: Gunakan **Named Volume** (`mysql_data:/var/lib/mysql`). Docker menyimpannya di file virtual disk WSL 2 (`.vhdx`) dengan format Linux native (ext4) yang cepat dan stabil.

### B. Trap `MYSQL_USER=root` di Docker
* Pada official image MySQL di Docker, user `root` otomatis dibuat dengan password dari `MYSQL_ROOT_PASSWORD`.
* Jika mendefinisikan `MYSQL_USER=root`, container akan mengalami crash loop dengan error:
  `MYSQL_USER="root", MYSQL_USER and MYSQL_PASSWORD are for configuring a regular user and cannot be used for the root user`.
* Variabel `MYSQL_USER` hanya boleh digunakan untuk membuat akun non-root tambahan.

---

## 7. Langkah Selanjutnya (Next Steps Roadmap)
1. **Repository Layer**: Membuat query database CRUD (misal: `UserRepository`, `PostRepository`) menggunakan `database/sql` atau query builder.
2. **Service Layer**: Menerapkan logika bisnis (hashing password bcrypt, validasi business rules).
3. **Handler / Controller Layer**: Menerima request HTTP dari Gin, validasi DTO struct (`binding:"required"`), dan mengembalikan response JSON.
4. **JWT Authentication**: Middleware autentikasi untuk memproteksi endpoint pembuatan tweet dan komentar.
