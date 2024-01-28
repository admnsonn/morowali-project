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
        <h3 class="title-warga">Data Kreatifitas Desa</h3>
        <p class="subtitle-warga">Management Content dan Layanan Kreatifitas</p>
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
            <img src="../../../../src/assets/img/search.svg" class="me-2" />
            Search...
          </button>
        </div>
      </div>

      <!-- Tabel -->
      <div class="tabel-container">
        <div class="table-scroll">
          <table class="tabel">
            <thead>
              <tr>
                <th>No.</th>
                <th>
                  Judul Kreatifitas
                  <button
                    type="button"
                    class="btn btn-link"
                    @click="sortByJudul()"
                  >
                    <img src="../../../../src/assets/img/sort.svg" />
                  </button>
                </th>
                <th>Total Pengunjung</th>
                <th>Foto Kreatifitas</th>
                <th>Action</th>
              </tr>
            </thead>
            <tbody>
              <tr v-for="(item, index) in displayedData" :key="index">
                <td>{{ (currentPage - 1) * itemsPerPage + index + 1 }}</td>
                <td>{{ item.judul_kreatifitas }}</td>
                <td>{{ item.total_pengunjung }}</td>
                <td>
                  <img
                    class="td-foto"
                    :src="`data:image/png;base64,${item.foto_kreatifitas}`"
                    alt="foto kreatifitas"
                    height="75"
                    width="100"
                  />
                </td>
                <td>
                  <router-link
                    :to="`/detail-kreatifitas/${item.id_kreatifitas}`"
                  >
                    <button class="btn btn-info m-1">
                      <img
                        src="../../../../src/assets/img/view.svg"
                        class="custom-icon"
                      />
                    </button>
                  </router-link>
                  <router-link
                    :to="`/update-kreatifitas/${item.id_kreatifitas}`"
                  >
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
                    @click.prevent="
                      deleteData(item.id_kreatifitas, item.judul_kreatifitas)
                    "
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
        .get("http://localhost:8080/kreatifitas/list/1")
        .then(({ data }) => {
          this.tableData = data.data;
          this.filteredData = this.tableData; // Initialize filteredData
        })
        .catch((error) => {
          console.error("Error in Axios POST request:", error);
        });
    },
    async deleteData(id, judul_kreatifitas) {
      try {
        const result = await Swal.fire({
          title: `Hapus data ${judul_kreatifitas}?`,
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
            `http://localhost:8080/kreatifitas/delete/${id}`
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

    sortByJudul() {
      this.filteredData.sort((a, b) =>
        a.judul_kreatifitas.localeCompare(b.judul_kreatifitas)
      ); // Sort by judul_kreatifitas alphabetically
      // Toggle ascending/descending (optional):
      this.filteredData.sort((a, b) =>
        this.sortDirection === "asc"
          ? a.judul_kreatifitas.localeCompare(b.judul_kreatifitas)
          : b.judul_kreatifitas.localeCompare(a.judul_kreatifitas)
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
  width: 100%;
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
</style>
