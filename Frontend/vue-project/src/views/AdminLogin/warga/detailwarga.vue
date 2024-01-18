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
        <div class="field1">
          <div class="form-group">
            <label for="NamaLengkap">Nama Lengkap</label>
            <input type="text" v-model="model.warga.nama_lengkap" class="form-control" id="NamaLengkap" aria-label="nama"
              placeholder="masukan nama">
          </div>
        </div>

        <div class="field2">
          <div class="form-group">
            <label for="JenisKelamin">Jenis Kelamin</label>
            <input type="text" v-model="model.warga.jenis_kelamin" class="form-control" id="JenisKelamin"
              aria-label="kelamin" placeholder="masukan jenis kelamin">
          </div>
        </div>

        <div class="field3">
          <div class="form-group">
            <label for="NomorTelpon">Nomer Telepon</label>
            <input type="text" v-model="model.warga.no_telp" class="form-control" id="NomorTelpon" aria-label="hp"
              placeholder="masukan no hp">
          </div>
        </div>

        <div class="field4">
          <div class="form-group">
            <label for="NIK">Nomor Induk Kependudukan (NIK)</label>
            <input type="text" v-model="model.warga.nik" class="form-control" id="NIK" aria-label="nik"
              placeholder="masukan nik">
          </div>
        </div>

        <div class="field6">
          <div class="form-group">
            <label for="Alamat">Alamat</label>
            <input type="text" v-model="model.warga.alamat_pengguna" class="form-control" id="Alamat" aria-label="alamat"
              placeholder="masukan alamat">
          </div>
        </div>

        <div class="field7">
          <div class="form-group">
            <label for="Pendidikan">Pendidikan Terahir</label>
            <input type="text" v-model="model.warga.profesi" class="form-control" id="Pendidikan" aria-label="pendidikan"
              placeholder="masukan pendidikan">
          </div>
        </div>

        <div class="field8">
          <div class="form-group">
            <label for="KK">Nomor Kartu Keluarga (KK)</label>
            <input type="text" v-model="model.warga.kk" class="form-control" id="KK" aria-label="kk"
              placeholder="masukan kartu keluarga">
          </div>
        </div>

        <div class="field9">
          <div class="form-group">
            <label for="Agama">Agama</label>
            <input type="text" v-model="model.warga.agama" class="form-control" id="Agama" aria-label="agama"
              placeholder="masukan agama">
          </div>
        </div>

        <div class="field10">
          <div class="form-group">
            <label for="formFile" class="form-label">Foto Warga</label>
            <input class="form-control" v-on:change="onFileChange" type="file" id="formFile">
          </div>
        </div>

        <div class="field11">
          <div class="form-group">
            <label for="TTL">Tempat Lahir</label>
            <input type="text" v-model="model.warga.tempat_lahir" class="form-control" id="TTL" aria-label="ttl"
              placeholder="masukan tempat lahir">
          </div>
        </div>

        <div class="field12">
          <div class="form-group">
            <label for="Status">Status Perkawinan</label>
            <input type="text" v-model="model.warga.status_pernikahan" class="form-control" id="Status" aria-label="ttl"
              placeholder="masukan status">
          </div>
        </div>

        <div class="field14">
          <div class="form-group">
            <label for="Tanggal">Tanggal Lahir</label>
            <input type="text" v-model="model.warga.tanggal_lahir" class="form-control" id="Tanggal" aria-label="tanggal"
              placeholder="masukan tanggal lahir">
          </div>
        </div>

        <div class="field13">
          <button type="button" class="btn btn-success btn-simpan p-2 my-2" @click="addNewData">
            <div class="nav-link router-link-underline teks-tambah">+ Tambah Data</div>
          </button>
        </div>
      </div>
    </div>
  </div>
</template>
  
<script>
import axios from 'axios'

export default {
  name: 'wargaCreate',
  data() {
    return {
      model: {
        warga: {
          nama_lengkap: '',
          jenis_kelamin: '',
          no_telp: '',
          nik: '',
          alamat_pengguna: '',
          profesi: '',
          kk: '',
          agama: '',
          foto_pengguna: '',
          tanggal_lahir: '',
          tempat_lahir: '',
          status_pernikahan: '',
        },
      },
      tableData: [],
    };
  },
  methods: {
    addNewData() {
      axios.post("http://localhost:8080/warga/tambah", this.model.warga)
        .then(res => {
          console.log(res.data)
          alert(res.data.message);

          this.model.warga = {
            nama_lengkap: '',
            jenis_kelamin: '',
            no_telp: '',
            nik: '',
            alamat_pengguna: '',
            profesi: '',
            kk: '',
            agama: '',
            foto_pengguna: '',
            tanggal_lahir: '',
            tempat_lahir: '',
            status_pernikahan: '',
          };
        })
        .catch(error => {
          console.error(error);
          // Handle error, e.g., show an error message
        });
    },
    onFileChange(e) {
      var files = e.target.files || e.dataTransfer.files;
      if (!files.length)
        return;
      this.createImage(files[0]);
    },
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
  color: #5E5E5E;
}

.content-warga {
  background-color: #F7F7F7;
  padding-top: 20px;
  padding-bottom: 5px;
  padding-left: 20px;
  border-radius: 3px;
  border: 1px solid #C7C7C7;
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
  border: 1px solid #C7C7C7;
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