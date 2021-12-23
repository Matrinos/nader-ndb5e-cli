# brew install coreutils

# init operation time
./ncli --address $1 --slave 00 sp

function getRunStatus1(){
    current_st=$(./ncli --address $1  --slave $2 rs)
    json_str=`echo $current_st|awk  -F '{'  '{print $2}'|awk  -F '}'  '{print $1}'`
    arr_json=(${json_str//,/ })  
    st_code=`echo ${arr_json[4]}|awk -F ':' '{print $2}'`
    echo $st_code
}

function getWeekDay(){
    week_day_code_c=$(date +"%w")
    week_day_code=$(( $week_day_code_c+$1 ))
    if [[ $week_day_code == 1 ]];then
        week_day="Monday"
    fi
    if [[ $week_day_code == 2 ]];then
        week_day="Tuesday"
    fi
    if [[ $week_day_code == 3 ]];then
        week_day="Wednesday"
    fi
    if [[ $week_day_code == 4 ]];then
        week_day="Thursday"
    fi
    if [[ $week_day_code == 5 ]];then
        week_day="Friday"
    fi
    if [[ $week_day_code == 6 ]];then
        week_day="Saturday"
    fi
    if [[ $week_day_code == 7 ]];then
        week_day="Sunday"
    fi
    echo $week_day
}

function testCase7(){
    st_code_1=$(getRunStatus1 $1 $2)
    t_time=$(gdate +%H:%M:%S -d "3 minutes")
    cat ./sample.json| jq --arg t "$t_time" '.TimeOffTime0 |= $t'| jq --arg t "$t_time" '.TimeOnTime0 |= $t'| jq --arg t "$t_time" '.TimeOffTime1 |= $t'| jq --arg t "$t_time" '.TimeOnTime1 |= $t'| jq --arg t "$t_time" '.TimeOffTime2 |= $t'| jq --arg t "$t_time" '.TimeOnTime2 |= $t'| jq --arg t "$t_time" '.TimeOffTime3 |= $t'| jq --arg t "$t_time" '.TimeOnTime3 |= $t'| jq --arg t "$t_time" '.TimeOffTime4 |= $t'| jq --arg t "$t_time" '.TimeOnTime4 |= $t' > tmp_7.json
    ./ncli --address $1 --slave $2 st --jsonpath=./tmp_7.json
    sleep 30s
    ./ncli --address $1 --slave $2 tp
    sleep 160s
    st_code_2=$(getRunStatus1 $1 $2)
    [[ $st_code_1==$st_code_2 ]] || echo "** Failed, RunStatus1 should equal $st_code_1, actual reault is $st_code_2"
    [[ $st_code_1==$st_code_2 ]] && echo "* Pass to set the time"
}

function testCase8(){
    week_day_c=$(getWeekDay 0)
    week_day_n=$(getWeekDay 1)
    t_time_0=$(gdate +%H:%M:%S -d "1 minutes")
    t_time_1=$(gdate +%H:%M:%S -d "2 minutes")
    t_time_2=$(gdate +%H:%M:%S -d "3 minutes")
    t_time_3=$(gdate +%H:%M:%S -d "4 minutes")
    t_time_4=$(gdate +%H:%M:%S -d "5 minutes")
    t_time_5=$(gdate +%H:%M:%S -d "6 minutes")
    t_time_6=$(gdate +%H:%M:%S -d "7 minutes")
    t_time_7=$(gdate +%H:%M:%S -d "8 minutes")
    t_time_8=$(gdate +%H:%M:%S -d "9 minutes")
    t_time_9=$(gdate +%H:%M:%S -d "10 minutes")
    cat ./sample.json| jq --arg t "$t_time_0" '.TimeOffTime0 |= $t'| jq --arg t "$t_time_1" '.TimeOnTime0 |= $t'| jq --arg t "$t_time_2" '.TimeOffTime1 |= $t'| jq --arg t "$t_time_3" '.TimeOnTime1 |= $t'| jq --arg t "$t_time_4" '.TimeOffTime2 |= $t'| jq --arg t "$t_time_5" '.TimeOnTime2 |= $t'| jq --arg t "$t_time_6" '.TimeOffTime3 |= $t'| jq --arg t "$t_time_7" '.TimeOnTime3 |= $t'| jq --arg t "$t_time_8" '.TimeOffTime4 |= $t'| jq --arg t "$t_time_9" '.TimeOnTime4 |= $t' > tmp_8_1.json
    cat ./tmp_8_1.json | jq --arg wdc "$week_day_c" '.TimeOffDay0[0] |= $wdc'| jq --arg wdc "$week_day_c" '.TimeOffDay1[0] |= $wdc' | jq --arg wdc "$week_day_c" '.TimeOffDay2[0] |= $wdc'| jq --arg wdc "$week_day_c" '.TimeOffDay3[0] |= $wdc'| jq --arg wdc "$week_day_c" '.TimeOffDay4[0] |= $wdc'| jq --arg wdn "$week_day_n" '.TimeOffDay0[1] |= $wdn'| jq --arg wdn "$week_day_n" '.TimeOffDay1[1] |= $wdn'| jq --arg wdn "$week_day_n" '.TimeOffDay2[1] |= $wdn'| jq --arg wdn "$week_day_n" '.TimeOffDay3[1] |= $wdn'| jq --arg wdn "$week_day_n" '.TimeOffDay4[1] |= $wdn'> tmp_8_2.json && rm -rf ./tmp_8_1.json
    cat ./tmp_8_2.json | jq --arg wdc "$week_day_c" '.TimeOnDay0[0] |= $wdc'| jq --arg wdc "$week_day_c" '.TimeOnDay1[0] |= $wdc' | jq --arg wdc "$week_day_c" '.TimeOnDay2[0] |= $wdc'| jq --arg wdc "$week_day_c" '.TimeOnDay3[0] |= $wdc'| jq --arg wdc "$week_day_c" '.TimeOnDay4[0] |= $wdc'| jq --arg wdn "$week_day_n" '.TimeOnDay0[1] |= $wdn'| jq --arg wdn "$week_day_n" '.TimeOnDay1[1] |= $wdn'| jq --arg wdn "$week_day_n" '.TimeOnDay2[1] |= $wdn'| jq --arg wdn "$week_day_n" '.TimeOnDay3[1] |= $wdn'| jq --arg wdn "$week_day_n" '.TimeOnDay4[1] |= $wdn'> tmp_8.json && rm -rf ./tmp_8_2.json
    ./ncli --address $1 --slave $2 st --jsonpath=./tmp_8.json
    sleep 5s
    ./ncli --address $1 --slave $2 tp
    sleep 5s
    for ((i=1;i<=5;i++))
    do
        echo "************ Turn Off ************"
        sleep 60s
        st_code_8_1=$(getRunStatus1 $1 $2)
        [[ "$st_code_8_1" == "148" ]] || echo "** Failed, RunStatus1 should equal 148, actual reault is $st_code_8_1"
        [[ "$st_code_8_1" == "148" ]] && echo "* Pass to set the time"
        echo "************ Turn On ************"
        sleep 60s
        st_code_8_2=$(getRunStatus1 $1 $2)
        [[ "$st_code_8_2" == "180" ]] || echo "** Failed, RunStatus1 should equal 180, actual reault is $st_code_8_2"
        [[ "$st_code_8_2" == "180" ]] && echo "* Pass to set the time"
    done
}

function testCase9(){
    week_day_c=$(getWeekDay 0)
    week_day_n=$(getWeekDay 1)
    t_time_0=$(gdate +%H:%M:%S -d "1 minutes")
    t_time_1=$(gdate +%H:%M:%S -d "2 minutes")
    t_time_2=$(gdate +%H:%M:%S -d "3 minutes")
    t_time_3=$(gdate +%H:%M:%S -d "4 minutes")
    t_time_4=$(gdate +%H:%M:%S -d "5 minutes")
    t_time_5=$(gdate +%H:%M:%S -d "6 minutes")
    t_time_6=$(gdate +%H:%M:%S -d "7 minutes")
    t_time_7=$(gdate +%H:%M:%S -d "8 minutes")
    t_time_8=$(gdate +%H:%M:%S -d "9 minutes")
    t_time_9=$(gdate +%H:%M:%S -d "10 minutes")
    cat ./sample.json| jq --arg t "$t_time_0" '.TimeOffTime0 |= $t'| jq --arg t "$t_time_1" '.TimeOnTime0 |= $t'| jq --arg t "$t_time_2" '.TimeOffTime1 |= $t'| jq --arg t "$t_time_3" '.TimeOnTime1 |= $t'| jq --arg t "$t_time_4" '.TimeOffTime2 |= $t'| jq --arg t "$t_time_5" '.TimeOnTime2 |= $t'| jq --arg t "$t_time_6" '.TimeOffTime3 |= $t'| jq --arg t "$t_time_7" '.TimeOnTime3 |= $t'| jq --arg t "$t_time_8" '.TimeOffTime4 |= $t'| jq --arg t "$t_time_9" '.TimeOnTime4 |= $t' > tmp_8_1.json
    cat ./tmp_8_1.json | jq --arg wdc "$week_day_c" '.TimeOffDay0[0] |= $wdc'| jq --arg wdc "$week_day_c" '.TimeOffDay1[0] |= $wdc' | jq --arg wdc "$week_day_c" '.TimeOffDay2[0] |= $wdc'| jq --arg wdc "$week_day_c" '.TimeOffDay3[0] |= $wdc'| jq --arg wdc "$week_day_c" '.TimeOffDay4[0] |= $wdc'| jq --arg wdn "$week_day_n" '.TimeOffDay0[1] |= $wdn'| jq --arg wdn "$week_day_n" '.TimeOffDay1[1] |= $wdn'| jq --arg wdn "$week_day_n" '.TimeOffDay2[1] |= $wdn'| jq --arg wdn "$week_day_n" '.TimeOffDay3[1] |= $wdn'| jq --arg wdn "$week_day_n" '.TimeOffDay4[1] |= $wdn'> tmp_8_2.json && rm -rf ./tmp_8_1.json
    cat ./tmp_8_2.json | jq --arg wdc "$week_day_c" '.TimeOnDay0[0] |= $wdc'| jq --arg wdc "$week_day_c" '.TimeOnDay1[0] |= $wdc' | jq --arg wdc "$week_day_c" '.TimeOnDay2[0] |= $wdc'| jq --arg wdc "$week_day_c" '.TimeOnDay3[0] |= $wdc'| jq --arg wdc "$week_day_c" '.TimeOnDay4[0] |= $wdc'| jq --arg wdn "$week_day_n" '.TimeOnDay0[1] |= $wdn'| jq --arg wdn "$week_day_n" '.TimeOnDay1[1] |= $wdn'| jq --arg wdn "$week_day_n" '.TimeOnDay2[1] |= $wdn'| jq --arg wdn "$week_day_n" '.TimeOnDay3[1] |= $wdn'| jq --arg wdn "$week_day_n" '.TimeOnDay4[1] |= $wdn'> tmp_8.json && rm -rf ./tmp_8_2.json
    ./ncli --address $1 --slave $2 st --jsonpath=./tmp_8.json
    sleep 2s
    ./ncli --address $1 --slave $3 st --jsonpath=./tmp_8.json
    sleep 2s
    ./ncli --address $1 --slave $4 st --jsonpath=./tmp_8.json
    sleep 2s
    ./ncli --address $1 --slave $5 st --jsonpath=./tmp_8.json
    sleep 2s
    ./ncli --address $1 --slave $6 st --jsonpath=./tmp_8.json
    sleep 2s
    ./ncli --address $1 --slave $2 tp
    ./ncli --address $1 --slave $3 tp
    ./ncli --address $1 --slave $4 tp
    ./ncli --address $1 --slave $5 tp
    ./ncli --address $1 --slave $6 tp
    for ((i=1;i<=5;i++))
    do
        echo "************ Turn Off ************"
        sleep 60s
        st_code_9_1=$(getRunStatus1 $1 $2)
        st_code_9_3=$(getRunStatus1 $1 $3)
        st_code_9_5=$(getRunStatus1 $1 $4)
        st_code_9_7=$(getRunStatus1 $1 $5)
        st_code_9_9=$(getRunStatus1 $1 $6)
        [[ "$st_code_9_1" == "148" ]] || echo "** Slave $2 Failed, RunStatus1 should equal 148, actual reault is $st_code_9_1"
        [[ "$st_code_9_1" == "148" ]] && echo "* Slave $2 Pass to set the time"
        [[ "$st_code_9_3" == "148" ]] || echo "** Slave $3 Failed, RunStatus1 should equal 148, actual reault is $st_code_9_3"
        [[ "$st_code_9_3" == "148" ]] && echo "* Slave $3 Pass to set the time"
        [[ "$st_code_9_5" == "148" ]] || echo "** Slave $4 Failed, RunStatus1 should equal 148, actual reault is $st_code_9_5"
        [[ "$st_code_9_5" == "148" ]] && echo "* Slave $4 Pass to set the time"
        [[ "$st_code_9_7" == "148" ]] || echo "** Slave $5 Failed, RunStatus1 should equal 148, actual reault is $st_code_9_7"
        [[ "$st_code_9_7" == "148" ]] && echo "* Slave $5 Pass to set the time"
        [[ "$st_code_9_9" == "148" ]] || echo "** Slave $6 Failed, RunStatus1 should equal 148, actual reault is $st_code_9_9"
        [[ "$st_code_9_9" == "148" ]] && echo "* Slave $6 Pass to set the time"
        echo "************ Turn On ************"
        sleep 60s
        st_code_9_2=$(getRunStatus1 $1 $2)
        st_code_9_4=$(getRunStatus1 $1 $3)
        st_code_9_6=$(getRunStatus1 $1 $4)
        st_code_9_8=$(getRunStatus1 $1 $5)
        st_code_9_10=$(getRunStatus1 $1 $6)
        [[ "$st_code_9_2" == "180" ]] || echo "** Slave $2 Failed, RunStatus1 should equal 180, actual reault is $st_code_9_2"
        [[ "$st_code_9_2" == "180" ]] && echo "* Slave $2 Pass to set the time"
        [[ "$st_code_9_4" == "180" ]] || echo "** Slave $3 Failed, RunStatus1 should equal 180, actual reault is $st_code_9_4"
        [[ "$st_code_9_4" == "180" ]] && echo "* Slave $3 Pass to set the time"
        [[ "$st_code_9_6" == "180" ]] || echo "** Slave $4 Failed, RunStatus1 should equal 180, actual reault is $st_code_9_6"
        [[ "$st_code_9_6" == "180" ]] && echo "* Slave $4 Pass to set the time"
        [[ "$st_code_9_8" == "180" ]] || echo "** Slave $5 Failed, RunStatus1 should equal 180, actual reault is $st_code_9_8"
        [[ "$st_code_9_8" == "180" ]] && echo "* Slave $5 Pass to set the time"
        [[ "$st_code_9_10" == "180" ]] || echo "** Slave $6 Failed, RunStatus1 should equal 180, actual reault is $st_code_9_10"
        [[ "$st_code_9_10" == "180" ]] && echo "* Slave $6 Pass to set the time"
    done
}

function testCase11(){
    week_day=$(getWeekDay 0)
    t_time_0=$(gdate +%H:%M:%S -d "1 minutes")
    t_time_1=$(gdate +%H:%M:%S -d "2 minutes")
    cat ./sample.json| jq --arg t "$t_time_0" '.TimeOnTime0 |= $t'| jq --arg t "$t_time_1" '.TimeOnTime1 |= $t'| jq --arg wdc "$week_day" '.TimeOnDay0[0] |= $wdc'| jq --arg wdc "$week_day" '.TimeOnDay1[0] |= $wdc' > tmp_11.json
    ./ncli --address $1 --slave $2 st --jsonpath=./tmp_11.json
    sleep 5s
    ./ncli --address $1 --slave $2 tp
    echo "************ Turn On ************"
    sleep 60s
    st_code_11_1=$(getRunStatus1 $1 $2)
    [[ "$st_code_11_1" == "180" ]] || echo "** Failed, RunStatus1 should equal 180, actual reault is $st_code_11_1"
    [[ "$st_code_11_1" == "180" ]] && echo "* Pass to set the time"
    echo "************ Turn On Again ************"
    sleep 60s
    st_code_11_2=$(getRunStatus1 $1 $2)
    [[ "$st_code_11_2" == "180" ]] || echo "** Failed, RunStatus1 should equal 180, actual reault is $st_code_11_2"
    [[ "$st_code_11_2" == "180" ]] && echo "* Pass to set the time"
}

function testCase12(){
    week_day=$(getWeekDay 0)
    t_time_0=$(gdate +%H:%M:%S -d "1 minutes")
    t_time_1=$(gdate +%H:%M:%S -d "2 minutes")
    cat ./sample.json| jq --arg t "$t_time_0" '.TimeOffTime0 |= $t'| jq --arg t "$t_time_1" '.TimeOffTime1 |= $t'| jq --arg wdc "$week_day" '.TimeOffDay0[0] |= $wdc'| jq --arg wdc "$week_day" '.TimeOffDay1[0] |= $wdc' > tmp_12.json
    ./ncli --address $1 --slave $2 st --jsonpath=./tmp_12.json
    sleep 5s
    ./ncli --address $1 --slave $2 tp
    echo "************ Turn Off ************"
    sleep 60s
    st_code_12_1=$(getRunStatus1 $1 $2)
    [[ "$st_code_12_1" == "148" ]] || echo "** Failed, RunStatus1 should equal 180, actual reault is $st_code_12_1"
    [[ "$st_code_12_1" == "148" ]] && echo "* Pass to set the time"
    echo "************ Turn Off Again ************"
    sleep 60s
    st_code_12_2=$(getRunStatus1 $1 $2)
    [[ "$st_code_12_2" == "148" ]] || echo "** Failed, RunStatus1 should equal 180, actual reault is $st_code_12_2"
    [[ "$st_code_12_2" == "148" ]] && echo "* Pass to set the time"
}

echo "====================== * Start Test case 7 ======================"
testCase7 "192.168.124.249" "100"

echo "====================== * Start Test case 8 ======================"
testCase8 "192.168.124.149" "100"

echo "====================== * Start Test case 9 ======================"
testCase9 "192.168.124.249" "100" "101" "102" "103" "104"

echo "====================== * Start Test case 11 ======================"
testCase11 "192.168.124.149" "100"

echo "====================== * Start Test case 12 ======================"
testCase12 "192.168.124.149" "100"