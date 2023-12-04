<template>
  <div v-if="userRole === 'Admin'">
    <h1>Footer Buat Admin Desa</h1>
  </div>

  <div v-else-if="userRole === 'Warga'" >
    <h1 >Footer untuk warga desa (Layanan Desa)</h1>
  </div>

  <div v-else>
    <footer class="footer bg-transparent mt-5">
      <div class="container">
        <div class="row">
          <div class="col-lg-4 mb-4">
            <img
              src="../assets/img/Logo.png"
              alt="Logo Desa"
              height="50"
              class="me-3"
            />
            <div>
              <span class="fw-bold nama-desa text-white">Desa Bahomoleo</span>
              <span class="sub-text text-white-50">Kab. Morowali</span>
            </div>
          </div>
          <div class="col-lg-8 mb-4">
            <ul class="nav justify-content-end">
              <li class="nav-item">
                <h5 class="nav-link text-white tulisan">Alamat Kantor</h5>
              </li>
              <li class="nav-item">
                <h5 class="nav-link text-white tulisan">Profil Desa</h5>
              </li>
              <!-- Tambahkan menu lainnya sesuai kebutuhan -->
            </ul>
          </div>
        </div>
      </div>
    </footer>
  </div>
</template>

<script>
export default {
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
};
</script>

<style scoped>
/* Gaya untuk footer */
.footer {
  
  left: 0;
  width: 100%;
  background-color: #003366 !important;
  line-height: 2;
  padding: 140px;
}

/* Gaya untuk elemen-elemen dalam footer */
.nama-desa {
  font-weight: bold;
}

.sub-text {
  font-size: 0.7em;
  opacity: 0.7;
  display: block;
  margin-top: -10px;
}

.jarak {
  margin-right: 400px; /* Jika diperlukan */
}

.tulisan {
  font-weight: bold;
}

.text-white {
  color: #fff !important;
}

.text-white-50 {
  color: rgba(255, 255, 255, 0.5) !important;
}
</style>
