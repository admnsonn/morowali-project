<template>
  <nav class="navbar navbar-expand-lg navbar-dark bg-transparent">
    <div className="container">
      <router-link to="/" class="navbar-brand d-flex align-items-center">
        <img src="../assets/logo.png" alt="Logo Desa" height="50" class="me-3" />
        <div>
          <span class="font-weight-bold nama-desa">Desa Bahomoleo</span>
          <span class="sub-text">Kab. Morowali</span>
        </div>
      </router-link>

      <button class="navbar-toggler" type="button" data-bs-toggle="collapse" data-bs-target="#navbarNav"
        aria-controls="navbarNav" aria-expanded="false" aria-label="Toggle navigation">
        <span class="navbar-toggler-icon"></span>
      </button>

      <div class="collapse navbar-collapse"></div>
      <div class="collapse navbar-collapse justify-content-center" id="navbarNav">
        <ul class="navbar-nav">

          <li class="nav-item">
            <router-link to="/" class="nav-link router-link-underline white">Beranda</router-link>
          </li>

          <li class="nav-item dropdown">
            <a class="nav-link router-link-underline white dropdown-toggle" href="#" id="navbarDropdown" role="button"
              data-bs-toggle="dropdown" aria-expanded="false">
              Profil Desa
            </a>
            <ul class="dropdown-menu" aria-labelledby="navbarDropdown">
              <li>
                <router-link to="/sejarah-desa" class="dropdown-item router-link-underline white">Sejarah Desa</router-link>
              </li>
              <li>
                <router-link to="/wilayah-desa" class="dropdown-item router-link-underline white">Wilayah Desa</router-link>
              </li>
              <li>
                <router-link to="/potensi-desa" class="dropdown-item router-link-underline white">Potensi Desa</router-link>
              </li>
            </ul>
          </li>

          <li class="nav-item">
            <router-link to="/berita" class="nav-link router-link-underline white">Pemerintah Desa</router-link>
          </li>

          <li class="nav-item">
            <router-link to="/informasi" class="nav-link router-link-underline white">Informasi</router-link>
          </li>

          <li class="nav-item dropdown">
            <a class="nav-link router-link-underline white dropdown-toggle" href="#" id="navbarDropdown" role="button"
              data-bs-toggle="dropdown" aria-expanded="false">
              Media
            </a>
            <ul class="dropdown-menu" aria-labelledby="navbarDropdown">
              <li>
                <router-link to="/galeri-foto" class="dropdown-item router-link-underline white">Galeri Foto</router-link>
              </li>
              <li>
                <router-link to="/berita" class="dropdown-item router-link-underline white">Berita</router-link>
              </li>
              <li>
                <router-link to="/artikel" class="dropdown-item router-link-underline white">Artikel</router-link>
              </li>
            </ul>
          </li>

          <li class="nav-item dropdown">
            <a class="nav-link router-link-underline white dropdown-toggle" href="#" id="navbarDropdown" role="button"
              data-bs-toggle="dropdown" aria-expanded="false">
              Destinasi
            </a>
            <ul class="dropdown-menu" aria-labelledby="navbarDropdown">
              <li>
                <router-link to="/umkm" class="dropdown-item router-link-underline white">UMKM</router-link>
              </li>
              <li>
                <router-link to="/wisata" class="dropdown-item router-link-underline white">Wisata</router-link>
              </li>
            </ul>
          </li>

          <li class="nav-item">
            <router-link to="/idm" class="nav-link router-link-underline white">IDM</router-link>
          </li>

          <li class="nav-item">
            <router-link to="/data-desa" class="nav-link router-link-underline white">Data Desa</router-link>
          </li>
        </ul>
      </div>

      <div class="collapse navbar-collapse justify-content-end" id="navbarNav">
        <div class="d-flex">
          <button class="btn panjang-tombol-login btn-light me-2 text-tombol">
            <router-link to="/login" class="nav-link router-link-underline">Login</router-link>
          </button>
        </div>
      </div>
    </div>
  </nav>
</template>

<script>
export default {
  data() {
    return {
      isDropdownVisible: false, // Menyimpan status tampilan dropdown
    };
  },
  methods: {
    showDropdown() {
      this.isDropdownVisible = true; // Menampilkan dropdown saat kursor masuk
    },
    keepDropdown() {
      this.isDropdownVisible = true; // Tetapkan dropdown terbuka saat kursor berada di dalamnya
    },
    hideDropdown() {
      this.isDropdownVisible = false; // Menyembunyikan dropdown saat kursor keluar
    },
  },
};
</script>

<style scoped>
* {
  font-family: "Poppins", sans-serif !important;
}

.dropdown-menu {
  background-color: #003366;
  /* Atur warna latar belakang menjadi biru */
}

.dropdown-menu.show a:hover {
  color: black !important;
  /* Atur warna teks menjadi hitam saat dihover */
}

nav {
  box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
  /* Ubah nilai sesuai preferensi Anda */
}

.sub-text {
  font-size: 0.7em;
  /* Ukuran font yang lebih besar untuk sub-tulisan */
  opacity: 0.7;
  /* Opasitas untuk menampilkan sub-tulisan dengan lebih ringan */
  display: block;
  /* Menjadikan sub-tulisan sebagai blok terpisah */
  margin-top: -10px;
  /* Jarak antara teks utama dan sub-tulisan */
}

.nama-desa {
  font-weight: bold;
}

.text-tombol {
  font-weight: bold;
}

.panjang-tombol-login {
  width: 100px;
}

.router-link-underline {
  position: relative;
  text-decoration: none;
  transition: all 0.3s ease;
}

.router-link-underline::after {
  content: "";
  position: absolute;
  width: 100%;
  height: 1px;
  bottom: -0.5px;
  left: 0;
  background-color: transparent;
  transition: all 0.3s ease;
}

.router-link-underline:hover::after {
  background-color: white;
  /* Ganti dengan warna garis bawah yang diinginkan */
}

.white {
  color: white !important;
  /* Atur warna teks menjadi putih */
}
</style>
