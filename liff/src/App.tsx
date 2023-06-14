import { useEffect, useState } from 'react'
import './App.css'
import liff from '@line/liff'
import GoogleLogin from './components/GoogleLogin'
import InExternalBrowser from './components/InExternalBrowser'

export type LineUser = {
  id: string,
  displayName: string,
}

function App() {

  const [auth, setAuth] = useState<boolean>(false)
  const [lineUser, setLineUser] = useState<LineUser>({id: '', displayName: ''})

  useEffect(() => {
    liff
      .init({ liffId: import.meta.env.VITE_APP_LIFF_ID })
      .then(async () => {
        if (liff.isLoggedIn()) {
          const profile = await liff.getProfile()
          setAuth(true)
          setLineUser({
            id: profile.userId,
            displayName: profile.displayName
          })
        }
      })
      .catch(() => {
        setAuth(false)
      })
  })
  
  return (
    <>
      {auth ? <GoogleLogin user={lineUser} /> : <InExternalBrowser/>}
    </>
  )
}

export default App
