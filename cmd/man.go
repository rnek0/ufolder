package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var manCmd = &cobra.Command{
	Use:   "man",
	Short: "Une description détaillée de chaque dossier.",
	Long: `Le propos de man est de donner des informations complètes
	sur le dossier demandé. 
	
	Par example (pour le dossier boot):
	
	$ufolder man boot
	
	`,
	Args: func(cmd *cobra.Command, args []string) error {
		var err error
		if err := cobra.MinimumNArgs(1)(cmd, args); err != nil {
			cmd.Help()
			os.Exit(0)
			return err
		}

		folder = args[0]
		folder, err = SanitizeQueryFolder(folder)
		if err != nil {
			fmt.Printf("Found SanitizeQueryFolder error :\n" + err.Error())
		}

		infos, err = GetFolderDatas(man.String(), folder)
		if err != nil {
			return err
		}

		return err
	},
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("\n"+YELLOW+" [ %s %s ]\n"+RESET, man.String(), folder)
		//infos = wordwrap.WrapString(infos, 120)
		// https://codereview.stackexchange.com/questions/244435/word-wrap-in-go
		// https://github.com/mitchellh/go-wordwrap/blob/master/wordwrap.go
		infos = WordWrap(infos, 90)
		//fmt.Println(coco)
		fmt.Printf("\n%s\n", infos)
		fmt.Printf("\n")
	},
}

func init() {

	manCmd.SetHelpFunc(func(c *cobra.Command, s []string) {
		fmt.Printf("\n "+GREEN+" Aide %s %s\033[m"+RESET+" :", "ufolder", manCmd.Name())
		fmt.Printf("\n\tVous devez passer le "+BOLD+"<nom du dossier>"+RESET+" en argument à la commande %s.", manCmd.Name())
		fmt.Printf("\n\tExemple (avec le dossier root) :")
		fmt.Printf("\n\t"+BOLD+"$ufolder %s root"+RESET, manCmd.Name())
		fmt.Printf("\n\n")
	})

	rootCmd.AddCommand(manCmd)
}
