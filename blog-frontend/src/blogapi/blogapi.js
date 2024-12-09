import request from "./config";

export const blogapi = {
    getNetworkArticles() {
        return request({
            url: '/articles/network',
            method: 'GET'
        })
    },
    getAlgorithmArticles() {
        return request({
            url: '/articles/algorithm',
            method: 'GET'
        })
    },
    getArchitectureArticles() {
        return request({
            url: '/articles/architecture',
            method: 'GET'
        })
    },
    getOsArticles() {
        return request({
            url: '/articles/os',
            method: 'GET'
        })
    },
    getLinuxArticles() {
        return request({
            url: '/articles/linux',
            method: 'GET'
        })
    },
    getDatabaseArticles() {
        return request({
            url: '/articles/database',
            method: 'GET'
        })
    },
    getGolangArticles() {
        return request({
            url: '/articles/golang',
            method: 'GET'
        })
    },  
    getLeetcodeArticles() {
        return request({
            url: '/articles/leetcode',
            method: 'GET'
        })
    },
    getProjectArticles() {
        return request({
            url: '/articles/project',
            method: 'GET'
        })
    },


    // 获取验证码
    getCaptcha(uuid) {
        return request({
            url: `/home/captcha/${uuid}`,
            method: 'GET'
        })
    },
    // 获取手机验证码
    getPhoneCode(object) {
        return request({
            url:  `/home/register/${object.phone}/${object.captcha}/${object.uuid}`,
            method: 'GET'
        })
    },
    // 注册
    postRegister(object) {
        return request({
            url: '/home/register',
            method: 'POST',
            data: object
        })
    },
    // 登录
    postLogin(object) {
        return request({
            url: '/home/login',
            method: 'POST',
            data: object
        })
    },
    // 检查登录
    checkLogin(){
        return request({
            url: "/home/checkLogin",
            method: 'GET'
        })
    },


    // 获取文章详情
    getArticleById(id) {
        return request({
            url: `/articles/${id}`,
            method: 'GET'
        })
    },
    // 获取所有文章
    getAllArticles() {
        return request({
            url: '/articles/all',
            method: 'GET'
        })
    },
    // 创建文章
    createArticle(article) {
        return request({
            url: '/articles',
            method: 'POST',
            data: article
        })
    },
    // 更新文章
    updateArticle(article) {
        return request({
            url: `/articles/${article.ArticleId}`,
            method: 'PUT',
            data: article
        })
    },
    // 删除文章
    deleteArticle(id) {
        return request({
            url: `/articles/${id}`,
            method: 'DELETE'
        })
    }
}