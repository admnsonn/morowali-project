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
        <h3 class="title-warga">Ubah Data UMKM</h3>
        <p class="subtitle-warga">Management Content dan Layanan UMKM</p>
      </div>
    </div>

    <div class="content-warga">
      <p class="header-title">Form Input</p>
    </div>
    <div class="isi-tambahdata">
      <div class="grid-container">
        <div class="field1">
          <div class="form-group">
            <label for="NamaLengkap">Nama UMKM</label>
            <input
              type="text"
              v-model="this.tableData[0].nama_umkm"
              class="form-control"
              id="JudulBerita"
              aria-label="nama"
              placeholder="Nama UMKM"
            />
          </div>
        </div>

        <div class="field2">
          <div class="form-group">
            <label for="JenisKelamin">Konten UMKM</label>
            <QuillEditor
              toolbar="essential"
              v-model:content="this.tableData[0].konten_umkm"
              theme="snow"
              content-type="html"
            />
          </div>
        </div>

        <div class="field3">
          <div class="form-group">
            <label for="NomorTelpon">Alamat UMKM</label>
            <input
              type="text"
              v-model="this.tableData[0].alamat_umkm"
              class="form-control"
              id="AlamatUMKM"
              aria-label="hp"
              placeholder="Alamat dari UMKM"
            />
          </div>
        </div>

        <div class="field10">
          <div class="form-group">
            <label for="formFile" class="form-label">Foto UMKM</label>
            <input
              class="form-control"
              v-on:change="onFileChange"
              type="file"
              id="formFile"
              accept="image/*"
            />
          </div>
        </div>

        <div class="field11">
          <div class="form-group">
            <label for="NomorTelpon">No. Telp. UMKM</label>
            <input
              type="text"
              v-model="this.tableData[0].no_telp_umkm"
              class="form-control"
              id="NoTelpUMKM"
              aria-label="hp"
              placeholder="Nomor telepon UMKM"
            />
          </div>
        </div>

        <div class="field10">
          <div class="form-group">
            <label for="formFile" class="form-label">Kategori UMKM</label>
            <br />
            <select
              ref="kategoriSelect"
              v-model="this.tableData[0].kategori_umkm_id"
              class="form-control"
              id="kategori_id"
              aria-label="category"
            >
              <option value="" disabled selected>--Pilih Kategori--</option>
              <option
                v-for="item in kategoriList"
                :value="item.id_kategori_umkm"
              >
                {{ item.umkm_kategori }}
              </option>
            </select>
          </div>
        </div>

        <div class="form-group-foto">
          <label for="foto">Preview foto</label>
          <img
            :src="`data:image/png;base64,${this.tableData[0].foto_umkm}`"
            alt="foto berita"
            height="300"
            width="400"
            class="td-foto"
          />
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
import { QuillEditor } from "@vueup/vue-quill";
import "@vueup/vue-quill/dist/vue-quill.snow.css";

export default {
  components: {
    QuillEditor,
  },
  name: "umkmCreate",
  data() {
    return {
      tableData: [],
      kategoriList: [],
    };
  },

  methods: {
    fetchKategori() {
      axios
        .get("http://localhost:8080/umkm/umkm_kategori")
        .then(({ data }) => {
          this.kategoriList = data.data;
        })
        .catch((error) => {
          console.error("Error in Axios GET request:", error);
        });
    },
    fetchData() {
      axios
        .get(`http://localhost:8080/umkm/${this.$route.params.id}`)
        .then(({ data }) => {
          this.tableData = data.data;

          this.filteredData = this.tableData; // Initialize filteredData
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
          .put("http://localhost:8080/umkm/update_umkm", this.tableData[0])
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
      const imageFile = files[0];
      const validTypes = ["image/jpeg", "image/png", "image/gif"]; // Adjust as needed
      if (!validTypes.includes(imageFile.type)) {
        // Display an error message or alert
        alert("Silakan unggah file gambar.");
        return;
      }

      const reader = new FileReader();
      reader.readAsDataURL(files[0]);

      reader.onload = (e) => {
        const imageUrl = e.target.result;
        const img = new Image();

        img.onload = () => {
          const canvas = document.createElement("canvas");
          const scaleX = 0.5; // Resize to 50%
          const scaleY = 0.5;
          const width = img.width * scaleX;
          const height = img.height * scaleY;

          canvas.width = width;
          canvas.height = height;
          const ctx = canvas.getContext("2d");
          ctx.drawImage(img, 0, 0, width, height);

          canvas.toBlob((resizedBlob) => {
            const reader = new FileReader();
            reader.readAsDataURL(resizedBlob);

            reader.onloadend = () => {
              const base64String = reader.result.split(",")[1]; // Extract base64-encoded data
              this.tableData[0].foto_umkm = base64String;
            };
          }, "image/jpeg"); // Adjust the image type as needed
        };

        img.src = imageUrl;
      };
    },
  },
  mounted() {
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

.form-group-foto {
  display: flex;
  flex-direction: column;
}

.field13 {
  display: flex;
  grid-column: span 2;
  justify-content: flex-end;
  padding-bottom: 40%;
}

.td-foto {
  border-radius: 0.375rem;
}
</style>
