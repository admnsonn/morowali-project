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
        <h3 class="title-warga">Data Warga Desa</h3>
        <p class="subtitle-warga">Management Content dan Layanan Warga</p>
      </div>
      <div class="col-auto">
        <button type="button" class="btn btn-light btn-excel my-2" @click="exportToExcel">
          Export Excel
        </button>
      </div>
    </div>

    <div class="bungkus-tabel">
      <div class="row">
        <div class="col">
          <button type="button" class="btn btn-search w-100 my-2">
            <img src="../../../../src/assets/img/search.svg" class="me-2" /> Search...
          </button>
        </div>
        <div class="col-auto">
          <button type="button" class="btn btn-success btn-tambah my-2">
            <router-link to="/tambah-warga" class="nav-link router-link-underline">+ Tambah Data</router-link>
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
                  ID
                  <button type="button" class="btn btn-link" @click="sortById()">
                    <img src="../../../../src/assets/img/sort.svg" />
                  </button>
                </th>
                <th>
                  NIK
                  <button type="button" class="btn btn-link" @click="sortByNIK()">
                    <img src="../../../../src/assets/img/sort.svg" />
                  </button>
                </th>
                <th>Alamat</th>
                <th>Nama</th>
                <th>No HP</th>
                <th>Kartu Keluarga</th>
                <th>
                  Jenis Kelamin
                  <select v-model.number="selectedKategorijk" @change="filterByKategorijk" class="btn btn-light p-1 my-1">
                    <option value="0">All</option>
                    <option v-for="item in kategoriJK" :value="item.id_jenis_kelamin">
                      {{ item.jenis_kelamin }}</option>
                  </select>
                </th>
                <th>
                  Agama
                  <select v-model.number="selectedKategoriag" @change="filterByKategoriag" class="btn btn-light p-1 my-1">
                    <option value="0">All</option>
                    <option v-for="item in kategoriAG" :value="item.id_agama">
                      {{ item.Agama }}</option>
                  </select>
                </th>
                <th>
                  Pendidikan
                  <select v-model.number="selectedKategoripn" @change="filterByKategoripn" class="btn btn-light p-1 my-1">
                    <option value="0">All</option>
                    <option v-for="item in kategoriPN" :value="item.id_pendidikan">
                      {{ item.kategori_pendidikan }}</option>
                  </select>
                </th>
                <th>
                  Finansial
                  <select v-model.number="selectedKategorifn" @change="filterByKategorifn" class="btn btn-light p-1 my-1">
                    <option value="0">All</option>
                    <option v-for="item in kategoriFN" :value="item.id_financial">
                      {{ item.financial }}</option>
                  </select>
                </th>
                <th>Umur</th>
                <th>Action</th>
              </tr>
            </thead>
            <tbody>
              <tr v-for="(item, index) in displayedData" :key="index">
                <td>{{ item.id_pengguna }}</td>
                <td>{{ item.nik }}</td>
                <td>{{ item.alamat_pengguna }}</td>
                <td>{{ item.nama_lengkap }}</td>
                <td>{{ item.no_telp }}</td>
                <td>{{ item.kk }}</td>
                <td>{{ item.jenis_kelamin }}</td>
                <td>{{ item.nama_agama }}</td>
                <td>{{ item.nama_pendidikan }}</td>
                <td>{{ item.kategori_financial }}</td>
                <td>{{ item.umur }}</td>
                <td>
                  <router-link :to="`/detail-warga/${item.id_pengguna}`">
                    <button class="btn btn-info m-1">
                      <img src="../../../../src/assets/img/view.svg" class="custom-icon" />
                    </button>
                  </router-link>
                  <router-link :to="`/update-warga/${item.id_pengguna}`">
                    <button class="btn btn-warning m-1">
                      <img src="../../../../src/assets/img/edit.svg" class="custom-icon" />
                    </button>
                  </router-link>
                  <button type="button" @click.prevent="
                    deleteData(item.id_pengguna, item.nama_lengkap)
                    " class="btn btn-danger m-1">
                    <img src="../../../../src/assets/img/delete.svg" class="custom-icon" />
                  </button>
                </td>
              </tr>
            </tbody>
          </table>
        </div>
      </div>
      <div class="pagination">
        <button @click="prevPage" :disable="currentPage === 1">Previous</button>
        <span> {{ currentPage }} dari {{ totalPages }}</span>
        <button @click="nextPage" :enable="currentPage === totalPages">
          Next
        </button>
      </div>
    </div>
  </div>
</template>

<script>
import axios from "axios";
import Swal from "sweetalert2";
import * as XLSX from "xlsx";

export default {
  data() {
    return {
      tableData: [],
      currentPage: 1,
      selectedKategorijk: 0,
      selectedKategorifn: 0,
      selectedKategoriag: 0,
      selectedKategoripn: 0,
      selectedKategorist: 0,
      itemsPerPage: 7,
      sortDirection: "asc",
      filteredData: [],
      kategoriJK: [],
      kategoriFN: [],
      kategoriPN: [],
      kategoriAG: [],
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
    exportToExcel() {
      const ws = XLSX.utils.json_to_sheet(this.tableData);
      const wb = XLSX.utils.book_new();
      XLSX.utils.book_append_sheet(wb, ws, "Sheet1");

      XLSX.writeFile(wb, "warga_data.xlsx");
    },
    fetchKategoriJK() {
      axios
        .get("http://localhost:8080/warga/jk")
        .then(({ data }) => {
          this.kategoriJK = data.data;
          console.log(this.kategoriJK);
        })
        .catch((error) => {
          console.error("Error in Axios GET request:", error);
        });
    },

    fetchKategoriAG() {
      axios
        .get("http://localhost:8080/warga/agama")
        .then(({ data }) => {
          this.kategoriAG = data.data;
          console.log(this.kategoriAG);
        })
        .catch((error) => {
          console.error("Error in Axios GET request:", error);
        });
    },

    fetchKategoriPN() {
      axios
        .get("http://localhost:8080/warga/pendidikan")
        .then(({ data }) => {
          this.kategoriPN = data.data;
          console.log(this.kategoriPN);
        })
        .catch((error) => {
          console.error("Error in Axios GET request:", error);
        });
    },

    fetchKategoriFN() {
      axios
        .get("http://localhost:8080/warga/financial")
        .then(({ data }) => {
          this.kategoriFN = data.data;
          console.log(this.kategoriFN);
        })
        .catch((error) => {
          console.error("Error in Axios GET request:", error);
        });
    },

    fetchData() {
      const payload = { id_desa: localStorage.getItem("desa_id") };

      axios
        .post("http://localhost:8080/warga/list", payload)
        .then(({ data }) => {
          this.tableData = data.data;
          this.filteredData = this.tableData;
          this.filterByKategorijk();
          this.filterByKategorifn();
          this.filterByKategoriag();
          this.filterByKategoripn();
        })
        .catch((error) => {
          console.error("Error in Axios POST request:", error);
        });
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
          cancelButtonText: "Batal",
        });

        if (result.isConfirmed) {
          const response = await axios.delete(
            `http://localhost:8080/warga/delete/${id}`
          );
          if (response.data.status) {
            await Swal.fire(
              "Data berhasil dihapus!",
              response.data.message,
              "success"
            );
            this.fetchData();
          } else {
            await Swal.fire(
              "Data gagal dihapus.",
              response.data.message,
              "error"
            );
          }
        }
      } catch (error) {
        console.error("Error in Axios DELETE request:", error);
      }
    },
    sortById() {
      this.filteredData.sort((a, b) => a.id_pengguna - b.id_pengguna); // Sort by ID ascending
      // If you want to toggle ascending/descending order:
      this.filteredData.sort((a, b) =>
        this.sortDirection === "asc"
          ? a.id_pengguna - b.id_pengguna
          : b.id_pengguna - a.id_pengguna
      );

      this.sortDirection = this.sortDirection === "asc" ? "desc" : "asc"; // Toggle direction

      this.displayedData = this.filteredData.slice(startIndex, endIndex); // Recalculate displayedData
    },
    sortByNIK() {
      this.filteredData.sort((a, b) => a.nik.localeCompare(b.nik)); // Sort by judul alphabetically
      // Toggle ascending/descending (optional):
      this.filteredData.sort((a, b) =>
        this.sortDirection === "asc"
          ? a.nik.localeCompare(b.nik)
          : b.nik.localeCompare(a.nik)
      );

      this.sortDirection = this.sortDirection === "asc" ? "desc" : "asc"; // Toggle direction

      this.displayedData = this.filteredData.slice(startIndex, endIndex); // Recalculate displayedData
    },
    fetchDetailData() {
      const residentId = 1; // You can change this to the desired resident ID
      axios
        .get(`http://localhost:8080/warga/detail/${residentId}`)
        .then((res) => {
          this.detailWarga = res.data.data;
        })
        .catch((error) => {
          console.error("Error in Axios GET request:", error);
        });
    },
    filterByKategorijk() {
      this.filteredData = this.tableData.filter(
        (item) =>
          item.jk_id === this.selectedKategorijk ||
          this.selectedKategorijk === 0
      );
      this.currentPage = 1;
    },
    filterByKategoriag() {
      this.filteredData = this.tableData.filter(
        (item) =>
          item.agama_id === this.selectedKategoriag ||
          this.selectedKategoriag === 0
      );
      this.currentPage = 1;
    },
    filterByKategoripn() {
      this.filteredData = this.tableData.filter(
        (item) =>
          item.pendidikan_id === this.selectedKategoripn ||
          this.selectedKategoripn === 0
      );
      this.currentPage = 1;
    },
    filterByKategorifn() {
      this.filteredData = this.tableData.filter(
        (item) =>
          item.kategori_financial_id === this.selectedKategorifn ||
          this.selectedKategorifn === 0
      );
      this.currentPage = 1;
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
  created() {
    this.fetchData();
    this.fetchKategoriJK();
    this.fetchKategoriFN();
    this.fetchKategoriPN();
    this.fetchKategoriAG();
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

.btn-excel {
  background-color: #33b949;
  color: #ffffff;
  border: none;
  font-size: 14px;
  padding-top: 10%;
  padding-bottom: 10%;
}

.btn-excel:hover {
  background-color: #33b949;
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
</style>
