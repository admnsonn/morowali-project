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
            <input v-model="nambahData.nama_lengkap" type="text" class="form-control" id="NamaLengkap" aria-label="nama" placeholder="masukan nama">
          </div>
        </div>

        <div class="field2">
          <div class="form-group">
            <label for="JenisKelamin">Jenis Kelamin</label>
            <input v-model="nambahData.jenis_kelamin" type="text" class="form-control" id="JenisKelamin" aria-label="kelamin"
              placeholder="masukan jenis kelamin">
          </div>
        </div>

        <div class="field3">
          <div class="form-group">
            <label for="NomorTelpon">Nomer Telepon</label>
            <input v-model="nambahData.no_telp" type="text" class="form-control" id="NomorTelpon" aria-label="hp" placeholder="masukan no hp">
          </div>
        </div>

        <div class="field4">
          <div class="form-group">
            <label for="NIK">Nomor Induk Kependudukan (NIK)</label>
            <input v-model="nambahData.nik" type="text" class="form-control" id="NIK" aria-label="nik" placeholder="masukan nik">
          </div>
        </div>

        <div class="field6">
          <div class="form-group">
            <label for="Alamat">Alamat</label>
            <input v-model="nambahData.alamat_pengguna" type="text" class="form-control" id="Alamat" aria-label="alamat" placeholder="masukan alamat">
          </div>
        </div>

        <div class="field7">
          <div class="form-group">
            <label for="Pendidikan">Pendidikan Terahir</label>
            <input v-model="nambahData.profesi" type="text" class="form-control" id="Pendidikan" aria-label="pendidikan"
              placeholder="masukan pendidikan">
          </div>
        </div>

        <div class="field8">
          <div class="form-group">
            <label for="KK">Nomor Kartu Keluarga (KK)</label>
            <input v-model="nambahData.kk" type="text" class="form-control" id="KK" aria-label="kk" placeholder="masukan kartu keluarga">
          </div>
        </div>

        <div class="field9">
          <div class="form-group">
            <label for="Agama">Agama</label>
            <input v-model="agama" type="text" class="form-control" id="Agama" aria-label="agama" placeholder="masukan agama">
          </div>
        </div>

        <div class="field10">
          <div class="form-group">
            <label for="formFile" class="form-label">Foto Warga</label>
            <input v-on="nambahData.foto_pengguna" class="form-control" type="file" id="formFile">
          </div>
        </div>

        <div class="field11">
          <div class="form-group">
            <label for="TTL">Tempat Lahir</label>
            <input v-model="nambahData.tempat_lahir" type="text" class="form-control" id="TTL" aria-label="ttl" placeholder="masukan tempat lahir">
          </div>
        </div>

        <div class="field12">
          <div class="form-group">
            <label for="Status">Status Perkawinan</label>
            <input v-model="nambahData.status_pernikahan" type="text" class="form-control" id="Status" aria-label="ttl" placeholder="masukan status">
          </div>
        </div>

        <div class="field13">
          <button type="button" class="btn btn-success btn-simpan p-2 my-2" @click="addNewData">
            <div class="nav-link router-link-underline teks-tambah">+ Tambah Data</div>
          </button>
        </div>

        <div class="field14">
          <div class="form-group">
            <label for="Tanggal">Tanggal Lahir</label>
            <input v-model="nambahData.tanggal_lahir" type="text" class="form-control" id="Tanggal" aria-label="tanggal" placeholder="masukan tanggal lahir">
          </div>
        </div>
      </div>
    </div>
  </div>
</template>
  
<script>
import axios from 'axios'

export default {
  data() {
    return {
      nambahData: {
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
    };
  },
  methods: {
    addNewData() {
      // Create the payload for the POST request
      const payload = {
        id_desa: localStorage.getItem("desa_id"),

        no_telp: this.nambahData.no_telp,
        alamat_pengguna: this.nambahData.alamat_pengguna,
        profesi: this.nambahData.profesi,
        nik: this.nambahData.nik,
        agama: this.nambahData.kk,
        nama_lengkap: this.nambahData.nama_lengkap,
        kk: this.nambahData.kk,
        foto_pengguna: this.nambahData.foto_pengguna,
        tanggal_lahir: this.nambahData.tanggal_lahir,
        status_pernikahan: this.nambahData.status_pernikahan,
        tempat_lahir: this.nambahData.tempat_lahir,
        jenis_kelamin: this.nambahData.jenis_kelamin,
      };

      // Make a POST request to add new data
      axios
        .post("http://localhost:8080/warga/list", payload)
        .then((response) => {
          // Handle success, for example, refresh the data or show a success message
          console.log("Data added successfully:", response.data);
          // Refresh the data by calling fetchData or any other appropriate method
          this.fetchData();
          this.$router.push('/warga-desa')
          // Clear the form fields
          this.clearForm();
        })
        .catch((error) => {
          // Handle errors, for example, show an error message
          console.error("Error in Axios POST request:", error);
        });
    },

    clearForm() {
      // Reset the form fields
      this.newData = {
        nama_lengkap: "",
        jenis_kelamin: "",
        no_telp: "",
        nik: "",
        alamat_pengguna: "",
        profesi: "",
        kk: "",
        agama: "",
        foto_pengguna: "",
        tanggal_lahir: "",
        tempat_lahir: "",
        status_pernikahan: "",
      };
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