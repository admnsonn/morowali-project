<template>
  <div>
    <div class="container lg mt-3 content-desa">
      <h1 class="mb-3 mt-4 judul-berita">Form Permohonan Layanan Desa</h1>

      <form class="form-container">
        <div class="form-group">
          <label for="nama">Nama</label>
          <input
            v-model="model.nama"
            type="text"
            class="form-control"
            id="nama"
            placeholder="Nama"
          />
        </div>

        <div class="form-group">
          <label for="nama">NIK</label>
          <input
            v-model="model.nik"
            type="text"
            class="form-control"
            id="nama"
            placeholder="Nama"
          />
        </div>

        <div class="form-group">
          <label for="nohp">No. HP</label>
          <input
            v-model="model.nomor_hp"
            type="text"
            class="form-control"
            id="nohp"
            placeholder="No. HP"
          />
        </div>
        <div class="form-group">
          <label for="alamat">Alamat</label>
          <input
            v-model="model.alamat"
            type="text"
            class="form-control"
            id="alamat"
            placeholder="Alamat"
          />
        </div>
        <div class="form-group">
          <label for="layanan">Layanan</label>
          <select
            v-model="model.layanan"
            ref="kategoriSelect"
            class="form-control"
            id="kategori_id"
            aria-label="category"
          >
            <option value="" disabled selected>--Pilih Kategori--</option>
            <option value="Permohonan Surat Keterangan Beda Identitas">
              Surat Keterangan Beda Identitas
            </option>
            <option value="Permohonan Surat Keterangan Belum Menikah">
              Surat Keterangan Belum Menikah
            </option>
            <option value="Permohonan Surat Keterangan Domisili">
              Surat Keterangan Domisili
            </option>
            <option value="Permohonan Surat Keterangan Izin Keramaian">
              Surat Keterangan Izin Keramaian
            </option>
            <option value="Permohonan Surat Keterangan Keterlambatan Menikah">
              Surat Keterangan Keterlambatan Menikah
            </option>
            <option value="Permohonan Surat Keterangan Keluarga Tidak Mampu">
              Surat Keterangan Keluarga Tidak Mampu
            </option>
            <option value="Permohonan Surat Keterangan Mendapatkan SKCK">
              Surat Keterangan Mendapatkan SKCK
            </option>
            <option value="Permohonan Surat Keterangan Pemakaman">
              Surat Keterangan Pemakaman
            </option>
            <option value="Permohonan Surat Keterangan Pindah Domisili">
              Surat Keterangan Pindah Domisili
            </option>
            <option value="Permohonan Surat Keterangan Rekomendasi Nikah">
              Surat Keterangan Rekomendasi Nikah
            </option>
            <option value="Permohonan Surat Keterangan Sudah Menikah">
              Surat Keterangan Sudah Menikah
            </option>
            <option value="Permohonan Surat Keterangan Tidak Mampu">
              Surat Keterangan Tidak Mampu
            </option>
            <option value="Permohonan Surat Keterangan Tempat Usaha">
              Surat Keterangan Tempat Usaha
            </option>
            <option value="Permohonan Surat Keterangan Usaha">
              Surat Keterangan Usaha
            </option>
            <option value="Permohonan Surat Pengantar Izin Penutupan Jalan">
              Surat Pengantar Izin Penutupan Jalan
            </option>
          </select>
        </div>
      </form>

      <div class="wrapper-button">
        <button @click="addNewData" class="btn btn-primary">Submit</button>
      </div>
    </div>
  </div>
</template>

<script>
import axios from "axios";
import Swal from "sweetalert2";

export default {
  data() {
    return {
      model: {
        desa_id: 1,
        nama: "",
        nomor_hp: "",
        alamat: "",
        layanan: "",
        status: "Dalam Proses",
        nik: "",
      },
      tableData: [],
    };
  },
  methods: {
    async addNewData() {
      console.log(this.model);
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
          .post("http://localhost:8080/layanan/input", {
            desa_id: 1,
            nama: this.model.nama,
            nomor_hp: this.model.nomor_hp,
            alamat: this.model.alamat,
            layanan: this.model.layanan,
            status: "Dalam Proses",
            nik: this.model.nik,
          })
          .then((res) => {
            this.model.warga = {
              desa_id: 1,
              nama: "",
              nomor_hp: "",
              alamat: "",
              layanan: "",
              status: "Dalam Proses",
              nik: "",
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
          });
      }
    },
  },
};
</script>

<style scoped>
.judul-berita {
  font-weight: bold;
}

.form-container {
  display: grid;
  grid-template-columns: 1fr 1fr;
  width: 100%;
  gap: 20px;
  margin-top: 50px;
}

.wrapper-button {
  margin-top: 20px;
}

@media (max-width: 425px) {
  .form-container {
    grid-template-columns: 1fr;
  }
}
</style>
