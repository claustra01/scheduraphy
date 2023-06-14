import { useEffect, useState } from 'react'
import './App.css'
import liff from '@line/liff'

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

  const redirectToWeb = () => {
    liff.openWindow({
      url: import.meta.env.VITE_APP_REDIRECT_WEB_URL,
      external: true,
    });
  }
  
  return (
    <button onClick={() => redirectToWeb()}>
        Login with Google
    </button>
  )
}

export default App
