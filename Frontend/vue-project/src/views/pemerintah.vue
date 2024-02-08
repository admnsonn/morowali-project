<template>
    <div class="p-5 bg-utama">
        <h1 class="text-title mb-4 mt-0 pb-10">Pemerintahan Desa Bahomoleo</h1>
        <div class="container-lg mt-5">
            <div class="row">
                <div class="col-md-6">
                    <img
                    class="rounded mx-auto d-block custom-ukuran"
                    :src="`data:image/png;base64,${this.imageUrl}`"
                    alt="Desa"
                    height="75"
                    width="100"
                  />
                </div>
                <div class="col-md-6">
                    <p class="penjelasan-title">{{ this.sambutan }}</p>
                </div>
            </div>
        </div>
    </div>

    <div class="container-lg mt-5">
        <h4 class="text-visi mb-4">Visi</h4>
        <p class="penjelasan-vismis">{{ this.visi }}</p>
    </div>
    <br>
    <div>
        <h4 class="text-visi mb-4">Misi</h4>
        <p class="penjelasan-vismis">{{ this.misi }}</p>
    </div>

    <br>
    <div class="mt-5 p-5 bg-kedua">
        <h2 class="text-subtitle mb-4 mt-0 pb-10">Program Kerja</h2>
        <div class="container-lg mt-5">
            <p class="penjelasan-subtitle">{{ this.proker_desa }}</p>
        </div>
    </div>

    <div class="mt-5 p-5">
        <h2 class="text-subtitle mb-4 mt-0 pb-10">Peraturan Kerja</h2>
        <div class="container-lg mt-5">
            <p class="penjelasan-subtitle">{{ this.peraturan }}</p>
        </div>
    </div>

</template>

<script>
import axios from 'axios';

export default {
    data() {
        return {
            paragraphs: "",
            visi_misi: "",

            imageUrl: "",
            id_desa: "", 
            sambutan: "",
            visi: "",
            misi: "",
            proker_desa: "",
            peraturan: "",
        };
    },
    methods: {
        async kepaladesa() {
            try {
                const response = await axios.get(`http://localhost:8080/pemerintah/kepdes/${this.id_desa}`, {});
                if (response.data.status) { 
                    this.imageUrl = response.data.kepala_desa_saat_ini[0].foto_kepala_desa
                } else {
                    console.log("Data Kosong atau Terjadi Kesalahan")
                }
            } catch (error) {
                console.error('Error fetching data:', error);
            }
        },

        async katasambutan() {
            try {
                const response = await axios.get(`http://localhost:8080/pemerintah/sambutan/${this.id_desa}`, {});
                if (response.data.status) { 
                    this.sambutan = response.data.data[0].sambutan
                } else {
                    console.log("Data Kosong atau Terjadi Kesalahan")
                }
            } catch (error) {
                console.error('Error fetching data:', error);
            }
        },

        async visimisi() {
            try {
                const response = await axios.get(`http://localhost:8080/pemerintah/visi-misi/${this.id_desa}`, {});
                if (response.data.status) { 
                    this.visi = response.data.data[0].visi_desa
                    this.misi = response.data.data[0].misi_desa
                } else {
                    console.log("Data Kosong atau Terjadi Kesalahan")
                }
            } catch (error) {
                console.error('Error fetching data:', error);
            }
        },

        async proker() {
            try {
                const response = await axios.get(`http://localhost:8080/pemerintah/proker/${this.id_desa}`, {});
                if (response.data.status) { 
                    this.proker_desa = response.data.data[0].proker
                } else {
                    console.log("Data Kosong atau Terjadi Kesalahan")
                }
            } catch (error) {
                console.error('Error fetching data:', error);
            }
        },

        async peraturanDesa() {
            try {
                const response = await axios.get(`http://localhost:8080/pemerintah/peraturan/${this.id_desa}`, {});
                if (response.data.status) { 
                    this.peraturan = response.data.data[0].peraturan
                } else {
                    console.log("Data Kosong atau Terjadi Kesalahan")
                }
            } catch (error) {
                console.error('Error fetching data:', error);
            }
        }
    },
    created() {
        this.id_desa = localStorage.getItem("id_desa");
        this.kepaladesa();
        this.katasambutan();
        this.visimisi();
        this.proker();
        this.peraturanDesa();
    },
};
</script>

<style scoped>
.container {
    padding-left: 20%;
    padding-right: 20%;
}

.bg-utama {
    background-image: linear-gradient(#003366, white);
}

.text-title {
    color: #fff;
    font-weight: bold;
    text-align: center;
}

.custom-ukuran {
    width: 380px;
    max-width: 100%;
    display: block;
    height: auto;
}

.penjelasan-title {
    font-weight: 500;
    color: black;
    font-family: 'Poppins', sans-serif;
}

.text-visi {
    color: #003366;
    font-weight: bold;
    text-align: center;
}

.penjelasan-vismis {
    text-align: center;
}

.text-subtitle {
    color: #003366;
    font-weight: bold;
    text-align: center;
}

.bg-kedua {
    background-color: #D1D9E0;
}

.penjelasan-subtitle {
    text-align: center;
}

.judul-artikel {
    font-weight: 600;
    color: #003366 !important;
}

.subjudul-artikel {
    font-size: 100%;
}

.btn-secondary {
    background-color: #003366;
    border: none;
    padding: 10px 30px;
    font-family: 'Poppins', sans-serif;
    font-weight: 600;
    font-size: 14px;
    cursor: pointer;
    border-radius: 10px;
    display: grid;
    margin: auto;
}

.with-shadow {
    transition: box-shadow 0.3s ease-in-out;
    box-shadow: none;
}

.with-shadow:hover,
.with-shadow.active {
    box-shadow: 0 0 10px 0 rgba(0, 51, 102, 0.5);
}

.text-center{
    margin-left: 3px;
}
</style>