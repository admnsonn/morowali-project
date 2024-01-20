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
        <button type="button" class="btn btn-light btn-grey p-3 my-2">
          Import Excel
        </button>
      </div>
    </div>

    <div class="content-warga">
      <div class="row">
        <div class="col-auto">
          <select v-model="selectedKategori" @change="filterByKategori" class="btn btn-light btn-grey p-3 my-2">
            <option value="">All</option>
            <option value="Laki-laki">Laki-Laki</option>
            <option value="Perempuan">Perempuan</option>
          </select>
        </div>
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
              <th>Jenis Kelamin</th>
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
                    <img src="src/assets/img/view.svg" />
                  </button>
                </router-link>
                <router-link to="/edit-detail">
                  <button class="btn btn-warning m-1" @click="fetchDetailData">
                    <img src="src/assets/img/edit.svg" />
                  </button>
                </router-link>
                <button type="button" @click.prevent="
                  deleteData(item.id_pengguna, item.nama_lengkap)
                  " class="btn btn-danger m-1">
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
    fetchData() {
      const payload = { id_desa: localStorage.getItem("desa_id") };

      axios
        .post("http://localhost:8080/warga/list", payload)
        .then(({ data }) => {
          this.tableData = data.data;
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
  background-color: #33b949;
}

.btn-excel:hover {
  background-color: #33b949;
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

.bartipis {
  background-color: black;
  height: 100%;
  width: 3px;
  grid-row: span 2;
  top: 0;
  left: 0;
  align-self: center;
}


/* Style for the button icon/image */
.color-button img {
  width: 20px; /* Adjust the width as needed */
  height: 20px; /* Adjust the height as needed */
  margin-right: 10px; /* Add some space between the text and the icon */
}

/* Add a subtle shadow effect on hover */
.color-button:hover::before {
  content: "";
  position: absolute;
  top: 50%;
  left: 50%;
  width: 100%;
  height: 100%;
  background: rgba(255, 255, 255, 0.1);
  border-radius: 5px;
  transform: translate(-50%, -50%);
  z-index: -1;
}

/* Optional: Add some margin to the button for spacing */
.m-1 {
  margin: 5px;
}

.color-button{
  background-color: #3498db;
  color: #fff;
  border: none;
  padding: 10px 20px;
  border-radius: 5px;
  text-decoration: none;
  cursor: pointer;
  display: inline-block;
  transition: background-color 0.3s ease;
  position: relative;
  overflow: hidden;
}
</style>
