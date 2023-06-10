import { useEffect, useState } from "react"
import liff from "@line/liff"
import "./App.css"

function App() {
  const [message, setMessage] = useState<string>("")
  const [error, setError] = useState<string>("")

  useEffect(() => {
    liff
      .init({
        liffId: import.meta.env.VITE_APP_LIFF_ID,
      })
      .then(async () => {
        if (liff.isLoggedIn()) {
          const profile = await liff.getProfile()
          setMessage(`Hello, ${profile.displayName}!`)
        }
      })
      .catch((e: Error) => {
        setMessage('LIFF init failed.')
        setError(`${e}`)
      })
  })

  return (
    <div className="App">
      <h1>create-liff-app</h1>
      {message && <p>{message}</p>}
      {error && (
        <p>
          <code>{error}</code>
        </p>
      )}
      <a
        href="https://developers.line.biz/ja/docs/liff/"
        target="_blank"
        rel="noreferrer"
      >
        LIFF Documentation
      </a>
    </div>
  )
}

export default App