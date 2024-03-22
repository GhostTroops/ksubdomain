doTag $1
cd $HOME/MyWork/goSqlite_gorm
go get "github.com/GhostTroops/go-utils@$1"
go mod vendor
cd $HOME/MyWork/ksubdomain
go get "github.com/GhostTroops/go-utils@$1"
go mod vendor

