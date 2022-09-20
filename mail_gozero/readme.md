# 生成步骤 
~~~
$1 goctl api new mail_gozero

$2 cd mail_zoero

$3 go mod tidy

$4 go run mailgozero.go -f etc/mailgozero-api.yaml 

编写api文件后，执行该命令生成api

$5 goctl api go -api mail_gozero.api -dir . -style gozero
