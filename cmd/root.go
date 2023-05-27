package cmd

import (
	"fmt"
	"os"
	"time"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const (
	GREEN  = "\033[34m"
	GRAY   = "\033[90m"
	YELLOW = "\033[33m"
	BOLD   = "\033[1m"
	RESET  = "\033[m"
)

var cfgFile string
var heure string
var author string
var license string

var rootCmd = &cobra.Command{
	Use:   "ufolder",
	Short: "Permet de connaitre le rôle des dossiers gnu-linux.",
	Long:  ` Description des dossiers qui composent l'arborescence dans les distributions gnu-linux:`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	cobra.CheckErr(err)
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $pwd/.ufolder.yaml so 'app dir')")
	rootCmd.Flags().StringVar(&heure, "time", time.Now().Format("15:04:05"), "Time.")
	rootCmd.PersistentFlags().StringVar(&author, "author", "rnek0", "Author name for copyright attribution")
	rootCmd.PersistentFlags().StringVar(&license, "license", "GPLv3", "copyright attribution")

	// TODO : set completions
	rootCmd.CompletionOptions.DisableDefaultCmd = true
	rootCmd.CompletionOptions.HiddenDefaultCmd = true
	rootCmd.CompletionOptions.DisableNoDescFlag = true
	rootCmd.CompletionOptions.DisableDescriptions = true

	var appBanner string = fmt.Sprintf(` .==========.
 |  .==========.  
 | /           /  %[1]s  
 |/           /   %[2]s
 .===========.    `, GREEN+"ufolder"+RESET, GREEN+"Dossiers de l'arborescence Gnu-Linux."+RESET)
	var appHelp string = fmt.Sprintf(`
 ufolder [commande] <dossier>

 Commandes disponibles:
    %[1]s  <dossier>  Description sommaire de l'utilité du dossier.
    %[2]s   <dossier>  Description détaillée de l'utilité du dossier.
    %[3]s             Liste les dossiers connus.
    %[4]s             Aide sur n'importe quelle commande.

 Flags:
    %[5]s, --help       Aide de ufolder

 Utilisez "ufolder [commande] -h" pour plus d'informations sur une commande.
`, GREEN+"info"+RESET, GREEN+"man"+RESET, GREEN+"list"+RESET, GREEN+"help"+RESET, GREEN+"-h"+RESET)

	rootCmd.SetHelpFunc(func(c *cobra.Command, s []string) {
		fmt.Printf("%s                                  [%s]                                   %s\n",
			GRAY,
			"ufolder",
			RESET)
		fmt.Printf("%s", appBanner)
		fmt.Printf("\n\n "+GREEN+"Aide %s\033[m"+RESET+" ", "ufolder :")
		fmt.Printf("" + BOLD + appHelp + RESET + "")
		fmt.Printf("\n Exemples:")
		fmt.Printf("\n    - " + BOLD + "$ufolder info -h" + RESET)
		fmt.Printf("\n    - " + BOLD + "$ufolder man home" + RESET)
		fmt.Printf("\n    - " + BOLD + "$ufolder info /" + RESET)
		fmt.Printf("\n\n")
		time, _ := rootCmd.Flags().GetString("time")
		if time != "" {
			fmt.Printf("%s             %s - [%s] - license: %s                %s\n",
				GRAY,
				time,
				author,
				license,
				RESET)
		}
	})
}

// initConfig reads config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		home, err := os.Getwd()
		cobra.CheckErr(err)
		viper.SetConfigName(".ufolder")
		viper.SetConfigType("yaml")
		viper.AddConfigPath(".")
		viper.AddConfigPath(home)
		viper.SetConfigFile(home + "/.ufolder.yaml")
	}

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		if viper.InConfig("app.author") {
			author = viper.GetString("app.author")
		}
		if viper.InConfig("app.license") {
			license = viper.GetString("app.license")
		}
	}
}
