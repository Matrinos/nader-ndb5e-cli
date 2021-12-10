# test case 1
echo "====================== * Start Test case 1 ======================"
./ncli --address 192.168.124.149 --slave 100 to
sleep  10s
./ncli --address 192.168.124.149 --slave 100 to
sleep  10s

# test case 2
echo "====================== * Start Test case 2 ======================"
./ncli --address 192.168.124.149 --slave 100 tf
sleep  10s
./ncli --address 192.168.124.149 --slave 100 tf
sleep  10s

# test case 3
echo "====================== * Start Test case 3 ======================"
./nader_shell_test_tof.sh 192.168.124.149 100  50 4s 

# test case 4
echo "====================== * Start Test case 4 ======================"
./nader_shell_test_tof.sh 192.168.124.149 100  20 60s

# test case 5
echo "====================== * Start Test case 5 ======================"
./nader_shell_test_tof.sh 192.168.124.149 100  50 15s > test_case5_output_1.log &
sleep 5s
./nader_shell_test_tof.sh 192.168.124.149 100  50 15s > test_case5_output_2.log &
sleep 5s
./nader_shell_test_tof.sh 192.168.124.149 100  50 15s > test_case5_output_3.log &

# test case 6
echo "====================== * Start Test case 6 ======================"
./nader_shell_test_tof.sh 192.168.124.249 100  50 5s > test_case6_output_1.log &
sleep 5s
./nader_shell_test_tof.sh 192.168.124.249 101  50 5s > test_case6_output_2.log &
sleep 5s
./nader_shell_test_tof.sh 192.168.124.249 102  50 5s > test_case6_output_3.log &
sleep 5s
./nader_shell_test_tof.sh 192.168.124.249 103  50 5s > test_case6_output_4.log &
sleep 5s
./nader_shell_test_tof.sh 192.168.124.249 104  50 5s > test_case6_output_5.log &
