<template>
  <div v-if="userRole === 'Admin'">
    <!-- <h1>Footer Buat Admin Desa</h1> -->
  </div>

  <div v-else-if="userRole === 'Warga'">
    <!-- <h1>Footer untuk warga desa (Layanan Desa)</h1> -->
  </div>

  <div v-else>
    <br />
    <footer class="footer py-5">
      <div class="container">
        <div class="row">
          <div class="col-md-4 mb-4">
            <br />
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

          <div class="col-md-4 mb-4">
            <br />
            <ul class="nav justify-content-start">
              <li class="nav-item">
                <h2 class="nav-link text-white tulisan">Alamat Kantor</h2>
                <p class="nav-link text-white tulisan1 py-1 mb-2">
                  {{ alamat }}
                </p>
              </li>
            </ul>
          </div>

          <div class="col-md-4 mb-4">
            <br />
            <ul class="nav justify-content-end">
              <li class="nav-item">
                <h2 class="nav-link text-white tulisan">Kontak Person</h2>
                <p class="nav-link text-white tulisan1 py-1 mb-2">
                  <svg
                    xmlns="http://www.w3.org/2000/svg"
                    width="16"
                    height="16"
                    fill="currentColor"
                    class="bi bi-telephone-fill"
                    viewBox="0 0 16 16"
                  >
                    <path
                      fill-rule="evenodd"
                      d="M1.885.511a1.745 1.745 0 0 1 2.61.163L6.29 2.98c.329.423.445.974.315 1.494l-.547 2.19a.678.678 0 0 0 .178.643l2.457 2.457a.678.678 0 0 0 .644.178l2.189-.547a1.745 1.745 0 0 1 1.494.315l2.306 1.794c.829.645.905 1.87.163 2.611l-1.034 1.034c-.74.74-1.846 1.065-2.877.702a18.634 18.634 0 0 1-7.01-4.42 18.634 18.634 0 0 1-4.42-7.009c-.362-1.03-.037-2.137.703-2.877L1.885.511z"
                    />
                  </svg>
                  {{ kontak }}
                </p>
                <h2 class="nav-link text-white tulisan">Situs Terkait</h2>
                <p class="nav-link text-white link_web py-1 mb-2">
                  <a href="https://bungkutengah.com/"
                    >Kecamatan Bungku Tengah</a
                  >
                </p>
                <p class="nav-link text-white link_web mb-2">
                  <a href="https://morowalikab.go.id/">Kabupaten Morowali</a>
                </p>
              </li>

              <!-- Add other menu items as needed -->
            </ul>
          </div>
        </div>
      </div>
    </footer>
  </div>
</template>

<script>
export default {
  data() {
    return {
      alamat: "",
      kontak: "",
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
    fetchData() {
      const apiResponse = {
        data: {
          alamat:
            "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Vivamus ac scelerisque velit. Nam sit amet porta lectus, feugiat accumsan nisi. Sed tristique sem placerat velit luctus, sed sodales lectus consequat.",
          kontak: "(021) 2941 968323",
        },
      };
      this.alamat = apiResponse.data.alamat;
      this.kontak = apiResponse.data.kontak;
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
  },
  mounted() {
    this.fetchData();
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
  font-size: x-large;
}

.sub-text {
  font-size: 0.7em;
  opacity: 0.7;
  display: block;
  margin-top: -10px;
  font-size: medium;
}

.tulisan {
  font-weight: bold;
  font-size: x-large;
}

.tulisan1 {
  font-weight: 300;
  font-size: medium;
  text-overflow: clip ellipsis;
}

.jarak {
  margin-right: 400px;
  /* Jika diperlukan */
}

.text-white {
  color: #fff !important;
}

.text-white-50 {
  color: rgba(255, 255, 255, 0.5) !important;
}

.link_web {
  text-decoration: underline;
}

a {
  color: #fff;
}

a:hover {
  color: grey;
}

@media (max-width: 767px) {
  .footer {
    padding: 109px;
    margin-right: 10px;
  }

  .jarak {
    margin-right: 0;
  }
}

@media (max-width: 425px) {
  .footer {
    padding: 10px;
    margin-right: 10px;
  }

  .jarak {
    margin-right: 0;
  }

  .justify-content-end {
    justify-content: start !important;
  }
}
</style>
