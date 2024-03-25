import { createBrowserRouter } from 'react-router-dom'

import MainRoutes from '../routes/MainRoutes'
import SimpleLayout from '../layouts/SimpleLayout'
import NotFound from '../pages/NotFound'

const router = createBrowserRouter([
  MainRoutes,
  {
    element: <SimpleLayout />,
    children: [{ path: '*', element: <NotFound /> }],
  },
])

export default router
