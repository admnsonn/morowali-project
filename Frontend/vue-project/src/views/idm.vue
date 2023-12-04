<template>
  <div>
    <div class="p-5 text-center bg-hero mt-0 pb-10">
      <h1 class="mb-3 text-title">Status Indeks Desa Membangun</h1>
    </div>

    <div class="container py-5">
      <div class="row">
        <div class="col">

          <div class="row row-cols-4 mb-4">
            <div class="col text-white">
              <div class="gradient-skor-saat-ini p-2 pb-4 rounded">
                <p><strong>0.0000</strong><br>SKOR SAAT INI</p>
              </div>
            </div>
            <div class="col">
              <img src="src/assets/img/skor-saat-ini.svg" alt="Image Skor Saat Ini" class="img-fluid text-center" />
            </div>
            <div class="col text-white">
              <div class="gradient-status-idm p-2 pb-4 rounded">
                <p><strong>MANDIRI</strong><br>STATUS IDM</p>
              </div>
            </div>
            <div class="col">
              <img src="src/assets/img/status-idm.svg" alt="Image Status IDM" class="img-fluid text-center" />
            </div>
          </div>

          <div class="row row-cols-4">
            <div class="col text-white">
              <div class="gradient-skor-minimal p-2 pb-4 rounded">
                <p><strong>0.0000</strong><br>SKOR MINIMAL</p>
              </div>
            </div>
            <div class="col">
              <img src="src/assets/img/skor-minimal.svg" alt="Image Skor Minimal" class="img-fluid text-center" />
            </div>
            <div class="col text-white">
              <div class="gradient-target-status p-2 pb-4 rounded">
                <p><strong>MANDIRI</strong><br>TARGET STATUS</p>
              </div>
            </div>
            <div class="col">
              <img src="src/assets/img/target-status.svg" alt="Image Target Status" class="img-fluid text-center" />
            </div>
          </div>

          <div class="row">
            <div class="col my-4">
              <table class="table table-striped">
                <tbody>
                  <tr>
                    <td>Tahun</td>
                    <td>: 2023</td>
                  </tr>
                  <tr>
                    <td>Desa</td>
                    <td>: Sukrawetan</td>
                  </tr>
                  <tr>
                    <td>Kecamatan</td>
                    <td>: Sukra</td>
                  </tr>
                  <tr>
                    <td>Kabupaten</td>
                    <td>: Indramayu</td>
                  </tr>
                  <tr>
                    <td>Provinsi</td>
                    <td>: Jawa Barat</td>
                  </tr>
                </tbody>
              </table>
            </div>
          </div>
        </div>

        <div class="col py-3">
          <p class="text-center">Indeks Desa Membangun (IDM) 2023</p>
          <p class="text-center text-chart-heading">SKOR : IKS, IKE, IKL</p>
          <div class="container-lg mt-4">
            <Pie class="text-chart" :options="chartOptions" :data="chartData" />
          </div>
        </div>
      </div>

      <!-- <div class="row pt-3">
        <div class="col p-3">
          <table class="table table-bordered">
            <thead>
              <tr>
                <th scope="col">No.</th>
                <th scope="col">First</th>
                <th scope="col">Last</th>
                <th scope="col">Handle</th>
              </tr>
            </thead>
            <tbody>
              <tr>
                <th scope="row">1.</th>
                <td>Mark</td>
                <td>Otto</td>
                <td>@mdo</td>
              </tr>
              <tr>
                <th scope="row">2.</th>
                <td>Jacob</td>
                <td>Thornton</td>
                <td>@fat</td>
              </tr>
              <tr>
                <th scope="row">3.</th>
                <td>Larry</td>
                <td>the Bird</td>
                <td>@twitter</td>
              </tr>
            </tbody>
          </table>
        </div>
      </div> -->

      <div class="row py-3 ">
        <div class="col mx-0">
          <button class="btn btn-secondary float-start mx-0">Download Pdf</button>
        </div>
      </div>
    </div>
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
        labels: ['IKS', 'IKE', 'IKL'],
        datasets: [
          {
            backgroundColor: ['#90ED7D', '#434348', '#7CB5EC'],
            data: [31.5, 29.7, 38.8]
          }
        ]
      },
      chartOptions: {
        responsive: true,
        maintainAspectRatio: false
      }
    };
  }
}
</script>
  
<style scoped>
.bg-hero {
  background-color: #003366;
}

/* Remove margin-top */
.mt-0 {
  margin-top: 0;
}

/* Text title */
.text-chart-heading {
  color: #6666;
}

.text-title {
  color: #fff;
  font-weight: bold;
  text-shadow: 2px 2px #252525;
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
}

.text-blue {
  color: #003366 !important;
}

.gradient-skor-saat-ini {
  background-image: linear-gradient(to right, #4A19C8, #7979FF);
}

.gradient-skor-minimal {
  background-image: linear-gradient(to right, #CF0723, #F63B55);
}

.gradient-status-idm {
  background-image: linear-gradient(to right, #07BA7D, #0FD793);
}

.gradient-target-status {
  background-image: linear-gradient(to right, #FF7800, #F5A811);
}

/* Set card background color and spacing */
</style>
  