


# custom Handlerを作成
func new -l Custom -t HttpTrigger -n hello -a anonymous


# ローカルで実行
func start

# ローカルでazure functionを起動するためにserver/　内でbuildしてそのファイルをhostingに移動
go build main.go && mv main ./hosting