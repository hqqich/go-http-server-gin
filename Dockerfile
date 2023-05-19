# 引入基础运行环境
FROM alpine:latest

# 设置维护者信息
LABEL org.opencontainers.image.authors="hqqich1314@outlook.com"

# 设置工作目录：/app
WORKDIR /app

# 执行shell指令
RUN mkdir /app/static

# 卷标挂载
VOLUME G:\\opt /app/static

# 将本机文件移动到docker工作目录
COPY goHttpServerGin-linux-64 /app
COPY config.ini /app

# 移动文件
COPY index.html /app/static
COPY ./dist/ /app/dist

# 修改可执行文件权限
RUN chmod 777 /app/goHttpServerGin-linux-64

# # 执行shell指令
# RUN mkdir /app/static

# 指定端口，docker对外的端口
EXPOSE 8888

# 设置工作目录：/app
WORKDIR /app

#需要运行的命令
CMD ["/app/goHttpServerGin-linux-64"]