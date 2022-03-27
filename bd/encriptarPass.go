package bd

import "golang.org/x/crypto/bcrypt"

//EncriptarPass Rutina de encriptamineto de la pass
func EncriptarPass(pass string) (string, error) {

	costo := 6
	bytes, err := bcrypt.GenerateFromPassword([]byte(pass), costo)
	return string(bytes), err
}
