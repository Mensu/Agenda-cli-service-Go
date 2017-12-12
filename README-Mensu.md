# fork 项目的位置

[https://github.com/Mensu/Agenda-cli-service-Go](https://github.com/Mensu/Agenda-cli-service-Go)

# 个人工作摘要（每次提交）

- [``b72ae71d1d``](https://github.com/Mensu/Agenda-cli-service-Go/commit/b72ae71d1d) 搭建好作业的框架，准备开始同步开发
- [``b72ae71d1d``](https://github.com/Mensu/Agenda-cli-service-Go/commit/b72ae71d1d) 连接 [``Travis CI``](https://travis-ci.org/Mensu/Agenda-cli-service-Go) 为项目持续集成
- [``42ee4feab7``](https://github.com/Mensu/Agenda-cli-service-Go/commit/42ee4feab7) 在 [``Docker Hub``](https://hub.docker.com/r/mensu/agenda-cli-service-go/) 上生成镜像
- [``be159ae036``](https://github.com/Mensu/Agenda-cli-service-Go/commit/be159ae036) [service] 初步完成登录API
- [``679c8cbf81``](https://github.com/Mensu/Agenda-cli-service-Go/commit/679c8cbf81) [service] 初步完成用户登出API
- [``d3d517bb36``](https://github.com/Mensu/Agenda-cli-service-Go/commit/d3d517bb36) [service] 初步完成删除用户账户的API
- [``a22a422456``](https://github.com/Mensu/Agenda-cli-service-Go/commit/a22a422456) [service] 初步完成用户注册的API
- [``2658441b00``](https://github.com/Mensu/Agenda-cli-service-Go/commit/2658441b00) [service] 完善 user API 的单元测试
- [``b780a9a737``](https://github.com/Mensu/Agenda-cli-service-Go/commit/b780a9a737) [service] 初步完成用户列表的API
- [``1c9b9ee1e5``](https://github.com/Mensu/Agenda-cli-service-Go/commit/1c9b9ee1e5) [service] 初步完成创建会议的API
- [``b5072c579a``](https://github.com/Mensu/Agenda-cli-service-Go/commit/b5072c579a) [service] 初步完成添加会议成员的API
- [``825fd6c9a2``](https://github.com/Mensu/Agenda-cli-service-Go/commit/825fd6c9a2) [service] 初步完成删除会议成员的API
- [``f13c61aeb4``](https://github.com/Mensu/Agenda-cli-service-Go/commit/f13c61aeb4) [service] 初步完成退出会议的API
- [``5b77e971b4``](https://github.com/Mensu/Agenda-cli-service-Go/commit/5b77e971b4) [service] 初步完成搜索会议的API
- [``ccebe92fdd``](https://github.com/Mensu/Agenda-cli-service-Go/commit/ccebe92fdd) [service] 完善 meeting API 的单元测试
- [``d7f401638b``](https://github.com/Mensu/Agenda-cli-service-Go/commit/d7f401638b) [service] 完善 meeting_participator API 的单元测试

# 项目小结

通过本次项目，我学到了很多。

首先是体验了服务程序开发实战，对命令行客户端调用远端服务这一服务开发的重要内容有了更深理解。

接着是学习了 API 设计工具，尝试实现了从资源（领域）建模，到 API 设计的过程，体验了 RESTful API 设计的优势。更重要的，是实践了前后端通过 API mock server 同步开发的过程。我和另一团队成员一人负责服务端，一人负责客户端，双方都遵守 github 上商定好的 API 接口，便可以不受对方进度的影响，专心进行客户端或服务端的开发，提高了开发的生产力。另外，两个团队还可以独立测试 API，更有利于模块的解耦。

更重要的是，这次还实践了 Docker，初次尝试在 docker hub 上生成镜像，并使用镜像进行综合测试。
