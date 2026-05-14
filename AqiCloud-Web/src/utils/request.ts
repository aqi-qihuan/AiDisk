/**
 * Spring Boot 服务 Axios 实例
 * 端口: 8080 (开发) / pan.aqi125.cn (生产)
 * 用于: 文件管理、用户管理、分享、回收站等核心功能
 */
import axios from "axios";
import { ElMessage } from "element-plus";
import JSONBIG from "json-bigint";
import { useLoginUserStore } from "@/store/user";

axios.defaults.transformResponse = [];
// 创建自定义 Axios 实例
const myAxios = axios.create({
  // 使用相对路径，由 Vite 代理到后端服务（开发环境）
  // 生产环境由 Nginx 代理
  baseURL: "/api",
  timeout: 10000,
  withCredentials: false,
  // 直接在创建时设置 transformResponse
  transformResponse: [
    function (data, headers) {
      try {
        // 检查是否是二进制数据（Blob、ArrayBuffer 等）
        if (data instanceof Blob || data instanceof ArrayBuffer) {
          return data;
        }

        // 检查 Content-Type 是否是二进制类型
        const contentType = headers?.["content-type"] || "";
        if (
          contentType.includes("application/octet-stream") ||
          contentType.includes("application/zip") ||
          contentType.includes("application/x-zip-compressed")
        ) {
          return data;
        }

        if (typeof data !== "string") {
          // 如果数据不是字符串（例如，对于二进制数据），则直接返回
          return data;
        }

        // 使用 json-bigint 解析 JSON 数据，storeAsString 参数确保大整数被存储为字符串
        const jsonParser = JSONBIG({ storeAsString: true });
        const parsedData = jsonParser.parse(data);

        return parsedData;
      } catch (error) {
        console.error("Error parsing response:", error);
        // 对于解析失败的情况，可能是二进制数据，直接返回
        return data;
      }
    },
  ],
});

// 添加请求拦截器
myAxios.interceptors.request.use(
  (config) => {
    // 在每次请求前添加token
    // 注意：不能解构store，否则会丢失响应式
    const loginUserStore = useLoginUserStore();
    const token = loginUserStore.token;
    if (token) {
      config.headers["token"] = `${token}`;
    }
    return config;
  },
  (error) => {
    return Promise.reject(error);
  },
);

// 添加响应拦截器
myAxios.interceptors.response.use(
  function (response) {
    console.log("Raw response:", response);
    // localStorage.setItem("time", response.data.data.createTime)

    // 检查并处理数据（如果需要）
    if (typeof response.data === "object") {
      console.log("Parsed data in response interceptor:", response.data);
    }

    // 继续原有的逻辑...
    const { data } = response;
    if (data.code === 20004) {
      if (
        !response.request.responseURL.includes("/account/v1/detail") &&
        !window.location.pathname.includes("/account/v1/login")
      ) {
        ElMessage.warning("请先登录");
        window.location.href = `/account/v1/login?redirect=${window.location.href}`;
      }
    }

    return response;
  },
  function (error) {
    return Promise.reject(error);
  },
);

export default myAxios;
