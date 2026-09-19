# Catatan Pembelajaran #02: Standar Konvensi Pesan Commit Git (Conventional Commits)

Dokumentasi ini menjelaskan standar industri penulisan pesan commit Git yang dikenal sebagai **Conventional Commits**. Standar ini digunakan oleh ribuan open-source project (seperti Angular, Vue, Kubernetes) dan perusahaan teknologi modern agar riwayat perubahan (*git log*) rapi, mudah ditelusuri, dan bisa diotomatisasi untuk rilis versi (Changelog & Semantic Versioning).

---

## 1. Struktur Dasar Pesan Commit

Format standar Conventional Commits adalah:

```text
<type>[optional scope]: <deskripsi singkat imperative>

[optional body / penjelasan lebih detail]

[optional footer / issue tracker, misal: Closes #12]
```

### Contoh:
```text
feat(auth): implementasi login menggunakan jwt token
fix(mysql): perbaiki timeout koneksi database saat startup
docs: perbarui panduan instalasi di readme
```

---

## 2. Daftar Prefix / Type Commit & Situasi Penggunaannya

Berikut daftar tipe pesan commit yang diakui secara standar industri:

| Type | Arti / Kategori | Kapan Harus Dipakai? (Situasi Nyata) |
| :--- | :--- | :--- |
| **`feat:`** | **Feature** | Menambahkan fitur, kapabilitas, atau endpoint baru yang sebelumnya belum ada bagi user/sistem. |
| **`fix:`** | **Bug Fix** | Memperbaiki bug, error, atau logic yang salah sehingga aplikasi kembali berfungsi normal. *(Catatan: standar resminya pakai `fix:`, bukan `bug:`)*. |
| **`refactor:`** | **Refactoring** | Mengubah struktur kode tanpa mengubah fungsi atau perilakunya (misal: merapikan kodingan, memecah fungsi panjang jadi fungsi kecil). |
| **`perf:`** | **Performance** | Mengubah kode khusus untuk meningkatkan kecepatan, menghemat RAM, atau optimasi query database. |
| **`docs:`** | **Documentation** | Mengubah file dokumentasi (README.md, file di folder `docs/`, inline comments) tanpa menyentuh logic kode. |
| **`test:`** | **Testing** | Menambah, memperbaiki, atau melengkapi unit test / integration test. |
| **`chore:`** | **Maintenance** | Tugas pemeliharaan rutin yang tidak mengubah kode produksi (misal: update dependencies, konfigurasi linter, update `.gitignore`). |
| **`build:`** | **Build System** | Mengubah konfigurasi build, package manager, atau compiler (misal: `go.mod`, `package.json`, Dockerfile, docker-compose). |
| **`ci:`** | **Continuous Integration**| Mengubah file pipeline automasi (misal: GitHub Actions, GitLab CI, `.github/workflows/`). |
| **`style:`** | **Code Style** | Mengubah format penulisan tanpa merubah logika (spasi, indentasi, titik koma, trailing comma). |

---

## 3. Koreksi Penting: Kenapa `add:` dan `bug:` Kurang Tepat?

Banyak developer pemula sering menulis:
* ❌ `add: user model`
* ❌ `bug: fix login error`

### Mengapa perlu dikoreksi?
1. **`add:` vs `feat:`**:
   - Jika kamu "menambah fitur" baru -> gunakan **`feat:`**.
   - Jika kamu "menambah file konfigurasi / dependency" -> gunakan **`chore:`** atau **`build:`**.
   - Kata `add` terlalu ambigu: apakah menambah fitur, menambah bugfix, atau menambah file gambar?
2. **`bug:` vs `fix:`**:
   - Standar Semantic Versioning memakai kata kerja tindakan: **`fix:`** (artinya: *"commit ini memperbaiki..."*).
   - Tool automasi (seperti *Semantic Release*) mengenali kata `fix:` untuk otomatis menaikkan versi PATCH (misal: `v1.0.0` -> `v1.0.1`).

---

## 4. Studi Kasus & Contoh Situasi Nyata di Project `go-tweets`

Berikut panduan memilih prefix berdasarkan skenario yang sedang kamu kerjakan:

### Skenario 1: Membuat Endpoint Register User Baru
* **Type**: `feat`
* **Pesan**: `feat(auth): implementasi endpoint register user dengan hashing bcrypt`

### Skenario 2: Memperbaiki Crash MySQL Karena `MYSQL_USER=root`
* **Type**: `fix`
* **Pesan**: `fix(docker): hapus MYSQL_USER root untuk mencegah container crash loop`

### Skenario 3: Membuat File `.gitignore` dan `.env.example`
* **Type**: `chore`
* **Pesan**: `chore: tambahkan gitignore dan template env example`

### Skenario 4: Menulis Dokumentasi Belajar di Folder `docs/`
* **Type**: `docs`
* **Pesan**: `docs: tambahkan rangkuman pembelajaran setup migration dan docker`

### Skenario 5: Mempercepat Query Tweet dengan Menambahkan Index di MySQL
* **Type**: `perf` atau `feat(db)`
* **Pesan**: `perf(db): tambahkan index pada kolom user_id di tabel posts`

### Skenario 6: Memecah Fungsi `main.go` yang Kepanjangan ke Layer Handler & Service
* **Type**: `refactor`
* **Pesan**: `refactor: pisahkan inisialisasi routing dari file main.go`

---

## 5. Tips & Best Practice Menulis Commit

1. **Gunakan Huruf Kecil untuk Type**:
   Gunakan `feat:`, bukan `Feat:` atau `FEAT:`.
2. **Gunakan Bentuk Kalimat Perintah (*Imperative Mood*)**:
   - ✅ `feat: add user authentication`
   - ❌ `feat: added user authentication`
   - ❌ `feat: adding user authentication`
3. **Singkat dan Jelas (Maksimal 50-72 karakter di baris pertama)**:
   Jangan menulis novel di baris pertama. Jika butuh penjelasan panjang, beri jarak 1 baris kosong lalu tulis detailnya di body commit.
4. **Pisahkan Hal Berbeda ke Commit Terpisah (Atomic Commit)**:
   Hindari menggabungkan perbaikan bug, penambahan fitur, dan perapian format kode ke dalam satu commit raksasa. Buatlah commit kecil-kecil yang fokus pada satu tujuan.

---

## 6. Contoh Commit Pertama untuk Project Ini

Untuk perubahan yang sudah kita lakukan sejauh ini di branch `main`, kamu bisa membuat commit seperti:

```bash
git add .
git commit -m "feat: inisialisasi project, docker mysql, dbmate migrations, dan gin server"
```
