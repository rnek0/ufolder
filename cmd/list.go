package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

//TODO: Corriger la longueur de l'entete et fin de tableau.

// datasColumn displays input in a column with 10 chars length
func datasColumn(input string, size int) string {

	var tailleCol int = size + 2
	tailleEntree := len(input)

	if tailleEntree > tailleCol {
		fmt.Println("Erreur L'entrée de la colonne doit être inferieure a la taille de la colonne.")
		// TODO : soulever l'erreur et la placer dans la signature de la func.
		os.Exit(1)
	}

	espaces := tailleCol - tailleEntree // calcul des espaces nécéssaires après la string (pour finir la colonne)
	affichage := make([]rune, espaces)

	entree := make([]rune, len(input))
	entree2 := []rune(input)
	entree = append(entree, entree2...)
	if tailleEntree < tailleCol {
		entree = append(entree, affichage...)
	}

	for i := 0; i < espaces; i++ {
		entree = append(entree, ' ') // On remplis les espaces restants.
	}

	return string(entree) //+ fmt.Sprint(espaces)
}

// Return a line of n runes.
//
// n (int) lenght of line
// c (rune) for line content.
func line(n int, c rune) string {
	result := ""
	for i := 0; i < n; i++ {
		result = string(c) + result
	}
	return result
}

// Calcul de la longueur des colonnes
func columnLenght(listeDossiers []GnuFolder) (int, int) {
	longueurMax1 := 0
	longueurMax2 := 0

	for i := 0; i < len(listeDossiers); i++ {
		nbchars1 := len(listeDossiers[i].Info)
		nbchars2 := len(listeDossiers[i].Folder)

		if nbchars1 > longueurMax1 {
			longueurMax1 = nbchars1
		}
		if nbchars2 > longueurMax2 {
			longueurMax2 = nbchars2
		}
	}
	return longueurMax1, longueurMax2
}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Liste les dossiers connus.",
	Long: ` 
 Fournit la liste des dossiers dans l'arborescence de gnu-linux avec une brève description. 
	`,
	Run: func(cmd *cobra.Command, args []string) {
		var err error
		if err := cobra.NoArgs(cmd, args); err != nil {
			cmd.Help()
			os.Exit(0)
		}

		listeDossiers, err := deserialize()
		if err != nil {
			fmt.Printf(" Erreur de recuperation de la liste : \n%s", err)
			os.Exit(0)
		}

		// CALCUL DE LA TAILLE DE L'ENSEMBLE DU TABLEAU
		left, right := columnLenght(listeDossiers)
		fmt.Printf(">>> %d - %d - %d\n", left, right, 2)

		lineTitle := line((left+right+2)-11, '-')
		fmt.Printf("%s\n", lineTitle)
		fmt.Printf("\n  -- Dossiers gnu-linux %s ", lineTitle)
		//____________________________________
		// TODO: extraire cette fonctionalité \
		//---------------------------------------------------------------------------------

		longueurMax := 0
		maxRightLengthCol := 0
		for i := 0; i < len(listeDossiers); i++ {
			nbchars := len(listeDossiers[i].Info)
			if nbchars > longueurMax {
				maxRightLengthCol = nbchars
				longueurMax = nbchars
				//fmt.Println(maxRightLengthCol)
			}
		}
		//fmt.Printf("La chaine plus grande A DROITE possède %d chars", maxRightLengthCol)
		// TODO: extraire cette fonctionalité
		longueurLeftMax := 0
		maxLeftLengthCol := 0
		for i := 0; i < len(listeDossiers); i++ {
			nbchars := len(listeDossiers[i].Folder)
			if nbchars > longueurLeftMax {
				maxLeftLengthCol = nbchars
				longueurLeftMax = nbchars
				//fmt.Println(maxLeftLengthCol)
			}
		}
		//fmt.Printf("La chaine plus grande A GAUCHE possède %d chars", maxLeftLengthCol)
		//---------------------------------------------------------------------------------

		for i := 0; i < len(listeDossiers); i++ {
			left := datasColumn(listeDossiers[i].Folder, maxLeftLengthCol)
			right := datasColumn(listeDossiers[i].Info, maxRightLengthCol)
			end := datasColumn("\t|", 2)
			//fmt.Printf("\n %s", left+right+end)
			fmt.Printf("\n | %s | %s \t%s", GREEN+left+RESET, right, end)
		}
		author = "rnek0" //viper.GetString("app.author")
		fmt.Printf("\n   -----------------------------------------------------------------------------%s--- ", author)
		fmt.Printf("\n\n")
		author = viper.GetString("app.author")
	},
}

func init() {

	// Help on command.
	listCmd.SetHelpFunc(func(c *cobra.Command, s []string) {
		fmt.Printf("\n "+GREEN+" Aide %s %s\033[m"+RESET+" :", "ufolder", listCmd.Name())
		fmt.Printf("\n\tVous devez passer seulement la commande " + BOLD + "<list>" + RESET + " sans arguments.")
		fmt.Printf("\n\tExemple :")
		fmt.Printf("\n\t" + BOLD + "$ufolder list" + RESET)
		fmt.Printf("\n\n")

	})

	rootCmd.AddCommand(listCmd)
}
