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
                <h3 class="title-warga">Data Wisata Desa</h3>
                <p class="subtitle-warga">Management Content dan Layanan Warga</p>
            </div>
        </div>

        <div class="bungkus-tabel">
            <div class="row">
                <div class="col">
                    <input v-model="searchKeyword" @input="filterData" type="text" class="form-control w-100 my-3"
                        placeholder="Search...">
                </div>
                <div class="col-auto">
                    <button type="button" class="btn btn-success btn-tambah my-2">
                        <router-link to="/tambah-wisata" class="nav-link router-link-underline">+ Tambah Data</router-link>
                    </button>
                </div>
            </div>

            <!-- Tabel -->
            <div class="tabel-container">
                <div class="table-scroll">
                    <table class="tabel">
                        <thead>
                            <tr>
                                <th>
                                    No.
                                </th>
                                <th>
                                    Nama Wisata
                                    <button type="button" class="btn btn-link" @click="sortByWisata()">
                                        <img src="../../../../src/assets/img/sort.svg" />
                                    </button>
                                </th>
                                <th>Alamat</th>
                                <th>Foto Wisata</th>
                                <th class="">
                                    Kategori:
                                    <select v-model.number="selectedKategori" @change="filterByKategori"
                                        class="btn btn-light btn-grey p-1 my-2 dropdown-kategori">
                                        <option value="0">All</option>
                                        <option v-for="item in kategoriList" :value="item.id_wisata_kategori">
                                            {{ item.wisata_kategori }}
                                        </option>
                                    </select>
                                </th>
                                <th>No Telp</th>
                                <th>Action</th>
                            </tr>
                        </thead>
                        <tbody>
                            <tr v-for="(item, index) in displayedData" :key="index">
                                <td>{{ (currentPage - 1) * itemsPerPage + index + 1 }}</td>
                                <td>{{ item.nama_wisata }}</td>
                                <td>{{ item.alamat }}</td>
                                <td>
                                    <img class="td-foto" :src="`data:image/png;base64,${item.foto_wisata}`"
                                        alt="foto wisata" />
                                </td>
                                <td>{{ item.kategori }}</td>
                                <td>{{ item.no_telp }}</td>
                                <td>
                                    <router-link :to="`/detail-wisata/${item.id_wisata}`">
                                        <button class="btn btn-info m-1">
                                            <img src="../../../../src/assets/img/view.svg" class="custom-icon" />
                                        </button>
                                    </router-link>
                                    <router-link :to="`/update-wisata/${item.id_wisata}`">
                                        <button type="button" class="btn btn-warning m-1">
                                            <!-- edit button -->
                                            <img src="../../../../src/assets/img/edit.svg" class="custom-icon" />
                                        </button>
                                    </router-link>
                                    <button type="button" @click.prevent="deleteData(item.id_wisata, item.nama_wisata)"
                                        class="btn btn-danger m-1">
                                        <img src="../../../../src/assets/img/delete.svg" class="custom-icon" />
                                    </button>
                                </td>
                            </tr>
                        </tbody>
                    </table>
                </div>
            </div>
            <div class="pagination">
                <button @click="prevPage" :disabled="currentPage === 1">
                    Previous
                </button>
                <span> {{ currentPage }} dari {{ totalPages }}</span>
                <button @click="nextPage" :disabled="currentPage === totalPages">
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
            selectedKategori: 0,
            itemsPerPage: 7, // Sesuaikan item table perhalaman
            sortDirection: "asc",
            filteredData: [],
            kategoriList: [],
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
            const payload = { id_desa: localStorage.getItem("desa_id") };
            axios
                .post("http://localhost:8080/wisata/list", payload)
                .then(({ data }) => {
                    this.tableData = data.data;
                    this.filteredData = this.tableData;
                    this.filterByKategori();
                })
                .catch((error) => {
                    console.error("Error in Axios POST request:", error);
                });
        },
        fetchKategori() {
            axios
                .get("http://localhost:8080/wisata/kategori_wisata")
                .then(({ data }) => {
                    this.kategoriList = data.data;
                })
                .catch((error) => {
                    console.error("Error in Axios GET request:", error);
                });
        },
        async deleteData(id, nama_wisata) {
            try {
                const result = await Swal.fire({
                    title: `Hapus data ${nama_wisata}?`,
                    text: "Data yang sudah dihapus tidak dapat dikembalikan lagi.",
                    icon: "warning",
                    showCancelButton: true,
                    confirmButtonColor: "#C03221",
                    cancelButtonColor: "#4F4F4F",
                    confirmButtonText: "Hapus",
                    cancelButtonText: "Batal"
                });

                if (result.isConfirmed) {
                    const response = await axios.delete(`http://localhost:8080/wisata/delete_wisata/${id}`);
                    if (response.data.status) {
                        await Swal.fire("Data berhasil dihapus!", response.data.message, "success");
                        this.fetchData();
                    } else {
                        await Swal.fire("Data gagal dihapus.", response.data.message, "error");
                    }
                }
            } catch (error) {
                console.error("Error in Axios DELETE request:", error);
            }
        },
        filterData() {
            this.filteredData = this.tableData.filter((item) => {
                return (
                    item.nama_wisata.toLowerCase().includes(this.searchKeyword.toLowerCase()) ||
                    item.alamat.toString().includes(this.searchKeyword) ||
                    item.no_telp.toString().includes(this.searchKeyword) ||
                    item.kategori.toLowerCase().includes(this.searchKeyword.toLowerCase())
                );
            });
            this.currentPage = 1; // Reset halaman ke 1 setiap kali pencarian berubah
        },
        sortByWisata() {
            this.filteredData.sort((a, b) => a.nama_wisata.localeCompare(b.nama_wisata)); // Sort by judul alphabetically
            // Toggle ascending/descending (optional):
            this.filteredData.sort((a, b) =>
                this.sortDirection === "asc"
                    ? a.nama_wisata.localeCompare(b.nama_wisata)
                    : b.nama_wisata.localeCompare(a.nama_wisata)
            );

            this.sortDirection = this.sortDirection === "asc" ? "desc" : "asc"; // Toggle direction

            this.displayedData = this.filteredData.slice(startIndex, endIndex); // Recalculate displayedData
        },
        prevPage() {
            if (this.currentPage > 1) {
                this.currentPage--;
            }
        },
        nextPage() {
            if (this.currentPage < this.totalPages) {
                this.currentPage++;
            }
        },
        filterByKategori() {
            this.filteredData = this.tableData.filter(
                (item) =>
                    item.id_kategori === this.selectedKategori ||
                    this.selectedKategori === 0
            );
            this.currentPage = 1; // Reset pagination
        },
    },
    created() {
        this.fetchKategori();
        this.fetchData();
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

select {
    font-size: 14px;
}

h3 {
    font-weight: bold;
}

.td-foto {
    max-width: 100%;
    height: auto;
    margin: 10px; 
}
</style>
  