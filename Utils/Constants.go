package Utils

const COMMAND_FOR_CPU = "mpstat |awk  '{if ($4 != \"CPU\") print $5 \" \" $7 \" \" $8 \" \" $14}'"

const COMMAND_FOR_MEMORY = "free -m | grep \"Mem\" | awk '{print $3/$2*100}' && free -m | grep \"Mem\" | awk '{print $4/$2*100}'"

const COMMAND_FOR_DISK = "df . | awk 'NR==2 {print $5}'"

const COMMAND_FOR_UPTIME = "awk '{print $1}' /proc/uptime"

const COMMAND_FOR_SYSTEM_INFO = "hostnamectl | grep \"Static\"| awk '{print $3}' && hostnamectl | grep \"Operating\"| awk '{print $3}' && hostnamectl |grep \"Operating\"| awk '{print $4}'"

const COMMAND_FOR_IFCONFIG = "cat /proc/net/dev | awk 'NR>2 {gsub(/:/,\"\"); print $1, $2, $10}'"
