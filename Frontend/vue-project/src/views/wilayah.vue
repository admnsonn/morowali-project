<template>
    <div>
      <div class="p-5 text-center bg-hero mt-0 pb-10">
        <h1 class="mb-3 text-white">Wilayah Desa</h1>
        <a class="btn btn-primary" href="" role="button">Profil Desa</a>
      </div>
      <div class="container mt-5">
        <h2 class="text-blue mb-4 wilayah-text" >Wilayah</h2>
  
        <!-- Grid untuk 3 baris dan 3 kolom -->
        <div class="row">
        <div v-for="(item, index) in wilayahData" :key="index" class="col-md-4 mb-4">
          <div class="wilayah-item text-center">
            <p class="wilayah-number" :class="'color-' + (index % 3 + 1)">{{ item.number }}</p>
            <p class="wilayah-description" :class="'color-' + (index % 3 + 1)">{{ item.description }}</p>
          </div>
        </div>
      </div>
      <h2 class="text-blue mb-4 wilayah-text">Peta Wilayah</h2>
      <div id="map" class="leaflet-map"></div>
    </div>
  </div>
</template>

  <script>
  import L from 'leaflet';
  import 'leaflet/dist/leaflet.css';

  export default {
    data() {
      return {
        wilayahData: [],
        map: null,
        
      };
    },
    
    mounted() {
      // Panggil fetchData() saat komponen dimuat
        this.fetchData();
        this.initializeMap();
    },
    
    methods: {
    fetchData() {
      // Simulasi pengambilan data dari API (dalam hal ini, data respons API)
      const apiResponse = {
        data: [
          {
            number: 123.123,
            description: "Hektare Wilayah",
          },
          {
            number: 56.565,
            description: "Jumlah Warga Desa",
          },
          {
            number: 89.789,
            description: "Luas Perkebunan",
          },
          {
            number: 123.123,
            description: "Hektare Wilayah",
          },
          {
            number: 56.565,
            description: "Jumlah Warga Desa",
          },
          {
            number: 89.789,
            description: "Luas Perkebunan",
          },
          {
            number: 123.123,
            description: "Hektare Wilayah",
          },
          {
            number: 56.565,
            description: "Jumlah Warga Desa",
          },
          {
            number: 89.789,
            description: "Luas Perkebunan",
          },
        ],
      };

      // Ambil wilayahData dengan fetched data
      this.wilayahData = apiResponse.data;
    },
    initializeMap() {
      // Initialize the map
      this.map = L.map('map').setView([-1.8500, 121.7167], 12); // Set lokasi

      // Add a simple tile layer
      L.tileLayer('https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png', {
        attribution: '© OpenStreetMap contributors',
      }
      ).addTo(this.map);
        // Edit Marker
        const customIcon = L.icon({
            iconUrl:'src/assets/marker.png',
            iconSize: [32, 32],
            iconAnchor: [16, 32],
            popupAnchor: [0, -32],
     }
     );
    const marker = L.marker([-1.8500, 121.7167], { icon: customIcon }).addTo(this.map);
        marker.bindPopup('Bahomoleo, Morowali, Sulawesi');
        },
    },
    };
</script>
  
  <style scoped>
    .bg-hero {
      background-color: #003366;
    }
  
    .text-white {
      color: white;
      font-weight: bold;
      text-shadow: 2px 2px #252525;
    }
    .text-blue {
      color: #003366 !important;
    }
    .wilayah-text{
        text-align: center;
        font-size: 32px;
        font-weight: bold;
        color: #003366; 
        font-family: 'Poppins', sans-serif;
    }
    .btn-primary{
        background-color: white;
        color: #003366;
        font-weight: bold;
        border: none; 
        padding: 10px 30px; 
        font-family: 'Poppins', sans-serif;
        font-weight: 700  ;
        font-size: 14px; 
        cursor: pointer; 
        border-radius: 30px; 
    }
    .wilayah-number {
        font-family: 'Poppins', sans-serif;
        font-size: 48px;
        font-weight: 600;
        color: #003366;
    }

  /* Styling wilayah description */
    .wilayah-description {
        font-family: 'Poppins', sans-serif;
        color: #003366;
        font-size: 24px;
        font-weight: 500;
    }
    .color-1 {
        color: #F25555; /* Ubah warna sesuai keinginan Anda */
    }

    .color-2 {
            color: #003366; /* Ubah warna sesuai keinginan Anda */
        }

    .color-3 {
            color: #C0D442; /* Ubah warna sesuai keinginan Anda */
        }
    /* Leaflet map container style */
    .leaflet-map {
    height: 400px; /* Adjust the height as needed */
    margin-top: 20px;
    margin-bottom: 20px;
    }

  </style>