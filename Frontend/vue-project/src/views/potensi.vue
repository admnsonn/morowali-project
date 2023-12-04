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
              <img src="src/assets/img/Artikel.png" alt="Latest Image" class="img-fluid" />
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
                  <br>
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
              <img src="src/assets/img/Artikel.png" alt="Latest Image" class="img-fluid" />
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
      // Simulasi pengambilan data dari API (dalam hal ini, data respons API)
      const apiResponse = {
        data: [
          {
            id_potensi: 1,
            judul_potensi: "Pertanian Organik",
            date: "28-04-2001",
            deskripsi:
              "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Duis sodales consequat dictum. Praesent fermentum blandit ipsum et ultricies. Nunc ultrices neque ac velit aliquet, in iaculis nisi pellentesque. Pengembangan pertanian organik di daerah dengan tanah subur.",
            foto_potensi_desa: "gambar1.jpg",
            sub_judul: "Deskripsi tambahan untuk potensi desa 1",
          },
          {
            id_potensi: 2,
            judul_potensi: "Pemberdayaan Wanita",
            date: "28-04-2001",
            deskripsi:
              "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Duis sodales consequat dictum. Praesent fermentum blandit ipsum et ultricies. Nunc ultrices neque ac velit aliquet, in iaculis nisi pellentesque. Program pelatihan untuk pemberdayaan ekonomi wanita di desa.",
            foto_potensi_desa: "gambar2.jpg",
            sub_judul: "Deskripsi tambahan untuk potensi desa 2",
          },
          {
            id_potensi: 3,
            judul_potensi: "Usaha Kecil Menengah (UKM)",
            date: "28-04-2001",
            deskripsi:
              "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Duis sodales consequat dictum. Praesent fermentum blandit ipsum et ultricies. Nunc ultrices neque ac velit aliquet, in iaculis nisi pellentesque. Dukungan pembinaan UKM di bidang kerajinan tangan.",
            foto_potensi_desa: "gambar3.jpg",
            sub_judul: "Deskripsi tambahan untuk potensi desa 3",
          },
          {
            id_potensi: 4,
            judul_potensi: "Edukasi Anak Muda",
            date: "28-04-2001",
            deskripsi:
              "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Duis sodales consequat dictum. Praesent fermentum blandit ipsum et ultricies. Nunc ultrices neque ac velit aliquet, in iaculis nisi pellentesque. Program edukasi dan pelatihan keterampilan untuk anak muda.",
            foto_potensi_desa: "gambar4.jpg",
            sub_judul: "Deskripsi tambahan untuk potensi desa 4",
          },
          {
            id_potensi: 5,
            judul_potensi: "Program Lingkungan",
            date: "28-04-2001",
            deskripsi:
              "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Duis sodales consequat dictum. Praesent fermentum blandit ipsum et ultricies. Nunc ultrices neque ac velit aliquet, in iaculis nisi pellentesque. Inisiatif untuk menjaga lingkungan alam sekitar.",
            foto_potensi_desa: "gambar5.jpg",
            sub_judul: "Deskripsi tambahan untuk potensi desa 5",
          },
          {
            id_potensi: 6,
            judul_potensi: "Pariwisata Lokal",
            date: "28-04-2001",
            deskripsi:
              "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Duis sodales consequat dictum. Praesent fermentum blandit ipsum et ultricies. Nunc ultrices neque ac velit aliquet, in iaculis nisi pellentesque. Pengembangan atraksi pariwisata lokal.",
            foto_potensi_desa: "gambar6.jpg",
            sub_judul: "Deskripsi tambahan untuk potensi desa 6",
          },
          {
            id_potensi: 7,
            judul_potensi: "Kesehatan Masyarakat",
            date: "28-04-2001",
            deskripsi:
              "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Duis sodales consequat dictum. Praesent fermentum blandit ipsum et ultricies. Nunc ultrices neque ac velit aliquet, in iaculis nisi pellentesque. Program kesehatan masyarakat di desa.",
            foto_potensi_desa: "gambar7.jpg",
            sub_judul: "Deskripsi tambahan untuk potensi desa 7",
          },
          {
            id_potensi: 8,
            judul_potensi: "Program Literasi",
            date: "28-04-2001",
            deskripsi:
              "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Duis sodales consequat dictum. Praesent fermentum blandit ipsum et ultricies. Nunc ultrices neque ac velit aliquet, in iaculis nisi pellentesque. Peningkatan literasi penduduk desa melalui kegiatan membaca.",
            foto_potensi_desa: "gambar8.jpg",
            sub_judul: "Deskripsi tambahan untuk potensi desa 8",
          },
          {
            id_potensi: 9,
            judul_potensi: "Pelatihan Petani",
            date: "28-04-2001",
            deskripsi:
              "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Duis sodales consequat dictum. Praesent fermentum blandit ipsum et ultricies. Nunc ultrices neque ac velit aliquet, in iaculis nisi pellentesque. Pelatihan teknik pertanian modern untuk petani desa.",
            foto_potensi_desa: "gambar9.jpg",
            sub_judul: "Deskripsi tambahan untuk potensi desa 9",
          },
          {
            id_potensi: 10,
            judul_potensi: "Infrastruktur Desa",
            date: "28-04-2001",
            deskripsi:
              "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Duis sodales consequat dictum. Praesent fermentum blandit ipsum et ultricies. Nunc ultrices neque ac velit aliquet, in iaculis nisi pellentesque. Pengembangan infrastruktur dasar di desa.",
            foto_potensi_desa: "gambar10.jpg",
            sub_judul: "Deskripsi tambahan untuk potensi desa 10",
          },
        ],
      };

      // Setelah data diambil dari API (simulasi pengambilan data), atur ke latestData
      this.latestData = apiResponse.data;
      this.showLatestData = true;

      // Set halaman pertama ketika data diambil
      this.paginatedData = this.latestData.slice(0, this.itemsPerPage);
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
    // Panggil fetchData() saat komponen dimuat
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

.item-potensi{
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
  font-family: 'Poppins', sans-serif;
  font-weight: 700;
  font-size: 14px;
  cursor: pointer;
  border-radius: 30px;
}

.btn-secondary{
  background-color:#003366;
  border: none;
  padding: 10px 30px;
  font-family: 'Poppins', sans-serif;
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

