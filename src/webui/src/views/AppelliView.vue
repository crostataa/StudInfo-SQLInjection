<template>
  <div class="appelli-container">
    <h2>Ricerca Appelli d'Esame</h2>
    <p class="subtitle">Cerca l'insegnamento per visualizzare le date disponibili e prenotarti.</p>

    <!-- BARRA DI RICERCA (Qui avverrà la magia oscura in Fase 4) -->
    <div class="search-box">
      <input
          type="text"
          v-model="ricerca"
          placeholder="Es: Sicurezza, Programmazione, Basi di Dati..."
          @keyup.enter="cercaAppelli"
      />
      <button @click="cercaAppelli" class="btn-search">Cerca</button>
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
          <th>Codice</th>
          <th>Insegnamento</th>
          <th>Docente</th>
          <th>Data</th>
          <th>Azione</th>
        </tr>
        </thead>
        <tbody>
        <tr v-for="appello in risultati" :key="appello.id">
          <td>{{ appello.codice }}</td>
          <td><strong>{{ appello.materia }}</strong></td>
          <td>{{ appello.docente }}</td>
          <td>{{ appello.data }}</td>
          <td>
            <button class="btn-book" @click="prenota(appello.id)">Prenota</button>
          </td>
        </tr>
        </tbody>
      </table>
    </div>

    <div v-else-if="haCercato && !errore" class="no-results">
      Nessun appello trovato per la ricerca effettuata.
    </div>
  </div>
</template>

<script>
export default {
  name: 'AppelliView',
  data() {
    return {
      ricerca: '',
      errore: null,
      haCercato: false,
      // DATI FINTI: simuliamo la risposta del database
      tuttiAppelli: [
        { id: 1, codice: 'INF01', materia: 'Sicurezza Informatica', docente: 'Prof. Hacker', data: '15/10/2026' },
        { id: 2, codice: 'INF02', materia: 'Basi di Dati', docente: 'Prof. Query', data: '22/10/2026' },
        { id: 3, codice: 'MAT01', materia: 'Analisi Matematica', docente: 'Prof. Integrale', data: '10/11/2026' }
      ],
      risultati: []
    }
  },
  mounted() {
    // Mostriamo subito tutti gli appelli di default
    this.risultati = this.tuttiAppelli;
  },
  methods: {
    cercaAppelli() {
      this.haCercato = true;
      this.errore = null;

      // MOCKUP: Simuliamo una ricerca finta filtrando l'array Javascript
      // Nella fase 3 sostituiremo questo con la vera chiamata Axios al server Go!
      if (this.ricerca.trim() === '') {
        this.risultati = this.tuttiAppelli;
      } else {
        const query = this.ricerca.toLowerCase();
        this.risultati = this.tuttiAppelli.filter(a =>
            a.materia.toLowerCase().includes(query)
        );
      }
    },
    prenota(id) {
      alert("Mockup: Hai prenotato l'appello con ID: " + id + ". Nella versione finale, questo salverà i dati nel DB!");
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