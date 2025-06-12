package users

import (
	"fmt"
	"time"

	"github.com/jesvaca/godesde0/modelos"
)

func AltaUsuario() {
	usr := new(modelos.Usuario)
	usr.AddUser(1, "Jesús", time.Now(), true)
	fmt.Println(usr)

}
