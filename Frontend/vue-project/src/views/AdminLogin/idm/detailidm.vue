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
    <ul class="list-group list-group-flush">
      <li class="list-group-item card-idm">
        <div>
          <h1>IDM</h1>
          <table class="tabel">
            <thead>
              <tr>
                <th>No.</th>
                <th>Tahun</th>
                <th>Indikator</th>
                <th>Skor</th>
                <th>Keterangan</th>
                <th>Kegiatan</th>
                <th>Nilai</th>
                <th>Pusat</th>
                <th>Provinsi</th>
                <th>Kabupaten</th>
                <th>Desa</th>
                <th>CSR</th>
                <th>Lainnya</th>
              </tr>
            </thead>
            <tbody>
              <tr v-for="(row, index) in tableDataIDM" :key="index">
                <td>{{ index + 1 }}</td>
                <td>{{ row.tahun }}</td>
                <td>{{ row.indikator }}</td>
                <td>{{ row.skor }}</td>
                <td>{{ row.keterangan }}</td>
                <td>{{ row.kegiatan }}</td>
                <td>{{ row.nilai }}</td>
                <td>{{ row.pusat }}</td>
                <td>{{ row.provinsi }}</td>
                <td>{{ row.kabupaten }}</td>
                <td>{{ row.desa }}</td>
                <td>{{ row.csr }}</td>
                <td>{{ row.lainnya }}</td>
              </tr>
            </tbody>
          </table>
        </div>
      </li>
      <li class="list-group-item card-idm">
        <div>
          <h1>IKE</h1>
          <table class="tabel">
            <thead>
              <tr>
                <th>No.</th>
                <th>Tahun</th>
                <th>Indikator</th>
                <th>Skor</th>
                <th>Keterangan</th>
                <th>Kegiatan</th>
                <th>Nilai</th>
                <th>Pusat</th>
                <th>Provinsi</th>
                <th>Kabupaten</th>
                <th>Desa</th>
                <th>CSR</th>
                <th>Lainnya</th>
              </tr>
            </thead>
            <tbody>
              <tr v-for="(row, index) in tableDataIKE" :key="index">
                <td>{{ index + 1 }}</td>
                <td>{{ row.tahun }}</td>
                <td>{{ row.indikator }}</td>
                <td>{{ row.skor }}</td>
                <td>{{ row.keterangan }}</td>
                <td>{{ row.kegiatan }}</td>
                <td>{{ row.nilai }}</td>
                <td>{{ row.pusat }}</td>
                <td>{{ row.provinsi }}</td>
                <td>{{ row.kabupaten }}</td>
                <td>{{ row.desa }}</td>
                <td>{{ row.csr }}</td>
                <td>{{ row.lainnya }}</td>
              </tr>
            </tbody>
          </table>
        </div>
      </li>
      <li class="list-group-item card-idm">
        <div>
          <h1>IKS</h1>
          <table class="tabel">
            <thead>
              <tr>
                <th>No.</th>
                <th>Tahun</th>
                <th>Indikator</th>
                <th>Skor</th>
                <th>Keterangan</th>
                <th>Kegiatan</th>
                <th>Nilai</th>
                <th>Pusat</th>
                <th>Provinsi</th>
                <th>Kabupaten</th>
                <th>Desa</th>
                <th>CSR</th>
                <th>Lainnya</th>
              </tr>
            </thead>
            <tbody>
              <tr v-for="(row, index) in tableDataIKS" :key="index">
                <td>{{ index + 1 }}</td>
                <td>{{ row.tahun }}</td>
                <td>{{ row.indikator }}</td>
                <td>{{ row.skor }}</td>
                <td>{{ row.keterangan }}</td>
                <td>{{ row.kegiatan }}</td>
                <td>{{ row.nilai }}</td>
                <td>{{ row.pusat }}</td>
                <td>{{ row.provinsi }}</td>
                <td>{{ row.kabupaten }}</td>
                <td>{{ row.desa }}</td>
                <td>{{ row.csr }}</td>
                <td>{{ row.lainnya }}</td>
              </tr>
            </tbody>
          </table>
        </div>
      </li>
    </ul>

    <div class="excelbutton-wrapper">
      <div class="add-clear-buttons">
        <router-link to="/idm-beranda" class="btn btn-primary">
          Kembali
        </router-link>
      </div>
    </div>
  </div>
</template>

<script>
import axios from "axios";

export default {
  name: "idm-detail",
  data() {
    return {
      tableDataIDM: [],
      tableDataIKE: [],
      tableDataIKS: [],
    };
  },
  methods: {
    fetchData() {
      const payload = {
        desa_id: localStorage.getItem("desa_id"),
        tahun: this.$route.params.tahun,
      };
      axios
        .post("http://localhost:8080/idmiksike/detail", payload)
        .then(({ data }) => {
          console.log(data);
          this.tableDataIDM = data.idm;
          this.tableDataIKE = data.ike;
          this.tableDataIKS = data.iks;
        })
        .catch((error) => {
          console.error("Error in Axios POST request: ", error);
        });
    },
  },
  created() {
    this.fetchData();
  },
};
</script>

<style>
tr {
  font-size: small;
}

td {
  font-size: small;
}

.excelbutton-wrapper {
  display: flex;
  width: 100%;
  justify-content: space-between;

  margin-top: 100px;
}

.card-idm {
  height: auto;
  display: flex;
}

.add-clear-buttons {
  display: flex;
  width: 100%;
  flex-direction: column;
  max-width: fit-content;
  gap: 10px;
}
</style>
