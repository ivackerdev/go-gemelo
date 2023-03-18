//Programa para realizar copias de seguridad de archivos
// Ivacker 20230126

package main

import (
	"fmt"
	"os"

	"github.com/ncw/rclone"
	"github.com/ncw/rclone/fs"
)

func main() {
	// Define la configuración de origen y destino
	src := "local:~/mi-carpeta"
	dst := "usb:~/mi-carpeta"

	// Inicializar el cliente rclone
	f, err := rclone.NewFs(src)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	d, err := rclone.NewFs(dst)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Crear un indicador de progreso para mostrar el progreso de la copia diferencial
	bar := fs.NewBar(fs.BarCopy, f.Name(), f.Size())
	bar.Start()
	defer bar.Finish()

	// Copia la carpeta usando la función "Copy" y la opción "--update"
	err = fs.CopyDir(f, d, fs.Config{
		NoTraverse:     false,
		IgnoreExisting: true,
		Progress:       bar,
	})
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("\nLa copia diferencial ha finalizado!")
}
