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
        <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" fill="currentColor" class="bi bi-person-circle"
          viewBox="0 0 16 16">
          <path d="M11 6a3 3 0 1 1-6 0 3 3 0 0 1 6 0" />
          <path fill-rule="evenodd"
            d="M0 8a8 8 0 1 1 16 0A8 8 0 0 1 0 8m8-7a7 7 0 0 0-5.468 11.37C3.242 11.226 4.805 10 8 10s4.757 1.225 5.468 2.37A7 7 0 0 0 8 1" />
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
      <div class="row">
        <div class="col">
          <input v-model="searchKeyword" @input="filterData" type="text" class="form-control w-100 my-3"
            placeholder="Search...">
        </div>
        <div class="col-auto">
          <button type="button" class="btn btn-success btn-tambah my-2">
            <router-link to="/tambah-umkm" class="nav-link router-link-underline">+ Tambah Data</router-link>
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
                  Judul
                  <button type="button" class="btn btn-link m-1" @click="sortByNama()">
                    <img src="../../../../src/assets/img/sort.svg" class="custom-icon" />
                  </button>
                </th>
                <th>Sub Judul</th>
                <th>Deskripsi</th>
                <th>Foto Potensi</th>
                <th>Action</th>
              </tr>
            </thead>
            <tbody>
              <tr v-for="(item, index) in displayedData" :key="index">
                <td>{{ (currentPage - 1) * itemsPerPage + index + 1 }}</td>
                <td>{{ item.judul_potensi }}</td>
                <td>{{ item.sub_judul }}</td>
                <td>{{ item.deskripsi }}</td>
                <td>
                  <img class="td-foto" :src="`data:image/png;base64,${item.foto_potensi_desa}`" alt="foto_potensi"
                    height="75" width="100" />
                </td>
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
                  <button type="button" @click.prevent="deleteData(item.id_potensi, item.judul_potensi)"
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
      // Start with filteredData
      const startIndex = (this.currentPage - 1) * this.itemsPerPage;
      const endIndex = startIndex + this.itemsPerPage;
      return this.filteredData.slice(startIndex, endIndex);
    },
  },
  methods: {
    fetchData() {
      axios
        .get("http://localhost:8080/potensi_desa/list/1")
        .then(({ data }) => {
          this.tableData = data.data;
          this.filteredData = this.tableData; // Initialize filteredData
        })
        .catch((error) => {
          console.error("Error in Axios POST request:", error);
        });
    },
    async deleteData(id, judul_potensi) {
      try {
        const result = await Swal.fire({
          title: `Hapus data ${judul_potensi}?`,
          text: "Data yang sudah dihapus tidak dapat dikembalikan lagi.",
          icon: "warning",
          showCancelButton: true,
          confirmButtonColor: "#C03221",
          cancelButtonColor: "#4F4F4F",
          confirmButtonText: "Hapus",
          cancelButtonText: "Batal",
        });

        if (result.isConfirmed) {
          const response = await axios.delete(`http://localhost:8080/potensi_desa/delete_potensi/${id}`);
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
    filterData() {
      this.filteredData = this.tableData.filter((item) => {
        return (
          item.judul_potensi.toLowerCase().includes(this.searchKeyword.toLowerCase()) ||
          item.deskripsi.toString().includes(this.searchKeyword) ||
          item.sub_judul.toString().includes(this.searchKeyword.toLowerCase())
        );
      });
      this.currentPage = 1; // Reset halaman ke 1 setiap kali pencarian berubah
    },
    sortByNama() {
      this.filteredData.sort((a, b) => a.judul_potensi.localeCompare(b.judul_potensi)); // Sort by nama alphabetically
      // Toggle ascending/descending (optional):
      this.filteredData.sort((a, b) =>
        this.sortDirection === "asc"
          ? a.judul_potensi.localeCompare(b.judul_potensi)
          : b.judul_potensi.localeCompare(a.judul_potensi)
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
  background-color: #003366;
  color: #fff;
  border: none;
  font-size: 14px;
  padding-top: 10%;
  padding-bottom: 10%;
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
  height: 90%;
  display: flex;
  align-items: center;
  margin-right: -13%;
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
