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
    <div class="wrapper-col">
      <div class="col">
        <h3 class="title-warga">Data Indeks Desa Membangun (IDM)</h3>
        <p class="subtitle-warga">Management Content dan Data IDM</p>
      </div>
      <div class="wrapper-button">
        <router-link to="/tambah-idm" class="btn btn-success">
          <div>Tambah Data</div>
        </router-link>
      </div>
    </div>

    <div class="">
      <div class="card" style="width: 100%">
        <ul
          class="list-group list-group-flush"
          v-for="(item, index) in tableData"
          :key="index"
        >
          <li class="list-group-item card-idm">
            <div>IDM {{ item.tahun }}</div>
            <div>
              <router-link
                :to="`/detail-idm/${item.tahun}`"
                class="btn btn-primary"
              >
                Detail
              </router-link>
            </div>
          </li>
        </ul>
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
    };
  },
  methods: {
    fetchData() {
      const payload = { desa_id: localStorage.getItem("desa_id") };
      axios
        .post("http://localhost:8080/idmiksike/list", payload)
        .then(({ data }) => {
          this.tableData = data.idm;
          console.log(data.idm[0].tahun);
        })
        .catch((error) => {
          console.error("Error in Axios POST request:", error);
        });
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

.btn-excel {
  background-color: #33b949;
  color: #ffffff;
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

.card-idm {
  height: 5rem;
  display: flex;
  justify-content: space-between;
  padding-left: 50px;
  padding-right: 50px;
  align-items: center;
  border: none;
  border-color: black;
}

.wrapper-col {
  display: flex;
  justify-content: space-between;
  margin-bottom: 10px;
}

select {
  font-size: 14px;
}

h3 {
  font-weight: bold;
}
</style>
