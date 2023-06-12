import { useEffect } from "react"
import { getAuth, GoogleAuthProvider, signInWithRedirect, getRedirectResult } from "firebase/auth"

function Login() {
  
  const auth = getAuth()
  const provider = new GoogleAuthProvider()

  const loginWithGoogle = () => signInWithRedirect(auth, provider)

  useEffect(() => {
    getRedirectResult(auth)
    .then((result) => {
      if(result !== null){
        console.log(result)
      }
    }).catch((error) => {
      console.error(error);
    });    
  });
  
  return (
    <button onClick={() => loginWithGoogle()}>
      Login with Google
    </button>
  )
}

export default Login