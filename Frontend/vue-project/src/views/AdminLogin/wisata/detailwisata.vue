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
                <h3 class="title-warga">Detail Data Wisata</h3>
                <p class="subtitle-warga">Management Content dan Layanan Berita</p>
            </div>
        </div>

        <div class="content-warga">
            <p class="header-title">Form Detail</p>
        </div>
        <div class="isi-tambahdata">
            <div class="grid-container">
                <div class="field5">
                    <div class="form-group-foto">
                        <label for="Fpto">Foto Berita</label>
                        <img :src="`data:image/png;base64,${detail[0].foto_wisata}`" alt="foto wisata" />
                    </div>
                </div>

                <div class="field1">
                    <div class="form-group">
                        <label for="NamaWisata">Nama Wisata</label>
                        <input type="text" v-model="detail[0].nama_wisata" class="form-control" id="NamaWisata"
                            aria-label="nama" disabled />
                    </div>
                </div>

                <div class="field2">
                    <div class="form-group">
                        <label for="Konten">Konten Wisata</label>
                        <textarea type="text" v-model="detail[0].konten_wisata" class="form-control" id="Konten"
                            aria-label="konten" disabled />
                    </div>
                </div>

                <div class="field3">
                    <div class="form-group">
                        <label for="NoTelp">No Telp</label>
                        <input type="text" v-model="detail[0].no_telp" class="form-control" id="NoTelp" aria-label="telp"
                            disabled />
                    </div>
                </div>

                <div class="field4">
                    <div class="form-group">
                        <label for="Alamat">Konten Wisata</label>
                        <textarea type="text" v-model="detail[0].alamat" class="form-control" id="Alamat"
                            aria-label="alamat" disabled />
                    </div>
                </div>

                <div class="field5">
                    <div class="form-group">
                        <label for="Kategori">Kategori</label>
                        <input type="text" v-model="detail[0].kategori_wisata" class="form-control" id="Kategori"
                            aria-label="kategori" disabled />
                    </div>
                </div>
            </div>
        </div>
    </div>
</template>
  
<script>
import axios from "axios";

export default {
    data() {
        return {
            id: this.$route.params.id,
            detail: [],
        };
    },
    methods: {
        getDetail() {
            axios
                .get(`http://localhost:8080/wisata/${this.id}`)
                .then((response) => {
                    const data = response.data.data;
                    this.detail = data;
                    console.log(this.detail[0]);
                })
                .catch((error) => {
                    console.error("Error fetching detail:", error);
                });
        },
    },
    created() {
        this.getDetail();
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
  