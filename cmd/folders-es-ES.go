package cmd

var StrJsonesES string = `[
        {
            "folder":"/",
            "info":"Jerarquía primaria, la raíz o root del sistema Gnu-Linux.",
            "description":"est la racine de l'arborescence des fichiers, sur gnu-linux tout est fichier !!!. \n C'est à partir de cette racine que sont montés tous les autres repertoires."
        },
        {
            "folder":"bin",
            "info":"Aplicaciones binarias de comando que son esenciales a una sesion de usuario.",
            "description":" Aplicaciones binarias de comando que son esenciales para que estén disponibles para una sesión de usuario único, o bien, para todos los usuarios (multiusuario). Incluyen, por ejemplo, cat, ls, cp, rm, mkdir, etc...\n La seule chose qui distingue le dossier bin des autres est qu'il semblerait que de nombreux fichiers\n binaires dans ce dossier semblent être très importants en ce qui concerne le démarrage et l'utilisation\n de Linux lui-même. Sans le dossier bin, il n'y aurait aucun moyen qu'il y ait un système de fonctions,\n cependant de nombreux fichiers binaires dans d'autres dossiers tels que le dossier /usr/sbin/ ne seraient\n pas si critiques pour le fonctionnement de Linux lui-même.\n"
        },
        {
            "folder":"boot",
            "info":"Archivos cargadores de arranque (por ejemplo, los núcleos, el bootloader).",
            "description":"possède les éléments indispensables au démarrage de gnu-linux,\n comme le noyau ou la configuration de grub. "
        },
        {
            "folder":"dev",
            "info":"Contiene archivos especiales de bloques y caracteres asociados a dispositivos hardware.",
            "description":"fournit les fichiers qui sont des points d'accès aux péripheriques du système.\n Pour accèder au lecteur de disquettes par exemple, le noyau utilisera /dev/fd0 "
        },
        {
            "folder":"etc",
            "info":"Contiene archivos de configuración del sistema específicos del Host de todo el sistema.",
            "description":" C'est ici que sont stockés les differents fichiers de configuration du système."
        },
        {
            "folder":"root",
            "info":"c'est le dossier de l'administrateur du système.",
            "description":" - "
        },
        {
            "folder":"home",
            "info":"directorio que contiene los usuarios del sistéma.",
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
