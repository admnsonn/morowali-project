<template>
  <div>
    <Navbar />
  </div>

  <div class="main">
    <RouterView />
  </div>

  <div class="kaki">
    <footerComponent />
  </div>


</template>

<script>
import Navbar from './components/Navbar.vue'
import footerComponent from './components/Footer.vue'
// import SideBard from './components/SideBar.vue'

export default {
  components: {
    Navbar,
    footerComponent
  }
}
</script>

<style>
body {
  margin: 0;
}
</style>

<style scoped>


.kaki {
  margin-top: auto;
  left: 0;
  width: 100%;
  background-color: #003366;
  line-height: 2;
  padding: 10px;
}

</style>