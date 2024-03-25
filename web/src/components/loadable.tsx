import React, { Suspense, type ElementType } from 'react'

import Loader from './loader'

// eslint-disable-next-line @typescript-eslint/no-explicit-any
const Loadable = (Component: ElementType) => (props: any) => (
  <Suspense fallback={<Loader />}>
    <Component {...props} />
  </Suspense>
)

export default Loadable
