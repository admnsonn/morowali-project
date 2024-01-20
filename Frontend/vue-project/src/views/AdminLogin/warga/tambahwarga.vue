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
        <h3 class="title-warga">Tambah Data Desa</h3>
        <p class="subtitle-warga">Management Content dan Layanan Warga</p>
      </div>
    </div>

    <div class="content-warga">
      <p class="header-title">Form Nama</p>
    </div>
    <div class="isi-tambahdata">
      <div class="grid-container">
        <div class="field4">
          <div class="form-group">
            <label for="NIK">Nomor Induk Kependudukan (NIK)</label>
            <input type="text" v-model="model.warga.nik" class="form-control" id="NIK" aria-label="nik"
              placeholder="masukan nik" />
          </div>
        </div>

        <div class="field1">
          <div class="form-group">
            <label for="NamaLengkap">Nama Lengkap</label>
            <input type="text" v-model="model.warga.nama_lengkap" class="form-control" id="NamaLengkap" aria-label="nama"
              placeholder="masukan nama" />
          </div>
        </div>

        <div class="field8">
          <div class="form-group">
            <label for="KK">Nomor Kartu Keluarga (KK)</label>
            <input type="text" v-model="model.warga.kk" class="form-control" id="KK" aria-label="kk"
              placeholder="masukan kartu keluarga" />
          </div>
        </div>

        <div class="field10">
          <div class="form-group">
            <label for="formFile" class="form-label">Jenis Kelamin</label>
            <br />
            <select
              ref="kategoriSelect"
              v-model="model.warga.jenis_kelamin"
              class="form-control"
              id="kategori_id"
              aria-label="category"
            >
              <option value="" disabled selected>--Pilih Jenis Kelamin--</option>
              <option value="1">Laki-Laki</option>
              <option value="2">Perempuan</option>
            </select>
          </div>
        </div>

        <div class="field10">
          <div class="form-group">
            <label for="formFile" class="form-label">Foto Warga</label>
            <input class="form-control" v-on:change="onFileChange" type="file" id="formFile">
          </div>
        </div>

        <div class="field13">
          <button type="button" class="btn btn-success btn-simpan p-2 my-2" @click="addNewData">
            <div class="nav-link router-link-underline teks-tambah">
              + Tambah Data
            </div>
          </button>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import axios from "axios";

export default {
  name: "wargaCreate",
  data() {
    return {
      model: {
        warga: {
          nik: "",
          nama_lengkap: "",
          kk: "",
          jenis_kelamin: "",
          foto_pengguna: "",
        },
      },
    };
  },
  methods: {
    addNewData() {
      // console.log(this.model.warga);
      axios.post("http://localhost:8080/warga/tambah", this.model.warga)
        .then(res => {
          console.log(res.data)
          alert(res.data.message);

          this.model.warga = {
            nik: "",
            nama_lengkap: "",
            kk: "",
            jenis_kelamin: "",
            foto_pengguna: "",
          };
        })
        .catch((error) => {
          console.error("Error in Axios POST request:", error);
        });
    },
    onFileChange(e) {
      // console.log('Event:', e);
      var files = e.target.files || e.dataTransfer.files;
      if (!files.length) return;

      this.model.warga.foto_pengguna = files[0].name;
    },
  },
  mounted() {
    console.log(this.$refs.kategoriSelect.value);
    this.$refs.kategoriSelect.focus();
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

.title-warga {
  font-size: 20px;
}

.subtitle-warga {
  font-size: 15px;
  color: #5e5e5e;
}

.content-warga {
  background-color: #f7f7f7;
  padding-top: 20px;
  padding-bottom: 5px;
  padding-left: 20px;
  border-radius: 3px;
  border: 1px solid #c7c7c7;
}

.header-title {
  font-weight: bold;
  color: #000000;
}

.isi-tambahdata {
  background-color: #ffffff;
  padding-top: 20px;
  padding-bottom: 5px;
  padding-left: 20px;
  border-radius: 3px;
  border: 1px solid #c7c7c7;
}

.grid-container {
  display: grid;
  grid-template-columns: auto auto auto;
  grid-gap: 40px;
  padding: 10px;
  padding-right: 30px;
}

.btn-simpan {
  background-color: #003366;
  color: #fff;
  border: none;
}

.btn-simpan:hover {
  background-color: #003366;
  color: #fff;
  border: none;
}

.teks-tambah {
  font-size: 15px;
  padding-top: 2%;
  padding-bottom: 2%;
  padding-left: 5px;
  padding-right: 5px;
}
</style>
