import liff from "@line/liff";
import { LineUser } from "../App";

type Props = {
  user: LineUser
}

function GoogleLogin(props: Props) {

  const redirectToWeb = () => {
    liff.openWindow({
      url: import.meta.env.VITE_APP_REDIRECT_WEB_URL + '?id=' + props.user.id,
      external: true,
    });
  }

  return (
    <>
      <p>
        ようこそ、{props.user.displayName}
      </p>
      <button onClick={() => redirectToWeb()}>
          Login with Google
      </button>
    </>
  )
}

export default GoogleLogin