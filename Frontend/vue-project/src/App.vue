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


export default {

  components: {
    Navbar,
    footerComponent,
  },



}
</script>

