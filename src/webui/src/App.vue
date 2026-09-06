<template>
  <div id="app">
    <!-- BARRA DI NAVIGAZIONE (Nascosta se siamo sulla pagina di Login: '/') -->
    <nav v-if="$route.path !== '/'" class="navbar">
      <div class="nav-brand">
        <strong>StudInfo</strong>
      </div>
      <div class="nav-links">
        <router-link to="/dashboard" exact-active-class="active">Home</router-link>
        <router-link to="/dashboard/appelli" exact-active-class="active">Appelli</router-link>
        <router-link to="/dashboard/prenotazioni" exact-active-class="active">Prenotazioni</router-link>
        <router-link to="/dashboard/libretto" exact-active-class="active">Libretto</router-link>
        <router-link to="/dashboard/profilo" exact-active-class="active">Profilo</router-link>
      </div>
      <div class="nav-logout">
        <router-link to="/" class="btn-logout">Esci</router-link>
      </div>
    </nav>

    <!-- AREA DEI CONTENUTI -->
    <!-- Il router-view "inietta" qui dentro la pagina corrispondente all'URL -->
    <main :class="{ 'main-content': $route.path !== '/' }">
      <router-view />
    </main>
  </div>
</template>

<script>
export default {
  name: 'App'
}
</script>

<style>
/* STILI GLOBALI E DELLA NAVBAR */
body {
  margin: 0;
  font-family: 'Segoe UI', Tahoma, Geneva, Verdana, sans-serif;
  background-color: #f5f7fa;
  color: #333;
}

.navbar {
  display: flex;
  justify-content: space-between;
  align-items: center;
  background-color: #812936; /* Rosso sapienza*/
  color: white;
  padding: 0 20px;
  height: 60px;
  box-shadow: 0 2px 4px rgba(0,0,0,0.1);
}

.nav-brand {
  font-size: 1.2rem;
  letter-spacing: 0.5px;
}

.nav-links {
  display: flex;
  gap: 20px;
}

.nav-links a {
  color: #a0aec0;
  text-decoration: none;
  font-weight: 500;
  padding: 5px 10px;
  border-radius: 4px;
  transition: all 0.3s;
}

.nav-links a:hover {
  color: white;
  background-color: rgba(255, 255, 255, 0.1);
}

/* Evidenzia la pagina corrente */
.nav-links a.active {
  color: white;
  background-color: #812936;
}

.btn-logout {
  color: #fc8181;
  text-decoration: none;
  font-weight: bold;
  border: 1px solid #fc8181;
  padding: 5px 15px;
  border-radius: 20px;
  transition: all 0.3s;
}

.btn-logout:hover {
  background-color: #fc8181;
  color: #812936;
}

.main-content {
  max-width: 1200px;
  margin: 40px auto;
  padding: 0 20px;
}
</style>