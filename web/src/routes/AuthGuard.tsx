import React, { useEffect, useState } from 'react'
import { useDispatch } from 'react-redux'

import { isAdmiralError } from '../services/errors.ts'
import { setUser } from '../store/slices/user'
import { services } from '../services'

function FallbackComponent(): JSX.Element {
  return <div>An error has occurred</div>
}

const AuthGuard = ({ children }: { children: React.ReactElement }): JSX.Element => {
  const dispatch = useDispatch()
  const [isLoading, setLoading] = useState(true)
  const [isError, setError] = useState(false)

  useEffect(() => {
    const fetchData = async (): Promise<void> => {
      try {
        setLoading(true)

        const user = await services.user.get()
        dispatch(setUser(user))

        setLoading(false)
      } catch (error) {
        console.error('Error fetching user:', error)
        if (isAdmiralError(error) && error.status?.code !== 401) {
          setLoading(false)
          setError(true)
        }
      }
    }

    void fetchData()
  }, [dispatch])

  if (isLoading) {
    return (
      <div>
        <p>Loading...</p>
      </div>
    )
  } else {
    if (isError) {
      return <FallbackComponent />
    } else {
      return <>{children}</>
    }
  }
}

export default AuthGuard
