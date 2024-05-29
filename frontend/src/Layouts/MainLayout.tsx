import Topbar from '../components/Topbar/Topbar'
import { Outlet } from 'react-router-dom';

function MainLayout() {
  return (
    <>
    <Topbar />
    <div className="flex w-full h-full">
        {/* <Sidebar /> */}
        <div className="w-full h-full overflow-auto">
            <Outlet />
        </div>
    </div>
</>
  )
}

export default MainLayout
