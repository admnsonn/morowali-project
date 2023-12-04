<template>
  <div>
    <div class="p-5 text-center bg-hero mt-0 pb-10">
      <h1 class="mb-3 text-white">Informasi</h1>
    </div>
    <div>
      <h2 class="text-blue mb-3 perdes-judul">{{ title }}</h2>
      <p v-for="(paragraph, index) in paragraphs" :key="index" class="perdes-paragraph">{{ paragraph }}</p>

      <!-- Tabel -->
      <div class="table-container">
      <table class="table">
        <thead>
          <tr>
            <th>No</th>
            <th>Jenis Pembangunan</th>
            <th>Waktu</th>
            <th>Lokasi</th>
            <th>Pelaksana</th>
            <th>Kategori</th>
            <th>File</th>
          </tr>
        </thead>
        <tbody>
          <tr v-for="(item, index) in displayedData" :key="index">
            <td>{{ index + 1 + (currentPage - 1) * itemsPerPage }}</td>
            <td>{{ item.jenisPembangunan }}</td>
            <td>{{ item.waktu }}</td>
            <td>{{ item.lokasi }}</td>
            <td>{{ item.pelaksana }}</td>
            <td>{{ item.kategori }}</td>
            <td>{{ item.file }}</td>
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
  </div>
</template>

<script>
export default {
  data() {
    return {
      title: "",
      paragraphs: "",
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
      
      const apiResponse = {
        data: {
          title: "Perencanaan Desa",
          paragraphs: [
           "Lorem ipsum dolor sit amet, consectetur adipiscing elit. In felis urna, vehicula ut felis non, venenatis lacinia urna. Nullam eget nibh ante. Phasellus sed mi id turpis luctus congue ut sed neque. Pellentesque quis pretium justo, convallis imperdiet felis. Proin porttitor tristique varius. Lorem ipsum dolor sit amet, consectetur adipiscing elit. In felis urna, vehicula ut felis non, venenatis lacinia urna. Nullam eget nibh ante. Phasellus sed mi id turpis luctus congue ut sed neque. Pellentesque quis pretium justo, convallis imperdiet felis. Proin porttitor tristique varius.",
          ],
        },
      };

      const tableResponse = {
        data: [
          { jenisPembangunan: "Pembangunan A", waktu: "2023-01-01", lokasi: "Lokasi A", pelaksana: "Pelaksana A", kategori: "Kategori A", file: "File A" },
          { jenisPembangunan: "Pembangunan B", waktu: "2023-02-01", lokasi: "Lokasi B", pelaksana: "Pelaksana B", kategori: "Kategori B", file: "File B" },
          { jenisPembangunan: "Pembangunan A", waktu: "2023-01-01", lokasi: "Lokasi A", pelaksana: "Pelaksana A", kategori: "Kategori A", file: "File A" },
          { jenisPembangunan: "Pembangunan B", waktu: "2023-02-01", lokasi: "Lokasi B", pelaksana: "Pelaksana B", kategori: "Kategori B", file: "File B" },
          { jenisPembangunan: "Pembangunan A", waktu: "2023-01-01", lokasi: "Lokasi A", pelaksana: "Pelaksana A", kategori: "Kategori A", file: "File A" },
          { jenisPembangunan: "Pembangunan B", waktu: "2023-02-01", lokasi: "Lokasi B", pelaksana: "Pelaksana B", kategori: "Kategori B", file: "File B" },
        ],
      };

      this.title = apiResponse.data.title;
      this.paragraphs = apiResponse.data.paragraphs;
      this.tableData = tableResponse.data;
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
  .bg-hero {
    background-color: #003366;
  }

  /* Set text color to white */
  .text-white {
    color: white;
    font-weight: bold;
    text-shadow: 2px 2px #252525;
  }
  .text-blue {
    color: #003366 !important;
  }
  
  .btn-primary{
      background-color: white;
      color: #003366;
      font-weight: bold;
      border: none; 
      padding: 10px 30px; 
      font-family: 'Poppins', sans-serif;
      font-weight: 700  ;
      font-size: 14px; 
      cursor: pointer; 
      border-radius: 30px; 
  }
  .perdes-judul {
      text-align: center;
      font-size: 32px;
      font-weight: bold;
      color: #003366; 
      font-family: 'Poppins', sans-serif;
      margin-top: 50px;
  }
  .perdes-paragraph{
      color: #000;
      text-align: center;
      font-family: 'Poppins', sans-serif;
      font-size: 20px;
      font-style: normal;
      font-weight: 500;
      line-height: normal;
      margin-left: 180px;
      margin-right: 180px;
  }
  .table {
      width: calc(100% - 140px); 
      margin-left: 70px;
      margin-right: 70px;
      border-collapse: collapse;
      margin-top: 108px;
  }
  .table th, .table td {
      border: 1px solid #ddd;
      padding: 8px;
      text-align: left;
  }
  .table th {
      background-color: #003366;
      color: white;
  }
  .table tbody tr:nth-child(even) {
      background-color: #f2f2f2;
  }
  .pagination {
      margin-right: 80px ;
      margin-top: 20px;
      display: flex;
      justify-content: flex-end; /* Memindahkan pagination ke sebelah kanan */
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
@media (max-width: 768px) {
  .bg-hero {
    padding: 3rem 1rem;
  }
  .perdes-judul {
    font-size: 24px;
    margin-top: 30px;
  }
  .perdes-paragraph {
    margin-left: 10px;
    margin-right: 10px;
  }
  .table {
    width: 100%;
    margin-left: 10px;
    margin-right: 10px;
  }
  .pagination {
    margin-right: 10px;
    margin-top: 10px;
  }
  .pagination button,
  .pagination span {
    font-size: 12px;
  }
}
</style>
