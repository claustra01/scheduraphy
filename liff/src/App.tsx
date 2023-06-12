import { useEffect, useState } from "react"
import liff from "@line/liff"
import "./App.css"

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
    <div className="App">
      <h1>Scheduraphy</h1>
      {auth ? (
        <button>Googleでログイン</button>
      ) : (
        <p>LINEアプリ内で開いてください</p>
      )}
    </div>
  )
}

export default App