package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type PBM struct {
	data          [][]bool
	width, height int
	magicNumber   string
}

func main() {
	ReadPBM("test.pbm")
}

// ReadPBM lit une image PBM à partir d'un fichier et renvoie une structure représentant l'image.
func ReadPBM(filename string) {
	data, err := os.ReadFile(filename)
	if err != nil {
		fmt.Print(err)
		return
	}

	scanner := bufio.NewScanner(bytes.NewReader(data))
	Width := 0

	for scanner.Scan() {
		Width++

		// Si c'est la troisième ligne, imprimez-la
		if Width == 3 {
			ligne := scanner.Text()
			chiffres := strings.Fields(ligne)
			for _, chiffreStr := range chiffres {
				chiffre, err := strconv.Atoi(chiffreStr)
				if err == nil {
					if chiffre <= 6 {
						fmt.Println("Width : ", chiffre)
					} else {
						fmt.Println()
					}
				}
			}
			break // Sortir de la boucle après avoir imprimé la troisième ligne
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Erreur :", err)
	}
}
