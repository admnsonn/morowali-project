<template>
  <div>
    <Navbar />
  </div>

  <div class="main" :style="{
    backgroundColor: '#EEF1F3',
    marginLeft: posisi !== 0 ? '260px' : '0',
    paddingBottom: posisi !== 0 ? '15%' : '15%',
  }">
    <RouterView />
  </div>

  <div class="kaki" :style="{
    backgroundColor: '#EEF1F3',
    marginLeft: posisi !== 0 ? '260px' : '0',
  }">
    <footerComponent />
  </div>
</template>

<script>
import Navbar from "./components/Navbar.vue";
import footerComponent from "./components/Footer.vue";

export default {
  components: {
    Navbar,
    footerComponent,
  },

  data() {
    return {
      posisi: 0,
    };
  },

  created() {
    this.checkUserRole();
  },

  methods: {
    checkUserRole() {
      this.userRole = localStorage.getItem("role_pengguna");
      if (this.userRole === "Admin") {
        this.posisi = 1;
      }

      if (this.userRole === "Warga") {
        this.posisi = 2;
      }
    },
  },
};
</script>

<style>
/* admin */
.custom-icon {
  width: 18px;
  height: 18px;
}

.bungkus-tabel {
  background-color: #ffffff;
  padding: 20px;
  width: 100%;
  border-radius: 5px;
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

.teks-admin {
  font-weight: bold;
  font-size: 24px;
  margin: 0;
}

.teks-kabupaten {

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
  background-color: #FFFFFF;
  height: 131px;
  display: flex;
  align-items: center;
  padding-left: 30px;
  padding-top: 10px;
  padding-bottom: 10px;
}

.tabel {
  border-collapse: collapse;
  margin-top: 15px;
  border-radius: 8px;
  overflow: hidden;
  font-size: 14px;
  min-width: 100%;
}

.tabel th,
.tabel td {
  border-bottom: 1px solid #eef1f3;
  padding: 10px;
  padding-right: 49px;
  text-align: center;
}

.tabel th {
  background-color: #eef1f3;
  padding-right: 49px;
}

.pagination {
  margin-right: 30px;
  margin-top: 20px;
  display: flex;
  justify-content: flex-end;
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
  font-size: 14px;
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
  font-size: 14px;
  font-weight: bold;
}

.tabel-container {
  overflow-x: scroll;
}

h3{
  font-weight: bold;
}

.td-foto {
  border-radius: 0.375rem;
}

.form-group-foto {
  display: flex;
  flex-direction: column;
}
</style>