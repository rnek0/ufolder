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

	"github.com/spf13/viper"
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
//
// First, check on locale system, after on viper config.
// Return fr-FR if err.
func GetLocale() (string, error) {
	//var locale string = "fr-FR"
	//var userLocale string = ""

	// DetectIETF will return the current locale as a string.
	// The format of the locale will be the ISO 639 two-letter language code,
	// a DASH, then an ISO 3166 two-letter country code.
	// userLocale, err := jibber_jabber.DetectIETF()
	// if err != nil {
	// 	fmt.Printf("%s", err)
	// 	userLocale = locale
	// }
	//userLocale = viper.GetString("app.lang")

	return viper.GetString("app.lang"), nil // nil BECAUSE TESTING WITH VIPER VARS !
	//return userLocale, err
}

// Just a display for debug.
//
// Take a string to display on debug tests.
// Seulement pour debug. (feignant !)
func Debug(blabla string) {
	fmt.Printf("\n%s", blabla)
}

// Vérifie que folder est un dossier valide.
//
// Retourne une chaîne vide s'il y a une erreur sur l'existence
// du dossier demandé.
func SanitizeQueryFolder(folder string) (string, error) {

	// [ok] Enlève les antislash.
	folder = strings.ReplaceAll(folder, "\\", "")

	// [ok] Teste la chaîne folder.
	r, err := regexp.Compile(`^(\.{1,2})|(\/)*\/?$`)
	if err != nil {
		return "", errors.New(" [!] Nous avons un problème de regexp. ")
	}
	if r.MatchString(folder) {
		if len(folder) > 1 && folder[len(folder)-1:] == "/" {
			folder = folder[:len(folder)-1] // Enleve les / de la fin de chaine. A revoir si on fais les sousdossiers.
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

	userLocale, err := GetLocale()
	if err != nil {
		log.Printf("Problème de lecture de locale: %s", err)
		os.Exit(0)
	}

	// folders, err := os.ReadFile(jsonFile)
	// ...
	// 	// Pour faire une install avec un fichier de resources .json c'est une grosse douleur au cul !!!
	// 	// Je dois le mettre dans un fichier go, même si j'augmente la taille de l'executable.
	// 	// Ou alors faire une recherche dans un dossier qui serait dans le $PATH comme $GOPATH (pas coseillé)
	// 	// Ou alors le passer en paramettre à l'appel de l'executable :/
	// 	// Ou alors le placer en ligne ??? j'ai pas envie :P
	//  // La solution actuelle est pas terrible mais elle marche.

	var jsonLocale string = ""
	switch userLocale {
	case "es-ES":
		jsonLocale = StrJsonesES // var StrJsonesES string in folders-es-ES.go
	case "fr-FR":
		jsonLocale = StrJsonfrFR
	case "en-US":
		jsonLocale = StrJsonenUS
	default:
		jsonLocale = StrJsonfrFR
	}
	folders := []byte(jsonLocale)

	// ATTENTION !!! La definition de StrJsonfrFR est dans folders-fr-FR.go !!!
	// Voir le problème d'install en local. Ou alors la solution sera d'aller chercher les infos sur le net,
	// un peu comme le fait tldr avec github.
	//rawJson := json.Unmarshal(folders, &gnuFolders)
	//rawJson := json.Unmarshal([]byte(StrJsonesES), &gnuFolders)
	//rawJson := json.Unmarshal([]byte(StrJsonfrFR), &gnuFolders)

	rawJson := json.Unmarshal(folders, &gnuFolders)
	if rawJson != nil {
		err := errors.New(" Faut voir ce qu'il se passe ici !!! ... ")
		log.Printf("Problème de deserialization de : %s ", err)
		return nil, err
	}
	return gnuFolders, nil
}

// getFolderDatas returns information about GnuFolder struct.
//
// info (string) : short description.
// man  (string) : long description.
func getFolderDatas(cmd, arg string) (string, error) {
	var gnuFolders []GnuFolder

	// TODO : Doing i18n with error messages !!!
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

// Returns infos about one folder.
func RequestDatas(folder string, what string) (string, string, error) {
	// cleaning the user request
	var err error
	folder, err = SanitizeQueryFolder(folder)
	if err != nil {
		fmt.Printf("Found SanitizeQueryFolder error :\n" + err.Error())
	}

	// Get datas
	infos, err = getFolderDatas(what, folder)
	if err != nil {
		return "", "", err
	}

	return folder, infos, err
}
