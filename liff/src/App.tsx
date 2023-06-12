import { useEffect, useState } from "react"
import { liff } from "@line/liff"
import { initializeApp } from "firebase/app"

import Login from "./auth/Login"
import "./App.css"

const firebaseConfig = {
  apiKey: import.meta.env.VITE_APP_FIREBASE_API_KEY,
  authDomain: import.meta.env.VITE_APP_FIREBASE_AUTH_DOMAIN,
  projectId: import.meta.env.VITE_APP_FIREBASE_PROJECT_ID,
  storageBucket: import.meta.env.VITE_APP_FIREBASE_STORAGE_BUCKET,
  messagingSenderId: import.meta.env.VITE_APP_FIREBASE_MESSAGIN_SENDER_ID,
  appId: import.meta.env.VITE_APP_FIREBASE_APP_ID
}
const firebaseApp = initializeApp(firebaseConfig)

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
        <Login/>
      ) : (
        <p>LINEアプリ内で開いてください</p>
      )}
    </div>
  )
}

export default App