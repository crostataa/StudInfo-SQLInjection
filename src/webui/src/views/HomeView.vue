<template>
  <div style="font-family: Arial, sans-serif; max-width: 800px; margin: 40px auto; text-align: center;">

    <h2>Ricerca Dipendenti Aziendali</h2>
    <p>Inserisci l'username del dipendente per visualizzare i suoi dati.</p>

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

    <!-- Visualizzazione errore SQL (Error-Based SQLi) -->
    <div v-if="errore" style="background-color: #ffe6e6; color: #d8000c; padding: 15px; border: 1px solid #d8000c; border-radius: 5px; margin-bottom: 20px; text-align: left;">
      <strong>⚠️ ATTENZIONE - Errore del Database:</strong><br><br>
      <span style="font-family: monospace;">{{ errore }}</span>
    </div>

    <table v-if="utenti.length > 0" style="width: 100%; border-collapse: collapse; text-align: left;">
      <thead>
      <tr style="background-color: #f2f2f2;">
        <th style="padding: 12px; border: 1px solid #ddd;">ID</th>
        <th style="padding: 12px; border: 1px solid #ddd;">Username</th>
        <th style="padding: 12px; border: 1px solid #ddd;">Email</th>
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

      this.utenti = [];
      this.errore = null;
      this.haCercato = true;

      try {
        const response = await api.get('/api/users', {
          params: { username: this.username }
        });

        if (response.data.error) {
          // Espone l'errore SQL restituito dal server
          this.errore = response.data.error;
        } else if (response.data.users) {
          this.utenti = response.data.users;
        }
      } catch (err) {
        console.error("Errore API:", err);
        this.errore = "Impossibile contattare il backend. Verifica che il server Go sia avviato.";
      }
    }
  }
};
</script>