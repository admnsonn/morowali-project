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
                <div v-for="item in paginatedData" :key="item.id_berita"
                    class="border rounded row align-items-center pt-3 pb-3 mb-3 with-shadow"
                    :class="{ 'with-shadow': isHovered }" @mouseenter="addShadow" @mouseleave="removeShadow">
                    <!-- Display the data here -->

                    <div class="container mt-4">
                        <div class="row">
                            <div class="col-md-6">
                                <!-- Gambar berita -->
                                <img :src="`data:image/png;base64,${item.foto_berita}`" alt="Gambar Berita"
                                    class="img-fluid">
                            </div>
                            <div class="col-md-6">
                                <div class="d-none d-md-block">
                                    <h2 class="mb-2 warna-judul-berita">
                                        {{ item.judul }}
                                    </h2>
                                    <br>
                                </div>
                                <h5 class="mb-4 sub-berita">{{ item.sub_judul }}</h5>
                                <!-- <h5 class="mb-4 item-berita">{{ item.date }}</h5> -->
                                <p class="mb-4">{{ item.deskripsi }}</p>
                                <a href="/contentberita">
                                    <button class="btn btn-secondary" @click="showDetail(item)">
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
import axios from 'axios';

export default {
    data() {
        return {
            id_desa: "",
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
        async fetchBerita() {
            try {
                const response = await axios.post("http://localhost:8080/berita/list", { id_desa: this.id_desa });
                if (response.data.status) {
                    this.latestData = response.data.data;
                    this.showLatestData = true;
                    // Set halaman pertama ketika data diambil
                    this.paginatedData = this.latestData.slice(0, this.itemsPerPage);
                } else {
                    console.log("Data Kosong atau Terjadi Kesalahan")
                }
            } catch (error) {
                console.error('Error fetching data:', error);
            }
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

    created() {
        this.id_desa = localStorage.getItem("id_desa");
        this.fetchBerita();
    },
};
</script>

<style scoped>
.bg-hero {
    background-color: #003366;
}

.text-white {
    color: white;
    font-weight: bold;
    text-shadow: 2px 2px #252525;
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

.btn-secondary {
    background-color: #003366;
    border: none;
    padding: 10px 30px;
    font-family: 'Poppins', sans-serif;
    font-weight: 600;
    font-size: 14px;
    cursor: pointer;
    border-radius: 10px;
}

.sub-berita {
    font-weight: bold;
}

.item-berita {
    color: #003366;
    font-size: 18px;
    font-weight: bold;
}

.warna-judul-berita {
    font-weight: bold;
    color: #003366 !important;
}

.with-shadow {
    transition: box-shadow 0.3s ease-in-out;
    box-shadow: none;
}

.with-shadow:hover,
.with-shadow.active {
    box-shadow: 0 0 10px 0 rgba(0, 51, 102, 0.5);
}
</style>