<template>
  <div class="dashboard">

    <!-- Avviso in caso di errore -->
    <div v-if="errore" class="error-box" style="color: red; margin-bottom: 15px;">
      <strong>⚠️ Errore:</strong> {{ errore }}
    </div>

    <!-- Mostriamo la dashboard solo a caricamento completato -->
    <div v-if="!caricamento">
      <header class="dashboard-header">
        <h2>Benvenuta/o, {{ studente.nome }} {{ studente.cognome }}</h2>
        <p class="matricola">Matricola: {{ studente.matricola }}</p>
      </header>

      <!-- Riepilogo Statistiche -->
      <div class="stats-grid">
        <div class="stat-card">
          <h3>Media Ponderata</h3>
          <div class="stat-value">{{ statistiche.media }}</div>
        </div>
        <div class="stat-card">
          <h3>CFU Acquisiti</h3>
          <div class="stat-value">{{ statistiche.cfu }} / 120</div>
        </div>
        <div class="stat-card">
          <h3>Esami Sostenuti</h3>
          <div class="stat-value">{{ statistiche.esamiSostenuti }}</div>
        </div>
      </div>

      <!-- Prossimi impegni (Prenotazioni) -->
      <div class="upcoming-section">
        <h3>I tuoi prossimi appelli</h3>
        <table class="data-table" v-if="prossimiAppelli.length > 0">
          <thead>
          <tr>
            <th>Insegnamento</th>
            <th>Data Esame</th>
            <th>Aula</th>
            <th>Stato</th>
          </tr>
          </thead>
          <tbody>
          <tr v-for="appello in prossimiAppelli" :key="appello.id_prenotazione">
            <td><strong>{{ appello.insegnamento }}</strong></td>
            <td>{{ appello.data_esame }}</td>
            <td>{{ appello.aula }}</td>
            <td>{{ appello.stato }}</td>
          </tr>
          </tbody>
        </table>
        <p v-else class="no-data">Non hai appelli prenotati a breve.</p>
      </div>
    </div>

    <div v-else>
      <p>Caricamento della dashboard in corso...</p>
    </div>

  </div>
</template>

<script>
export default {
  name: 'DashboardView',
  data() {
    return {
      caricamento: true,
      errore: null,
      studente: {
        nome: '',
        cognome: '',
        matricola: ''
      },
      statistiche: {
        media: '0.00',
        cfu: 0,
        esamiSostenuti: 0
      },
      prossimiAppelli: []
    }
  },
  async mounted() {
    const matricola = localStorage.getItem('utente_loggato');

    if (!matricola) {
      this.errore = "Nessun utente loggato. Effettua l'accesso.";
      this.caricamento = false;
      return;
    }

    try {
      // 1. PRENDIAMO I DATI DEL PROFILO
      const resProfilo = await fetch(`http://127.0.0.1:3000/api/profilo/${matricola}`);
      const dataProfilo = await resProfilo.json();
      if (dataProfilo.studente && dataProfilo.studente.length > 0) {
        this.studente = dataProfilo.studente[0];
      }

      // 2. PRENDIAMO IL LIBRETTO E CALCOLIAMO LE STATISTICHE
      const resLibretto = await fetch(`http://127.0.0.1:3000/api/libretto/${matricola}`);
      const dataLibretto = await resLibretto.json();

      if (dataLibretto.libretto) {
        const esami = dataLibretto.libretto;
        this.statistiche.esamiSostenuti = esami.length;

        let sommaCfu = 0;
        let sommaVotiPonderati = 0;

        // Calcolo della media ponderata: (Voto * CFU) / (Totale CFU)
        esami.forEach(esame => {
          sommaCfu += esame.cfu;
          sommaVotiPonderati += (esame.voto * esame.cfu);
        });

        this.statistiche.cfu = sommaCfu;
        if (sommaCfu > 0) {
          // Arrotonda a due cifre decimali
          this.statistiche.media = (sommaVotiPonderati / sommaCfu).toFixed(2);
        }
      }

      // 3. PRENDIAMO LE PRENOTAZIONI FUTURE
      const resPrenotazioni = await fetch(`http://127.0.0.1:3000/api/prenotazioni?q=&matricola=${matricola}`);
      const dataPrenotazioni = await resPrenotazioni.json();

      if (dataPrenotazioni.prenotazione) {
        this.prossimiAppelli = dataPrenotazioni.prenotazione;
      }

    } catch (err) {
      this.errore = "Impossibile caricare la dashboard: " + err.message;
    } finally {
      this.caricamento = false; // Togliamo la scritta di caricamento
    }
  }
}
</script>

<style scoped>
.dashboard {
  display: flex;
  flex-direction: column;
  gap: 30px;
}

.dashboard-header h2 {
  margin-bottom: 5px;
  color: #812936;
}

.matricola {
  color: #718096;
  font-size: 1.1rem;
  margin-top: 0;
}

.stats-grid {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(200px, 1fr));
  gap: 20px;
}

.stat-card {
  background: white;
  padding: 20px;
  border-radius: 8px;
  box-shadow: 0 2px 4px rgba(0,0,0,0.05);
  text-align: center;
  border-top: 4px solid #812936;
}

.stat-card h3 {
  font-size: 1rem;
  color: #4a5568;
  margin-bottom: 10px;
  margin-top: 0;
}

.stat-value {
  font-size: 2rem;
  font-weight: bold;
  color: #812936;
}

.upcoming-section h3 {
  color: #812936;
  border-bottom: 2px solid #e2e8f0;
  padding-bottom: 10px;
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
  color: #812936;
}

.data-table tr:last-child td {
  border-bottom: none;
}

.no-data {
  color: #718096;
  font-style: italic;
}
</style>