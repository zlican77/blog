import { createRouter, createWebHistory } from 'vue-router';
import HomeView from '../views/HomeView.vue';
import ComputerView from '../views/ComputerView.vue';
import GolangView from '../views/GolangView.vue';
import LeetCode from '../views/LeetCode.vue';
import ProjectView from '../views/ProjectView.vue';
import ComputerNetwork from '../views/computer/ComputerNetwork.vue';
import DataAlgorithm from '../views/computer/DataAlgorithm.vue';
import ComputerArchitecture from '../views/computer/ComputerArchitecture.vue';
import ComputerOS from '../views/computer/ComputerOS.vue';
import LinuxServer from '../views/computer/LinuxServer.vue';
import DataBases from '../views/computer/DataBases.vue';
import AdminView from '../views/AdminView.vue';
import CreateArticle from '../views/admin/CreateArticle.vue';
import ManageArticle from '../views/admin/ManageArticle.vue';
import ManageTag from '../views/admin/ManageTag.vue';
import ManageMessage from '../views/admin/ManageMessage.vue';
import ArticleDetail from '../components/ArticleDetail.vue';
import LoginBoard from '../views/home/LoginBoard.vue';
import RegisterBoard from '../views/home/RegisterBoard.vue';



const routes = [
    {
        path: '/',
        redirect: '/home'
    },
    { 
        path: '/home', 
        name: 'HomeView',
        component: HomeView,
        children: [
            {
                path: 'login',
                name: 'LoginBoard',
                component: LoginBoard
            },
            {
                path: 'register',
                name: 'RegisterBoard',
                component: RegisterBoard
            }
        ]
    },
    {
        path: '/computer',
        name: 'ComputerView',
        component: ComputerView,
        children: [
            {
                path: '',
                redirect: '/computer/network'
            },
            {
                path: 'network',
                name: 'ComputerNetwork',
                component: ComputerNetwork
            },
            {
                path: 'algorithm',
                name: 'DataAlgorithm',
                component: DataAlgorithm
            },
            {
                path: 'architecture',
                name: 'ComputerArchitecture',
                component: ComputerArchitecture
            },
            {
                path: 'os',
                name: 'ComputerOS',
                component: ComputerOS
            },
            {
                path: 'linux',
                name: 'LinuxServer',
                component: LinuxServer
            },
            {
                path: 'database',
                name: 'DataBases',
                component: DataBases
            }
        ]
    },
    {
        path: '/golang',
        name: 'GolangView',
        component: GolangView
    },
    {
        path: '/leetcode',
        name: 'LeetCode',
        component: LeetCode
    },
    {
        path: '/project',
        name: 'ProjectView',
        component: ProjectView
    },
    {
        path: '/admin',
        name: 'AdminView',
        component: AdminView,
        children: [
            {
                path: '',
                redirect: '/admin/create'
            },
            {
                path: 'create',
                name: 'CreateArticle',
                component: CreateArticle
            },
            {
                path: 'articles',
                name: 'ManageArticle',
                component: ManageArticle
            },
            {
                path: 'tags',
                name: 'ManageTag',
                component: ManageTag
            },
            {
                path: 'messages',
                name: 'ManageMessage',
                component: ManageMessage
            }
        ]
    },
    {
        path: '/articles/:id',  // 文章详情路由
        name: 'ArticleDetail',
        component: ArticleDetail
    }
]

const router = createRouter({
    history: createWebHistory(),
    routes,
})


export default router;
