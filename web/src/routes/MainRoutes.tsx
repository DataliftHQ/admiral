import React, { lazy } from 'react'

import Loadable from '../components/loadable'
import MainLayout from '../layouts/MainLayout'
import AuthGuard from './AuthGuard'

const Dashboard = Loadable(lazy(async () => await import('../pages/Dashboard')))

const MainRoutes = {
  path: '/',
  element: (
    <AuthGuard>
      <MainLayout />
    </AuthGuard>
  ),
  children: [{ path: '/', element: <Dashboard /> }],
}

export default MainRoutes
