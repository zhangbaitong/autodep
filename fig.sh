 #!/bin/sh

function scandir() {
    local cur_dir workdir
    workdir=$1

    if [ ${workdir} = "/" ]
    then
        cur_dir=""
    else
        cur_dir=$(pwd)
    fi

    for dirlist in $(ls ${cur_dir})
    do
        if test -d ${dirlist}
        then
           cd ${dirlist}
           if test -f "fig.yml"
           then
                case "$2" in
                    start)  fig stop && fig rm --force && fig up -d ;;
                     stop)  fig stop ;;
                       rm)  fig stop && fig rm --force ;;
                esac
           fi
           cd ..
        fi
   done
}

scandir $1 $2