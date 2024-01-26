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
                        <input type="text" v-model="detail[0].nama_lengkap " class="form-control" id="NamaLengkap"
                            aria-label="nama" disabled />
                    </div>
                </div>

                <div class="field2">
                    <div class="form-group">
                        <label for="Umur">Umur</label>
                        <input type="text" v-model="detail[0].umur" class="form-control" id="Umur" aria-label="umur"
                            disabled />
                    </div>
                </div>

                <div class="field3">
                    <div class="form-group">
                        <label for="NIK">Nomer Induk Keluarga (NIK)</label>
                        <input type="text" v-model="detail[0].nik" class="form-control" id="NIK" aria-label="nik" disabled />
                    </div>
                </div>

                <div class="field4">
                    <div class="form-group">
                        <label for="KK">Kartu Keluarga (KK)</label>
                        <input type="text" v-model="detail[0].kk" class="form-control" id="KK" aria-label="kk" disabled />
                    </div>
                </div>

                <div class="field4">
                    <div class="form-group">
                        <label for="TL">Tempat Lahir</label>
                        <input type="text" v-model="detail[0].tempat_lahir" class="form-control" id="TL" aria-label="tl"
                            disabled />
                    </div>
                </div>

                <div class="field5">
                    <div class="form-group">
                        <label for="LT">Tanggal Lahir</label>
                        <input type="text" v-model="detail[0].tanggal_lahir" class="form-control" id="LT" aria-label="lt"
                            disabled />
                    </div>
                </div>

                <div class="field6">
                    <div class="form-group">
                        <label for="JenisKelamin">Jenis Kelamin</label>
                        <input type="text" v-model="detail[0].jenis_kelamin" class="form-control" id="JenisKelamin"
                            aria-label="jeniskelamin" disabled />
                    </div>
                </div>

                <div class="field7">
                    <div class="form-group">
                        <label for="NamaAgama">Agama</label>
                        <input type="text" v-model="detail[0].nama_agama" class="form-control" id="NamaAgama" aria-label="namaagama"
                            disabled />
                    </div>
                </div>

                <div class="field8">
                    <div class="form-group">
                        <label for="StatusPernikahan">Status Pernikahan</label>
                        <input type="text" v-model="detail[0].status_perkawinan" class="form-control" id="StatusPenikahan"
                            aria-label="statuspernihakan" disabled />
                    </div>
                </div>

                <div class="field9">
                    <div class="form-group">
                        <label for="Profesi">Profesi</label>
                        <input type="text" v-model="detail[0].profesi" class="form-control" id="Profesi" aria-label="profesi"
                            disabled />
                    </div>
                </div>

                <div class="field10">
                    <div class="form-group">
                        <label for="NoHP">No Telp</label>
                        <input type="text" v-model="detail[0].no_telp" class="form-control" id="NoHP" aria-label="nohp"
                            disabled />
                    </div>
                </div>

                <div class="field11">
                    <div class="form-group">
                        <label for="TP">Tingkat Pendidikan</label>
                        <input type="text" v-model="detail[0].tingkat_pendidikan" class="form-control" id="TP" aria-label="tp"
                            disabled />
                    </div>
                </div>

                <div class="field12">
                    <div class="form-group">
                        <label for="FotoPengguna">Foto Pengguna</label>
                        <input type="text" v-model="detail[0].foto_pengguna" class="form-control" id="FotoPengguna"
                            aria-label="fotopengguna" disabled />
                    </div>
                </div>

                <div class="field13">
                    <div class="form-group">
                        <label for="AlamatPengguna">Alamat Pengguna</label>
                        <textarea type="text" v-model="detail[0].alamat_pengguna" class="form-control" id="AlamatPengguna"
                            aria-label="alamatpengguna" disabled />
                    </div>
                </div>

                <div class="field14">
                    <div class="form-group">
                        <label for="RT">RT</label>
                        <input type="text" v-model="detail[0].rt" class="form-control" id="RT" aria-label="rt" disabled />
                    </div>
                </div>

                <div class="field15">
                    <div class="form-group">
                        <label for="RW">RW</label>
                        <input type="text" v-model="detail[0].rw" class="form-control" id="RW" aria-label="rw" disabled />
                    </div>
                </div>

                <div class="field16">
                    <div class="form-group">
                        <label for="NamaDesa">Nama Desa</label>
                        <input type="text" v-model="detail[0].nama_desa" class="form-control" id="NamaDesa"
                            aria-label="namadesa" disabled />
                    </div>
                </div>

                <div class="field17">
                    <div class="form-group">
                        <label for="NamaKecamatan">Nama Kecamatan</label>
                        <input type="text" v-model="detail[0].nama_kecamatan" class="form-control" id="NamaKecamatan"
                            aria-label="namakecamatan" disabled />
                    </div>
                </div>
            </div>
        </div>
    </div>
</template>
  
<script>
import axios from "axios";
import { array } from "prop-types";

export default {
    data() {
        return {
            id: this.$route.params.id,
            detail: [],
        };
    },
    methods: {
        getDetail() {
            axios.get(`http://localhost:8080/warga/detail/${this.id}`)
                .then(response => {
                    const data = response.data.data;
                    this.detail = data;
                    console.log(this.detail[0]);
                })
                .catch(error => {
                    console.error('Error fetching detail:', error);
                });
        },
    },
    created() {
        this.getDetail();
        console.log(this.id);
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
}</style>
  