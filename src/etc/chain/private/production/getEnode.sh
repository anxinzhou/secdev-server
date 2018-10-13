node=$1

nodekey="${node}/geth/nodekey"

pubAdd=$(bootnode --nodekey $nodekey -writeaddress)

ip=$(/sbin/ifconfig -a|grep inet|grep -v 127.0.0.1|grep -v inet6|awk '{print $2}'|tr -d "addr:")

enode="enode://${pubAdd}@${ip}:30311"

echo $enode
