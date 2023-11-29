<template>
    <div>
      <div class="p-5 text-center bg-hero mt-0 pb-10">
        <h1 class="mb-3 text-white">Sejarah Desa</h1>
        <a class="btn btn-primary" href="" role="button">Profil Desa</a>
      </div>
      <div class="container-lg mt-5 sejarah-desa">
        <div class="row">
          <div class="col-md-6">
            <!-- Gambar Sejarah Desa -->
            <img :src="imageUrl" alt="Desa Lamberea" class="img-fluid" />
          </div>
          <div class="col-md-6">
            <!-- Judul Sejarah Desa -->
            <h2 class="text-blue mb-3 judul-sejarah">{{ title }}</h2>
            <!-- Teks Sejarah Desa -->
            <p v-for="(paragraph, index) in paragraphs" :key="index" class="sejarah-paragraph">{{ paragraph }}</p>
          </div>
        </div>
      </div>
      <div class="container-lg mt-5 silsilah-kepala-desa">
      <h2 class="text-blue mb-3 judul-sejarah">Silsilah Kepala Desa</h2>
      
    </div>
  </div>
</template>
  
  <script>
  export default {
    data() {
      return {
        imageUrl: "", 
        title: "",    
        paragraphs: "",     
      };
    },
    methods: {
      fetchData() {
        const apiResponse = {
          data: {
            imageUrl: "src/assets/sdesa.png",
            title: "Desa Lamberea",
            paragraphs :[
              "Lorem ipsum dolor sit amet, consectetur adipiscing elit. In felis urna, vehicula ut felis non, venenatis lacinia urna. Nullam eget nibh ante. Phasellus sed mi id turpis luctus congue ut sed neque. Pellentesque quis pretium justo, convallis imperdiet felis. Proin porttitor tristique varius.",
              "Integer non convallis dui, ac cursus neque. Curabitur facilisis scelerisque lacus quis bibendum. Mauris placerat dolor sem, sit amet vestibulum magna venenatis eu. Aliquam consectetur metus sapien, in gravida enim elementum id. Nulla et purus magna. Vestibulum porta vitae ipsum nec facilisis. Phasellus efficitur luctus enim vitae condimentum. Curabitur commodo purus ac mi luctus, vitae lobortis ante semper. Duis quam orci, rhoncus vitae ullamcorper quis, tincidunt non erat. Maecenas dui ante, eleifend gravida sagittis quis, congue vestibulum enim.",
              ],
            },
        };
  
        // Update data properties with the fetched data
        this.imageUrl = apiResponse.data.imageUrl;
        this.title = apiResponse.data.title;
        this.paragraphs = apiResponse.data.paragraphs;
      },
    },
    mounted() {
      
      this.fetchData();
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
    .sejarah-desa{
        margin-bottom: 20px;
    }
    .judul-sejarah{
        font-size: 32px;
        font-weight: bold;
        color: #003366; 
        font-family: 'Poppins', sans-serif;
    }
    .sejarah-paragraph{
        font-size: 20px;
        font-weight: 500;
        color: black; 
        font-family: 'Poppins', sans-serif;
        
    }
  </style>
