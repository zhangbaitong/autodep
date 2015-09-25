 #!/bin/sh

#读取配置文件
getKey(){ 
   
    section=$(echo $1 | cut -d '.' -f 1)    
    key=$(echo $1 | cut -d '.' -f 2)    
    sed -n "/\[$section\]/,/\[.*\]/{    
     /^\[.*\]/d    
     /^[ \t]*$/d    
     /^$/d    
     /^#.*$/d    
     s/^[ \t]*$key[ \t]*=[ \t]*\(.*\)[ \t]*/\1/p    
    }" /home/docker/fig/fig.ini    
}


#操作fig项目
function scandir() {
    local workdir onoff include
    workdir=$1

    onoff=$(getKey "fig.on-off")
    include=$(getKey "fig.include")

    if test -d ${workdir}
    then
	    cd ${workdir}
	    for dirlist in $(ls ${workdir})
	    do
            if  [ "${onoff}"="0" ] && [[ ! ",${include}," =~ ",${dirlist}," ]]
	    then
	         continue
	    fi 

	        if test -d ${dirlist}
	        then
	           cd ${dirlist}
	           if test -f "fig.yml"
	           then
		        echo ${dirlist}
	                case "$2" in
	                    start) echo "start" &&  /usr/local/bin/fig stop && /usr/local/bin/fig rm --force && /usr/local/bin/fig up -d ;;
	                     stop)  /usr/local/bin/fig stop ;;
	                       rm)  /usr/local/bin/fig stop && /usr/local/bin/fig rm --force ;;
	                esac
	           fi
	           cd ..
	        fi
	   done
	fi
}

scandir $1 $2