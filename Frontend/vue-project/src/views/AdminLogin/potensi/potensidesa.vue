<template>
  <div class="kontainer-admin">
    <div class="kontainer">
      <div class="bartipis" />
      <div class="">
        <h1 class="teks-admin">Admin desa Bahomoleo</h1>
        <p class="teks-kabupaten">Kabupaten, Morowali</p>
      </div>
    </div>
    <div class="container-userbtn">
      <button class="btn user-button">
        User
        <svg
          xmlns="http://www.w3.org/2000/svg"
          width="16"
          height="16"
          fill="currentColor"
          class="bi bi-person-circle"
          viewBox="0 0 16 16"
        >
          <path d="M11 6a3 3 0 1 1-6 0 3 3 0 0 1 6 0" />
          <path
            fill-rule="evenodd"
            d="M0 8a8 8 0 1 1 16 0A8 8 0 0 1 0 8m8-7a7 7 0 0 0-5.468 11.37C3.242 11.226 4.805 10 8 10s4.757 1.225 5.468 2.37A7 7 0 0 0 8 1"
          />
        </svg>
      </button>
    </div>
  </div>

  <div class="container">
    <div class="row">
      <div class="col">
        <h3 class="title-warga">Data Potensi Desa</h3>
        <p class="subtitle-warga">Management Content dan Layanan Warga</p>
      </div>
    </div>

    <!-- tabel -->
    <div class="bungkus-tabel">
      <div class="wrapper-tabletop">
        <div class="col-auto">
          <button type="button" class="btn btn-tambah my-2">
            <router-link
              to="/potensi-desa"
              class="nav-link router-link-underline"
            >
              <svg
                xmlns="http://www.w3.org/2000/svg"
                width="16"
                height="16"
                fill="currentColor"
                class="bi bi-trash-fill"
                viewBox="0 0 16 16"
              >
                <path
                  d="M2.5 1a1 1 0 0 0-1 1v1a1 1 0 0 0 1 1H3v9a2 2 0 0 0 2 2h6a2 2 0 0 0 2-2V4h.5a1 1 0 0 0 1-1V2a1 1 0 0 0-1-1H10a1 1 0 0 0-1-1H7a1 1 0 0 0-1 1zm3 4a.5.5 0 0 1 .5.5v7a.5.5 0 0 1-1 0v-7a.5.5 0 0 1 .5-.5M8 5a.5.5 0 0 1 .5.5v7a.5.5 0 0 1-1 0v-7A.5.5 0 0 1 8 5m3 .5v7a.5.5 0 0 1-1 0v-7a.5.5 0 0 1 1 0"
                />
              </svg>
              Delete
            </router-link>
          </button>
        </div>
        <div class="container-searchbar">
          <button type="button" class="btn btn-search">
            <img src="src/assets/img/search.svg" class="me-2" /> Search...
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
                    class="btn btn-link m-1"
                    @click="sortById()"
                  >
                    <img src="src/assets/img/sort.svg" class="custom-icon" />
                  </button>
                </th>

                <th>
                  Judul
                  <button
                    type="button"
                    class="btn btn-link m-1"
                    @click="sortByNama()"
                  >
                    <img src="src/assets/img/sort.svg" class="custom-icon" />
                  </button>
                </th>
                <th>Deskripsi</th>
                <th>Sub Judul</th>
                <th>Foto Potensi</th>
                <th>Action</th>
              </tr>
            </thead>
            <tbody>
              <tr v-for="(item, index) in displayedData" :key="index">
                <td>AWAITING API</td>
                <td>AWAITING API</td>
                <td>AWAITING API</td>
                <td>AWAITING API</td>
                <td>AWAITING API</td>
                <td>
                  <!-- <button type="button" class="btn btn-primary m-1">
                    <img src="src/assets/img/view.svg" />
                  </button> -->
                  <button type="button" class="btn btn-primary button-detail">
                    <svg
                      xmlns="http://www.w3.org/2000/svg"
                      width="16"
                      height="16"
                      fill="currentColor"
                      class="bi bi-card-list"
                      viewBox="0 0 16 16"
                    >
                      <path
                        d="M14.5 3a.5.5 0 0 1 .5.5v9a.5.5 0 0 1-.5.5h-13a.5.5 0 0 1-.5-.5v-9a.5.5 0 0 1 .5-.5zm-13-1A1.5 1.5 0 0 0 0 3.5v9A1.5 1.5 0 0 0 1.5 14h13a1.5 1.5 0 0 0 1.5-1.5v-9A1.5 1.5 0 0 0 14.5 2z"
                      />
                      <path
                        d="M5 8a.5.5 0 0 1 .5-.5h7a.5.5 0 0 1 0 1h-7A.5.5 0 0 1 5 8m0-2.5a.5.5 0 0 1 .5-.5h7a.5.5 0 0 1 0 1h-7a.5.5 0 0 1-.5-.5m0 5a.5.5 0 0 1 .5-.5h7a.5.5 0 0 1 0 1h-7a.5.5 0 0 1-.5-.5m-1-5a.5.5 0 1 1-1 0 .5.5 0 0 1 1 0M4 8a.5.5 0 1 1-1 0 .5.5 0 0 1 1 0m0 2.5a.5.5 0 1 1-1 0 .5.5 0 0 1 1 0"
                      />
                    </svg>
                    Detail
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
    async deleteData(id, nama_umkm) {
      try {
        const result = await Swal.fire({
          title: `Hapus data ${nama_umkm}?`,
          text: "Data yang sudah dihapus tidak dapat dikembalikan lagi.",
          icon: "warning",
          showCancelButton: true,
          confirmButtonColor: "#C03221",
          cancelButtonColor: "#4F4F4F",
          confirmButtonText: "Hapus",
          cancelButtonText: "Batal",
        });

        if (result.isConfirmed) {
          const response = await axios
            .delete
            // `http://localhost:8080/berita/delete/${id}`
            ();
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
      this.filteredData.sort((a, b) => a.id_umkm - b.id_umkm); // Sort by ID ascending
      // If you want to toggle ascending/descending order:
      this.filteredData.sort((a, b) =>
        this.sortDirection === "asc"
          ? a.id_umkm - b.id_umkm
          : b.id_umkm - a.id_umkm
      );

      this.sortDirection = this.sortDirection === "asc" ? "desc" : "asc"; // Toggle direction

      this.displayedData = this.filteredData.slice(startIndex, endIndex); // Recalculate displayedData
    },

    sortByNama() {
      this.filteredData.sort((a, b) => a.nama_umkm.localeCompare(b.nama_umkm)); // Sort by nama alphabetically
      // Toggle ascending/descending (optional):
      this.filteredData.sort((a, b) =>
        this.sortDirection === "asc"
          ? a.nama_umkm.localeCompare(b.nama_umkm)
          : b.nama_umkm.localeCompare(a.nama_umkm)
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
      if (this.currentPage > 1) {
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
  background-color: #b50511;
  color: #fff;
  border: none;
  font-size: 14px;
  padding-top: 10%;
  padding-bottom: 10%;
  height: 100%;
}

.btn-tambah:hover {
  background-color: #c41e1e;
  color: #fff;
  border: none;
}

.btn-excel {
  background-color: #33b949;
  color: #000;
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
  width: 100%;
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

.wrapper-tabletop {
  display: flex;
  flex-direction: row;
  justify-content: space-between;
  width: 100%;
  align-items: center;
}

.container-searchbar {
  width: 23%;
}

.kontainer-admin {
  justify-content: space-between;
  padding-right: 11rem;
}

.container-userbtn {
  height: 100%;
  display: flex;
  align-items: center;
}

.user-button {
  background-color: #003366;
  color: white;
  height: 50%;
  padding-left: 2rem;
  padding-right: 2rem;
  display: flex;
  flex-direction: row;
  justify-items: center;
  align-items: center;
  gap: 10px;
  font-weight: 500;
}

.button-detail {
  width: 100%;
  font-weight: bold;
}
</style>
