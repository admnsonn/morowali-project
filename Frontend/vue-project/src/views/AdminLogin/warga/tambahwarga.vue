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

        <div class="field0">
          <div class="form-group">
            <label for="formFile" class="form-label">Foto Profile</label>
            <input class="form-control" v-on:change="onFileChange" type="file" id="formFile" accept="image/*" />
          </div>
          <div class="form-group-foto" v-if="this.model.warga.foto_profile">
            <label for="foto">Preview foto</label>
            <img :src="`data:image/png;base64,${this.model.warga.foto_profile}`" alt="foto profile" height="300"
              width="400" class="td-foto" />
          </div>
          <div v-else>Tidak ada gambar yang dipilih</div>
        </div>

        <div class="field1">
          <div class="form-group">
            <label for="Nama">Nama Lengkap</label>
            <input type="text" v-model="model.warga.nama_lengkap" class="form-control" id="Nama" aria-label="nama"
              placeholder="Nama Lengkap" />
          </div>
        </div>

        <div class="field2">
          <div class="form-group">
            <label for="NIK">Nomor Induk Keluarga</label>
            <input type="text" v-model="model.warga.nik" class="form-control" id="NIK" aria-label="nik"
              placeholder="NIK" />
          </div>
        </div>

        <div class="field3">
          <div class="form-group">
            <label for="KK">Kartu Keluarga</label>
            <input type="text" v-model="model.warga.kk" class="form-control" id="KK" aria-label="kk" placeholder="KK" />
          </div>
        </div>

        <div class="field4">
          <div class="form-group">
            <label for="Tempat">Tempat Lahir</label>
            <input type="text" v-model="model.warga.tempat_lahir" class="form-control" id="Tempat" aria-label="tempat"
              placeholder="Tempat Lahir" />
          </div>
        </div>

        <div class="field5">
          <div class="form-group">
            <label for="Tanggal">Tanggal Lahir</label>
            <input type="text" v-model="model.warga.tanggal_lahir" class="form-control" id="Tanggal" aria-label="tanggal"
              placeholder="Tanggal Lahir" />
          </div>
        </div>

        <div class="field6">
          <div class="form-group">
            <label for="Alamat">Alamat</label>
            <QuillEditor toolbar="essential" v-model:content="model.warga.alamat_pengguna" theme="snow"
              content-type="html" />
          </div>
        </div>

        <div class="field7">
          <div class="form-group">
            <label for="Profesi">Profesi</label>
            <input type="text" v-model="model.warga.profesi" class="form-control" id="Profesi" aria-label="profesi"
              placeholder="Profesi" />
          </div>
        </div>

        <div class="field8">
          <div class="form-group">
            <label for="No">No Telp</label>
            <input type="text" v-model="model.warga.no_telp" class="form-control" id="No" aria-label="no"
              placeholder="No Telp" />
          </div>
        </div>

        <div class="field9">
          <div class="form-group">
            <label for="RW">RW</label>
            <input type="text" v-model="model.warga.rw" class="form-control" id="RW" aria-label="rw" placeholder="RW" />
          </div>
        </div>

        <div class="field10">
          <div class="form-group">
            <label for="RT">RT</label>
            <input type="text" v-model="model.warga.rt" class="form-control" id="RT" aria-label="rt" placeholder="RT" />
          </div>
        </div>

        <div class="field11">
          <div class="form-group">
            <label for="POS">Kode Pos</label>
            <input type="text" v-model="model.warga.kode_pos" class="form-control" id="POS" aria-label="pos"
              placeholder="POS" />
          </div>
        </div>

        <div class="field12">
          <div class="form-group">
            <label for="Kewarganegaraan">Kewarganegaraan</label>
            <input type="text" v-model="model.warga.kewarganegaraan" class="form-control" id="Kewarganegaraan"
              aria-label="kewarganegaraan" placeholder="Kewarganegaraan" />
          </div>
        </div>

        <div class="field13">
          <div class="form-group">
            <label for="formFile" class="form-label">Jenis Kelamin</label>
            <br />
            <select ref="selectJK" v-model="model.warga.jenis_kelamin_id" class="form-control" id="jenis_kelamin_id"
              aria-label="category">
              <option value="" disabled selected>--Pilih Kategori--</option>
              <option v-for="item in kategoriJK" :value="item.id_jenis_kelamin">
                {{ item.jenis_kelamin }}
              </option>
            </select>
          </div>
        </div>

        <div class="field14">
          <div class="form-group">
            <label for="formFile" class="form-label">Agama</label>
            <br />
            <select ref="selectAG" v-model="model.warga.agama_id" class="form-control" id="agama_id"
              aria-label="category">
              <option value="" disabled selected>--Pilih Kategori--</option>
              <option v-for="item in kategoriAG" :value="item.id_agama">
                {{ item.Agama }}
              </option>
            </select>
          </div>
        </div>

        <div class="field15">
          <div class="form-group">
            <label for="formFile" class="form-label">Pendidikan</label>
            <br />
            <select ref="selectPN" v-model="model.warga.pendidikan_terakhir_id" class="form-control"
              id="pendidikan_terakhir_id" aria-label="category">
              <option value="" disabled selected>--Pilih Kategori--</option>
              <option v-for="item in kategoriPN" :value="item.id_pendidikan">
                {{ item.kategori_pendidikan }}
              </option>
            </select>
          </div>
        </div>

        <div class="field16">
          <div class="form-group">
            <label for="formFile" class="form-label">Finansial</label>
            <br />
            <select ref="selectFN" v-model="model.warga.kategori_financial_id" class="form-control"
              id="kategori_financial_id" aria-label="category">
              <option value="" disabled selected>--Pilih Kategori--</option>
              <option v-for="item in kategoriFN" :value="item.id_financial">
                {{ item.financial }}
              </option>
            </select>
          </div>
        </div>

        <div class="field8">
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
import Swal from "sweetalert2";
import { QuillEditor } from "@vueup/vue-quill";
import "@vueup/vue-quill/dist/vue-quill.snow.css";

export default {
  components: {
    QuillEditor
  },
  name: "wargaCreate",
  data() {
    return {
      model: {
        warga: {
          nama_lengkap: "",
          nik: "",
          kk: "",
          tempat_lahir: "",
          tanggal_lahir: "",
          jenis_kelamin_id: "",
          alamat_pengguna: "",
          agama_id: "",
          profesi: "",
          no_telp: "",
          pendidikan_terakhir_id: "",
          foto_profile: "",
          rw: "",
          rt: "",
          id_desa: "1",
          kategori_financial_id: "",
          kode_pos: "",
          kewarganegaraan: "",
        },
      },
      tableData: [],
      kategoriJK: [],
      kategoriAG: [],
      kategoriFN: [],
      kategoriPN: [],
    };
  },
  methods: {
    fetchKategoriJK() {
      axios
        .get("http://localhost:8080/warga/jk")
        .then(({ data }) => {
          this.kategoriJK = data.data;
        })
        .catch((error) => {
          console.error("Error in Axios GET request:", error);
        });
    },
    fetchKategoriAG() {
      axios
        .get("http://localhost:8080/warga/agama")
        .then(({ data }) => {
          this.kategoriAG = data.data;
        })
        .catch((error) => {
          console.error("Error in Axios GET request:", error);
        });
    },
    fetchKategoriPN() {
      axios
        .get("http://localhost:8080/warga/pendidikan")
        .then(({ data }) => {
          this.kategoriPN = data.data;
        })
        .catch((error) => {
          console.error("Error in Axios GET request:", error);
        });
    },
    fetchKategoriFN() {
      axios
        .get("http://localhost:8080/warga/financial")
        .then(({ data }) => {
          this.kategoriFN = data.data;
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
        const plainTextDescription = this.model.warga.alamat_pengguna;

        axios
          .post("http://localhost:8080/warga/tambah", {
            nama_lengkap: this.model.warga.nama_lengkap,
            nik: this.model.warga.nik,
            alamat_pengguna: plainTextDescription,
            kk: this.model.warga.kk,
            id_desa: this.model.warga.id_desa,
            tempat_lahir: this.model.warga.tempat_lahir,
            tanggal_lahir: this.model.warga.tanggal_lahir,
            jenis_kelamin_id: this.model.warga.jenis_kelamin_id,
            agama_id: this.model.warga.agama_id,
            profesi: this.model.warga.profesi,
            no_telp: this.model.warga.no_telp,
            pendidikan_terakhir_id: this.model.warga.pendidikan_terakhir_id,
            foto_profile: this.model.warga.foto_profile,
            rw: this.model.warga.rw,
            rt: this.model.warga.rt,
            kategori_financial_id: this.model.warga.kategori_financial_id,
            kode_pos: this.model.warga.kode_pos,
            kewarganegaraan: this.model.warga.kewarganegaraan,
          })
          .then((res) => {
            this.model.warga = {
              nama_lengkap: "",
              nik: "",
              kk: "",
              tempat_lahir: "",
              tanggal_lahir: "",
              jenis_kelamin_id: "",
              alamat_pengguna: "",
              agama_id: "",
              profesi: "",
              no_telp: "",
              pendidikan_terakhir_id: "",
              foto_profile: "",
              rw: "",
              rt: "",
              id_desa: "1",
              kategori_financial_id: "",
              kode_pos: "",
              kewarganegaraan: "",
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
      const imageFile = files[0];
      const validTypes = ["image/jpeg", "image/png", "image/gif"]; // Adjust as needed
      if (!validTypes.includes(imageFile.type)) {
        // Display an error message or alert
        alert("Please select an image file.");
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
              this.model.warga.foto_profile = base64String;
              // Now you can send this.model.berita to your API
            };
          }, "image/jpeg"); // Adjust the image type as needed
        };

        img.src = imageUrl;
      };
    },
  },
  created(){
    this.fetchKategoriAG();
    this.fetchKategoriJK();
    this.fetchKategoriFN();
    this.fetchKategoriPN();
  },
  mounted() {
    this.$refs.selectAG.focus();
    this.$refs.selectFN.focus();
    this.$refs.selectPN.focus();
    this.$refs.selectJK.focus();
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
