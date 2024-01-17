<template>
  <div class="kontainer-admin">
    <h3>Data Warga Desa</h3>
    <p>Management Content dan Layanan Warga</p>

    <!-- Tabel -->
    <div class="table-container">
      <table class="table">
        <thead>
          <tr>
            <th>ID <button type="button" class="btn btn-link"><img src="src/assets/img/sort.svg" /></button></th>
            <th>NIK <button type="button" class="btn btn-link"><img src="src/assets/img/sort.svg" /></button></th>
            <th>Nama <button type="button" class="btn btn-link"><img src="src/assets/img/sort.svg" /></button></th>
            <th>KK <button type="button" class="btn btn-link"><img src="src/assets/img/sort.svg" /></button></th>
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
              <button type="button" class="btn btn-primary m-1"><img src="src/assets/img/view.svg" /></button>
              <button type="button" class="btn btn-warning"><img src="src/assets/img/edit.svg" /></button>
              <button type="button" class="btn btn-danger m-1"><img src="src/assets/img/delete.svg" /></button>
            </td>
          </tr>
        </tbody>
      </table>

      <!-- Pagination -->
      <div class="pagination">
        <button @click="prevPage" :disabled="currentPage === 1">Previous</button>
        <span> {{ currentPage }} dari {{ totalPages }}</span>
        <button @click="nextPage" :disabled="currentPage === totalPages">Next</button>
      </div>
    </div>
  </div>
</template>

<script>
import axios from 'axios'

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
      const payload = { id_desa: parseInt(localStorage.getItem('desa_id')) };

      axios.post('http://localhost:8080/warga/list', payload)
        .then(({ data }) => {
          this.tableData = data.data;
        })
        .catch(error => {
          console.error('Error in Axios POST request:', error);
        });
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

.kontainer-admin {
  margin-top: 32px;
}

.table {
  width: calc(100% - 70px);
  border-collapse: collapse;
  margin-top: 32px;
}

.table th,
.table td {
  border-bottom: 1px solid #ddd;
  padding: 20px;
  text-align: center;
}

.table th {
  background-color: #EEF1F3;
}

.table tbody tr:nth-child(even) {
  background-color: #f2f2f2;
}

.pagination {
  margin-right: 80px;
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
  color: #003366;
}

.table-container {
  overflow-x: auto;
}
</style>