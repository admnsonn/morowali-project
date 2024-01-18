<template>
  <div>
    <Navbar />
  </div>

  <div class="main" :style="{ marginLeft: posisi !== 0 ? '260px' : '0' }">
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

  data() {
    return {
      posisi: 0,
    };
  },

  created() {
    this.checkUserRole();
  },

  methods: {
    checkUserRole() {
      this.userRole = localStorage.getItem("role_pengguna");
      if (this.userRole === "Admin") {
        this.posisi = 1;
      }

      if (this.userRole === "Warga") {
        this.posisi = 2;
      }
    },
  },
}
</script>
