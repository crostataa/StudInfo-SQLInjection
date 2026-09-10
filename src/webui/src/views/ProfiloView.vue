<template>
  <div class="profilo-container">
    <h2>Profilo Personale</h2>
    <p class="subtitle">I tuoi dati anagrafici e di immatricolazione.</p>

    <div class="profile-card">
      <div class="profile-row"><strong>Nome:</strong> Mario</div>
      <div class="profile-row"><strong>Cognome:</strong> Rossi</div>
      <div class="profile-row"><strong>Matricola:</strong> 123456</div>
      <div class="profile-row"><strong>Data di Nascita:</strong> 15/05/2001</div>
      <div class="profile-row"><strong>Indirizzo:</strong> Via delle Scienze 42, Roma</div>
      <div class="profile-row"><strong>Email Istituzionale:</strong> mario.rossi@studenti.uni.it</div>
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