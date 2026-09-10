<template>
  <div class="prenotazioni-container">
    <h2>Le tue Prenotazioni</h2>
    <p class="subtitle">Esami a cui sei attualmente iscritta/o.</p>

    <!-- BARRA DI RICERCA -->
    <div class="search-box">
      <input
          type="text"
          v-model="ricerca"
          placeholder="Es: Sicurezza, Programmazione, Basi di Dati..."
          @keyup.enter="cercaPrenotazioni"
      />
      <button @click="cercaPrenotazioni" class="btn-search">Cerca</button>
    </div>

    <!-- BOX ERRORE SQL INJECTION (Nascosto, si attiverà in Fase 4) -->
    <div v-if="errore" class="error-box">
      <strong>⚠️ Errore di sistema:</strong><br>
      <span>{{ errore }}</span>
    </div>

    <!-- TABELLA RISULTATI -->
    <div v-if="risultati.length > 0" class="results-section">
      <table class="data-table">
        <thead>
        <tr>
          <th>Id prenotazione</th>
          <!-- COLONNA MATRICOLA RIMOSSA -->
          <th>Id appello</th>
          <th>Insegnamento</th>
          <th>Data esame</th>
          <th>Aula</th>
          <th>Stato</th>
          <th>Azioni</th>
        </tr>
        </thead>
        <tbody>
        <tr v-for="prenotazione in risultati" :key="prenotazione.id_prenotazione">
          <td>{{ prenotazione.id_prenotazione }}</td>
          <!-- DATO MATRICOLA RIMOSSO -->
          <td>{{ prenotazione.id_appello }}</td>
          <td>{{ prenotazione.insegnamento }}</td>
          <td>{{ prenotazione.data_esame }}</td>
          <td>{{ prenotazione.aula }}</td>
          <td>{{ prenotazione.stato }}</td>
          <td>
            <!-- (Nota: se queste sono le prenotazioni già effettuate, potresti voler rinominare il bottone in "Annulla" anziché "Prenota"!) -->
            <button class="btn-book" @click="annulla(prenotazione.id)">Annulla</button>
          </td>
        </tr>
        </tbody>
      </table>
    </div>

    <div v-else-if="haCercato && !errore" class="no-results">
      Nessuna prenotazione trovata per la ricerca effettuata.
    </div>
  </div>
</template>

<script>
export default {
  name: 'PrenotazioniView',
  data() {
  return {
    ricerca: '',
    errore: null,
    haCercato: false,
    risultati: []
  }
},
mounted() {
    this.cercaPrenotazioni();
},
  methods: {
    async cercaPrenotazioni() {
      // 1. Prendi la matricola e resetta gli errori precedenti
      const matricola = localStorage.getItem('utente_loggato');
      this.errore = null;
      this.haCercato = true;

      if (!matricola) {
        this.errore = "Devi effettuare l'accesso per vedere le prenotazioni!";
        return;
      }

      try {
        // 2. Chiamata al server Go
        const url = `http://127.0.0.1:3000/api/prenotazioni?q=${this.ricerca}&matricola=${matricola}`;
        const response = await fetch(url);
        const data = await response.json();

        // 3. Gestione dei dati
        if (data.error) {
          // Se Go ci manda un errore (es. SQL Injection fail), lo mostriamo
          this.errore = data.error;
          this.risultati = [];
        } else if (data.prenotazione) {
          // Se ci sono risultati, popoliamo la tabella
          this.risultati = data.prenotazione;
        } else {
          this.risultati = [];
        }

      } catch (err) {
        this.errore = "Errore di connessione al server: " + err.message;
      }
    },
  annulla(id) {
    alert("funzione di cancellazione in costruzione! ID: " + id);
  }
}
}
</script>

<style scoped>
.appelli-container {
  display: flex;
  flex-direction: column;
  gap: 20px;
}

h2 {
  color: #812936;
  margin-bottom: 0;
}

.subtitle {
  color: #718096;
  margin-top: 5px;
}

.search-box {
  display: flex;
  gap: 10px;
  background: white;
  padding: 20px;
  border-radius: 8px;
  box-shadow: 0 2px 4px rgba(0,0,0,0.05);
}

.search-box input {
  flex: 1;
  padding: 12px;
  border: 1px solid #cbd5e0;
  border-radius: 4px;
  font-size: 16px;
}

.search-box input:focus {
  outline: none;
  border-color: #812936;
  box-shadow: 0 0 0 3px rgba(43, 108, 176, 0.2);
}

.btn-search {
  background-color: #6a1a21;
  color: white;
  border: none;
  padding: 0 25px;
  border-radius: 4px;
  font-size: 16px;
  font-weight: bold;
  cursor: pointer;
  transition: background-color 0.3s;
}

.btn-search:hover {
  background-color: #812936;
}

.error-box {
  background-color: #ffe6e6;
  color: #d8000c;
  padding: 15px;
  border-left: 5px solid #d8000c;
  border-radius: 4px;
  font-family: monospace;
}

.data-table {
  width: 100%;
  border-collapse: collapse;
  background: white;
  box-shadow: 0 1px 3px rgba(0,0,0,0.1);
  border-radius: 8px;
  overflow: hidden;
}

.data-table th, .data-table td {
  padding: 15px;
  text-align: left;
  border-bottom: 1px solid #e2e8f0;
}

.data-table th {
  background-color: #f7fafc;
  font-weight: 600;
  color: #4a5568;
}

.btn-book {
  background-color: #48bb78; /* Verde per azioni positive */
  color: white;
  border: none;
  padding: 8px 15px;
  border-radius: 4px;
  font-weight: bold;
  cursor: pointer;
}

.btn-book:hover {
  background-color: #38a169;
}

.no-results {
  text-align: center;
  padding: 30px;
  color: #a0aec0;
  background: white;
  border-radius: 8px;
}
</style>