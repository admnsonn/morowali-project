<template>
    <div class="container">
      <div class="row">
        <div class="col">
          <h3 class="title-warga">Detail Warga Desa</h3>
          <p class="subtitle-warga">Management Content dan Layanan Warga</p>
        </div>
      </div>
  
      <div class="content-warga">
        <p class="header-title">Form Nama</p>
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
        const payload = { id_desa: (localStorage.getItem('desa_id')) };
  
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
  
  .container {
    margin-top: 50px;
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
    border-bottom: 1px solid #EEF1F3;
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
  
  .content-warga {
    background-color: #F7F7F7;
    padding: 20px;
    width: 100%;
    margin-bottom: 20%;
  }
  
  .title-warga {
    font-size: 20px;
  }
  
  .subtitle-warga{
    font-size: 15px;
    color: #5E5E5E;
  }

  .header-title{
    font-weight: 500;
    color: #003366;
    align-items: center;
  }
  </style>