import { RouteObject } from 'react-router-dom';
import MainLayout from '../../Layouts/MainLayout';
import Env from '../../pages/Env';
import Error from '../../pages/Error';

export const routes:RouteObject[] = [
    {
        element:<MainLayout/>,
        path:"/",
        hasErrorBoundary:true,
        ErrorBoundary:Error,
        errorElement:<Error/>,
        children:[
            {
                path: "/",
                element: <Env/>
            },
        ]
    },
]