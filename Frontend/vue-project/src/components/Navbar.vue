<template>
  <div v-if="userRole === 'Admin'">
    <nav class="sidebar">
      <br />
      <ul class="sidebar-top sidebar-bg">
        <li class="sidebar-toggle navbar-brand d-flex align-items-center">
          <img src="../assets/../assets/img/Logo.png" alt="Logo Desa" height="50" class="me-3" />
          <div>
            <span class="font-weight-bold nama-desa">Desa Bahomoleo</span>
            <span class="sub-text">Kab. Morowali</span>
          </div>
        </li>
      </ul>

      <div class="sidebar-content">
        <ul class="sidebar-list">
          <li class="sidebar-link">
            <router-link to="/beranda-admin" class="menu-item">Beranda</router-link>
            <router-link to="/warga-desa" class="menu-item">
              Warga Desa</router-link>
            <router-link to="/berita-desa" class="menu-item">Berita Desa</router-link>
            <router-link to="/umkm-desa" class="menu-item">UMKM Desa</router-link>

            <router-link to="/wisata-desa" class="menu-item">Wisata Desa</router-link>

            <div class="menu-item1 dropdown">
              <button @click="togglePemerintahanDropdown" class="btn dropdown-toggle" type="button"
                id="dropdownPemerintahan" :aria-expanded="isPemerintahanDropdownOpen ? 'true' : 'false'">
                Pemerintahan
              </button>
              <div v-if="isPemerintahanDropdownOpen" class="dropdown-menu" aria-labelledby="dropdownPemerintahan">
                <router-link to="/pemerintahan/kepala-desa" class="dropdown-item">Kepala Desa</router-link>
                <router-link to="/pemerintahan/sambutan-kades" class="dropdown-item">Sambutan Kepala Desa</router-link>
                <router-link to="/pemerintahan/sejarah-desa" class="dropdown-item">Sejarah Desa</router-link>
                <router-link to="/pemerintahan/visi-misi" class="dropdown-item">Visi Misi</router-link>
                <router-link to="/pemerintahan/program-kerja" class="dropdown-item">Program Kerja</router-link>
                <router-link to="/pemerintahan/peraturan" class="dropdown-item">Peraturan</router-link>
                <router-link to="/pemerintahan/pegawai" class="dropdown-item">Pegawai</router-link>
                <router-link to="/pemerintahan/struktur-organisasi" class="dropdown-item">Struktur Organisasi</router-link>
              </div>
            </div>

            <router-link to="/potensi-desa" class="menu-item">Potensi Desa</router-link>

            <router-link to="/kreatifitas-desa" class="menu-item">Kreatifitas Desa</router-link>

            <div class="menu-item1 dropdown">
              <button @click="togglePemerintahanDropdown" class="btn dropdown-toggle" type="button"
                id="dropdownPemerintahan" :aria-expanded="isPemerintahanDropdownOpen ? 'true' : 'false'">
                Informasi
              </button>
              <div v-if="isPemerintahanDropdownOpen" class="dropdown-menu" aria-labelledby="dropdownPemerintahan">
                <router-link to="/pemerintahan/halaman1" class="dropdown-item">Pembangunan</router-link>
                <router-link to="/wisata-desa" class="dropdown-item">Produk Hukum</router-link>
              </div>
            </div>

            <router-link to="/wilayah-desa" class="menu-item">Wilayah Desa</router-link>

            <router-link to="/warga-desa" class="menu-item">Layanan Desa</router-link>
          </li>
        </ul>

        <div class="sidebar-bottom">
          <button @click="clearLocalStorage" class="btn panjang-tombol-login1 btn-light text-tombol text-start">
            <span class="tulisan"><v-icon name="ai-academiaSquare" />Logout</span>
          </button>
        </div>
      </div>
    </nav>
  </div>

  <div v-else-if="userRole === 'Warga'">
    <nav class="sidebar">
      <br />
      <router-link to="/beranda-warga">
        <ul class="sidebar-top sidebar-bg">
          <li class="sidebar-toggle navbar-brand d-flex align-items-center">
            <img src="../assets/../assets/img/Logo.png" alt="Logo Desa" height="50" class="me-3" />
            <div>
              <span class="font-weight-bold nama-desa">Desa Bahomoleo</span>
              <span class="sub-text">Kab. Morowali</span>
            </div>
          </li>
        </ul>
      </router-link>

      <div class="sidebar-content">
        <ul class="sidebar-list">
          <li class="sidebar-link">
            <router-link to="/skbi" class="menu-item-warga">Surat Keterangan Beda Identitas</router-link>
            <hr class="tambah-border" />
            <router-link to="/skm-skck" class="menu-item-warga">
              Surat Keterangan Mendapatkan SKCK</router-link>
            <hr class="tambah-border" />

            <router-link to="/skp" class="menu-item-warga">Surat Keterangan Pemakaman</router-link>
            <hr class="tambah-border" />

            <router-link to="/skpd" class="menu-item-warga">Surat Keterangan Pindah Domisili</router-link>
            <hr class="tambah-border" />

            <router-link to="/skrn" class="menu-item-warga">Surat Keterangan Rekomendasi Nikah</router-link>
            <hr class="tambah-border" />

            <router-link to="/sktu" class="menu-item-warga">Surat Keterangan Tempat Usaha</router-link>
            <hr class="tambah-border" />

            <router-link to="/sktm" class="menu-item-warga">Surat Keterangan Tidak Mampu</router-link>
            <hr class="tambah-border" />

            <router-link to="/sku" class="menu-item-warga">Surat Keterangan Usaha</router-link>
            <hr class="tambah-border" />

            <router-link to="/skbm" class="menu-item-warga">Surat Keterangan Belum Menikah</router-link>
            <hr class="tambah-border" />

            <router-link to="/skd" class="menu-item-warga">Surat Keterangan Domisili</router-link>
            <hr class="tambah-border" />

            <router-link to="/sksm" class="menu-item-warga">Surat Keterangan Sudah Menikah</router-link>
            <hr class="tambah-border" />

            <router-link to="/skktm" class="menu-item-warga">Surat Keterangan Keluarga Tidak Mampu</router-link>
            <hr class="tambah-border" />

            <router-link to="/skkm" class="menu-item-warga">Surat Keterangan Keterlambatan Menikah</router-link>
            <hr class="tambah-border" />

            <router-link to="/skik" class="menu-item-warga">Surat Keterangan Izin Keramaian</router-link>
            <hr class="tambah-border" />

            <router-link to="/spipj" class="menu-item-warga">Surat Pengantar Izin Penutupan Jalan</router-link>
          </li>
        </ul>

        <div class="sidebar-bottom">
          <button @click="clearLocalStorage" class="btn panjang-tombol-login1 btn-light text-tombol text-start">
            <span class="tulisan"><v-icon name="ai-academiaSquare" />Logout</span>
          </button>
        </div>
      </div>
    </nav>
  </div>

  <div v-else>
    <nav class="menu">
      <div class="navbar navbar-expand-lg navbar-dark bg-transparent">
        <div className="container">
          <router-link to="/" class="navbar-brand d-flex align-items-center">
            <img src="../assets/../assets/img/Logo.png" alt="Logo Desa" height="50" class="me-3" />
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
            <div class="menu-item"><a href="/">Beranda</a></div>
            <Dropdown title="Profil Desa" :items="profil" />
            <div class="menu-item">
              <a href="/pemerintah-desa">Pemerintah Desa</a>
            </div>
            <div class="menu-item"><a href="/informasi">Informasi</a></div>
            <Dropdown title="Media" :items="media" />
            <Dropdown title="Destinasi" :items="destinasi" />
            <div class="menu-item"><a href="/idm">IDM</a></div>
            <div class="menu-item"><a href="/data-desa">Data Desa</a></div>
          </div>

          <div class="collapse navbar-collapse justify-content-end" id="navbarNav">
            <div class="d-flex">
              <button class="btn panjang-tombol-login btn-light me-2 text-tombol">
                <router-link to="/login" class="nav-link router-link-underline">Login</router-link>
              </button>
            </div>
          </div>
        </div>
      </div>
    </nav>
  </div>
</template>

<script>
import Dropdown from "./Dropdown.vue";

export default {
  name: "navbar",
  components: {
    Dropdown,
  },
  data() {
    return {
      userRole: 'Admin',
      isPemerintahanDropdownOpen: true,
      profil: [
        {
          title: "Sejarah Desa",
          link: "/sejarah",
        },
        {
          title: "Wilayah Desa",
          link: "/wilayah",
        },
        {
          title: "Potensi Desa",
          link: "/potensi",
        },
      ],
      media: [
        {
          title: "Galeri Foto",
          link: "/galeri-foto",
        },
        {
          title: "Berita",
          link: "/berita",
        },
        {
          title: "Artikel",
          link: "/artikel",
        },
      ],
      destinasi: [
        {
          title: "UMKM",
          link: "/umkm",
        },
        {
          title: "Wisata",
          link: "/wisata",
        },
      ],
    };
  },

  created() {
    this.checkUserRole();
    window.addEventListener("storage", this.handleStorageChange);
  },
  destroyed() {
    window.removeEventListener("storage", this.handleStorageChange);
  },
  methods: {
    clearLocalStorage() {
      localStorage.clear();
      console.log("Localstorage dibersihin!");
      this.userRole = "";
      location.reload();
    },

    checkUserRole() {
      this.userRole = localStorage.getItem("role_pengguna");
      if (this.userRole === "Admin") {
        this.$router.push("/beranda-admin");
      }

      if (this.userRole === "Warga") {
        this.$router.push("/beranda-warga");
      }
    },
    handleStorageChange(event) {
      if (event.key === "role_pengguna") {
        this.checkUserRole();
      }
    },
    togglePemerintahanDropdown() {
      this.isPemerintahanDropdownOpen = !this.isPemerintahanDropdownOpen;
      console.log("Dropdown status:", this.isPemerintahanDropdownOpen);
    },
  },
  watch: {
    userRole(newRole) {
      if (newRole === "Admin") {
        this.$router.push("/beranda-admin");
      }

      if (newRole === "Warga") {
        this.$router.push("/beranda-warga");
      }
    },
    isPemerintahanDropdownOpen(newVal) {
      console.log("Dropdown status:", newVal);
    },
  },
};
</script>

<style>
.sidebar {
  width: 285px;
  height: 100vh;
  position: fixed;
  justify-content: center;
  top: 0;
  left: 0;
  z-index: 999;
  background-color: #003366;
  margin-left: -1%;
  padding: 0px 30px;
  overflow: auto;
}

.sidebar-top {
  width: 100%;
  height: 90px;
  line-height: 30px;
}

.sidebar-toggle {
  width: 90%;
  height: 90px;
  display: flex;
  justify-content: center;
  align-items: center;
  border: none;
}

.sidebar-list {
  display: flex;
  padding: 5px 15px;
  row-gap: 1.5rem;
  flex-direction: column;
  border: none;
}

ul {
  list-style-type: none;
}

/* User */
* {
  font-family: "Poppins", sans-serif !important;
}

.menu {
  background-color: #003366;
  line-height: 2;
  box-shadow: inset 0px 0px 20px 0px rgba(0, 0, 0, 0.2);
}

.sub-text {
  font-size: 0.7em;
  opacity: 0.7;
  display: block;
  margin-top: -10px;
  color: #ffffff;
}

.nama-desa {
  font-weight: bold;
  color: rgb(255, 255, 255);
}

.text-tombol {
  font-weight: bold;
  color: rgba(255, 255, 255, 0.2);
  font-size: 15px;
}

.panjang-tombol-login {
  width: 100px;
}

.menu-item a {
  color: white;
  text-decoration: none;
}

.menu-item {
  color: white;
  padding: 10px 15px;
  position: relative;
  text-align: center;
  border-bottom: 3px solid transparent;
  display: flex;
  transition: all 0.3s ease;
  text-decoration: none;
  font-size: 14px;
}

.menu-item.active,
.menu-item:hover {
  background-color: white;
  border-bottom-color: #003366;
  border-radius: 20px;
  color: #003366;
}

.menu-item.active a,
.menu-item:hover a {
  color: #003366;
}

.sidebar-bottom {
  color: white;
  padding: 80px 25px;
  position: relative;
  text-align: center;
  border-bottom: 3px solid transparent;
  display: flex;
  transition: all 0.3s ease;
}

.sidebar-buttom span {
  color: inherit;
  text-decoration: none;
}

.sidebar-buttom.active,
.sidebar-buttom:hover {
  background-color: white;
  border-bottom-color: #003366;
  border-radius: 20px;
}

.sidebar-buttom.active span,
.sidebar-buttom:hover span {
  color: #003366;
}

.panjang-tombol-login1 {
  width: 250px;
}

.menu-item-warga a {
  color: white;
  text-decoration: none;
}

.menu-item-warga {
  color: white;
  padding: 10px 15px;
  position: relative;
  text-align: left;
  border-bottom: 3px solid transparent;
  display: flex;
  transition: all 0.3s ease;
  text-decoration: none;
}

.menu-item-warga.active,
.menu-item-warga:hover {
  background-color: white;
  border-bottom-color: #003366;
  border-radius: 20px;
  color: #003366;
}

.menu-item-warga.active a,
.menu-item-warga:hover a {
  color: #003366;
}

.tambah-border {
  border-color: #ffffff;
  opacity: 1;
  border-width: 1px;
}

hr {
  margin-top: 0.8rem;
}
.menu-item1.dropdown {
  position: relative;
  display: inline-block;
}

.menu-item1.dropdown button {
  background-color: #003366;
  color: white;
  padding: 10px 15px;
  border: none;
  border-radius: 5px;
  cursor: pointer;
  transition: background-color 0.3s;
  font-size: 14px;
}

.menu-item1.dropdown button:hover {
  background-color: #ffffff;
  color: #003366;
  border-radius: 20px;
}

.menu-item1.dropdown .dropdown-menu {
  display: none;
  position: relative;
  top: 100%; /* Sesuaikan dengan tinggi tombol dropdown */
  left: 0;
  background-color: #003366;
  min-width: 160px;
  box-shadow: 0px 8px 16px 0px rgba(0,0,0,0.2);
  z-index: 1;
  border-radius: 15px;
  font-size: 14px;
}

.menu-item1.dropdown .dropdown-menu a {
  color: white;
  padding: 12px 18px;
  text-decoration: none;
  display: block;
  transition: background-color 0.3s;
}

.menu-item1.dropdown .dropdown-menu a:hover {
  background-color: #ffffff;
  color: #003366;
}

.menu-item1.dropdown:hover .dropdown-menu {
  display: block;
}

.dropdown-toggle{
  color: white;
}

.tulisan{
  font-size: 14px;
}
</style>
