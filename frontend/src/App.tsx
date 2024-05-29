import { RouterProvider, createBrowserRouter } from 'react-router-dom';
import { routes } from './ config/routes/routes.config';


function App() {
  const basename = '/';
  const router = createBrowserRouter(routes, {
    basename,
});

  return (
    <>
      <main className=' bg-gray-900 w-screen h-screen fixed'>
        <RouterProvider router={router}></RouterProvider>
      </main>
    </>
  )
}

export default App
