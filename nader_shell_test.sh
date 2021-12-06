# parameter 1: IP Address, parameter 2: slave node, parameter 3: execute times
function getRunStatus1(){
    current_st=$(./ncli --address "192.168.124.249"  --slave 104 rs)
    json_str=`echo $current_st|awk  -F '{'  '{print $2}'|awk  -F '}'  '{print $1}'`
    arr_json=(${json_str//,/ })  
    st_code=`echo ${arr_json[4]}|awk -F ':' '{print $2}'`
    echo $st_code
}


st_code=$(getRunStatus1)
for ((i=1;i<=$3;i++))
do
    echo "***************************$i***************************" 
    if [[ "$st_code" == "160" ]]  
    then
        ./ncli --address "$1"  --slave $2 tf
        sleep 5s
        st_code_1=$(getRunStatus1)
        [[ "$st_code_1" == "157" ]] || echo "Failed! RunStatus1 should equal '157', actual reault is $st_code_1"  
        echo "* Pass to trun off"
        ./ncli --address "$1"  --slave $2 to
        sleep 5s
        st_code_2=$(getRunStatus1)
        [[ "$st_code_2" == "160" ]] || echo "Failed! RunStatus1 should equal '160', actual reault is $st_code_2"  
        echo "* Pass to trun on"
    else
        ./ncli --address "$1"  --slave $2 to
        sleep 5s
        st_code_3=$(getRunStatus1)
        [[ "$st_code_3" == "160" ]] || echo "Failed! RunStatus1 should equal '160', actual reault is $st_code_3" 
        echo "* Pass to trun on" 
        ./ncli --address "$1"  --slave $2 tf
        sleep 5s
        st_code_4=$(getRunStatus1)
        [[ "$st_code_4" == "157" ]] || echo "Failed! RunStatus1 should equal '157', actual reault is $st_code_4" 
        echo "* Pass to trun off" 
    fi
done

 
