CREATE DATABASE IF NOT EXISTS vulnerabile_db;
USE vulnerabile_db;

-- 1. TABELLA STUDENTI
CREATE TABLE IF NOT EXISTS students (
                                        matricola INT PRIMARY KEY,
                                        nome VARCHAR(15) NOT NULL,
    cognome VARCHAR(50) NOT NULL,
    data_di_nascita DATE NOT NULL,
    indirizzo VARCHAR(50) NOT NULL,
    password VARCHAR(20) NOT NULL,
    email VARCHAR(100) NOT NULL
    );

INSERT INTO students (matricola, nome, cognome, data_di_nascita, indirizzo, password, email) VALUES
                                                                                                 (2112808, 'Cristina', 'Porceddu', '2004-10-14', 'via scarpa 71', 'Sapienza123$', 'porceddu.2112808@studenti.uniroma1.it'),
                                                                                                 (2082152, 'Marco', 'Mazzillo', '2003-10-23', 'piazzale aldo moro 30', 'rubberDuck09?', 'mazzillo.2082152@studenti.uniroma1.it'),
                                                                                                 (2014782, 'Roberto','Sacchetti','2002-02-03', 'via del castro laurenziano 5', 'passworD1234!', 'sacchetti.2014782@studenti.uniroma1.it');

-- 2. TABELLA APPELLI
CREATE TABLE IF NOT EXISTS appelli (
                                       id INT AUTO_INCREMENT PRIMARY KEY,
                                       insegnamento VARCHAR(50) NOT NULL,
    docente VARCHAR(50) NOT NULL,
    data DATE NOT NULL
    );

INSERT INTO appelli (id, insegnamento, data, docente) VALUES
                                                          (01, 'Basi di Dati Modulo II', '2026-09-11', 'Toni Mancini'),
                                                          (02, 'Sicurezza', '2026-09-16', 'Emiliano Casalicchio'),
                                                          (03, 'Analisi Modulo 1', '2026-09-21', 'Luigi Orsina'),
                                                          (04, 'Sistemi Operativi', '2026-09-14', 'Fabio de Gaspari'),
                                                          (05, 'Progettazione di Sistemi Digitali', '2026-09-19', 'Annalisa Massini'),
                                                          (06, 'Calcolo delle Probabilità', '2026-09-20', 'Alessandra Faggionato'),
                                                          (07, 'Ingegneria del Software', '2026-09-11', 'Enrico Tronci'),
                                                          (08, 'Programmazione di Sistemi Embedded e Multicore', '2026-09-06', 'Daniele De Sensi'),
                                                          (09, 'Reti di Calcolatori', '2026-09-22', 'Riccardo Beraldi'),
                                                          (10, 'Linguaggi di Programmazione', '2026-09-23', 'Ivano Salvo'),
                                                          (11, 'Interazione Uomo-Macchina', '2026-09-24', 'Tiziana Catarci'),
                                                          (12, 'Algoritmi e Strutture Dati', '2026-09-25', 'Camil Demetrescu'),
                                                          (13, 'Basi di Dati Modulo I', '2026-09-26', 'Riccardo Rosati'),
                                                          (14, 'Fisica', '2026-09-27', 'Paolo Mataloni'),
                                                          (15, 'Ricerca Operativa', '2026-09-28', 'Massimo Roma'),
                                                          (16, 'Intelligenza Artificiale', '2026-09-29', 'Giuseppe De Giacomo'),
                                                          (17, 'Metodi Matematici per l''Informatica', '2026-10-01', 'Rossella Petreschi'),
                                                          (18, 'Programmazione di Rete', '2026-10-02', 'Emiliano Casalicchio'),
                                                          (19, 'Sistemi Distribuiti', '2026-10-03', 'Roberto Baldoni'),
                                                          (20, 'Architettura dei Calcolatori', '2026-10-04', 'Irene Finocchi'),
                                                          (21, 'Sviluppo Web', '2026-10-05', 'Toni Mancini'),
                                                          (22, 'Machine Learning', '2026-10-06', 'Roberto Navigli'),
                                                          (23, 'Computer Vision', '2026-10-07', 'Barbara Caputo'),
                                                          (24, 'Fondamenti di Informatica', '2026-10-08', 'Silvia Chiusano');

-- 3. TABELLA PRENOTAZIONI
CREATE TABLE IF NOT EXISTS prenotazioni_esami (
                                                  id_prenotazione INT AUTO_INCREMENT PRIMARY KEY,
                                                  matricola INT NOT NULL,
                                                  id_appello INT NOT NULL,
                                                  insegnamento VARCHAR(50) NOT NULL,
    data_esame DATE NOT NULL,
    aula VARCHAR(50) NOT NULL,
    stato ENUM('Confermato', 'In attesa', 'Rifiutato') NOT NULL,
    FOREIGN KEY (id_appello) REFERENCES appelli (id),
    FOREIGN KEY (matricola) REFERENCES students (matricola)
    );

INSERT INTO prenotazioni_esami (matricola, id_appello, insegnamento, data_esame, aula, stato) VALUES
                                                                                                  (2112808, 01, 'Basi di Dati Modulo II', '2026-09-11', 'Aula Cabibbo', 'Confermato'),
                                                                                                  (2014782, 03, 'Analisi Modulo 1', '2026-09-21', 'Aula 3', 'In attesa'),
                                                                                                  (2082152, 04, 'Sistemi Operativi', '2026-09-14', 'Laboratori 2', 'Confermato'),
                                                                                                  (2112808, 02, 'Sicurezza', '2026-09-16', 'Aula Magna', 'Confermato'),
                                                                                                  (2112808, 12, 'Algoritmi e Strutture Dati', '2026-09-25', 'Aula 4', 'In attesa'),
                                                                                                  (2082152, 08, 'Programmazione di Sistemi Embedded e Multicore', '2026-09-06', 'Laboratori 1', 'Confermato'),
                                                                                                  (2082152, 09, 'Reti di Calcolatori', '2026-09-22', 'Aula Cabibbo', 'In attesa'),
                                                                                                  (2014782, 07, 'Ingegneria del Software', '2026-09-11', 'Aula Magna', 'Confermato'),
                                                                                                  (2014782, 11, 'Interazione Uomo-Macchina', '2026-09-24', 'Aula 3', 'Confermato');

-- 4. TABELLA LIBRETTO
CREATE TABLE IF NOT EXISTS libretto_esami (
                                              id_registrazione INT AUTO_INCREMENT PRIMARY KEY,
                                              matricola INT NOT NULL,
                                              id_appello INT NOT NULL,
                                              insegnamento VARCHAR(50) NOT NULL,
    data_registrazione DATE NOT NULL,
    voto INT NOT NULL,
    CFU INT NOT NULL,
    FOREIGN KEY (matricola) REFERENCES students (matricola)
    );

INSERT INTO libretto_esami (matricola, id_appello, insegnamento, data_registrazione, voto, CFU) VALUES
                                                                                                    (2014782, 01, 'Basi di Dati Modulo II', '2026-09-20', 23, 6),
                                                                                                    (2082152, 04, 'Sistemi Operativi', '2026-09-16', 30, 9),
                                                                                                    (2112808, 13, 'Basi di Dati Modulo I', '2025-06-15', 30, 9),
                                                                                                    (2112808, 14, 'Fisica', '2025-07-02', 27, 9),
                                                                                                    (2082152, 06, 'Calcolo delle Probabilità', '2025-02-12', 24, 6),
                                                                                                    (2082152, 15, 'Ricerca Operativa', '2025-06-25', 28, 6),
                                                                                                    (2014782, 05, 'Progettazione di Sistemi Digitali', '2025-09-10', 22, 9),
                                                                                                    (2014782, 10, 'Linguaggi di Programmazione', '2025-07-20', 26, 6);