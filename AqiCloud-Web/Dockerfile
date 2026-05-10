# ================================================================
# 🏗️ AQI-AIPan 前端 Dockerfile（腾讯云新加坡服务器 2C2G 优化版）
# 多阶段构建：Vue.js 静态文件 + Nginx 静态服务
#
# 🎯 服务器配置：腾讯云轻量应用服务器（新加坡）
#   - CPU：2 核
#   - 内存：2GB RAM
#   - 实例类型：S2-2C2G 或类似配置
# ================================================================

# ================================================================
# 📦 第一阶段：构建阶段（Node.js 18 Alpine）
# ================================================================
FROM swr.cn-north-4.myhuaweicloud.com/ddn-k8s/docker.io/node:18-alpine AS build-stage

# 设置工作目录
WORKDIR /app

# 安装构建依赖
RUN apk add --no-cache git

# 复制依赖描述文件（利用 Docker 层缓存）
COPY package*.json ./

# 配置 npm 镜像源
RUN npm config set registry https://registry.npmmirror.com && \
    npm config set fetch-retries 3 && \
    npm config set fetch-retry-mintimeout 20000 && \
    npm config set fetch-retry-maxtimeout 120000

# 安装依赖
RUN npm install --legacy-peer-deps --force && \
    echo "✅ Dependencies installed successfully"

# 复制源代码
COPY . .

# 构建生产环境静态文件
RUN npm run build && \
    echo "✅ Build completed successfully" && \
    ls -lh dist

# ================================================================
# 🚀 第二阶段：运行阶段（Nginx Alpine）
# ================================================================
FROM swr.cn-north-4.myhuaweicloud.com/ddn-k8s/docker.io/nginx:1.25-alpine AS server-stage

# 添加标签元数据
LABEL maintainer="AQI Team" \
      version="1.0.0" \
      description="AQI-AIPan Frontend - Static Files Serving (Tencent Cloud Singapore 2C2G Optimized)" \
      vendor="Tencent Cloud" \
      region="ap-singapore"

# 设置时区为新加坡
ENV TZ=Asia/Singapore \
    LANG=C.UTF-8 \
    LC_ALL=C.UTF-8

RUN apk add --no-cache tzdata && \
    cp /usr/share/zoneinfo/$TZ /etc/localtime && \
    echo $TZ > /etc/timezone && \
    echo "✅ Timezone set to $TZ" && \
    date

# 复制 Nginx 配置文件
COPY nginx.conf /etc/nginx/conf.d/default.conf

# 复制构建产物
COPY --from=build-stage /app/dist /usr/share/nginx/html

# 创建必要的目录并设置权限
RUN mkdir -p /var/log/nginx /var/cache/nginx && \
    chown -R nginx:nginx /var/log/nginx /var/cache/nginx /usr/share/nginx/html && \
    echo "✅ Permissions set"

# 声明暴露端口
EXPOSE 80

# 健康检查
HEALTHCHECK --interval=30s --timeout=10s --start-period=10s --retries=3 \
  CMD wget --quiet --tries=1 --spider http://localhost/ || exit 1

# 使用非 root 用户运行（安全最佳实践）
USER nginx

# 启动 Nginx
CMD ["nginx", "-g", "daemon off;"]

# ================================================================
# 📝 使用说明
# ================================================================
# 1️⃣ 构建镜像
# ------------------------------------------------
#    docker build -t aqi-aipan-frontend:latest .
#
# ================================================================
# 2️⃣ 运行容器（2C2G 配置）
# ------------------------------------------------
#    docker run -d \
#      --name aqi-aipan-frontend \
#      --restart unless-stopped \
#      -p 80:80 \
#      --memory=256m \
#      --memory-swap=512m \
#      --cpus=0.5 \
#      --read-only \
#      --tmpfs /var/cache/nginx:rw,noexec,nosuid,size=10m \
#      --tmpfs /var/run:rw,noexec,nosuid,size=1m \
#      aqi-aipan-frontend:latest
#
# ================================================================
# 3️⃣ 查看日志
# ------------------------------------------------
#    docker logs -f aqi-aipan-frontend
#
# ================================================================
# 4️⃣ 性能优化说明
# ================================================================
# - 使用 Alpine 镜像，体积小
# - 静态资源开启 Gzip 压缩
# - SPA 路由支持（try_files）
# - 静态文件缓存配置
# - 只读文件系统增强安全性
# - 非 root 用户运行提升安全性
#