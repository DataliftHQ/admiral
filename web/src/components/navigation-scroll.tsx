import { type ReactElement, useEffect } from 'react'

const NavigationScroll = ({ children }: { children: ReactElement | null }): ReactElement | null => {
  useEffect(() => {
    window.scrollTo({
      top: 0,
      left: 0,
      behavior: 'smooth',
    })
  }, [])

  return children ?? null
}

export default NavigationScroll
