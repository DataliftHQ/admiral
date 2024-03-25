import React from 'react'
import { RouterProvider } from 'react-router-dom'

import router from './routes'
import NavigationScroll from './components/navigation-scroll'

function App() {
  return (
    <NavigationScroll>
      <RouterProvider router={router} />
    </NavigationScroll>
  )
}

export default App
