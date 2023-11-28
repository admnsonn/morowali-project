<template>
  <nav className="navbar navbar-expand-lg navbar-dark bg-transparent">
    <div className="container">
      <router-link to="/" className="navbar-brand d-flex align-items-center">
        <img src="../assets/logo.png" alt="Logo Desa" height="50" className="me-3" />
        <div>
          <span className="font-weight-bold nama-desa">Desa Bahomoleo</span>
          <span className="sub-text">Kab. Morowali</span>
        </div>
      </router-link>

      <button className="navbar-toggler" type="button" data-bs-toggle="collapse" data-bs-target="#navbarNav"
        aria-controls="navbarNav" aria-expanded="false" aria-label="Toggle navigation">
        <span className="navbar-toggler-icon"></span>
      </button>

      <div className="collapse navbar-collapse"></div>
      <div className="collapse navbar-collapse justify-content-center" id="navbarNav">
        <ul className="navbar-nav">

          <li className="nav-item">
            <router-link to="/" className="nav-link router-link-underline white">Beranda</router-link>
          </li>

          <li className="nav-item dropdown">
            <a className="nav-link router-link-underline white dropdown-toggle" href="#" id="navbarDropdown" role="button"
              data-bs-toggle="dropdown" aria-expanded="false">
              Profil Desa
            </a>
            <ul className="dropdown-menu" aria-labelledby="navbarDropdown">
              <li>
                <router-link to="/sejarah-desa" className="dropdown-item router-link-underline white">Sejarah Desa</router-link>
              </li>
              <li>
                <router-link to="/wilayah-desa" className="dropdown-item router-link-underline white">Wilayah Desa</router-link>
              </li>
              <li>
                <router-link to="/potensi-desa" className="dropdown-item router-link-underline white">Potensi Desa</router-link>
              </li>
            </ul>
          </li>

          <li className="nav-item">
            <router-link to="/pemerintah" className="nav-link router-link-underline white">Pemerintah Desa</router-link>
          </li>

          <li className="nav-item">
            <router-link to="/informasi" className="nav-link router-link-underline white">Informasi</router-link>
          </li>

          <li className="nav-item dropdown">
            <a className="nav-link router-link-underline white dropdown-toggle" href="#" id="navbarDropdown" role="button"
              data-bs-toggle="dropdown" aria-expanded="false">
              Media
            </a>
            <ul className="dropdown-menu" aria-labelledby="navbarDropdown">
              <li>
                <router-link to="/galeri-foto" className="dropdown-item router-link-underline white">Galeri Foto</router-link>
              </li>
              <li>
                <router-link to="/berita" className="dropdown-item router-link-underline white">Berita</router-link>
              </li>
              <li>
                <router-link to="/artikel" className="dropdown-item router-link-underline white">Artikel</router-link>
              </li>
            </ul>
          </li>

          <li className="nav-item dropdown">
            <a className="nav-link router-link-underline white dropdown-toggle" href="#" id="navbarDropdown" role="button"
              data-bs-toggle="dropdown" aria-expanded="false">
              Destinasi
            </a>
            <ul className="dropdown-menu" aria-labelledby="navbarDropdown">
              <li>
                <router-link to="/umkm" className="dropdown-item router-link-underline white">UMKM</router-link>
              </li>
              <li>
                <router-link to="/wisata" className="dropdown-item router-link-underline white">Wisata</router-link>
              </li>
            </ul>
          </li>

          <li className="nav-item">
            <router-link to="/idm" className="nav-link router-link-underline white">IDM</router-link>
          </li>

          <li className="nav-item">
            <router-link to="/data-desa" className="nav-link router-link-underline white">Data Desa</router-link>
          </li>
        </ul>
      </div>

      <div className="collapse navbar-collapse justify-content-end" id="navbarNav">
        <div className="d-flex">
          <button className="btn panjang-tombol-login btn-light me-2 text-tombol">
            <router-link to="/login" className="nav-link router-link-underline">Login</router-link>
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
