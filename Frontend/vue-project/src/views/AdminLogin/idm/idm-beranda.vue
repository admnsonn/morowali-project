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

  <div class="wrapper-h1">
    <div class="wrapper-button">
      <router-link to="/tambah-idm" class="btn btn-success">
        <div>Tambah Data</div>
      </router-link>
    </div>
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

<style>
.wrapper-h1 {
  padding: 100px;
}

.wrapper-button {
  display: flex;
  justify-content: flex-end;
  margin-bottom: 10px;
}

.card-idm {
  height: 5rem;
  display: flex;
  justify-content: space-between;
  padding-left: 50px;
  padding-right: 50px;
  align-items: center;
  border: solid;
  border-color: black;
}

.excelbutton-wrapper {
  display: flex;
  width: 100%;
  justify-content: space-between;
  margin-top: 100px;
}
</style>
