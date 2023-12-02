<template>
  <div v-if="userRole === 'Admin'">
    <nav class="sidebar-admin">
      <h1>INI JADI SIDE BAR</h1>
    </nav>
  </div>

  <div v-else-if="userRole === 'Warga'">
    <nav class="sidebar-admin">
      <h1>INI JADI SIDEBAR WARGA</h1>
    </nav>
  </div>

  <div v-else>
    <nav class="menu">
      <div class="navbar navbar-expand-lg navbar-dark bg-transparent">
        <div className="container">
          <router-link to="/" class="navbar-brand d-flex align-items-center">
            <img
              src="../assets/Logo.png"
              alt="Logo Desa"
              height="50"
              class="me-3"
            />
            <div>
              <span class="font-weight-bold nama-desa">Desa Bahomoleo</span>
              <span class="sub-text">Kab. Morowali</span>
            </div>
          </router-link>

          <button
            class="navbar-toggler"
            type="button"
            data-bs-toggle="collapse"
            data-bs-target="#navbarNav"
            aria-controls="navbarNav"
            aria-expanded="false"
            aria-label="Toggle navigation"
          >
            <span class="navbar-toggler-icon"></span>
          </button>

          <div class="collapse navbar-collapse"></div>
          <div
            class="collapse navbar-collapse justify-content-center"
            id="navbarNav"
          >
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

          <div
            class="collapse navbar-collapse justify-content-end"
            id="navbarNav"
          >
            <div class="d-flex">
              <button
                class="btn panjang-tombol-login btn-light me-2 text-tombol"
              >
                <router-link to="/login" class="nav-link router-link-underline"
                  >Login</router-link
                >
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
          link: "#",
        },
      ],
      destinasi: [
        {
          title: "UMKM",
          link: "#",
        },
        {
          title: "Wisata",
          link: "#",
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
  },
};
</script>

<style>
* {
  font-family: "Poppins", sans-serif !important;
}

.sidebar-admin {
  background-color: red;
  line-height: 2;
  box-shadow: inset 0px 0px 20px 0px rgba(0, 0, 0, 0.2);
}

.menu {
  background-color: #003366;
  line-height: 2;
  box-shadow: inset 0px 0px 20px 0px rgba(0, 0, 0, 0.2);
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

.menu-item {
  color: white;
  padding: 10px 15px;
  position: relative;
  text-align: center;
  border-bottom: 3px solid transparent;
  display: flex;
  transition: all 0.3s ease;
}

.menu-item a {
  color: inherit;
  text-decoration: none;
}

.menu-item.active,
.menu-item:hover {
  background-color: white;
  border-bottom-color: #003366;
  border-radius: 20px;
}

.menu-item.active a,
.menu-item:hover a {
  color: #003366;
}
</style>
