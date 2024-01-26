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
        <h3 class="title-warga">Tambah Data UMKM</h3>
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
              v-model="model.umkm.nama_umkm"
              class="form-control"
              id="NamaUmkm"
              aria-label="nama"
              placeholder="Nama UMKM"
            />
          </div>
        </div>

        <div class="field2">
          <div class="form-group">
            <label for="KontenUMKM">Konten UMKM</label>
            <input
              type="text"
              v-model="model.umkm.konten_umkm"
              class="form-control"
              id="KontenUMKM"
              aria-label="Konten-UMKM"
              placeholder="Deskripsi UMKM"
            />
          </div>
        </div>

        <div class="field32">
          <div class="form-group">
            <label for="NomorTelpon">Alamat UMKM</label>
            <input
              type="text"
              v-model="model.umkm.alamat_umkm"
              class="form-control"
              id="AlamatUMKM"
              aria-label="hp"
              placeholder="Alamat dari UMKM"
            />
          </div>
        </div>

        <div class="field11">
          <div class="form-group">
            <label for="NomorTelpon">No. Telp. UMKM</label>
            <input
              type="text"
              v-model="model.umkm.no_telp_umkm"
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
              v-model="model.umkm.kategori_umkm_id"
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

        <div class="field10">
          <div class="form-group">
            <label for="formFile" class="form-label">Foto UMKM</label>
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
            @click="addNewData"
          >
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
import Swal from "sweetalert2";

export default {
  name: "umkmCreate",
  data() {
    return {
      model: {
        umkm: {
          nama_umkm: "",
          konten_umkm: "",
          kategori_umkm_id: "",
          foto_umkm: "", // Store only the file name here
          desa_id: 1,
          no_telp_umkm: "",
          alamat_umkm: "",
        },
      },
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
    async addNewData() {
      const result = await Swal.fire({
        title: "Apakah anda yakin?",
        text: "Cek kembali data yang atkan ditambahkan!",
        icon: "warning",
        showCancelButton: true,
        confirmButtonColor: "#003366",
        cancelButtonColor: "#d33",
        confirmButtonText: "Ya, tambahkan!",
      });
      if (result.isConfirmed) {
        axios
          .post("http://localhost:8080/umkm/tambah_umkm", this.model.umkm)
          .then((res) => {
            this.model.umkm = {
              nama_umkm: "",
              konten_umkm: "",
              kategori_umkm_id: "",
              desa_id: 1,
              foto_umkm: "", // Reset for next upload
              no_telp_umkm: "",
              alamat_umkm: "", //
            };
            if (res.data.status) {
              Swal.fire(
                "Data berhasil ditambahkan.",
                res.data.message,
                "success"
              );
            } else {
              Swal.fire("Data gagal ditambahkan.", res.data.message, "error");
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
              this.model.umkm.foto_umkm = base64String;
            };
          }, "image/jpeg"); // Adjust the image type as needed
        };

        img.src = imageUrl;
      };
    },
  },
  created() {
    this.fetchKategori();
  },
  mounted() {
    this.$refs.kategoriSelect.focus(); // Focus the element
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
