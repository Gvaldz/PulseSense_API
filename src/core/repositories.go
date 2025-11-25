package core

import (
	"database/sql"
	"fmt"
	"pulse_sense/src/internal/services/auth/domain"
	users "pulse_sense/src/internal/users/domain"
)

type AuthRepository struct {
	DB *sql.DB
}

func NewAuthRepository(DB *sql.DB) domain.AuthRepository {
	return &AuthRepository{DB: DB}
}

func (r *AuthRepository) FindUserByEmail(email string) (users.User, error) {
	var user users.User
	query := `SELECT idusuario, correo, contrasena, idtipo_usuario FROM usuarios WHERE correo = ?`
	err := r.DB.QueryRow(query, email).Scan(&user.IdUsuario, &user.Correo, &user.Contrasena, &user.Tipo)
	return user, err
}

func (r *AuthRepository) UpdateLastLogin(userID int32) error {
	query := `UPDATE usuarios SET ultimo_login = NOW() WHERE idusuario = ?`
	_, err := r.DB.Exec(query, userID)
	return err
}

func (r *AuthRepository) FindUserByID(userID int32) (users.User, error) {
	var user users.User
	query := `SELECT idusuario, correo, contrasena, idtipo_usuario FROM usuarios WHERE idusuario = ?`
	err := r.DB.QueryRow(query, userID).Scan(&user.IdUsuario, &user.Correo, &user.Contrasena, &user.Tipo)
	return user, err
}

type UserRepository struct {
	DB *sql.DB
}

func NewUserRepository(DB *sql.DB) users.UserRepository {
	return &UserRepository{DB: DB}
}

func (r *UserRepository) CreateUser(user users.User) (users.User, error) {
	result, err := r.DB.Exec(
		"INSERT INTO usuarios (nombres, correo, contrasena) VALUES (?, ?, ?)",
		user.Nombres, user.Correo, user.Contrasena,
	)
	if err != nil {
		return users.User{}, fmt.Errorf("error al crear usuario: %w", err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return users.User{}, fmt.Errorf("error al obtener ID: %w", err)
	}

	return users.User{
		IdUsuario: int32(id),
		Nombres:   user.Nombres,
		Correo:    user.Correo,
	}, nil
}

func (r *UserRepository) GetAllUsers() ([]users.User, error) {
	query := "SELECT idusuario, nombres, correo FROM usuarios"
	rows, err := r.DB.Query(query)
	if err != nil {
		return nil, fmt.Errorf("error al obtener usuarios: %w", err)
	}
	defer rows.Close()

	var usersList []users.User
	for rows.Next() {
		var user users.User
		if err := rows.Scan(&user.IdUsuario, &user.Nombres, &user.Correo); err != nil {
			return nil, fmt.Errorf("error al escanear user: %w", err)
		}
		usersList = append(usersList, user)
	}

	return usersList, nil
}

func (r *UserRepository) GetUserByID(iduser int32) (users.User, error) {
	if r.DB == nil {
		return users.User{}, fmt.Errorf("database connection is nil")
	}

	var user users.User
	query := "SELECT idusuario, nombres, correo, FCMtoken FROM usuarios WHERE idusuario = ?"
	err := r.DB.QueryRow(query, iduser).Scan(&user.IdUsuario, &user.Nombres, &user.Correo, &user.FCMToken)
	if err != nil {
		return user, fmt.Errorf("error al obtener usuario: %w", err)
	}
	return user, nil
}

func (r *UserRepository) GetNursePerHospital(idhospital int32) ([]users.User, error) {
	query := `SELECT usuarios.idusuario, usuarios.nombres, usuarios.apellido_p, usuarios.apellido_m, usuarios.correo 
	FROM usuarios 
	JOIN trabajadores ON usuarios.idusuario = trabajadores.idusuario 
	WHERE usuarios.idtipo_usuario = 2 AND trabajadores.idhospital = ?`

	rows, err := r.DB.Query(query, idhospital)
	if err != nil {
		return nil, fmt.Errorf("error al obtener usuarios: %w", err)
	}
	defer rows.Close()

	var usersList []users.User
	for rows.Next() {
		var user users.User
		if err := rows.Scan(&user.IdUsuario, &user.Nombres, &user.Apellido_p, &user.Apellido_m, &user.Correo); err != nil {
			return nil, fmt.Errorf("error al escanear user: %w", err)
		}
		usersList = append(usersList, user)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error después de iterar filas: %w", err)
	}

	return usersList, nil
}

func (r *UserRepository) GetNursePerPatient(idhospital int32) ([]users.User, error) {
	query := `SELECT usuarios.idusuario, usuarios.nombres, usuarios.apellido_p, usuarios.apellido_m, usuarios.correo 
	FROM usuarios 
	JOIN cuidadores ON usuarios.idusuario = cuidadores.idusuario 
	WHERE usuarios.idtipo_usuario = 2 AND cuidadores.idpaciente = ?`

	rows, err := r.DB.Query(query, idhospital)
	if err != nil {
		return nil, fmt.Errorf("error al obtener usuarios: %w", err)
	}
	defer rows.Close()

	var usersList []users.User
	for rows.Next() {
		var user users.User
		if err := rows.Scan(&user.IdUsuario, &user.Nombres, &user.Apellido_p, &user.Apellido_m, &user.Correo); err != nil {
			return nil, fmt.Errorf("error al escanear user: %w", err)
		}
		usersList = append(usersList, user)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error después de iterar filas: %w", err)
	}

	return usersList, nil
}

func (r *UserRepository) GetUserByEmail(email string) (users.User, error) {
	var user users.User
	query := "SELECT idusuario, nombres, correo, contrasena FROM usuarios WHERE correo = ?"
	err := r.DB.QueryRow(query, email).Scan(
		&user.IdUsuario,
		&user.Nombres,
		&user.Correo,
		&user.Contrasena,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return user, fmt.Errorf("usuario no encontrado")
		}
		return user, fmt.Errorf("error al obtener usuario por email: %w", err)
	}
	return user, nil
}

func (r *UserRepository) UpdateUser(id int32, user users.User) error {
	query := "UPDATE usuarios SET nombres = ?, correo = ? WHERE idusuario = ?"
	result, err := r.DB.Exec(query, user.Nombres, user.Correo, id)
	if err != nil {
		return fmt.Errorf("error al actualizar usuario: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("error al verificar actualización: %w", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("usuario no encontrado")
	}

	return nil
}

func (r *UserRepository) UpdatePassword(id int32, newHashedPassword string) error {
	query := "UPDATE usuarios SET contrasena = ? WHERE idusuario = ?"
	result, err := r.DB.Exec(query, newHashedPassword, id)
	if err != nil {
		return fmt.Errorf("error al actualizar contraseña: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("error al verificar actualización de contraseña: %w", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("usuario no encontrado")
	}

	return nil
}

func (r *UserRepository) DeleteUser(id int32) error {
	query := "DELETE FROM usuarios WHERE idusuario = ?"
	result, err := r.DB.Exec(query, id)
	if err != nil {
		return fmt.Errorf("error al eliminar usuario: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("error al verificar eliminación: %w", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("usuario no encontrado")
	}

	return nil
}

func (r *UserRepository) UpdateFCMToken(userID string, token string) error {
	query := "UPDATE usuarios SET fcm_token = ? WHERE idusuario = ?"
	_, err := r.DB.Exec(query, token, userID)
	return err
}

func (r *UserRepository) GetFCMToken(userID int32) (string, error) {
	var token string
	query := "SELECT fcm_token FROM usuarios WHERE idusuario = ?"
	err := r.DB.QueryRow(query, userID).Scan(&token)
	if err != nil {
		return "", err
	}
	return token, nil
}

func (r *UserRepository) GetDoctorsByPatientId(iduser int32) ([]users.User, error) {
    query := `SELECT 
        d.idusuario,
        d.nombres,
        d.apellido_m,
        d.apellido_p,
        d.correo,
        CONCAT(p.nombres, ' ', p.apellido_p, ' ', p.apellido_m) AS nombre_paciente,
        p.idpaciente
    FROM 
        cuidadores c
    JOIN 
        pacientes p ON c.idpaciente = p.idpaciente
    JOIN 
        usuarios d ON p.id_doctor = d.idusuario
    WHERE 
        c.idusuario = ?`

    rows, err := r.DB.Query(query, iduser)
    if err != nil {
        return nil, fmt.Errorf("error al obtener doctores por paciente: %w", err)
    }
    defer rows.Close()

    var doctors []users.User
    for rows.Next() {
        var doctor users.User
        var nombrePaciente string 
        var idPaciente int32 
        
        if err := rows.Scan(
            &doctor.IdUsuario,
            &doctor.Nombres,
            &doctor.Apellido_m,
            &doctor.Apellido_p,
            &doctor.Correo,
            &nombrePaciente,
            &idPaciente, 
        ); err != nil {
            return nil, fmt.Errorf("error al escanear doctor: %w", err)
        }
        doctors = append(doctors, doctor)
    }

    if err := rows.Err(); err != nil {
        return nil, fmt.Errorf("error después de iterar doctores: %w", err)
    }

    return doctors, nil
}