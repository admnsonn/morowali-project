<template>
    <div class="p-5 text-center bg-hero mt-0 pb-10">
        <h1 class="mb-3 tulisan_judul">Data Desa</h1>
    </div>
    <img :src="imageUrl" alt="Gambar Berita" class="img-fluid gambar" width="2000">
    <div class="container-lg mt-5">
        <p v-for="(pertama, index) in text_pertama" :key="index">{{ pertama }}</p>
        <div class="mt-4">
            <ol>
                <li>
                    <a href="#"> {{ text1 }}</a>
                </li>
                <li>
                    <a href="/data-pendidikan"> {{ text2 }}</a>
                </li>
                <li>
                    <a href="#"> {{ text3 }}</a>
                </li>
                <li>
                    <a href="#"> {{ text4 }}</a>
                </li>
                <li>
                    <a href="#"> {{ text5 }}</a>
                </li>
            </ol>
        </div>
        <div class="mt-4">
            <p v-for="(kedua, index) in text_kedua" :key="index">{{ kedua }}</p>
        </div>
    </div>
</template>

<script>
export default {
    data() {
        return {
            imageUrl: "",
            text_pertama: "",
            text1: "",
            text2: "",
            text3: "",
            text4: "",
            text5: "",
            text_kedua: "",
        };
    },
    methods: {
        fetchData() {
            const apiResponse = {
                data: {
                    imageUrl: "src/assets/img//desa.png",
                    text_pertama:
                        ["Halaman ini berisi tautan menuju informasi mengenai Basis Data Desa. Ada dua jenis data yang dimuat dalam sistem ini, yakni basis data kependudukan dan basis data sumber daya desa. Silahkan klik pada tautan berikut untuk mendapatkan tampilan data statistik per kategori :"],
                    text1: 'Data Wilayah Administratif ',
                    text2: 'Data Pendidikan Dalam KK',
                    text3: 'Data Pekerjaan ',
                    text4: 'Data Pendidikan Ditempuh ',
                    text5: 'Data Agama ',
                    text_kedua:
                        ["Data yang tampil adalah statistik yang didapatkan dari proses olah data dasar yang dilakukan secara offline di kantor desa secara rutin/harian. Data dasar di kantor desa diunggah ke dalam sistem online di website ini secara berkala. Silahkan hubungi kontak Pemerintah Desa untuk mendapatkan data dan informasi desa termutakhir."]
                },
            };
            this.imageUrl = apiResponse.data.imageUrl;
            this.text_pertama = apiResponse.data.text_pertama;
            this.text1 = apiResponse.data.text1;
            this.text2 = apiResponse.data.text2;
            this.text3 = apiResponse.data.text3;
            this.text4 = apiResponse.data.text4;
            this.text5 = apiResponse.data.text5;
            this.text_kedua = apiResponse.data.text_kedua;
        },
    },
    mounted() {
        this.fetchData();
    }
}
</script>

<style scoped>
.bg-hero {
    background-color: #003366;
}

.tulisan_judul {
    color: #fff;
    font-weight: bold;
    text-shadow: 2px 2px #252525;
}

.btn-primary {
    background-color: white;
    color: #003366;
    font-weight: bold;
    border: none;
    padding: 10px 30px;
    font-family: 'Poppins', sans-serif;
    font-weight: 700;
    font-size: 14px;
    cursor: pointer;
    border-radius: 30px;
}

img {
    max-width: 100%;
    height: auto;
}

ol {
    text-decoration: underline;
    color: #003366;
    font-size: 17px;
}

p{
    font-size: 17px;
}
</style>