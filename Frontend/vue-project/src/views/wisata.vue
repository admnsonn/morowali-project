<template>
  <!-- Hero Section -->
  <div class="p-5 text-center bg-hero mt-0 pb-10">
    <h1 class="mb-3 text-white">WISATA</h1>
    <a class="btn btn-primary" href="" role="button">Destinasi</a>
  </div>
  <div>
    <div>
      <!-- Filter Section -->
      <div class="filter-section p-4">
        <!-- Hamburger Icon for Filter -->
        <div class="filter-icon" @click="toggleFilter">
          <img v-if="!showFilter" src="@/assets/img/filter.png" alt="Filter Icon">
          <img v-if="showFilter" src="@/assets/img/filter.png" alt="Close Icon">
        </div>
        <div v-if="showFilter" class="row align-items-center">
          <div class="col-md-11 d-flex flex-wrap align-items-center">
            <!-- Four Dropdown Boxes -->
            <div class="mb-2">
              <select class="dropdown" v-model="dropdown1">
                <option value="" disabled selected>Select Option 1</option>
                <option value="option1">Option 1</option>
                <option value="option2">Option 2</option>
                <option value="option3">Option 3</option>
              </select>
            </div>

            <div class="mb-2">
              <select class="dropdown" v-model="dropdown2">
                <option value="" disabled selected>Select Option 2</option>
                <option value="optionA">Option A</option>
                <option value="optionB">Option B</option>
                <option value="optionC">Option C</option>
              </select>
            </div>

            <div class="mb-2">
              <select class="dropdown" v-model="dropdown3">
                <option value="" disabled selected>Select Option 3</option>
                <option value="valueX">Value X</option>
                <option value="valueY">Value Y</option>
                <option value="valueZ">Value Z</option>
              </select>
            </div>

            <div class="mb-2">
              <select class="dropdown" v-model="dropdown4">
                <option value="" disabled selected>Select Option 4</option>
                <option value="alpha">Alpha</option>
                <option value="beta">Beta</option>
                <option value="gamma">Gamma</option>
              </select>
            </div>
          </div>
        </div>
      </div>

      <!-- wisata Cards Section -->
      <div class="wisata-cards-section row mx-0">
        <div v-for="wisata in pagedwisatas" :key="wisata.id" class="col-md-6 mb-4">
          <div class="wisata-card">
            <img :src="wisata.imageUrl" alt="wisata image" class="img-fluid" />
            <h2 class="wisata-title">{{ wisata.title }}</h2>
            <p class="wisata-address">{{ wisata.address }}</p>
            <p class="wisata-phone">{{ wisata.phoneNumber }}</p>
          </div>
        </div>
      </div>
      <div class="pagination-controls text-center mt-4">
        <button @click="prevPage" :disabled="currentPage === 1">Previous Page</button>
        <span>{{ currentPage }} / {{ totalPages }}</span>
        <button @click="nextPage" :disabled="currentPage === totalPages">Next Page</button>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  data() {
    return {
      imageUrl: "",
      wisatas: [],
      secondColumnwisatas: [],
      dropdown1: '',
      dropdown2: '',
      dropdown3: '',
      dropdown4: '',
      showFilter: false,
      currentPage: 1,
      perPage: 6,
    };
  },
  computed: {
    totalPages() {
      return Math.ceil(this.wisatas.length / this.perPage);
    },
    startIndex() {
      return (this.currentPage - 1) * this.perPage;
    },
    endIndex() {
      return this.currentPage * this.perPage;
    },
    pagedwisatas() {
      return this.wisatas.slice(this.startIndex, this.endIndex);
    },
  },
  mounted() {

    setTimeout(() => {
      this.wisatas = [
        { id: 1, name: 'wisata 1', imageUrl: 'src/assets/img/Artikel.png', title: 'Mancing Mania Mantap', address: '633R+88F, Unnamed Road, Lalampu, Kec. Bahodopi, Kabupaten Morowali, Sulawesi Tengah 94974 ', phoneNumber: '+123456789' },
        { id: 2, name: 'wisata 2', imageUrl: 'src/assets/img/Artikel.png', title: 'Mancing Mania Mantap', address: '633R+88F, Unnamed Road, Lalampu, Kec. Bahodopi, Kabupaten Morowali, Sulawesi Tengah 94974 ', phoneNumber: '+123456789' },
        { id: 3, name: 'wisata 3', imageUrl: 'src/assets/img/Artikel.png', title: 'Mancing Mania Mantap', address: '633R+88F, Unnamed Road, Lalampu, Kec. Bahodopi, Kabupaten Morowali, Sulawesi Tengah 94974 ', phoneNumber: '+123456789' },
        { id: 4, name: 'wisata 4', imageUrl: 'src/assets/img/Artikel.png', title: 'Mancing Mania Mantap', address: '633R+88F, Unnamed Road, Lalampu, Kec. Bahodopi, Kabupaten Morowali, Sulawesi Tengah 94974 ', phoneNumber: '+123456789' },
        { id: 5, name: 'wisata 5', imageUrl: 'src/assets/img/Artikel.png', title: 'Mancing Mania Mantap', address: '633R+88F, Unnamed Road, Lalampu, Kec. Bahodopi, Kabupaten Morowali, Sulawesi Tengah 94974 ', phoneNumber: '+123456789' },
        { id: 6, name: 'wisata 6', imageUrl: 'src/assets/img/Artikel.png', title: 'Mancing Mania Mantap', address: '633R+88F, Unnamed Road, Lalampu, Kec. Bahodopi, Kabupaten Morowali, Sulawesi Tengah 94974 ', phoneNumber: '+123456789' },
        { id: 7, name: 'wisata 7', imageUrl: 'src/assets/img/Artikel.png', title: 'Mancing Mania Mantap', address: '633R+88F, Unnamed Road, Lalampu, Kec. Bahodopi, Kabupaten Morowali, Sulawesi Tengah 94974 ', phoneNumber: '+123456789' },
      ];
    }, 100);

  },
  methods: {
    toggleFilter() {
      this.showFilter = !this.showFilter; // Toggle filter visibility
    },
    nextPage() {
      if (this.currentPage < this.totalPages) {
        this.currentPage += 1;
      }
    },
    prevPage() {
      if (this.currentPage > 1) {
        this.currentPage -= 1;
      }
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

  .filter-section {
    margin-left: 5px;
    margin-right: 15px;
  }

  .filter-icon {
    cursor: pointer;
    font-size: 20px;
  }

  .dropdown {
    margin: 0 10px;
    padding: 8px;
    border: 1px solid #ccc;
    border-radius: 5px;
  }

  /* Style for the wisata Cards Section */
  .wisata-cards-section {
    display: flex;
    flex-wrap: wrap;
    justify-content: space-between;
    overflow-x: auto;
    padding: 10px;
    margin-left: 5px;
    margin-right: 5px;
  }

  .wisata-card {
    border: 1px solid #ccc;
    border-radius: 8px;
    margin-bottom: 10px;
    padding: 10px;
    background-color: #D9D9D9;
    width: 100%; /* Make cards take full width on mobile */
  }

  /* Style for wisata Card Images */
  .wisata-card img {
    width: 100%;
    height: auto;
    border-radius: 5px;
    margin-bottom: 8px;
  }

  .wisata-title {
    font-family: 'Poppins', sans-serif;
    color: #003366;
    font-size: 24px;
    font-weight: 500;
  }

  .wisata-address,
  .wisata-phone {
    font-family: 'Poppins', sans-serif;
    color: #003366;
    font-size: 15px;
    font-weight: 500;
  }

  .phone-info {
    display: flex;
    align-items: center;
    color: #555;
    /* Adjust the color as needed */
  }

  .pagination-controls {
    margin-top: 20px;
  }

  .pagination-controls button {
    background-color: #003366;
    color: white;
    border: none;
    padding: 8px 12px;
    margin: 0 5px;
    cursor: pointer;
    border-radius: 5px;
  }

  .pagination-controls button:disabled {
    background-color: #ccc;
    cursor: not-allowed;
  }

  .pagination-controls span {
    margin: 0 10px;
    font-weight: bold;
    color: #003366;
  }

  @media (max-width: 767px) {
    .filter-section {
      margin-left: 5px;
      margin-right: 5px;
    }

    .col-md-11 {
      flex: 0 0 100%;
      max-width: 100%;
    }

    .filter-section .row {
      display: flex;
      flex-wrap: wrap;
      justify-content: space-between;
      margin-top: 10px;
    }

    .filter-section .dropdown {
      margin-bottom: 10px;
      flex: 0 0 48%; /* Two dropdowns per row on mobile */
      max-width: 48%;
    }

    .wisata-cards-section {
      margin-left: 5px;
      margin-right: 5px;
    }

    .wisata-card {
      flex: 0 0 100%;
      max-width: 100%;
    }

    .pagination-controls {
      margin-top: 10px;
    }
  }
</style>

