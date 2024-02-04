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
        <h3 class="title-warga">Data Wilayah Desa</h3>
        <p class="subtitle-warga">Management Content dan Layanan Kreatifitas</p>
      </div>
    </div>

    <!-- tabel -->
    <div class="bungkus-tabel">
      <div class="row">
        <div class="col">
          <input v-model="searchKeyword" @input="filterData" type="text" class="form-control w-100 my-3" placeholder="Search...">
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
                  Luas Wilayah
                  <button type="button" class="btn btn-link" @click="sortByJudul()">
                    <img src="../../../../src/assets/img/sort.svg" />
                  </button>
                </th>
                <th>Longitude</th>
                <th>Latitude</th>
                <th>Status Wilayah</th>
                <th>Action</th>
              </tr>
            </thead>
            <tbody>
              <tr v-for="(item, index) in displayedData" :key="index">
                <td>{{ (currentPage - 1) * itemsPerPage + index + 1 }}</td>
                <td>{{ item.luas_wilayah }}</td>
                <td>{{ item.longitude }}</td>
                <td>{{ item.latitude }}</td>
                <td>{{ item.status_wilayah }}</td>
                <td>
                  <router-link :to="`/update-kreatifitas/${item.id_kreatifitas}`">
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
        .get("http://localhost:8080/wilayah_desa/desa/1")
        .then(({ data }) => {
          this.tableData = data.data;
          this.filteredData = this.tableData; // Initialize filteredData
        })
        .catch((error) => {
          console.error("Error in Axios POST request:", error);
        });
    },
    filterData() {
      this.filteredData = this.tableData.filter((item) => {
        return (
          item.luas_wilayah.toLowerCase().includes(this.searchKeyword.toLowerCase()) ||
          item.longitude.toString().includes(this.searchKeyword) ||
          item.latitude.toString().includes(this.searchKeyword) ||
          item.status_wilayah.toLowerCase().includes(this.searchKeyword.toLowerCase())
        );
      });
      this.currentPage = 1; // Reset halaman ke 1 setiap kali pencarian berubah
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
        this.sortDirection === "asc"
          ? a.luas_wilayah.localeCompare(b.luas_wilayah)
          : b.luas_wilayah.localeCompare(a.luas_wilayah)
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
  