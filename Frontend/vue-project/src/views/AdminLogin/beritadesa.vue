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

  <div class="kontainer-data">
    <div>
      <h3 class="teks-h3">Data Berita Desa</h3>
      <p class="teks-p">Management Content dan Layanan warga</p>
    </div>

    <!-- tabel -->
    <div class="content-berita">
      <div class="row">
        <div class="col">
          <button type="button" class="btn search w-100 p-2 my-2">
            <img src="src/assets/img/search.svg" class="me-2" /> Search...
          </button>
        </div>
        <div class="col-auto">
          <button type="button" class="btn btn-success btn-blue p-3 my-2">
            + Tambah Data
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
                <button type="button" class="btn btn-link" @click="sortById()">
                  <img src="src/assets/img/sort.svg" />
                </button>
              </th>

              <th>
                Judul
                <button
                  type="button"
                  class="btn btn-link"
                  @click="sortByJudul()"
                >
                  <img src="src/assets/img/sort.svg" />
                </button>
              </th>
              <th>Sub-Judul</th>
              <th>Deskripsi</th>
              <th>Foto Berita</th>
              <th>
                Kategori
                <select v-model="selectedKategori" @change="filterByKategori">
                  <option value="">All</option>
                  <option value="Politik">Politik</option>
                  <option value="Teknologi">Teknologi</option>
                  <option value="Ekonomi">Ekonomi</option>
                </select>
              </th>
              <th>Action</th>
            </tr>
          </thead>
          <tbody>
            <tr v-for="(item, index) in displayedData" :key="index">
              <td>{{ item.id_berita }}</td>
              <td>{{ item.judul }}</td>
              <td>{{ item.sub_judul }}</td>
              <td>{{ item.deskripsi }}</td>
              <td>{{ item.foto_berita }}</td>
              <td>{{ item.kategori }}</td>
              <td>
                <button type="button" class="btn btn-primary m-1">
                  <img src="src/assets/img/view.svg" />
                </button>
                <button type="button" class="btn btn-warning">
                  <img src="src/assets/img/edit.svg" />
                </button>
                <button type="button" class="btn btn-danger m-1">
                  <img src="src/assets/img/delete.svg" />
                </button>
              </td>
            </tr>
          </tbody>
        </table>

        <!-- Pagination -->
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
  </div>
</template>

<script>
import axios from "axios";
export default {
  data() {
    return {
      tableData: [],
      currentPage: 1,
      itemsPerPage: 7, // Sesuaikan item table perhalaman
      selectedKategori: "",
      sortDirection: "asc",
      filteredData: [],
    };
  },
  computed: {
    totalPages() {
      return Math.ceil(this.filteredData.length / this.itemsPerPage);
    },
    displayedData() {
      // Start with filteredData
      const startIndex = (this.currentPage - 1) * this.itemsPerPage;
      const endIndex = startIndex + this.itemsPerPage;
      return this.filteredData.slice(startIndex, endIndex);
    },
  },
  methods: {
    fetchData() {
      const payload = { id_desa: localStorage.getItem("desa_id") };

      axios
        .post("http://localhost:8080/berita/list", payload)
        .then(({ data }) => {
          this.tableData = data.data;
          this.filteredData = this.tableData; // Initialize filteredData
          this.filterByKategori(); // Apply initial filter
        })
        .catch((error) => {
          console.error("Error in Axios POST request:", error);
        });
    },
    sortById() {
      this.filteredData.sort((a, b) => a.id_berita - b.id_berita); // Sort by ID ascending
      // If you want to toggle ascending/descending order:
      this.filteredData.sort((a, b) =>
        this.sortDirection === "asc"
          ? a.id_berita - b.id_berita
          : b.id_berita - a.id_berita
      );

      this.sortDirection = this.sortDirection === "asc" ? "desc" : "asc"; // Toggle direction

      this.displayedData = this.filteredData.slice(startIndex, endIndex); // Recalculate displayedData
    },

    sortByJudul() {
      this.filteredData.sort((a, b) => a.judul.localeCompare(b.judul)); // Sort by judul alphabetically
      // Toggle ascending/descending (optional):
      this.filteredData.sort((a, b) =>
        this.sortDirection === "asc"
          ? a.judul.localeCompare(b.judul)
          : b.judul.localeCompare(a.judul)
      );

      this.sortDirection = this.sortDirection === "asc" ? "desc" : "asc"; // Toggle direction

      this.displayedData = this.filteredData.slice(startIndex, endIndex); // Recalculate displayedData
    },
    nextPage() {
      if (this.currentPage < this.totalPages) {
        this.currentPage++;
        // Fetch data for the next page if needed
      }
    },
    prevPage() {
      if (this.currentPage >= this.totalPages) {
        this.currentPage--;
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
  },
  created() {
    this.fetchData(); // Get original data
  },
};
</script>

<style scoped>
th {
  vertical-align: middle;
}

.teks-h3 {
  font-weight: bold;
  font-size: large;
}

.teks-p {
  font-size: medium;
}

.teks-admin {
  font-family: "Poppins", sans-serif;
  font-weight: bold;
  font-size: 24px;
  margin: 0;
}

.teks-kabupaten {
  font-family: "Poppins", sans-serif;
  font-size: 16px;
  margin: 0;
}

.kontainer-data {
  height: auto;
  display: flex;
  flex-direction: column;
  margin-left: 80px;
  margin-top: 32px;
  margin-bottom: 46px;
  margin-right: 79px;
}
.bartipis {
  background-color: black;
  height: 100%;
  width: 3px;
  grid-row: span 2;
  top: 0;
  left: 0;
  align-self: center;
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

.btn-grey {
  background-color: #ffffff;
  color: #000;
  border: none;
  font-size: small;
}

.search,
.search:hover {
  text-align: left;
  background-color: #fff;
  border: 1px solid #000;
  border-radius: 8px;
  padding: 10%;
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

.content-berita {
  background-color: #eef1f3;
  padding: 20px;
  width: 100%;
  border-radius: 5px;
}
</style>
