import liff from "@line/liff";

function GoogleLogin() {

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

export default GoogleLogin