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
        <h3 class="title-warga">Data Berita Desa</h3>
        <p class="subtitle-warga">Management Content dan Layanan Warga</p>
      </div>
    </div>

    <!-- tabel -->
    <div class="bungkus-tabel">
      <div class="row">
        <div class="col">
          <button type="button" class="btn btn-search w-100 my-2">
            <img src="src/assets/img/search.svg" class="me-2" /> Search...
          </button>
        </div>
        <div class="col-auto">
          <button type="button" class="btn btn-success btn-tambah my-2">
            <router-link
              to="/tambah-berita"
              class="nav-link router-link-underline"
              >+ Tambah Data</router-link
            >
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
                  <button
                    type="button"
                    class="btn btn-link"
                    @click="sortById()"
                  >
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
                  <select
                    v-model="selectedKategori"
                    @change="filterByKategori"
                    class="btn btn-light btn-grey p-1 my-2"
                  >
                    <option value="">All</option>
                    <option value="Politik">Politik</option>
                    <option value="Teknologi">Ekonomi</option>
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
                  <!-- <button type="button" class="btn btn-primary m-1">
                  <img src="src/assets/img/view.svg" />
                </button> -->
                  <router-link :to="`/update-berita/${item.id_berita}`">
                    <button type="button" class="btn btn-warning m-1">
                      <!-- edit button -->
                      <img src="src/assets/img/edit.svg" class="custom-icon" />
                    </button>
                  </router-link>
                  <button
                    type="button"
                    @click.prevent="deleteData(item.id_berita, item.judul)"
                    class="btn btn-danger m-1"
                  >
                    <!-- delete button -->
                    <img src="src/assets/img/delete.svg" class="custom-icon" />
                  </button>
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
    async deleteData(id, judul) {
      try {
        const result = await Swal.fire({
          title: `Hapus data ${judul}?`,
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
            `http://localhost:8080/berita/delete/${id}`
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
</style>
