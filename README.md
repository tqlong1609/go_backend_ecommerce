# go_backend_ecommerce

## Cấu trúc thư mục
- cmd: chứa các file main
+ cli: run background service
+ cronjob: run cronjob
+ server (main): run server
- configs: tham số và 3th party database
- docs: tài liệu
- global: các biến toàn cục, tham chiếu đến tất cả modal trong internal, vd: mongodb, sql
- internal:
+ dao: (data access object)
+ middleware
+ model
+ routes
+ service
- pkg
- storage
- script
- third_party

mermaid (link): https://mermaid.live/edit#pako:eNp9kT2PwjAMhv9K5RmE7tg6sNAbWQ5YIAy-1kcj2rhyHaQK8d8JafmQjiOD49fPKydOTpBzQZDCXrApk1VmXBLWuiX52F7jLhmPZ8k3eyV5sM83bPoP62NfiXDOToWr6mZ46IiXJEebU88G0TelhlurLN3Q8q4jzlDxB1t6CRdh2mprIO6TL6dWOwO7v4dEw8t7PZGnYYYqjKAmqdEW4U1PV5cBLakmA2lIC5SDAePOwYdeedm5HFIVTyMQ9vsS0l-s2qB8U6BSZjF8TH2vNug2zDd9vgCkmpAw
graph TD
    User1[User] --> Router
    User2[User] --> Router
    User3[User] --> Router
    
    Router --> Controller
    Controller --> Service
    Service --> Repository
    Repository --> Database
    Repository --> Model["Model/Entity"]
    Service --> Model
    Controller --> Model
    Router --> Model

# Initialization
```bash
go mod init
```

# Run
```bash
curl -H "Authorization: token" http://localhost:8080/v1/ping1
```