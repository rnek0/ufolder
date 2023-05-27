package cmd

var StrJsonfrFR string = `[
    {
        "folder":"/",
        "info":"dossier racine du système Gnu/Linux.",
        "description":"est la racine de l'arborescence des fichiers, sur gnu-linux tout est fichier !. \nC'est à partir de '/' que sont montés tous les autres repertoires. \n\nLe répertoire racine, où commence le système de fichiers, ne contiendra probablement que des sous-répertoires."
    },
    {
        "folder":".",
        "info":"répertoire où l'on se trouve. (répertoire courant)",
        "description":"represente le répertoire actuel. Il s'agit d'un alias du repertoir que l'on trouve avec la commande $pwd ; on doit le representer lorsque l'on veut executer un programme situé où nous sommes (./programme) Ceci a fin de ne pas confondre avec un programme qui serait situé dans le $PATH "
    },
    {
        "folder":"..",
        "info":"\"dossier parent\" du dossier où l'on se trouve.",
        "description":"represente le répertoire (parent), c'est le dossier qui contiens notre dossier dans l'arborescence. Ce repertoire ne se voit que si on passe le flag 'a' à la commande ls.  Les repertoires qui commencent par un point dont cachés par definition.  "
    },
    {
        "folder":"/bin",
        "info":"commandes exécutables utilisateur (binaires).",
        "description":"Le dossier /bin contient des programmes (exécutables) susceptibles d’être utilisés par tous les utilisateurs de la machine. Le dossier bin est l'un des nombreux dossiers typiques contenant des fichiers binaires, tous des programmes utilisés à partir de la ligne de commande bash. La seule chose qui distingue le dossier bin des autres est qu'il semblerait que de nombreux fichiers binaires dans ce dossier semblent être très importants en ce qui concerne le démarrage et l'utilisation de Linux lui-même. Sans le dossier bin, il n'y aurait aucun moyen qu'il y ait un système de fonctions, cependant de nombreux fichiers binaires dans d'autres dossiers tels que le dossier /usr/sbin/ ne seraient pas si critiques pour le fonctionnement de Linux lui-même."
    },
    {
        "folder":"/boot",
        "info":"fichiers du gestionnaire de démarrage (bootloader).",
        "description":"possède les éléments indispensables au démarrage de gnu-linux,\n\n comme le noyau ou la configuration de grub. "
    },
    {
        "folder":"/dev",
        "info":"fichiers des périphériques.",
        "description":"fournit les fichiers qui sont des points d'accès aux péripheriques du système. Pour accèder au lecteur de disquettes par exemple, le noyau utilisera /dev/fd0 "
    },
    {
        "folder":"/etc",
        "info":"fichiers de configuration du système.",
        "description":" C'est ici que sont stockés les differents fichiers de configuration du système."
    },
    {
        "folder":"/root",
        "info":"c'est le dossier de l'administrateur du système.",
        "description":" Dossier personnel de l'administrateur du système, tout comme chaque utilisateur possède son propre dossier."
    },
    {
        "folder":"/home",
        "info":"dossiers personnels des utilisateurs.",
        "description":"contient le répertoire de $USER (l'utilisateur loggué dans la machine).  C'est le repertoire principal où l'on trouve d'autres répertoires tels que Documents, Musique, Desktop etc...  Qui sont personnels à chaque utilisateur."
    },
    {
        "folder":"/lib",
        "info":"librairies partagées essentielles, modules du kernel.",
        "description":" - "
    },
    {
        "folder":"/media",
        "info":"points de montage des périphériques amovibles.",
        "description":" - "
    },
    {
        "folder":"/mnt",
        "info":"points de montage temporaires.",
        "description":" - "
    },
    {
        "folder":"/opt",
        "info":"applications tierces.",
        "description":" - "
    },
    {
        "folder":"/proc",
        "info":"informations sur les processus.",
        "description":" - "
    },
    {
        "folder":"/run",
        "info":"données temporaires (nouvelles versions de gnu/linux).",
        "description":"Représente un petit changement dans la façon dont Linux fonctionne par rapport aux données temporaires lors de l' exécution. "
    },
    {
        "folder":"/sbin",
        "info":"executables système.",
        "description":" - "
    },
    {
        "folder":"/srv",
        "info":"données des services de type serveur.",
        "description":" - "
    },
    {
        "folder":"/tmp",
        "info":"fichiers temporaires.",
        "description":" dossier accesible par tout le monde en  "
    },
    {
        "folder":"/var",
        "info":"fichiers où le système écrit au cours de son fonctionnement.",
        "description":" TODO: MAN /var A ECRIRE >> contient les fichiers variables. "
    },
    {
        "folder":"/usr",
        "info":"utilitaires et applications (multi)utilisateur.",
        "description":" TODO: MAN /usr A ECRIRE "
    },
    {
        "folder":"/usr/bin",
        "info":"commandes exécutables utilisateur (binaires) : voir dossier /bin.",
        "description":" Le dossier /usr/bin contient des programmes (exécutables)  susceptibles d’être utilisés par tous les utilisateurs de la machine. Le dossier /usr/bin est l'un des nombreux dossiers typiques contenant des fichiers binaires,  tous des programmes utilisés à partir de la ligne de commande bash. La seule chose qui distingue le dossier bin des autres est qu'il semblerait que de nombreux fichiers binaires dans ce dossier semblent être très importants en ce qui concerne le démarrage et l'utilisation de Linux lui-même. Sans le dossier bin, il n'y aurait aucun moyen qu'il y ait un système de fonctions, cependant de nombreux fichiers binaires dans d'autres dossiers tels que le dossier /usr/sbin/ ne seraient pas si critiques pour le fonctionnement de Linux lui-même. Sur Archlinux le dossier /bin est un lien vers /usr/bin"
    }
]
`
