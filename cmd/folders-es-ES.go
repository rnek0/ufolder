package cmd

var StrJsonesES string = `[
        {
            "folder":"/",
            "info":"Jerarquía primaria, la raíz o root del sistema Gnu-Linux.",
            "description":"est la racine de l'arborescence des fichiers, sur gnu-linux tout est fichier !!!. \n C'est à partir de cette racine que sont montés tous les autres repertoires."
        },
        {
            "folder":"/bin",
            "info":"Aplicaciones binarias de comando que son esenciales a una sesión de usuario.",
            "description":" El directorio /bin es un directorio estático y compartible en el que se almacenan archivos binarios/ejecutables necesarios para el funcionamiento del sistema. Estos archivos binarios los pueden usar la totalidad de usuarios del sistema operativo. Aplicaciones binarias de comando que son esenciales para que estén disponibles para una sesión de usuario único, o bien, para todos los usuarios (multiusuario). Incluyen, por ejemplo, cat, ls, cp, rm, mkdir, etc..."
        },
        {
            "folder":"/boot",
            "info":"Archivos cargadores de arranque (por ejemplo, los núcleos, el bootloader).",
            "description":"possède les éléments indispensables au démarrage de gnu-linux,\n comme le noyau ou la configuration de grub. "
        },
        {
            "folder":"/dev",
            "info":"Contiene archivos especiales de bloques y caracteres asociados a dispositivos hardware.",
            "description":"fournit les fichiers qui sont des points d'accès aux péripheriques du système.\n Pour accèder au lecteur de disquettes par exemple, le noyau utilisera /dev/fd0 "
        },
        {
            "folder":"/etc",
            "info":"Contiene archivos de configuración del sistema específicos del Host de todo el sistema.",
            "description":" C'est ici que sont stockés les differents fichiers de configuration du système."
        },
        {
            "folder":"/root",
            "info":"Directorio del administrador del sistema.",
            "description":" - "
        },
        {
            "folder":"/home",
            "info":"Directorio que contiene los usuarios del sistéma.",
            "description":"El directorio /home se trata de un directorio variable y compartible. Este directorio está destinado a alojar la totalidad de archivos personales de los distintos usuarios del sistema operativo a excepción del usuario root. Algunos de los archivos personales almacenados en la carpeta /home son fotografías, documentos de ofimática, vídeos, etc."
        },
        {
            "folder":"/lib",
            "info":"Bibliotecas esenciales compartidas de los programas alojados y módulos del kernel.",
            "description":" - "
        },
        {
            "folder":"/media",
            "info":"Contiene los puntos de montaje de los medios extraíbles de almacenamiento.",
            "description":" - "
        },
        {
            "folder":"/mnt",
            "info":"Sistema de archivos montados temporalmente.",
            "description":" - "
        },
        {
            "folder":"/opt",
            "info":"Programas opcionales de aplicaciones estáticas que pueden ser compartidas entre los usuarios.",
            "description":" - "
        },
        {
            "folder":"/proc",
            "info":"Sistema de archivos virtuales que documentan al núcleo y el estado de los procesos.",
            "description":" - "
        },
        {
            "folder":"/sbin",
            "info":"Sistema de binarios esencial, comandos y programas exclusivos del superusuario (root).",
            "description":" - "
        },
        {
            "folder":"/srv",
            "info":"contient les données des services de type serveur.",
            "description":" - "
        },
        {
            "folder":"/tmp",
            "info":"Aquí generalmente se guardan los archivos temporales guardados por apps.",
            "description":" - "
        },
        {
            "folder":"/var",
            "info":"Archivos variables, tales como logs, archivos spool, bases de datos, archivos de correo electrónico temporales.",
            "description":" - "
        },
        {
            "folder":"/usr",
            "info":"Contiene la mayoría de las utilidades y aplicaciones multiusuario.",
            "description":" - "
        },
        {
            "folder":".",
            "info":"Directorio actual.",
            "description":"represente le répertoire actuel. Il s'agit d'un alias du repertoir que l'on trouve\n avec la commande $pwd ; on doit le representer lorsque l'on veut executer un programme situé où nous sommes (./programme)\n Ceci a fin de ne pas confondre avec un programme qui serait situé dans le $PATH "
        },
        {
            "folder":"..",
            "info":"Directorio de nivel superior al actual.",
            "description":"represente le répertoire (parent), c'est le dossier qui contiens notre dossier dans l'arborescence.\n Ce repertoire ne se voit que si on passe le flag 'a' à la commande ls. \n Les repertoires qui commencent par un point dont cachés par definition.  "
        }
    ]
`
