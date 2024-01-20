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
            <div class="col-auto">
                <button type="button" class="btn btn-light btn-grey p-3 my-2">
                    Import Excel
                </button>
            </div>
        </div>

        <div class="content-warga">
            <div class="row">
                <div class="col">
                    <button type="button" class="btn search w-100 p-2 my-2">
                        <img src="src/assets/img/search.svg" class="me-2" /> Search...
                    </button>
                </div>
                <div class="col-auto">
                    <button type="button" class="btn btn-success btn-blue p-3 my-2">
                        <router-link to="/tambah-warga" class="nav-link router-link-underline">+ Tambah Data</router-link>
                    </button>
                </div>
            </div>

            <!-- Tabel -->
            <div class="table-container">
                <table class="table">
                    <thead>
                        <tr>
                            <th>
                                ID
                                <button type="button" class="btn btn-link">
                                    <img src="src/assets/img/sort.svg" />
                                </button>
                            </th>
                            <th>
                                Nama Wisata
                                <button type="button" class="btn btn-link">
                                    <img src="src/assets/img/sort.svg" />
                                </button>
                            </th>
                            <th>Alamat</th>
                            <th>Foto Wisata</th>
                            <th>No Telp</th>
                            <th>Action</th>
                        </tr>
                    </thead>
                    <tbody>
                        <tr v-for="(item, index) in displayedData" :key="index">
                            <td>{{ item.id_wisata }}</td>
                            <td>{{ item.nama_wisata }}</td>
                            <td>{{ item.alamat }}</td>
                            <td>{{ item.foto_wisata }}</td>
                            <td>{{ item.no_telp }}</td>
                            <td>
                                <!-- <button type="button" class="btn btn-primary m-1">
                                    <img src="src/assets/img/view.svg" />
                                </button> -->
                                <button type="button" class="btn btn-warning">
                                    <img src="src/assets/img/edit.svg" />
                                </button>
                                <button type="button" @click.prevent="deleteData(item.id_pengguna, item.nama_lengkap)"
                                    class="btn btn-danger m-1">
                                    <img src="src/assets/img/delete.svg" />
                                </button>
                            </td>
                        </tr>
                    </tbody>
                </table>

                <!-- Pagination -->
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
    </div>
</template>
  
<script>
import axios from "axios";
import Swal from "sweetalert2";

export default {
    data() {
        return {
            tableData: [],
            currentPage: 1,
            selectedKategori: "",
            itemsPerPage: 7, // Sesuaikan item table perhalaman
            sortDirection: "asc",
        };
    },
    computed: {
        totalPages() {
            return Math.ceil(this.tableData.length / this.itemsPerPage);
        },
        displayedData() {
            const startIndex = (this.currentPage - 1) * this.itemsPerPage;
            const endIndex = startIndex + this.itemsPerPage;
            return this.tableData.slice(startIndex, endIndex);
        },
    },
    methods: {
        async fetchData() {
            try {
                const response = await axios.get("http://localhost:8080/wisata/1");
                this.tableData = response.data.data;
            } catch (error) {
                console.error("Error in Axios GET request:", error);
            }
        },
        async deleteData(id, nama) {
            try {
                const result = await Swal.fire({
                    title: `Hapus data ${nama}?`,
                    text: "Data yang sudah dihapus tidak dapat dikembalikan lagi.",
                    icon: "warning",
                    showCancelButton: true,
                    confirmButtonColor: "#C03221",
                    cancelButtonColor: "#4F4F4F",
                    confirmButtonText: "Hapus",
                    cancelButtonText: "Batal"
                });

                if (result.isConfirmed) {
                    const response = await axios.delete(`http://localhost:8080/warga/delete/${id}`);
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

        filterByKategori() {
            this.filteredData = this.tableData.filter(
                (item) =>
                    item.kategori === this.selectedKategori ||
                    this.selectedKategori === ""
            );
            this.displayedData = this.filteredData.slice(
                (this.currentPage - 1) * this.itemsPerPage,
                this.currentPage * this.itemsPerPage
            );
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
    },
    mounted() {
        this.fetchData();
    },
};
</script>
  
<style scoped>
h3 {
    font-weight: 700;
}

.container {
    margin-top: 30px;
    margin-bottom: 50px;
    width: calc(100% - 100px);
}

.table {
    border-collapse: collapse;
    margin-top: 32px;
    border-radius: 8px;
    overflow: hidden;
}

.table th,
.table td {
    border-bottom: 1px solid #eef1f3;
    padding: 20px;
    text-align: center;
}

.table th {
    background-color: #eef1f3;
}

.table tbody tr:nth-child(even) {
    background-color: #f2f2f2;
}

.pagination {
    margin-right: 30px;
    margin-top: 20px;
    display: flex;
    justify-content: flex-end;
    /* Memindahkan pagination ke sebelah kanan */
    align-items: center;
}

.pagination button {
    margin: 0 10px;
    padding: 8px;
    cursor: pointer;
    border: 1px solid #ddd;
    background-color: #003366;
    color: white;
    border: none;
    padding: 8px 12px;
    margin: 0 5px;
    cursor: pointer;
    border-radius: 5px;
}

.pagination button:disabled {
    cursor: not-allowed;
    color: #aaa;
    border: 1px solid #ccc;
    background-color: #ccc;
    cursor: not-allowed;
}

.pagination span {
    margin: 0 10px;
    font-weight: bold;
}

.table-container {
    overflow-x: auto;
}

.btn-blue {
    background-color: #003366;
    color: #fff;
    border: none;
}

.btn-blue:hover {
    background-color: #003366;
    color: #fff;
    border: none;
}

.btn-grey {
    background-color: #ffffff;
    color: #000;
    border: none;
    font-size: small;
}

.btn-excel {
    background-color: #33B949;
}

.btn-excel:hover {
    background-color: #33B949;
}

.search,
.search:hover {
    text-align: left;
    background-color: #fff;
    border: 1px solid #000;
    border-radius: 8px;
    padding: 10%;
}

.content-warga {
    background-color: #eef1f3;
    padding: 20px;
    width: 100%;
    border-radius: 5px;
}

.title-warga {
    font-size: 20px;
}

.subtitle-warga {
    font-size: 15px;
    color: #5e5e5e;
}
</style>
  