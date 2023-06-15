import liff from "@line/liff"

type Props = {
  userId: string,
  displayName: string
}

function GoogleLogin({props}: {props: Props}) {

  const redirectToWeb = () => {
    liff.openWindow({
      url: import.meta.env.VITE_APP_REDIRECT_WEB_URL + '?id=' + props.userId,
      external: true,
    })
  }

  return (
    <>
      <p>
        ようこそ、{props.displayName}
      </p>
      <button onClick={() => redirectToWeb()}>
          Login with Google
      </button>
    </>
  )
}

export default GoogleLogin