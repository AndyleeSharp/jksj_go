ps -ef|grep shop_interface|grep -v grep|awk -F " " '{print $2}'|xargs kill
ps -ef|grep goods_service|grep -v grep|awk -F " " '{print $2}'|xargs kill
ps -ef|grep rate_service|grep -v grep|awk -F " " '{print $2}'|xargs kill
