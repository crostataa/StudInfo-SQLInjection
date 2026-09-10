<template>
  <div class="profilo-container">
    <h2>Profilo Personale</h2>
    <p class="subtitle">I tuoi dati anagrafici e di immatricolazione.</p>

    <!-- BOX ERRORE (Nascosto se va tutto bene) -->
    <div v-if="errore" class="error-box" style="color: red; margin-bottom: 15px;">
      <strong>⚠️ Errore:</strong> {{ errore }}
    </div>

    <!-- CARD DEL PROFILO (Si mostra solo quando i dati sono stati scaricati) -->
    <div v-if="profilo" class="profile-card">
      <div class="profile-row"><strong>Nome:</strong> {{ profilo.nome }}</div>
      <div class="profile-row"><strong>Cognome:</strong> {{ profilo.cognome }}</div>
      <div class="profile-row"><strong>Matricola:</strong> {{ profilo.matricola }}</div>
      <div class="profile-row"><strong>Data di Nascita:</strong> {{ profilo.data_di_nascita }}</div>
      <div class="profile-row"><strong>Indirizzo:</strong> {{ profilo.indirizzo }}</div>
      <div class="profile-row"><strong>Email Istituzionale:</strong> {{ profilo.email }}</div>
    </div>

    <!-- MESSAGGIO DI CARICAMENTO -->
    <div v-else-if="!errore">
      <p>Caricamento dati profilo in corso...</p>
    </div>
  </div>
</template>

<script>
export default {
  name: 'ProfiloView',
  data() {
    return { profilo: null, errore: null }
  },
  async mounted() {
    // Leggiamo chi è loggato
    const matricola = localStorage.getItem('utente_loggato');
    if (!matricola) {
      this.$router.push('/login'); // Se non sei loggato, via!
      return;
    }

    try {
      const response = await fetch(`http://127.0.0.1:3000/api/profilo/${matricola}`);
      const data = await response.json();
      if (data.studente && data.studente.length > 0) {
        this.profilo = data.studente[0]; // Salviamo i dati nel Vue
      }
    } catch (err) {
      this.errore = "Impossibile caricare il profilo";
    }
  }
}
</script>

<style scoped>
.profilo-container { display: flex; flex-direction: column; gap: 20px; }
h2 { color: #812936; margin-bottom: 0; }
.subtitle { color: #842029; margin-top: 5px; }
.profile-card { background: white; padding: 30px; border-radius: 8px; box-shadow: 0 2px 4px rgba(0,0,0,0.05); display: flex; flex-direction: column; gap: 15px; }
.profile-row { font-size: 1.1rem; padding-bottom: 10px; border-bottom: 1px solid #edf2f7; }
.profile-row strong { color: #b02a37; display: inline-block; width: 180px; }
</style>