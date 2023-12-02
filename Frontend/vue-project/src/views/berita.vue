<template>
    <section>
        <div>
            <div class="p-5 text-center bg-hero mt-0 pb-10">
                <h1 class="mb-3 text-white">Berita Desa</h1>
                <a class="btn btn-primary" href="" role="button">Media</a>
            </div>
        </div>

        <section v-if="showLatestData" class="py-5 bg-light">
            <div class="container">
                <div v-for="item in paginatedData" :key="item.id_potensi"
                    class="border rounded row align-items-center pt-3 pb-3 mb-3 with-shadow"
                    :class="{ 'with-shadow': isHovered }" @mouseenter="addShadow" @mouseleave="removeShadow">
                    <!-- Display the data here -->

                    <div class="container mt-4">
                        <div class="row">
                            <div class="col-md-6">
                                <!-- Gambar berita -->
                                <img src="~@/assets/Artikel.png" alt="Gambar Berita" class="img-fluid">
                            </div>
                            <div class="col-md-6">
                                <div class="d-none d-md-block">
                                    <h2 class="mb-2 warna-judul-artikel">
                                        {{ item.judul_potensi }}
                                    </h2>
                                </div>
                                <h5 class="mb-4">{{ item.sub_judul }}</h5>
                                <h5 class="mb-4">{{ item.date }}</h5>
                                <p class="mb-4">{{ item.deskripsi }}</p>
                                <a href="/contentberita">
                                <button class="btn btn-primary" @click="showDetail(item)">
                                    Lihat Selengkapnya
                                </button></a>
                            </div>
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
                            <button class="page-link" @click="goToPage(pageNumber)"
                                :class="{ active: currentPage === pageNumber }" :style="{
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
                        judul_potensi: "Warga Morowali Curhat Rumah Rusak-Anak Ispa gegara Debu Batu Bara PT IMIP",
                        date: "26-12-2023",
                        deskripsi:
                            "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Duis sodales consequat dictum. Praesent fermentum blandit ipsum et ultricies. Nunc ultrices neque ac velit aliquet, in iaculis nisi pellentesque. Pengembangan pertanian organik di daerah dengan tanah subur.",
                        foto_potensi_desa: "gambar1.jpg",
                        sub_judul: "Deskripsi tambahan untuk berita desa 1",
                    },
                    {
                        id_potensi: 2,
                        judul_potensi: "KKP Berhasil Amankan Tiga Pelaku Bom Ikan di Morowali",
                        date: "28-04-2001",
                        deskripsi:
                            "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Duis sodales consequat dictum. Praesent fermentum blandit ipsum et ultricies. Nunc ultrices neque ac velit aliquet, in iaculis nisi pellentesque. Program pelatihan untuk pemberdayaan ekonomi wanita di desa.",
                        foto_potensi_desa: "gambar2.jpg",
                        sub_judul: "Deskripsi tambahan untuk potensi desa 2",
                    },
                    {
                        id_potensi: 3,
                        judul_potensi: "Polsek Palu Barat, Buru Pelaku Pencurian Sampai di Morowali",
                        date: "28-04-2001",
                        deskripsi:
                            "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Duis sodales consequat dictum. Praesent fermentum blandit ipsum et ultricies. Nunc ultrices neque ac velit aliquet, in iaculis nisi pellentesque. Dukungan pembinaan UKM di bidang kerajinan tangan.",
                        foto_potensi_desa: "gambar3.jpg",
                        sub_judul: "Deskripsi tambahan untuk potensi desa 3",
                    },
                    {
                        id_potensi: 4,
                        judul_potensi: "Semarak HUT Morowali ke 24, Pemkab Gelar Berbagai Lomba, Turnamen Sepakbola Cup 2023 Dibuka Sekab",
                        date: "28-04-2001",
                        deskripsi:
                            "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Duis sodales consequat dictum. Praesent fermentum blandit ipsum et ultricies. Nunc ultrices neque ac velit aliquet, in iaculis nisi pellentesque. Program edukasi dan pelatihan keterampilan untuk anak muda.",
                        foto_potensi_desa: "gambar4.jpg",
                        sub_judul: "Deskripsi tambahan untuk potensi desa 4",
                    },
                    {
                        id_potensi: 5,
                        judul_potensi: "Pemkab Morut kolaborasi Kemenkumham Sulteng gali potensi KIK",
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

/* Remove margin-top */
.mt-0 {
    margin-top: 0;
}

/* Set text color to white */
.text-white {
    color: white;
}

.text-blue {
    color: #003366 !important;
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
</style>