<script>
export default {
  name: "LoginView"
}
</script>

<template>

</template>

<style scoped>

</style>



<template>
  <div style="font-family: Arial, sans-serif; max-width: 800px; margin: 40px auto; text-align: center;">

    <h2>Ricerca Dipendenti Aziendali</h2>
    <p>Inserisci l'username del dipendente per visualizzare i suoi dati.</p>

    <!-- BARRA DI RICERCA -->
    <div style="margin-bottom: 30px;">
      <input
          v-model="username"
          type="text"
          placeholder="Es: admin..."
          style="padding: 10px; width: 300px; font-size: 16px;"
          @keyup.enter="cercaUtente"
      />
      <button
          @click="cercaUtente"
          style="padding: 10px 20px; font-size: 16px; cursor: pointer; margin-left: 10px;">
        Cerca
      </button>
    </div>

    <!-- BOX ERRORE SQL INJECTION (Sfondo rosso) -->
    <div v-if="errore" style="background-color: #ffe6e6; color: #d8000c; padding: 15px; border: 1px solid #d8000c; border-radius: 5px; margin-bottom: 20px; text-align: left;">
      <strong>⚠️ ATTENZIONE - Errore del Database:</strong><br><br>
      <span style="font-family: monospace;">{{ errore }}</span>
    </div>

    <!-- TABELLA DEI RISULTATI -->
    <table v-if="utenti.length > 0" style="width: 100%; border-collapse: collapse; text-align: left;">
      <thead>
      <tr style="background-color: #f2f2f2;">
        <th style="padding: 12px; border: 1px solid #ddd;">ID</th>
        <th style="padding: 12px; border: 1px solid #ddd;">Username</th>
        <th style="padding: 12px; border: 1px solid #ddd;">Email (o Segreto!)</th>
      </tr>
      </thead>
      <tbody>
      <tr v-for="user in utenti" :key="user.id">
        <td style="padding: 12px; border: 1px solid #ddd;">{{ user.id }}</td>
        <td style="padding: 12px; border: 1px solid #ddd; font-weight: bold;">{{ user.username }}</td>
        <td style="padding: 12px; border: 1px solid #ddd;">{{ user.email }}</td>
      </tr>
      </tbody>
    </table>

    <div v-if="!errore && utenti.length === 0 && haCercato" style="color: gray;">
      Nessun utente trovato.
    </div>
  </div>
</template>

<script>
// Importiamo il file axios.js che abbiamo sistemato nel passo precedente!
import api from '../services/axios.js';

export default {
  data() {
    return {
      username: '',
      utenti: [],
      errore: null,
      haCercato: false
    };
  },
  methods: {
    async cercaUtente() {
      if (!this.username) return;

      // Resettiamo la pagina prima di ogni nuova ricerca
      this.utenti = [];
      this.errore = null;
      this.haCercato = true;

      try {
        // Chiamata HTTP al nostro backend Go!
        const response = await api.get('/api/users', {
          params: { username: this.username }
        });

        // Se il backend ci restituisce l'errore SQL, lo salviamo per mostrarlo nel box rosso
        if (response.data.error) {
          this.errore = response.data.error;
        }
        // Altrimenti salviamo gli utenti
        else if (response.data.users) {
          this.utenti = response.data.users;
        }
      } catch (err) {
        console.error("Errore API:", err);
        this.errore = "Impossibile contattare il backend. Hai acceso il server Go (go run)?";
      }
    }
  }
};
</script>