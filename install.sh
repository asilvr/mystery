#!/bin/sh
# Written by @nedimf (nedimcodes) 2021

# Colors
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[0;33m'
NC='\033[0m' # No Color
GLOBAL_DIR="mystery-release"

#check_if_dir_exists
check_if_dir_exists(){
    pwd=$(pwd)
    if [ -d "$pwd/$GLOBAL_DIR" ] 
    then
        return 0
    else
       return 1
    fi
}
#export downloaded archive
export_tar(){
    echo "Exporting... 📦"
    tar xvzf mystery_$1_darwin_$2.tar.gz
    cd mystery
    cp mystery /usr/local/bin/mystery
    #test if mystery installed successfully by calling mystery --version, clears output
    mystery --version > /dev/null 2>&1;
    status=$?
    if [[ status -eq 0 ]]
    then
        echo "Mystery CLI has been ${GREEN}successfully${NC} installed ✅\nLaunch by typing: ${YELLOW}mystery${NC} in your terminal"
    else
        echo "${RED}Mystery CLI failed to install. Copy above log and open new issue${NC}\nOpen: https://github.com/asilvr/mystery/issues/new/choose"
    fi
    
}
#download latest release
download(){
    if check_if_dir_exists; then
       cd $GLOBAL_DIR
    else
        mkdir $GLOBAL_DIR
        cd $GLOBAL_DIR
    fi
    arch_name=$1
    #Download latest release
    tag_version=$(curl -s https://api.github.com/repos/asilvr/mystery/releases/latest | grep "tag_name" | cut -d : -f 2,3 | tr -d \" | tr -d \,)
    echo "${GREEN}Downloading mystery $tag_version${NC}"
    curl -s https://api.github.com/repos/asilvr/mystery/releases/latest \
    | grep "browser_download_url.*_$arch_name.tar.gz" \
    | cut -d : -f 2,3 \
    | tr -d \" \
    | wget -qi -
    
    
    if [[ $? -eq 0 ]]
    then 
        export_tar $tag_version $arch_name
    else
        echo "${RED}Error: Mystery CLI failed to install!${NC}"
    fi    

}
#check of OS type and return values accordingly
if [[ "$OSTYPE" == "linux-gnu"* ]]; then
echo "${RED}Mystery has no Linux release build${NC}"
elif [[ "$OSTYPE" == "darwin"* ]]; then
 # Mac OS
arch_name="$(uname -m)"
if [ "${arch_name}" = "x86_64" ]; then
    if [ "$(sysctl -in sysctl.proc_translated)" = "1" ]; then
         echo "Unknown architecture: ${arch_name}, most likely Rosseta 2"
    else
        download "amd64"
    fi 
    elif [ "${arch_name}" = "arm64" ]; then
        download "arm64"
    else
    echo "Unknown architecture: ${arch_name}"
    fi
fi