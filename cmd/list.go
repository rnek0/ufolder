package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"text/tabwriter"
)

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
// n (int) length of line
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
		nbchars1 := len(listeDossiers[i].Folder)
		nbchars2 := len(listeDossiers[i].Info)

		if nbchars1 > longueurMax1 {
			longueurMax1 = nbchars1
		}
		if nbchars2 > longueurMax2 {
			longueurMax2 = nbchars2
		}
	}
	return longueurMax1, longueurMax2
}

// Displays the folder list table.
//
// Just the folder name and a brief description.
func displayResults(listeDossiers []GnuFolder) {
	fmt.Printf("\n  ┏━ " + YELLOW + "Dossiers Gnu-Linux" + RESET + " ━┓")

	left, right := columnLenght(listeDossiers)

	lineTitle := line((left + 6 + right), '━')
	fmt.Printf("\n ┏%s┓", lineTitle)

	// initialize tabwriter
	w := new(tabwriter.Writer)

	// https://pkg.go.dev/text/tabwriter#Writer.Init
	w.Init(os.Stdout, left, 8, 1, ' ', 0)

	for i := 0; i < len(listeDossiers); i++ {
		fmt.Fprintf(w, "\n ┃ %s\t ┃ %s \t%s", YELLOW+listeDossiers[i].Folder+RESET, listeDossiers[i].Info, "┃")
	}

	w.Flush()

	author = viper.GetString("app.author")
	lineTitle = line((left + right - 1), '━')
	fmt.Printf("\n ┗%s%s┛", lineTitle, author)
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

		displayResults(listeDossiers)
		fmt.Print("\n  More : https://es.wikipedia.org/wiki/Filesystem_Hierarchy_Standard \n")
		fmt.Println()
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
