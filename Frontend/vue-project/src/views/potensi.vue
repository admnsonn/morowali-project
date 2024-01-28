<template>
  <section>
    <div class="p-5 text-center bg-hero mt-0 pb-10">
      <h1 class="mb-3 tulisan_judul">Potensi Desa</h1>
      <a class="btn btn-primary" href="" role="button">Profil Desa</a>
    </div>

    <section v-if="showLatestData" class="py-5 bg-light">
      <div class="container">
        <div v-for="item in paginatedData" :key="item.id_potensi"
          class="border rounded row align-items-center pt-3 pb-3 mb-4 with-shadow" :class="{ 'with-shadow': isHovered }"
          @mouseenter="addShadow" @mouseleave="removeShadow">
          <!-- Display the data here -->

          <div class="row">
            <!-- INI UNTUK GAMBAR PADA MOBILE -->
            <div class="col-12 d-md-none mb-4">
              <img :src="item.foto_potensi_desa" alt="Latest Image" class="img-fluid" />
            </div>

            <!-- Kolom untuk teks pada kedua ukuran layar -->
            <div class="col-md-6 order-md-1">
              <div class="text-center text-md-start">
                <!-- Judul untuk tampilan mobile -->
                <h2 class="mb-3 d-md-none warna-judul-artikel">
                  {{ item.judul_potensi }}
                </h2>
                <!-- Container untuk judul pada tampilan desktop -->
                <div class="d-none d-md-block">
                  <h2 class="mb-2 warna-judul-artikel">
                    {{ item.judul_potensi }}
                  </h2>
                  <br />
                </div>
                <h5 class="mb-4 sub-potensi">{{ item.sub_judul }}</h5>
                <h5 class="mb-4 item-potensi">{{ item.date }}</h5>
                <p class="mb-4">{{ item.deskripsi }}</p>
                <button class="btn btn-secondary" @click="showDetail(item)">
                  Selengkapnya
                </button>
              </div>
            </div>

            <!-- INI UNTUK GAMBAR PADA DESKTOP -->
            <div class="col-md-6 order-md-2 d-none d-md-block">
              <img :src="item.foto_potensi_desa" alt="Latest Image" class="img-fluid" />
            </div>
          </div>
        </div>
      </div>

      <!-- Pagination -->
    </section>

    <section>
      <!-- Pagination -->
      <div class="text-center mt-4">
        <nav aria-label="Page navigation example">
          <ul class="pagination justify-content-center">
            <!-- Tombol Previous -->
            <li class="page-item" :class="{ disabled: currentPage === 1 }">
              <button class="page-link" @click="prevPage" aria-label="Previous">
                <span aria-hidden="true">&laquo;</span>
                <span class="visually-hidden">Previous</span>
              </button>
            </li>

            <!-- Tampilan Nomor Halaman -->
            <li v-for="pageNumber in totalPages" :key="pageNumber">
              <button class="page-link" @click="goToPage(pageNumber)" :class="{ active: currentPage === pageNumber }"
                :style="{
                  backgroundColor: currentPage === pageNumber ? '#003366' : '',
                  color: currentPage === pageNumber ? 'white' : '',
                }">
                {{ pageNumber }}
              </button>
            </li>

            <!-- Tombol Next -->
            <li class="page-item" :class="{ disabled: currentPage === totalPages }">
              <button class="page-link" @click="nextPage" aria-label="Next">
                <span aria-hidden="true">&raquo;</span>
                <span class="visually-hidden">Next</span>
              </button>
            </li>
          </ul>
        </nav>
      </div>
    </section>
  </section>
</template>

<script>
import axios from "axios";

export default {
  data() {
    return {
      showLatestData: false,
      latestData: [], // Data fetched from API will be stored here
      paginatedData: [], // Data to be displayed for the current page
      isHovered: false,
      itemsPerPage: 4, // Number of items per page
      currentPage: 1, // Current page
    };
  },

  computed: {
    totalPages() {
      return Math.ceil(this.latestData.length / this.itemsPerPage);
    },
  },

  methods: {
    goToPage(pageNumber) {
      this.currentPage = pageNumber;
      this.updatePaginatedData();
    },

    // Metode prevPage(), nextPage(), updatePaginatedData() dari sebelumnya tetap sama
    prevPage() {
      if (this.currentPage > 1) {
        this.currentPage--;
        this.updatePaginatedData();
      }
    },

    nextPage() {
      if (this.currentPage < this.totalPages) {
        this.currentPage++;
        this.updatePaginatedData();
      }
    },

    updatePaginatedData() {
      const startIndex = (this.currentPage - 1) * this.itemsPerPage;
      const endIndex = startIndex + this.itemsPerPage;

      // Update paginatedData based on the selected page
      this.paginatedData = this.latestData.slice(startIndex, endIndex);
    },

    addShadow() {
      this.isHovered = true;
    },
    removeShadow() {
      this.isHovered = false;
    },
    fetchData() {
      axios
        .get("http://localhost:8080/potensi_desa/list/1")
        .then(({ data }) => {
          this.latestData = data.data;
          this.showLatestData = true;
          // Set halaman pertama ketika data diambil
          this.paginatedData = this.latestData.slice(0, this.itemsPerPage);
        })
        .catch((error) => {
          console.error("Error in Axios POST request:", error);
        });
    },
    handlePageChange(page) {
      this.currentPage = page;
      const startIndex = (page - 1) * this.itemsPerPage;
      const endIndex = startIndex + this.itemsPerPage;

      // Update paginatedData based on the selected page
      this.paginatedData = this.latestData.slice(startIndex, endIndex);
    },
    showDetail(data) {
      console.log("Menampilkan detail informasi:", data);
    },
  },
  mounted() {
    this.fetchData();
  },
};
</script>

<style scoped>
.bg-hero {
  background-color: #003366;
}

.tulisan_judul {
  color: #fff;
  font-weight: bold;
  text-shadow: 2px 2px #252525;
}

.item-potensi {
  color: #003366;
  font-size: 18px;
  font-weight: bold;
}

.btn-primary {
  background-color: white;
  color: #003366;
  font-weight: bold;
  border: none;
  padding: 10px 30px;
  font-family: "Poppins", sans-serif;
  font-weight: 700;
  font-size: 14px;
  cursor: pointer;
  border-radius: 30px;
}

.btn-secondary {
  background-color: #003366;
  border: none;
  padding: 10px 30px;
  font-family: "Poppins", sans-serif;
  font-weight: 600;
  font-size: 14px;
  cursor: pointer;
  border-radius: 10px;
}

.warna-judul-artikel {
  font-weight: bold;
  color: #003366 !important;
  /* Warna biru (#003366) dengan menggunakan !important */
}

.sub-potensi {
  font-weight: bold;
}

.with-shadow {
  transition: box-shadow 0.3s ease-in-out;
  box-shadow: none;
}

.with-shadow:hover,
.with-shadow.active {
  box-shadow: 0 0 10px 0 rgba(0, 51, 102, 0.5);
}

.pagination .page-link {
  color: #003366;
  background-color: #fff;
  border-color: #003366;
}

.pagination .page-link:hover,
.pagination .page-link:focus {
  color: #fff;
  background-color: #003366;
  border-color: #003366;
}

.pagination .page-item.active .page-link {
  background-color: #fff;
  border-color: #003366;
  color: #003366;
}

.pagination .page-item.active .page-link:hover,
.pagination .page-item.active .page-link:focus {
  background-color: #003366;
  border-color: #003366;
  color: white;
}
</style>
