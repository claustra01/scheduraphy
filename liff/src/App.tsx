import { useEffect, useState } from 'react'
import './App.css'
import liff from '@line/liff'
import GoogleLogin from './components/GoogleLogin'
import InExternalBrowser from './components/InExternalBrowser'

function App() {

  const [auth, setAuth] = useState<boolean>(false)

  useEffect(() => {
    liff
      .init({ liffId: import.meta.env.VITE_APP_LIFF_ID })
      .then(() => {
        if (liff.isLoggedIn()) {
          setAuth(true)
        }
      })
      .catch(() => {
        setAuth(false)
      })
  })
  
  return (
    <>
      {auth ? <GoogleLogin/> : <InExternalBrowser/>}
    </>
  )
}

export default App
