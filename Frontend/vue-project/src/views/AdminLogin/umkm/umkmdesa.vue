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
        <h3 class="title-warga">Data UMKM Desa</h3>
        <p class="subtitle-warga">Management Content dan Layanan Warga</p>
      </div>
    </div>

    <!-- tabel -->
    <div class="bungkus-tabel">
      <div class="row">
        <div class="col">
          <button type="button" class="btn btn-search w-100 my-2">
            <img src="../../../../src/assets/img/search.svg" class="me-2" />
            Search...
          </button>
        </div>
        <div class="col-auto">
          <button type="button" class="btn btn-success btn-tambah my-2">
            <router-link
              to="/tambah-umkm"
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
                    <img
                      src="../../../../src/assets/img/sort.svg"
                      class="custom-icon"
                    />
                  </button>
                </th>

                <th>
                  Nama UMKM
                  <button
                    type="button"
                    class="btn btn-link"
                    @click="sortByNamaUmkm()"
                  >
                    <img
                      src="../../../../src/assets/img/sort.svg"
                      class="custom-icon"
                    />
                  </button>
                </th>
                <th>Alamat UMKM</th>
                <th>No. Telp</th>
                <th>Foto UMKM</th>
                <th class="">
                  Kategori:
                  <select
                    v-model.number="selectedKategori"
                    @change="filterByKategori"
                    class="btn btn-light btn-grey p-1 my-2 dropdown-kategori"
                  >
                    <option value="0">All</option>
                    <option
                      v-for="item in kategoriList"
                      :value="item.id_kategori_umkm"
                    >
                      {{ item.umkm_kategori }}
                    </option>
                  </select>
                </th>
                <th>Action</th>
              </tr>
            </thead>
            <tbody>
              <tr v-for="(item, index) in displayedData" :key="index">
                <td>{{ item.id_umkm }}</td>
                <td>{{ item.nama_umkm }}</td>
                <td>{{ item.alamat }}</td>
                <td>{{ item.no_telp_umkm }}</td>
                <!-- <td>{{ item.foto_umkm }}</td> -->
                <td>
                  <img
                    class="td-foto"
                    :src="`data:image/png;base64,${item.foto_umkm}`"
                    alt="foto berita"
                    height="75"
                    width="100"
                  />
                </td>
                <td>{{ item.kategori_umkm }}</td>
                <td>
                  <router-link :to="`/update-umkm/${item.id_umkm}`">
                    <button type="button" class="btn btn-warning m-1">
                      <!-- edit button -->
                      <img
                        src="../../../../src/assets/img/edit.svg"
                        class="custom-icon"
                      />
                    </button>
                  </router-link>
                  <button
                    type="button"
                    @click.prevent="deleteData(item.id_umkm, item.nama_umkm)"
                    class="btn btn-danger m-1"
                  >
                    <!-- delete button -->
                    <img
                      src="../../../../src/assets/img/delete.svg"
                      class="custom-icon"
                    />
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
      selectedKategori: 0,
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
    fetchKategori() {
      axios
        .get("http://localhost:8080/umkm/umkm_kategori")
        .then(({ data }) => {
          this.kategoriList = data.data;
        })
        .catch((error) => {
          console.error("Error in Axios GET request:", error);
        });
    },

    fetchData() {
      const payload = { id_desa: localStorage.getItem("desa_id") };

      axios
        .post("http://localhost:8080/umkm/list", payload)
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
          const response = await axios.delete(
            `http://localhost:8080/umkm/delete_umkm/${id}`
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

    sortByNamaUmkm() {
      this.filteredData.sort((a, b) => a.nama_umkm.localeCompare(b.nama_umkm)); // Sort by nama_umkm alphabetically
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
          item.kategori_umkm_id === this.selectedKategori ||
          this.selectedKategori === 0
      );
      this.currentPage = 1; // Reset pagination
    },
  },
  created() {
    this.fetchData(); // Get original data
    this.fetchKategori();
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

.td-foto {
  border-radius: 0.375rem;
}

select {
  font-size: 14px;
}

h3 {
  font-weight: bold;
}
</style>
