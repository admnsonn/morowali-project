<template>
  <div class="wrapper-h1">
    <h1>Selamat Datang</h1>
    <h2>{{ nama }}</h2>
  </div>
</template>

<script>
import axios from "axios";
export default {
  data() {
    return {
      nama: "",
    };
  },
  methods: {
    fetchNama() {
      const id_orang = localStorage.getItem("id_pengguna");
      axios
        .get(`http://localhost:8080/warga/detail/${id_orang}`)
        .then(({ data }) => {
          this.nama = data.data[0].nama_lengkap;
          console.log(this.nama);
        })
        .catch((error) => {
          console.error("Error in Axios GET request:", error);
        });
    },
  },
  created() {
    this.fetchNama();
  },
};
</script>

<style>
.wrapper-h1 {
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
  height: 100vh;
  background-color: #f5f5f5;
}
h1 {
  font-size: 80px;
  font-weight: bolder;
}

h2 {
  font-size: 40px;
}
</style>
