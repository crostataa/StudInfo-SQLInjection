<template>
  <div class="login-container">
    <div class="login-box">
      <div class="logo">🐦‍🔥</div>
      <h2>StudInfo</h2>
      <p class="subtitle">inserisci le tue credenziali</p>

      <form @submit.prevent="effettuaLogin" class="login-form">
        <div class="input-group">
          <label for="matricola">Matricola</label>
          <input
              type="text"
              id="matricola"
              v-model="matricola"
              placeholder="es. 123456"
              required
          />
        </div>

        <div class="input-group">
          <label for="password">Password</label>
          <input
              type="password"
              id="password"
              v-model="password"
              placeholder="••••••••"
              required
          />
        </div>

        <button type="submit" class="btn-login">Accedi</button>
      </form>

      <!-- Mostriamo un messaggio se sbagliano (per ora finto) -->
      <div v-if="errore" class="error-msg">
        Credenziali errate. Riprova.
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: 'LoginView',
  data() {
    return {
      matricola: '',
      password: '',
      errore: false
    }
  },
  methods: {
    async effettuaLogin() {
      try {
        const response = await fetch('http://127.0.0.1:3000/api/login', {
          method: 'POST',
          headers: {'Content-Type': 'application/json'},
          body: JSON.stringify({
            matricola: parseInt(this.matricola), // il tuo v-model
            password: this.password
          })
        });
        const data = await response.json();

        if (data.success) {
          // Magia: salviamo la matricola nel browser!
          localStorage.setItem('utente_loggato', data.matricola);
          alert(`Benvenut* ${data.nome}!`);
          this.$router.push('/dashboard'); // Vai alla pagina principale
        } else {
          this.errore = data.error;
        }
      } catch (err) {
        this.errore = "Errore di connessione";
      }
    }
  }
}
</script>

<style scoped>
.login-container {
  display: flex;
  justify-content: center;
  align-items: center;
  min-height: 100vh;
  background-color: #e2e8f0; /* Grigio chiaro di sfondo */
}

.login-box {
  background: white;
  padding: 40px;
  border-radius: 8px;
  box-shadow: 0 4px 6px rgba(0, 0, 0, 0.1);
  width: 100%;
  max-width: 400px;
  text-align: center;
}

.logo {
  font-size: 3rem;
  margin-bottom: 10px;
}

h2 {
  color: #812936;
  margin: 0;
}

.subtitle {
  color: #718096;
  margin-bottom: 30px;
  font-size: 0.9rem;
}

.login-form {
  display: flex;
  flex-direction: column;
  gap: 20px;
}

.input-group {
  text-align: left;
}

.input-group label {
  display: block;
  margin-bottom: 5px;
  font-weight: 500;
  color: #4a5568;
}

.input-group input {
  width: 100%;
  padding: 10px;
  border: 1px solid #cbd5e0;
  border-radius: 4px;
  font-size: 16px;
  box-sizing: border-box; /* Previene sbavature fuori dal form */
}

.input-group input:focus {
  outline: none;
  border-color: #812936;
  box-shadow: 0 0 0 3px rgba(43, 108, 176, 0.2);
}

.btn-login {
  background-color: #b02a37;
  color: white;
  border: none;
  padding: 12px;
  border-radius: 4px;
  font-size: 16px;
  font-weight: bold;
  cursor: pointer;
  transition: background-color 0.3s;
}

.btn-login:hover {
  background-color: #812936;
}

.error-msg {
  color: #e53e3e;
  margin-top: 15px;
  font-size: 0.9rem;
}
</style>