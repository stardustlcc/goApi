
# 目录结构
    app
        api         项目接口
        models      表model
        response    接口返回的结构体格式   
        routers     项目路由
    configs
        config.yaml 配置文件
    global          常量配置
    pkg             一些通用配置
    logs            日志
    main.go
    go.mod



# clone project
# 观察是否有配置文件 和 日志目录
go mod init dwd-api

go mod tidy

go run main.go