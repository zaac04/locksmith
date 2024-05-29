import { RouteObject } from 'react-router-dom';
import MainLayout from '../../Layouts/MainLayout';
import Env from '../../pages/Env';

export const routes:RouteObject[] = [
    {
        element:<MainLayout/>,
        path:"/",
        children:[
            {
                path: "/",
                element: <Env/>
            },
        ]
    },
]