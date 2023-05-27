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

// func SanitizeQueryFolder(folder string) string {
// 	if strings.Contains(folder, "\\") {
// 		warning := " >>> Les antislash [\\] ne sont pas utilisés sur gnu-linux mais sur des systèmes proprietaires."
// 		fmt.Println(warning)
// 		folder = removeSlashes(folder, '\\')
// 	}

// 	if strings.Contains(folder, ".") {
// 		fmt.Println("1." + folder)
// 		folder = removeDots(folder)
// 	}

// 	var folders = []string{"/", "/usr/bin", "/usr/sbin", "/usr/local", "/var/log"}

// 	sum := 0
// 	remove := true
// 	for i := 0; i < len(folders); i++ {
// 		if folder != folders[i] {
// 			remove = false
// 			break
// 		}
// 		sum += i
// 	}

// 	if remove {
// 		folder = removeSlashes(folder, '/')
// 	}

// 	return folder
// }

var infos string = "Nous n'avons pas d'infos sur ce dossier !!!"
var infoCmd = &cobra.Command{
	Use:   "info",
	Short: "Description sommaire de l'utilité du dossier passé en paramèttre.",
	Long: `Certains dossiers dans gnu-linux ont une utilité bien définie, cela aporte une structure au système . Par example:

 / est la racine du système qui contiendra les autres dossiers :
 /home contient le repertoire des utilisateurs.
 /boot les éléments nécéssaires au demarrage du système.
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

		// PRENDS L'ARGUMENT
		folder = args[0]

		// NETTOIE
		folder, err = SanitizeQueryFolder(folder)
		if err != nil {
			fmt.Printf("Error on SanitizeQueryFolder  :\n" + err.Error())
			os.Exit(1)
		}

		// RECUPERE DATAS
		infos, err = GetFolderDatas(info.String(), folder)
		if err != nil {
			return err
		}
		return err
	},
	Run: func(cmd *cobra.Command, args []string) {

		if infos == "" {
			infos = "n'est pas dans l'arborescence initiale de gnu-linux."
		}
		//fmt.Printf("\n [ %s ] ", folder)

		fmt.Printf(""+GREEN+"\n "+BOLD+"%s : "+RESET+"%s\n", folder, infos)
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
