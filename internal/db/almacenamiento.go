package db

import (
<<<<<<< HEAD
	"database/sql"
	"log"
	"os"
	"path/filepath"
	"pec2/internal/models"
	_ "modernc.org/sqlite"
)

var DB *sql.DB

func InitDB() {
	var err error
	os.MkdirAll(filepath.Join("internal", "db"), 0755)
	dbPath := filepath.Join("internal", "db", "platinum.db")

	DB, err = sql.Open("sqlite", dbPath)
	if err != nil {
		log.Fatal("Error abriendo base de datos SQLite:", err)
	}

	createTables()
}

func createTables() {
	userTableInfo := `
	CREATE TABLE IF NOT EXISTS usuarios (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		nombre TEXT NOT NULL,
		apellidos TEXT,
		fecha_nacimiento TEXT,
		direccion TEXT,
		telefono TEXT,
		correo TEXT,
		documento TEXT,
		metodo_pago TEXT,
		numero_pago TEXT,
		password TEXT NOT NULL,
		suscripcion_activa BOOLEAN DEFAULT 1
	);`

	_, err := DB.Exec(userTableInfo)
	if err != nil {
		log.Fatal("Error creando tabla usuarios: ", err)
	}

	resenaTableInfo := `
	CREATE TABLE IF NOT EXISTS resenas (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		autor TEXT NOT NULL,
		puntuacion INTEGER NOT NULL,
		texto TEXT NOT NULL
	);`

	_, err = DB.Exec(resenaTableInfo)
	if err != nil {
		log.Fatal("Error creando tabla resenas: ", err)
	}

	// No insertamos usuarios de prueba por defecto, el usuario se registrará.
	var count int

	DB.QueryRow("SELECT COUNT(*) FROM resenas").Scan(&count)
	if count == 0 {
		_, _ = DB.Exec(`INSERT INTO resenas (autor, puntuacion, texto) VALUES 
		('Carlos G.', 5, 'El mejor gimnasio de la ciudad. Las máquinas son increíbles.'),
		('María P.', 4, 'Muy buen ambiente, aunque a veces hay mucha gente en hora punta.'),
		('Luis R.', 5, 'Los fisioterapeutas me curaron una lesión de hombro que llevaba meses arrastrando.')`)
	}
}

func GuardarUsuario(u models.Usuario) error {
	var count int
	DB.QueryRow("SELECT COUNT(*) FROM usuarios WHERE nombre = ?", u.Nombre).Scan(&count)
	if count > 0 {
		// Update if exists or handle otherwise, for simplicity insert a new one if it has unique name? 
		// Actually let's just insert
	}

	stmt, err := DB.Prepare(`INSERT INTO usuarios (nombre, apellidos, fecha_nacimiento, direccion, telefono, correo, documento, metodo_pago, numero_pago, password) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(u.Nombre, u.Apellidos, u.FechaNacimiento, u.Direccion, u.Telefono, u.Correo, u.Documento, u.MetodoPago, u.NumeroPago, u.Password)
	return err
=======
	"encoding/json"
	"os"
	"path/filepath"
	"pec2/internal/models"
)

var dbFile = filepath.Join("internal", "db", "datos_registro.json")
var resenasFile = filepath.Join("internal", "db", "datos_resenas.json")

func InitDB() {
	// JSON no requiere inicialización estricta por adelantado a menos que queramos crear los ficheros
	// pero os.WriteFile lo creará si no existe en GuardarUsuario.
	// Solo creamos la carpeta base por seguridad.
	os.MkdirAll(filepath.Join("internal", "db"), 0755)
}

func GuardarUsuario(u models.Usuario) error {
	var usuarios []models.Usuario

	file, err := os.ReadFile(dbFile)
	if err == nil {
		json.Unmarshal(file, &usuarios)
	}

	usuarios = append(usuarios, u)

	data, err := json.MarshalIndent(usuarios, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(dbFile, data, 0644)
>>>>>>> d48f6ffdbdb90e0d503e476e6ffbce582ca54153
}

func ObtenerResenas() []models.Resena {
	var resenas []models.Resena
<<<<<<< HEAD
	rows, err := DB.Query("SELECT autor, puntuacion, texto FROM resenas ORDER BY id DESC")
	if err != nil {
		log.Println("Error obteniendo reseñas:", err)
		return resenas
	}
	defer rows.Close()

	for rows.Next() {
		var r models.Resena
		if err := rows.Scan(&r.Autor, &r.Puntuacion, &r.Texto); err == nil {
			resenas = append(resenas, r)
=======
	file, err := os.ReadFile(resenasFile)
	if err == nil {
		json.Unmarshal(file, &resenas)
	} else {
		// Mock data if file doesn't exist
		resenas = []models.Resena{
			{Autor: "Carlos G.", Puntuacion: 5, Texto: "El mejor gimnasio de la ciudad. Las máquinas son increíbles."},
			{Autor: "María P.", Puntuacion: 4, Texto: "Muy buen ambiente, aunque a veces hay mucha gente en hora punta."},
			{Autor: "Luis R.", Puntuacion: 5, Texto: "Los fisioterapeutas me curaron una lesión de hombro que llevaba meses arrastrando."},
>>>>>>> d48f6ffdbdb90e0d503e476e6ffbce582ca54153
		}
	}
	return resenas
}

func GuardarResena(r models.Resena) error {
<<<<<<< HEAD
	stmt, err := DB.Prepare("INSERT INTO resenas (autor, puntuacion, texto) VALUES (?, ?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(r.Autor, r.Puntuacion, r.Texto)
	return err
}

func ObtenerUsuarioPorCorreo(correo string) *models.Usuario {
	var u models.Usuario
	err := DB.QueryRow("SELECT id, nombre, apellidos, correo, direccion, telefono, metodo_pago FROM usuarios WHERE correo = ?", correo).
		Scan(&u.ID, &u.Nombre, &u.Apellidos, &u.Correo, &u.Direccion, &u.Telefono, &u.MetodoPago)
	if err != nil {
		return nil
	}
	return &u
=======
	resenas := ObtenerResenas()
	resenas = append(resenas, r)

	data, err := json.MarshalIndent(resenas, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(resenasFile, data, 0644)
>>>>>>> d48f6ffdbdb90e0d503e476e6ffbce582ca54153
}
