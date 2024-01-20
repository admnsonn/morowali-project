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
        <button type="button" class="btn btn-light btn-excel my-2">
          Import Excel
        </button>
      </div>
    </div>

    <div class="bungkus-tabel">
      <div class="row">
        <div class="col">
          <button type="button" class="btn btn-search w-100 my-2">
            <img src="src/assets/img/search.svg" class="me-2" /> Search...
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
                  <button type="button" class="btn btn-link">
                    <img src="src/assets/img/sort.svg" />
                  </button>
                </th>
                <th>
                  NIK
                  <button type="button" class="btn btn-link">
                    <img src="src/assets/img/sort.svg" />
                  </button>
                </th>
                <th>
                  Nama
                  <button type="button" class="btn btn-link">
                    <img src="src/assets/img/sort.svg" />
                  </button>
                </th>
                <th>
                  KK
                  <button type="button" class="btn btn-link">
                    <img src="src/assets/img/sort.svg" />
                  </button>
                </th>
                <th>
                  Jenis Kelamin
                  <select v-model="selectedKategori" @change="filterByKategori" class="btn btn-light p-2 my-1">
                    <option value="">All</option>
                    <option value="Laki-laki">Laki-Laki</option>
                    <option value="Perempuan">Perempuan</option>
                  </select>
                </th>
                <th>Action</th>
              </tr>
            </thead>
            <tbody>
              <tr v-for="(item, index) in displayedData" :key="index">
                <td>{{ item.id_pengguna }}</td>
                <td>{{ item.nik }}</td>
                <td>{{ item.nama_lengkap }}</td>
                <td>{{ item.kk }}</td>
                <td>{{ item.jenis_kelamin }}</td>
                <td>
                  <router-link to="/detail-warga">
                    <button class="btn btn-info m-1">
                      <img src="src/assets/img/view.svg" class="custom-icon" />
                    </button>
                  </router-link>
                  <router-link to="/edit-detail">
                    <button class="btn btn-warning m-1" @click="fetchDetailData">
                      <img src="src/assets/img/edit.svg" class="custom-icon" />
                    </button>
                  </router-link>
                  <button type="button" @click.prevent="
                    deleteData(item.id_pengguna, item.nama_lengkap)
                    " class="btn btn-danger m-1">
                    <img src="src/assets/img/delete.svg" class="custom-icon" />
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
      const payload = { id_desa: localStorage.getItem("desa_id") };

      axios
        .post("http://localhost:8080/warga/list", payload)
        .then(({ data }) => {
          this.tableData = data.data;
          this.filteredData = this.tableData;
          this.filterByKategori();
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
    fetchDetailData() {
      const residentId = 1; // You can change this to the desired resident ID
      axios.get(`http://localhost:8080/warga/detail/${residentId}`)
        .then((res) => {
          this.detailWarga = res.data.data;
        })
        .catch((error) => {
          console.error("Error in Axios GET request:", error);
        });
    },
    filterByKategori() {
      this.filteredData = this.tableData.filter(
        (item) =>
          item.jenis_kelamin === this.selectedKategori ||
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
  created() {
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
}

.btn-light option {
  font-size: small;
  padding-left: 11px;
}

select {
  font-size: 14px;
}
</style>
