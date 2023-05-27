# A CLI app to check the purpose of classic Gnu-Linux folders.

You can have a summary description of the directories as well as a more detailed description.  
A list is also available.

Just a little reminder that is not in the man pages.  

A trial to have fun with json, golang, cobra and viper.

TODO : play with .ufolder.yaml lang var to make an i18n app see [here](https://github.com/rnek0/ufolder/blob/857d97725cd09f625d30740c54e062d6285c1090/cmd/deserialize.go#L30)

Enjoy.

```bash
❯ ufolder -h
                                  [ufolder]                                   
 .==========.
 |  .==========.  
 | /           /  ufolder  
 |/           /   Dossiers de l'arborescence Gnu-Linux.
 .===========.    

 Aide ufolder : 
 ufolder [commande] <dossier>

 Commandes disponibles:
    info  <dossier>  Description sommaire de l'utilité du dossier.
    man   <dossier>  Description détaillée de l'utilité du dossier.
    list             Liste les dossiers connus.
    help             Aide sur n'importe quelle commande.

 Flags:
    -h, --help       Aide de ufolder

 Utilisez "ufolder [commande] -h" pour plus d'informations sur une commande.

 Exemples:
    - $ufolder info -h
    - $ufolder man home
    - $ufolder info /
```

![GPL-v3](gplv3-with-text-136x68.png)  
 [License](https://www.gnu.org/licenses/gpl-howto.html)