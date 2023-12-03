<template>
    <!-- Hero Section -->
        <div class="p-5 text-center bg-hero mt-0 pb-10">
                <h1 class="mb-3 text-white">UMKM</h1>
                <a class="btn btn-primary" href="" role="button">Destinasi</a>
            </div>
        <div>
    <div>
    <!-- Filter Section -->
    <div class="filter-section p-3 filter-umkm">
      <!-- Add your icon for filtering here -->
      <img src="@/assets/filter.png" alt="Filter Icon" class="filter-icon">

      <!-- Four Dropdown Boxes -->
      <select class="dropdown" v-model="dropdown1">
      <option value="" disabled selected>Select Option 1</option>
      <option value="option1">Option 1</option>
      <option value="option2">Option 2</option>
      <option value="option3">Option 3</option>
    </select>

    <select class="dropdown" v-model="dropdown2">
      <option value="" disabled selected>Select Option 2</option>
      <option value="optionA">Option A</option>
      <option value="optionB">Option B</option>
      <option value="optionC">Option C</option>
    </select>

    <select class="dropdown" v-model="dropdown3">
      <option value="" disabled selected>Select Option 3</option>
      <option value="valueX">Value X</option>
      <option value="valueY">Value Y</option>
      <option value="valueZ">Value Z</option>
    </select>

    <select class="dropdown" v-model="dropdown4">
      <option value="" disabled selected>Select Option 4</option>
      <option value="alpha">Alpha</option>
      <option value="beta">Beta</option>
      <option value="gamma">Gamma</option>
    </select>
    </div>

    <!-- UMKM Cards Section -->
    <div class="umkm-cards-section">
      <!-- First Column -->
      <div class="umkm-column">
        <!-- UMKM Card 1 -->
        <div v-for="umkm in umkms" :key="umkm.id" class="umkm-card">
            <img :src="umkm.imageUrl" alt="umkm image" class="img-fluid" />
            <h2 class="umkm-title">{{ umkm.title }}</h2>
          <p class="umkm-address">{{ umkm.address }}</p>
          <p class="umkm-phone">{{ umkm.phoneNumber }}</p>

        </div>
      </div>

      <!-- Second Column -->
      <div class="umkm-column">
        <!-- UMKM Card 5 -->
        <div v-for="umkm in secondColumnUmkms" :key="umkm.id" class="umkm-card">
            <img :src="umkm.imageUrl" alt="umkm image" class="img-fluid" />
            <h2 class="umkm-title">{{ umkm.title }}</h2>
            <p class="umkm-address">{{ umkm.address }}</p>
          <p class="umkm-phone">{{ umkm.phoneNumber }}</p>
         
         </div>
        </div>
      </div>
    </div>
  </div>

</template>

<script>
export default {
  data() {
    return {
      imageUrl: "", 
      umkms: [],
      secondColumnUmkms: [],
      dropdown1: '',
      dropdown2: '',
      dropdown3: '',
      dropdown4: '',
    };
  },
  mounted() {
   
    setTimeout(() => {
      this.umkms = [
        { id: 1, name: 'UMKM 1', imageUrl: 'src/assets/Artikel.png', title: 'Mancing Mania Mantap', address: '633R+88F, Unnamed Road, Lalampu, Kec. Bahodopi, Kabupaten Morowali, Sulawesi Tengah 94974 ', phoneNumber: '+123456789' },
        { id: 2, name: 'UMKM 2', imageUrl: 'src/assets/Artikel.png', title: 'Mancing Mania Mantap', address: '633R+88F, Unnamed Road, Lalampu, Kec. Bahodopi, Kabupaten Morowali, Sulawesi Tengah 94974 ', phoneNumber: '+123456789'  },
        { id: 3, name: 'UMKM 3', imageUrl: 'src/assets/Artikel.png', title: 'Mancing Mania Mantap', address: '633R+88F, Unnamed Road, Lalampu, Kec. Bahodopi, Kabupaten Morowali, Sulawesi Tengah 94974 ', phoneNumber: '+123456789'  },
        { id: 4, name: 'UMKM 4', imageUrl: 'src/assets/Artikel.png', title: 'Mancing Mania Mantap', address: '633R+88F, Unnamed Road, Lalampu, Kec. Bahodopi, Kabupaten Morowali, Sulawesi Tengah 94974 ', phoneNumber: '+123456789'  },
        { id: 5, name: 'UMKM 5', imageUrl: 'src/assets/Artikel.png', title: 'Mancing Mania Mantap', address: '633R+88F, Unnamed Road, Lalampu, Kec. Bahodopi, Kabupaten Morowali, Sulawesi Tengah 94974 ', phoneNumber: '+123456789'  },
        { id: 6, name: 'UMKM 6', imageUrl: 'src/assets/Artikel.png', title: 'Mancing Mania Mantap', address: '633R+88F, Unnamed Road, Lalampu, Kec. Bahodopi, Kabupaten Morowali, Sulawesi Tengah 94974 ', phoneNumber: '+123456789'  },
        // Add more simulated data as needed
      ];
    }, 1000); // Simulating a delay of 1000 milliseconds (1 second)
    setTimeout(() => {
      this.secondColumnUmkms = [
        { id: 7, name: 'UMKM 5', imageUrl: 'src/assets/Artikel.png', title: 'Mancing Mania Mantap', address: '633R+88F, Unnamed Road, Lalampu, Kec. Bahodopi, Kabupaten Morowali, Sulawesi Tengah 94974 ', phoneNumber: '+123456789'  },
        { id: 8, name: 'UMKM 5', imageUrl: 'src/assets/Artikel.png', title: 'Mancing Mania Mantap', address: '633R+88F, Unnamed Road, Lalampu, Kec. Bahodopi, Kabupaten Morowali, Sulawesi Tengah 94974 ', phoneNumber: '+123456789'  },
        { id: 9, name: 'UMKM 5', imageUrl: 'src/assets/Artikel.png', title: 'Mancing Mania Mantap', address: '633R+88F, Unnamed Road, Lalampu, Kec. Bahodopi, Kabupaten Morowali, Sulawesi Tengah 94974 ', phoneNumber: '+123456789'  },
        { id: 10, name: 'UMKM 5', imageUrl: 'src/assets/Artikel.png', title: 'Mancing Mania Mantap', address: '633R+88F, Unnamed Road, Lalampu, Kec. Bahodopi, Kabupaten Morowali, Sulawesi Tengah 94974 5', phoneNumber: '+123456789'  },
        { id: 11, name: 'UMKM 5', imageUrl: 'src/assets/Artikel.png', title: 'Mancing Mania Mantap', address: '633R+88F, Unnamed Road, Lalampu, Kec. Bahodopi, Kabupaten Morowali, Sulawesi Tengah 94974 ', phoneNumber: '+123456789'  },
        { id: 12, name: 'UMKM 5', imageUrl: 'src/assets/Artikel.png', title: 'Mancing Mania Mantap', address: '633R+88F, Unnamed Road, Lalampu, Kec. Bahodopi, Kabupaten Morowali, Sulawesi Tengah 94974 ', phoneNumber: '+123456789'  },
        
      ];
    }, 2000); // Simulating a delay of 2000 milliseconds (2 seconds)
  },
};
</script>

<style scoped>
  .bg-hero {
    background-color: #003366;
  }
  
  /* Remove margin-top */
  .mt-0 {
    margin-top: 0;
  }
  
  /* Set text color to white */
  .text-white {
    color: white;
  }
  
  .text-blue {
    color: #003366 !important;
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
    .filter-section {
    display: flex;
    align-items: center;
    justify-content: space-between;
    background-color: white;
    margin-left: 150px;
    margin-right: 150px;
    padding: 15px;
  }

  /* Style for the filter icon */
  .filter-icon {
    font-size: 20px;
    margin-right: 10px;
  }

  /* Style for the dropdown boxes */
  .dropdown {
    /* Add your styling for the dropdown boxes */
    margin: 0 10px;
    padding: 8px;
    border: 1px solid #ccc;
    border-radius: 5px;
  }

  /* Style for the UMKM Cards Section */
  .umkm-cards-section {
    display: flex;
    justify-content: space-between;
    overflow-x: auto;
    padding: 10px;
    margin-left: 150px;
    margin-right: 150px;
  }

  /* Style for the UMKM Columns */
  .umkm-column {
    flex: 0 0 48%; /* Adjust the width of the columns as needed */
  }

  /* Style for UMKM Cards */
  .umkm-card {
    border: 1px solid #ccc;
    border-radius: 8px;
    margin-bottom: 10px;
    padding: 10px;
    background-color: #D9D9D9;
    margin-top: 50px;
  }

  /* Style for UMKM Card Images */
  .umkm-card img {
    width: 100%;
    height: auto;
    border-radius: 5px;
    margin-bottom: 8px;
  }
  .umkm-title {
    font-family: 'Poppins', sans-serif;
    color: #003366;
    font-size: 24px;
    font-weight: 500;
}
.umkm-address {
    font-family: 'Poppins', sans-serif;
    color: #003366;
    font-size: 15px;
    font-weight: 500;
}
.umkm-phone {
    font-family: 'Poppins', sans-serif;
    color: #003366;
    font-size: 15px;
    font-weight: 500;
}
.phone-info {
  display: flex;
  align-items: center;
  color: #555; /* Adjust the color as needed */
}
</style>

  