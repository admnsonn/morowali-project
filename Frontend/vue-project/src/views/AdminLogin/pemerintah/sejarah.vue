<template>
    <div class="kontainer-admin">
        <div class="kontainer">
            <div class="bartipis" />
            <div class="">
                <h1 class="teks-admin">Admin desa Bahomoleo</h1>
                <p class="teks-kabupaten">Kabupaten, Morowali</p>
            </div>
        </div>
    </div>

    <div class="container">
        <div class="row">
            <div class="col">
                <h3 class="title-warga">Data Sejarah Desa</h3>
                <p class="subtitle-warga">Management Content dan Layanan Kreatifitas</p>
            </div>
        </div>

        <!-- tabel -->
        <div class="bungkus-tabel">
            <div class="row">
                <div class="col">
                    <input v-model="searchKeyword" @input="filterData" type="text" class="form-control w-100 my-3"
                        placeholder="Search...">
                </div>
            </div>

            <!-- Tabel -->
            <div class="tabel-container">
                <div class="table-scroll">
                    <table class="tabel">
                        <thead>
                            <tr>
                                <th>No.</th>
                                <th>Sejarah</th>
                                <th>Action</th>
                            </tr>
                        </thead>
                        <tbody>
                            <tr v-for="(item, index) in displayedData" :key="index">
                                <td>{{ (currentPage - 1) * itemsPerPage + index + 1 }}</td>
                                <td>{{ item.sejarah }}</td>
                                <td>
                                    <router-link :to="`/pemerintahan/update-sejarah-desa/1`">
                                        <button type="button" class="btn btn-warning m-1">
                                            <!-- edit button -->
                                            <img src="../../../../src/assets/img/edit.svg" class="custom-icon" />
                                        </button>
                                    </router-link>
                                </td>
                            </tr>
                        </tbody>
                    </table>
                </div>
            </div>
            <div class="pagination">
                <button @click="prevPage()" :disabled="currentPage === 1">
                    Previous
                </button>
                <span> {{ currentPage }} dari {{ totalPages }}</span>
                <button @click="nextPage()" :disabled="currentPage === totalPages">
                    Next
                </button>
            </div>
        </div>
    </div>
</template>
  
<script>
import axios from "axios";
import Swal from "sweetalert2";

export default {
    data() {
        return {
            searchKeyword: "",
            tableData: [],
            currentPage: 1,
            itemsPerPage: 7, // Sesuaikan item table perhalaman
            sortDirection: "asc",
            filteredData: [],
        };
    },
    computed: {
        totalPages() {
            return Math.ceil(this.filteredData.length / this.itemsPerPage);
        },
        displayedData() {
            const startIndex = (this.currentPage - 1) * this.itemsPerPage;
            const endIndex = startIndex + this.itemsPerPage;
            return this.filteredData.slice(startIndex, endIndex);
        },
    },
    methods: {
        fetchData() {
            axios
                .get("http://localhost:8080/pemerintah/sejarah/1")
                .then(({ data }) => {
                    this.tableData = data.data;
                    this.filteredData = this.tableData; // Initialize filteredData
                })
                .catch((error) => {
                    console.error("Error in Axios POST request:", error);
                });
        },
        nextPage() {
            if (this.currentPage < this.totalPages) {
                this.currentPage++;
                // Fetch data for the next page if needed
            }
        },
        filterData() {
            this.filteredData = this.tableData.filter((item) => {
                return (
                    item.sejarah.toLowerCase().includes(this.searchKeyword.toLowerCase())
                );
            });
            this.currentPage = 1; // Reset halaman ke 1 setiap kali pencarian berubah
        },
        prevPage() {
            if (this.currentPage > 1) {
                this.currentPage--;
            }
        },
        getImageSource(base64String) {
            return `data:image/png;base64,${base64String}`;
        },
    },
    created() {
        this.fetchData(); // Get original data
    },
};
</script>
  
<style scoped>
.container {
    margin-top: 30px;
    margin-bottom: 50px;
    width: calc(100% - 100px);
}

.title-warga {
    font-size: 20px;
}

.subtitle-warga {
    font-size: 15px;
    color: #5e5e5e;
}

.btn-tambah {
    background-color: #003366;
    color: #fff;
    border: none;
    font-size: 14px;
    padding-top: 10%;
    padding-bottom: 10%;
}

.btn-tambah:hover {
    background-color: #003366;
    color: #fff;
    border: none;
}

.btn-search,
.btn-search:hover {
    padding-top: 12px;
    padding-bottom: 12px;
    background-color: #fff;
    border: 1px solid #8b8a8a;
    text-align: left;
    font-size: 14px;
}

.btn-light option {
    font-size: small;
    padding-left: 11px;
}

.dropdown-kategori {
    width: max-content;
    margin-left: 20px;
}

select {
    font-size: 14px;
}

h3 {
    font-weight: bold;
}
</style>
  