<template>
    <div class="p-5 text-center bg-hero mt-0 pb-10">
        <h1 class="mb-3 tulisan_judul">Data Pendidikan</h1>
        <a class="btn btn-primary" href="" role="button">Data Desa</a>
    </div>
    <div class="container-lg mt-4">
        <Pie class="text-chart" :options="chartOptions" :data="chartData" />
    </div>
    <div class="container mt-4 ">
        <div class="row justify-content-between">
            <div v-for="(item, index) in pendidikanData" :key="index" class="col-4 mb-4 ml-2 border rounded ">
                <br>
                <p class="pendidikan-description mb-2">{{ index.toUpperCase() }}</p>
                <br>
                <p class="pendidikan-number mb-4" :class="'color-' + (index % 3 + 1)">{{ item }}</p>

                <div class="mb-2 gambar-orang">
                    <img src="../../src/assets/img/orang.png" alt="Latest Image" class="img-fluid" />
                </div>
            </div>
        </div>
    </div>
</template>

<script>
import { Chart as ChartJS, ArcElement, Tooltip, Legend } from 'chart.js'
import { Pie } from 'vue-chartjs'
import axios from 'axios'

ChartJS.register(ArcElement, Tooltip, Legend)

export default {
    name: 'PieChart',
    components: { Pie },
    data() {
        return {
            id_desa: "",
            pendidikanData: [],
            chartData: {
                labels: [],
                datasets: [
                    {
                        backgroundColor: ['#CD8924', '#4F0F13', '#4C0266'],
                        data: []
                    }
                ]
            },
            chartOptions: {
                responsive: true,
                maintainAspectRatio: false
            }
        };
    },

    methods: {
        async fetchPendidikan() {
            try {
                const response = await axios.post("http://localhost:8080/pendidikan", { id_desa: this.id_desa });
                if (response.data.status) {
                    this.pendidikanData = response.data.data[0]

                    // Update chartData with fetched data (masih blm bisa)
                    this.chartData.labels = Object.keys(response.data.data[0]);
                    this.chartData.datasets[0].data = Object.values(response.data.data[0]);
                    console.log(this.chartData)
                } else {
                    console.log("Data Kosong atau Terjadi Kesalahan")
                }
            } catch (error) {
                console.error('Error fetching data:', error);
            }
        },
    },

    created() {
        this.id_desa = localStorage.getItem("id_desa");
        this.fetchPendidikan();
    },
}
</script>

<style scoped>
.bg-hero {
    background-color: #003366;
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

.tulisan_judul {
    color: #fff;
    font-weight: bold;
    text-shadow: 2px 2px #252525;
}

.ml-2 {
    margin-left: -2%;
}

.gambar-orang {
    margin-left: 80%;
    margin-top: -32%;
}

.color-1 {
    color: #4C0266;
    /* Ubah warna sesuai keinginan Anda */
}

.color-2 {
    color: #4F0F13;
    /* Ubah warna sesuai keinginan Anda */
}

.color-3 {
    color: #CD8924;
    /* Ubah warna sesuai keinginan Anda */
}

.pendidikan-description {
    color: #959595;
    font-weight: bold;
}

.pendidikan-number {
    font-size: 30px;
    font-weight: bold;
}
</style>
