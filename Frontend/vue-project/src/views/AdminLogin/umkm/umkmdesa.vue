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
      <h3>Data Berita Desa</h3>
      <p>Management Content dan Layanan warga</p>
    </div>

    <!-- tabel -->
    <div class="content-berita">
      <div class="row">
        <div class="col-auto">
          <button type="button" class="btn btn-light btn-grey p-3 my-2">
            Cari Berdasarkan: Nama
          </button>
        </div>
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
export default {
  data() {
    return {
      tableData: [],
      currentPage: 1,
      itemsPerPage: 7, // Sesuaikan item table perhalaman
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
  },
  created() {
    this.fetchData();
  },
};

// export default {
//   data() {
//     return {
//       tableData: [],
//       currentPage: 1,
//       itemsPerPage: 7, // Sesuaikan item table perhalaman
//     };
//   },
//   computed: {
//     totalPages() {
//       return Math.ceil(this.tableData.length / this.itemsPerPage);
//     },
//     displayedData() {
//       const startIndex = (this.currentPage - 1) * this.itemsPerPage;
//       const endIndex = startIndex + this.itemsPerPage;
//       return this.tableData.slice(startIndex, endIndex);
//     },
//   },
//   methods: {
//     fetchData() {
//       const payload = { id_desa: "1" }; // Replace with the desired id_desa value

//       axios
//         .get("http://localhost:8080/berita/list", {
//           data: JSON.stringify(payload), // Send raw JSON body
//           headers: { "Content-Type": "application/json" },
//         })
//         .then(({ data }) => {
//           this.tableData = data.data;
//         })
//         .catch((error) => {
//           console.error("Error in Axios GET request:", error);
//         });
//     },
//     prevPage() {
//       if (this.currentPage > 1) {
//         this.currentPage--;
//       }
//     },
//     nextPage() {
//       if (this.currentPage < this.totalPages) {
//         this.currentPage++;
//       }
//     },
//   },
//   mounted() {
//     this.fetchData();
//   },
// };
</script>

<style scoped>
.kontainer-data {
  height: auto;
  display: flex;
  flex-direction: column;
  margin-left: 80px;
  margin-top: 32px;
  margin-bottom: 46px;
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

.kontainer {
  display: grid;
  grid-template-columns: 10px auto;
  grid-template-rows: auto;
}

.kontainer-admin {
  border: 22%;
  background-color: white;
  height: 131px;
  display: flex;
  align-items: center;
  padding-left: 30px;
  padding-top: 10px;
  padding-bottom: 10px;
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
