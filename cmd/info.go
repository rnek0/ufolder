package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var folder string = ""

// RequestForFolder is one Enum (string mode) for commands.
type RequestForFolder int

const (
	info RequestForFolder = iota
	man
)

func (rf RequestForFolder) String() string {
	return []string{"info", "man"}[rf]
}

var infos string = "Nous n'avons pas d'infos sur ce dossier !!!"
var infoCmd = &cobra.Command{
	Use:   "info",
	Short: "Description sommaire de l'utilité du dossier passé en paramètre.",
	Long: `Certains dossiers dans gnu-linux ont une utilité bien définie, cela aporte une structure au système . Par example:

 / est la racine du système qui contiendra les autres dossiers :
 /home contient le repertoire des utilisateurs.
 /boot les éléments nécessaires au démarrage du système.
Et bien d'autres.

Passez en argument le dossier dans la commande info et vous aurez une information sommaire à son sujet.

Exemple: $ufolder info bin

`,
	Args: func(cmd *cobra.Command, args []string) error {
		var err error

		if err := cobra.MinimumNArgs(1)(cmd, args); err != nil {
			cmd.Help()
			os.Exit(0)
		}

		folder = args[0]
		folder, infos, err = RequestDatas(folder, info.String())

		return err
	},
	Run: func(cmd *cobra.Command, args []string) {

		if infos == "" {
			infos = "n'est pas dans l'arborescence initiale de gnu-linux."
		}

		fmt.Printf(""+GREEN+"\n"+BOLD+"[info %s] : "+RESET+"%s\n", folder, infos)
		fmt.Printf("\n")

	},
}

func init() {

	// Help on command.
	infoCmd.SetHelpFunc(func(c *cobra.Command, s []string) {
		fmt.Printf("\n "+GREEN+" Aide %s %s\033[m"+RESET+" :", "ufolder", infoCmd.Name())
		fmt.Printf("\n\tVous devez passer le "+BOLD+"<nom du dossier>"+RESET+" en argument à la commande %s.", infoCmd.Name())
		fmt.Printf("\n\tExemple (avec le dossier root) :")
		fmt.Printf("\n\t" + BOLD + "$ufolder info root" + RESET)
		fmt.Printf("\n\n")
	})

	rootCmd.AddCommand(infoCmd)
}
