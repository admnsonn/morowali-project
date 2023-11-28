<template>
  <section>
    <div className="p-5 text-center bg-adit">
      <h1 className="mb-3">POTENSI DESA</h1>
      <button
        className="btn panjang-tombol-login btn-light me-2 text-tombol"
        @click="fetchData"
      >
        Lihat Terbaru
      </button>
    </div>

    <section v-if="showLatestData" className="py-5 bg-light">
      <div className="container">
        <div
          v-for="item in paginatedData"
          :key="item.id_potensi"
          className="border rounded row align-items-center pt-3 pb-3 mb-4 with-shadow"
          :className="{ 'with-shadow': isHovered }"
          @mouseenter="addShadow"
          @mouseleave="removeShadow"
        >
          <!-- Display the data here -->

          <div className="row">
            <!-- INI UNTUK GAMBAR PADA MOBILE -->
            <div className="col-12 d-md-none mb-4">
              <img
                src="~@/assets/Artikel.png"
                alt="Latest Image"
                className="img-fluid"
              />
            </div>

            <!-- Kolom untuk teks pada kedua ukuran layar -->
            <div className="col-md-6 order-md-1">
              <div className="text-center text-md-start">
                <!-- Judul untuk tampilan mobile -->
                <h2 className="mb-3 d-md-none warna-judul-artikel">
                  {{ item.judul_potensi }}
                </h2>
                <!-- Container untuk judul pada tampilan desktop -->
                <div className="d-none d-md-block">
                  <h1 className="mb-2 warna-judul-artikel">
                    {{ item.judul_potensi }}
                  </h1>
                </div>
                <h5 className="mb-4">{{ item.sub_judul }}</h5>
                <h5 className="mb-4">{{ item.date }}</h5>
                <p className="mb-4">{{ item.deskripsi }}</p>
                <button className="btn btn-primary" @click="showDetail(item)">
                  Lihat Selengkapnya
                </button>
              </div>
            </div>

            <!-- INI UNTUK GAMBAR PADA DESKTOP -->
            <div className="col-md-6 order-md-2 d-none d-md-block">
              <img
                src="~@/assets/Artikel.png"
                alt="Latest Image"
                className="img-fluid"
              />
            </div>
          </div>
        </div>
      </div>

      <!-- Pagination -->
    </section>

    <section>
      <!-- Pagination -->
      <div className="text-center mt-4">
        <nav aria-label="Page navigation example">
          <ul className="pagination justify-content-center">
            <!-- Tombol Previous -->
            <li className="page-item" :className="{ disabled: currentPage === 1 }">
              <button className="page-link" @click="prevPage" aria-label="Previous">
                <span aria-hidden="true">&laquo;</span>
                <span className="visually-hidden">Previous</span>
              </button>
            </li>

            <!-- Tampilan Nomor Halaman -->
            <li v-for="pageNumber in totalPages" :key="pageNumber">
              <button
                className="page-link"
                @click="goToPage(pageNumber)"
                :className="{ active: currentPage === pageNumber }"
                :style="{
                  backgroundColor: currentPage === pageNumber ? '#003366' : '',
                  color: currentPage === pageNumber ? 'white' : '',
                }"
              >
                {{ pageNumber }}
              </button>
            </li>

            <!-- Tombol Next -->
            <li
              className="page-item"
              :className="{ disabled: currentPage === totalPages }"
            >
              <button className="page-link" @click="nextPage" aria-label="Next">
                <span aria-hidden="true">&raquo;</span>
                <span className="visually-hidden">Next</span>
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
.bg-adit {
  background-color: #003366;
  color: white;
}

.bg-adit:hover,
.bg-adit:focus {
  background-color: #003366;
  color: white;
}

.btn {
  font-weight: 500;
}

/* Gaya CSS tambahan untuk komponen */
.bg-light {
  background-color: #f8f9fa;
  /* Warna latar belakang */
}

/* (Opsi) Gaya CSS khusus untuk tombol */
.btn-primary {
  background-color: #003366;
  border-color: #003366;
}

.warna-judul-artikel {
  color: #003366 !important;
  /* Warna biru (#003366) dengan menggunakan !important */
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

