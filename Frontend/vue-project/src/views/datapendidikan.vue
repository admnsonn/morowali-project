<template>
    <div class="p-5 text-center bg-hero mt-0 pb-10">
      <h1 class="mb-3 tulisan_judul">Data Pendidikan</h1>
      <a class="btn btn-primary" href="" role="button">Data Desa</a>
    </div>
    <div class=" container mt-4">
        <Pie class="text-chart" :options="chartOptions" :data="chartData" />
    </div>
</template>

<script>
import { Chart as ChartJS, ArcElement, Tooltip, Legend } from 'chart.js'
import { Pie } from 'vue-chartjs'

ChartJS.register(ArcElement, Tooltip, Legend)

export default {
    name: 'PieChart',
    components: { Pie },
    data() {
        return {
            chartData: {
                labels: ['SDN 04 Bungku', 'SDN 15 Bungku', 'SDN 1 Bahomoleo', 'SDN BAHOMANTE'],
                datasets: [
                    {
                        backgroundColor: ['#41B883', '#E46651', '#00D8FF', '#DD1B16'],
                        data: [40, 20, 80, 10]
                    }
                ]
            },
            chartOptions: {
                responsive: true,
                maintainAspectRatio: false
            }
        }
    }
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
</style>
