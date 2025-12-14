package main

import (
	"compress/gzip"
	"fmt"
	"io"
	"io/fs"
	"os"
	"path/filepath"
)

func main() {
	root := "./logs_prueba"    // Carpeta a vigilar
	var threshold int64 = 1024 // 1KB (Si pesa más de esto, se rota)

	fmt.Println("🛡️  Iniciando Go-Rotator...")

	err := filepath.WalkDir(root, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() {
			return nil
		}

		// 1. Verificar extensión y tamaño
		info, err := d.Info()
		if err != nil {
			return nil // Ignoramos errores de lectura por ahora
		}

		// Solo procesamos .log y que superen el tamaño
		if filepath.Ext(path) == ".log" && info.Size() > threshold {
			fmt.Printf("🔄 Rotando: %s (%d bytes)\n", path, info.Size())

			// 2. Comprimir
			if err := compressLog(path); err != nil {
				fmt.Printf("❌ Error comprimiendo %s: %v\n", path, err)
				return nil
			}

			// 3. Eliminar el original (Solo si la compresión funcionó)
			if err := os.Remove(path); err != nil {
				fmt.Printf("⚠️ No se pudo borrar original: %v\n", err)
			} else {
				fmt.Println("🗑️  Original eliminado.")
			}
		}

		return nil
	})

	if err != nil {
		fmt.Println("Error crítico:", err)
	}
	fmt.Println("✅ Tarea finalizada.")
}

// (Pega aquí la función compressLog que te di arriba)
func compressLog(path string) error {
	source, err := os.Open(path)
	if err != nil {
		return err
	}
	defer source.Close()

	dest, err := os.Create(path + ".gz")
	if err != nil {
		return err
	}
	defer dest.Close()

	gzipWriter := gzip.NewWriter(dest)
	// Ojo: defer cierra al salir de la función, pero a veces queremos cerrar antes
	// para asegurar que se guardó todo antes de borrar el original.
	// Por simplicidad lo dejaremos en defer, pero un SysAdmin paranoico lo haría manual.
	defer gzipWriter.Close()

	if _, err := io.Copy(gzipWriter, source); err != nil {
		return err
	}
	return nil
}
