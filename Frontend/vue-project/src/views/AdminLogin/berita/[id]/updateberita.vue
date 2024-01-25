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
        <h3 class="title-warga">Ubah Data Berita</h3>
        <p class="subtitle-warga">Management Content dan Layanan Berita</p>
      </div>
    </div>

    <div class="content-warga">
      <p class="header-title">Form Input</p>
    </div>
    <div class="isi-tambahdata">
      <div class="grid-container">
        <div class="field1">
          <div class="form-group">
            <label for="NamaLengkap">Judul</label>
            <input
              type="text"
              v-model="this.tableData[0].judul"
              class="form-control"
              id="JudulBerita"
              aria-label="nama"
              placeholder="Judul berita"
            />
          </div>
        </div>

        <div class="field2">
          <div class="form-group">
            <label for="JenisKelamin">Sub-Judul</label>
            <input
              type="text"
              v-model="this.tableData[0].sub_judul"
              class="form-control"
              id="SubJudulBerita"
              aria-label="kelamin"
              placeholder="Sub-Judul berita"
            />
          </div>
        </div>

        <div class="field3">
          <div class="form-group">
            <label for="NomorTelpon">Deskripsi</label>
            <input
              type="text"
              v-model="this.tableData[0].deskripsi"
              class="form-control"
              id="DeskripsiBerita"
              aria-label="hp"
            />
          </div>
        </div>

        <div class="field10">
          <div class="form-group">
            <label for="formFile" class="form-label">Kategori</label>
            <br />
            <select
              ref="kategoriSelect"
              v-model="this.tableData[0].kategori_id"
              class="form-control"
              id="kategori_id"
              aria-label="category"
            >
              <option value="" disabled selected>--Pilih Kategori--</option>
              <option
                v-for="item in kategoriList"
                :value="item.id_kategori_berita"
              >
                {{ item.berita_kategori }}
              </option>
            </select>
          </div>
        </div>

        <div class="field10">
          <div class="form-group">
            <label for="formFile" class="form-label">Foto Berita</label>
            <input
              class="form-control"
              v-on:change="onFileChange"
              type="file"
              id="formFile"
            />
          </div>
        </div>

        <div class="field13">
          <button
            type="button"
            class="btn btn-success btn-simpan p-2 my-2"
            @click="updateData"
          >
            <div class="nav-link router-link-underline teks-tambah">
              > Ganti Data
            </div>
          </button>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import axios from "axios";
import Swal from "sweetalert2";

export default {
  name: "beritaCreate",
  data() {
    return {
      tableData: [],
      kategoriList: [],
    };
  },

  methods: {
    fetchKategori() {
      axios
        .get("http://localhost:8080/berita/categori")
        .then(({ data }) => {
          this.kategoriList = data.data;
        })
        .catch((error) => {
          console.error("Error in Axios GET request:", error);
        });
      console.log(this.kategoriList);
    },
    fetchData() {
      axios
        .get(`http://localhost:8080/berita/${this.$route.params.id}`)
        .then(({ data }) => {
          this.tableData = data.data;
          console.log("ini this.tabledata", this.tableData);
          console.log("ini datadatadatajudul", fetchedDeskripsi);
          this.model.berita.deskripsi = fetchedDeskripsi;

          this.filteredData = this.tableData; // Initialize filteredData
          this.filterByKategori(); // Apply initial filter
        })
        .catch((error) => {
          console.error("Error in Axios POST request:", error);
        });
    },
    async updateData() {
      const result = await Swal.fire({
        title: "Apakah anda yakin?",
        text: "Cek kembali data yang akan diubah!",
        icon: "warning",
        showCancelButton: true,
        confirmButtonColor: "#003366",
        cancelButtonColor: "#d33",
        confirmButtonText: "Ya, ubah!",
      });
      if (result.isConfirmed) {
        axios
          .put("http://localhost:8080/berita/update", this.tableData[0])
          .then((res) => {
            if (res.data.status) {
              Swal.fire("Data berhasil diubah.", res.data.message, "success");
            } else {
              Swal.fire("Data gagal diubah.", res.data.message, "error");
            }
          })
          .catch((error) => {
            console.error(error);
            // Handle error, e.g., show an error message
          });
      }
    },
    onFileChange(e) {
      var files = e.target.files || e.dataTransfer.files;
      if (!files.length) return;

      // Update the model with only the file name
      this.model.berita.foto_berita = files[0].name;
    },
  },
  mounted() {
    console.log(this.$refs.kategoriSelect.value); // Access the selected value
    this.$refs.kategoriSelect.focus(); // Focus the element
  },
  created() {
    this.fetchData(); // Get original data
    this.fetchKategori();
  },
};
</script>

<style scoped>
h3 {
  font-weight: 700;
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
