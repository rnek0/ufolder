package cmd

// Version anglaise des datas json.
var StrJsonenUS string = `[
        {
            "folder":"/",
            "info":"is the root folder of the Gnu-Linux system.",
            "description":"est la racine de l'arborescence des fichiers, sur gnu-linux tout est fichier !!!. \n C'est à partir de cette racine que sont montés tous les autres repertoires."
        },
        {
            "folder":"bin",
            "info":"contains user executable commands (binaries).",
            "description":" Le dossier /bin contient des programmes (exécutables) susceptibles d’être utilisés par tous les utilisateurs de la machine.\n Le dossier bin est l'un des nombreux dossiers typiques contenant des fichiers binaires, \n tous des programmes utilisés à partir de la ligne de commande bash.\n La seule chose qui distingue le dossier bin des autres est qu'il semblerait que de nombreux fichiers\n binaires dans ce dossier semblent être très importants en ce qui concerne le démarrage et l'utilisation\n de Linux lui-même. Sans le dossier bin, il n'y aurait aucun moyen qu'il y ait un système de fonctions,\n cependant de nombreux fichiers binaires dans d'autres dossiers tels que le dossier /usr/sbin/ ne seraient\n pas si critiques pour le fonctionnement de Linux lui-même.\n"
        },
        {
            "folder":"boot",
            "info":"contient les fichiers du gestionnaire de démarrage (bootloader).",
            "description":"possède les éléments indispensables au démarrage de gnu-linux,\n comme le noyau ou la configuration de grub. "
        },
        {
            "folder":"dev",
            "info":"contient les fichiers des périphériques.",
            "description":"fournit les fichiers qui sont des points d'accès aux péripheriques du système.\n Pour accèder au lecteur de disquettes par exemple, le noyau utilisera /dev/fd0 "
        },
        {
            "folder":"etc",
            "info":"contient les fichiers de configuration du système.",
            "description":" C'est ici que sont stockés les differents fichiers de configuration du système."
        },
        {
            "folder":"root",
            "info":"c'est le dossier de l'administrateur du système.",
            "description":" - "
        },
        {
            "folder":"home",
            "info":"contient les dossiers personnels des utilisateurs.",
            "description":"contient le répertoire de $USER (l'utilisateur loggué dans la machine). \n C'est le repertoire principal où l'on trouve d'autres répertoires tels que Documents, Musique, Desktop etc... \n Qui sont personnels à chaque utilisateur."
        },
        {
            "folder":"lib",
            "info":"contient les librairies partagées essentielles, modules du kernel.",
            "description":" - "
        },
        {
            "folder":"media",
            "info":"contient les points de montage des périphériques amovibles.",
            "description":" - "
        },
        {
            "folder":"mnt",
            "info":"contient les points de montage temporaires.",
            "description":" - "
        },
        {
            "folder":"opt",
            "info":"contient les applications tierces.",
            "description":" - "
        },
        {
            "folder":"proc",
            "info":"contient les informations sur les processus.",
            "description":" - "
        },
        {
            "folder":"sbin",
            "info":"contient les executables système.",
            "description":" - "
        },
        {
            "folder":"srv",
            "info":"contient les données des services de type serveur.",
            "description":" - "
        },
        {
            "folder":"tmp",
            "info":"contient les fichiers temporaires.",
            "description":" - "
        },
        {
            "folder":"var",
            "info":"contient les fichiers variables.",
            "description":" - "
        },
        {
            "folder":"usr",
            "info":"contient les utilitaires et applications (multi)utilisateur.",
            "description":" - "
        },
        {
            "folder":".",
            "info":"est le répertoire où l'on se trouve.",
            "description":"represente le répertoire actuel. Il s'agit d'un alias du repertoir que l'on trouve\n avec la commande $pwd ; on doit le representer lorsque l'on veut executer un programme situé où nous sommes (./programme)\n Ceci a fin de ne pas confondre avec un programme qui serait situé dans le $PATH "
        },
        {
            "folder":"..",
            "info":"est le \"dossier parent\" du dossier où l'on se trouve.",
            "description":"represente le répertoire (parent), c'est le dossier qui contiens notre dossier dans l'arborescence.\n Ce repertoire ne se voit que si on passe le flag 'a' à la commande ls. \n Les repertoires qui commencent par un point dont cachés par definition.  "
        }
    ]
`
