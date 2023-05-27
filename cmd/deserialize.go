package cmd

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
	"unicode/utf8"

	"github.com/cloudfoundry/jibber_jabber"
)

// GnuFolder is a type that contains folders information.
//
// Folder string is a key because folders have always different names.
// Info string is a summary description.
// Description is more detailed description.
type GnuFolder struct {
	Folder      string `json:"folder"`
	Info        string `json:"info"`
	Description string `json:"description"`
}

// GetLocale gets the locale lang user.
// Return fr-FR if err.
func GetLocale() (string, error) {
	// TODO: check if lang is defined in .ufolder.yaml !
	locale, err := jibber_jabber.DetectIETF()
	if err != nil {
		locale = "fr-FR"
	}
	return locale, err
}

// Just a display for debug.
//
// Take a string to display on debug tests.
// Seulement pour debug. (feignant !)
func Debug(blabla string) {
	fmt.Printf("\n%s", blabla)
}

// Verifie que folder est un dossier valide.
// Retourne une chaine vide s'il y a un erreur sur l'existence
// du dissier demandé.
func SanitizeQueryFolder(folder string) (string, error) {
	// Enleve les antislash
	folder = strings.ReplaceAll(folder, "\\", "")
	// Teste la chaine folder.
	r, err := regexp.Compile(`^(\.{1,2})|(\/)*\/?$`)
	if err != nil {
		return "", errors.New(" [!] Nous avons un problème de regexp. ")
	}
	if r.MatchString(folder) {
		if len(folder) > 1 && folder[len(folder)-1:] == "/" {
			folder = folder[:len(folder)-1] // Enleve les / de la fin de chaine.
		}
	} else {
		return "", errors.New(" [!] Aucun fichier ou dossier de ce type: " + folder)
	}

	// Teste un 'Directory traversal'.
	traversal, err := regexp.Compile(`^((\/?|\.\/)\.{2})\/+`)
	if err != nil {
		return "", errors.New(" [!] Nous avons un problème de regexp pour 'Directory traversal'. ")
	}
	if traversal.MatchString(folder) {
		dots := 0
		count := 0
		for count < len(folder) {
			if folder[count] == '.' {
				dots += 1
			}
			count += 1
		}
		up := strconv.Itoa(dots / 2)
		fmt.Println("\n Remonte de " + up + " dossiers.")
		folder = ".."
	}
	return folder, nil
}

// deserialize creates a list of Gnufolder structs from json.
func deserialize() ([]GnuFolder, error) {
	var gnuFolders []GnuFolder

	// userLocale, err := GetLocale()
	// if err != nil {
	// 	log.Printf("Problème de lecture de locale: %s", err)
	// 	os.Exit(0)
	// }

	//TODO: Revoir le path après install !!!
	//
	// jsonFile := "./cmd/folders-" + userLocale + ".json"
	// _, err = os.ReadFile(jsonFile)
	// if err != nil {
	// 	//log.Fatal(err)
	// 	jsonFile = "/home/rnek0/go/bin/folders-" + userLocale + ".json"
	// }

	// folders, err := os.ReadFile(jsonFile)
	// if err != nil {
	// 	// Pour faire une install avec un fichier de resources .json c'est une grosse douleur au cul !!!
	// 	// Je dois le mettre dans un fichier go, même si j'augmente la taille de l'executable.
	// 	// Ou alors faire une recherche dans un dossier qui serait dans le $PATH comme $GOPATH (pas coseillé)
	// 	// Ou alors le passer en paramettre à l'appel de l'executable :/
	// 	// Ou alors le placer en ligne ??? j'ai pas envie :P
	// 	// go c'est de la daube.

	// 	//return nil, err

	// 	var jsonLocale string = "" // json dans le fichier go !!!
	// 	switch userLocale {
	// 	case "fr-FR":
	// 		jsonLocale = StrJsonfrFR
	// 	case "en-US":
	// 		jsonLocale = StrJsonenUS
	// 	default:
	// 		jsonLocale = StrJsonfrFR
	// 	}
	// 	folders = []byte(jsonLocale)

	// 	err = nil
	// }

	//rawJson := json.Unmarshal(folders, &gnuFolders)
	rawJson := json.Unmarshal([]byte(StrJsonfrFR), &gnuFolders)
	if rawJson != nil {
		err := errors.New(" Faut voir ce qu'il se passe ici !!! ... ")
		log.Printf("Problème de deserialization de : %s ", err)
		return nil, err
	}
	return gnuFolders, nil
}

// GetFolderDatas returns information about GnuFolder struct.
//
// info (string) : short description.
// man  (string) : long description.
func GetFolderDatas(cmd, arg string) (string, error) {
	var gnuFolders []GnuFolder

	var information string = "Désolé, nous n'avons pas d'information sur ce dossier !!!"

	gnuFolders, err := deserialize()
	if err != nil {
		log.Printf("Problème de recuperation des datas (déserialization): %s", err)
	}

	// check if we have arg (unix folder) in json file
	folderExists := false
	for i := 0; i < len(gnuFolders); i++ {
		if gnuFolders[i].Folder == arg {
			folderExists = true

			switch cmd {
			case info.String():
				information = gnuFolders[i].Info
			case man.String():
				information = gnuFolders[i].Description
			default:
				information = "Ce champ n'existe pas dans la structure json !"
				log.Println(information)
				os.Exit(1)
			}
		}
	}

	if folderExists {
		return information, nil
	}
	fmt.Printf("\n Argument invalide : %s\n Ce dossier n'existe pas.\n ", arg)
	//os.Exit(1)
	return "", err
}

func runesToUTF8(rs []rune) []byte {
	return []byte(string(rs))
}

func runeToByteArr(r rune) []byte {
	var bytes []byte               // Une rune peut avoir entre 1 et 4 bytes (cf: utf-8)
	var rn = []rune{r}             // Il me faut un arr de runes pour la fonction, même si j'en ai qu'une.
	var bt = runesToUTF8(rn)       // Je recupere un arr de bytes du codepoint
	for i := 0; i < len(bt); i++ { // Je peux avoir plusieurs bytes dans le codepoint alors je boucle.
		bytes = append(bytes, bt[i])
	}
	return bytes
}

func WordWrap(text string, lineWidth int) string {
	//var result = make([]string, 1)
	var res = ""
	//nombreDeCaracteres := utf8.RuneCountInString(text)
	//var textSlice = make([]rune, nombreDeCaracteres)
	// PARCOURS RUNE A RUNE LE TEXTE
	// -----------------------------
	// linePosition := 0
	// for _, runeValue := range text {
	// 	textSlice[runeValue] = runeValue
	// 	linePosition = linePosition + 1
	// 	//fmt.Printf("%c", runeValue)
	// 	if linePosition > lineWidth {
	// 		break
	// 	}
	// }

	// PARCOURS MOT A MOT (OU TAB & SEPARATEURS) LE TEXTE
	// --------------------------------------------------
	//var wordsSlice = make([]rune, nombreDeMots)
	wordsSlice := strings.Split(text, " ")
	wordPosition := 0

	if utf8.ValidString(text) {

		pos := 0
		for i := 0; i < len(wordsSlice); i++ { // POUR CHAQUE MOT

			var str = make([]string, 0)
			//str = append(str, " ")
			for _, wordValue := range wordsSlice[i] { // POUR CHAQUE LETTRE
				pos = pos + 1
				str = append(str, string(wordValue))
				if wordValue == '\n' {
					pos = 0
				}
				if pos >= lineWidth {
					fmt.Printf("%d\n", pos)
					pos = 0
				}
			}
			str = append(str, " ")
			res = strings.Join(str, "")
			for _, letter := range str {
				fmt.Printf("%s", letter)
			}
			//res = "-" + strings.Join(result, "") + strings.Join(str, "")
			str = nil // clear slice

			wordPosition = wordPosition + 1
			//fmt.Printf(" [%d]\n", wordPosition)
		}

	}

	//res := strings.Join(result, "")
	//return "\n------------------------------\nretour de la fonction WordWrap !!! \n"
	return res
}
