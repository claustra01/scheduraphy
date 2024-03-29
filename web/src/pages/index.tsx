import Head from 'next/head'
import axios from 'axios'
import { useEffect } from 'react'
import { useRouter } from 'next/router'
import { useCookies } from 'react-cookie'

export default function Home() {

  const router = useRouter()
  const [cookie, setCookie, removeCookie] = useCookies(['line_id'])
  
  useEffect(() => {

    const lineId = router.query.id
    if (!lineId) return

    const searchUser = async (id: string) => {
      const res = await axios.get('/api/users?lineId=' + id)
      if (res.status === 200 && res.data.line_id === id) {
        return true
      } else {
        return false
      }
    }

    const loginWithGoogle = async () => {
      const response = await axios.get('/api/google')
      const { authorizeUrl } = response.data
      location.href = authorizeUrl
    }

    const redirect = async () => {
      if (await searchUser(lineId as string)) {
        setCookie('line_id', lineId)
        loginWithGoogle()
      } else {
        router.replace('/unregistered')
      }
    }
    redirect().then()

  }, [router.query.id])

  return (
    <>
      <Head>
        <title>Scheduraphy</title>
        <meta name="description" content="Generated by create next app" />
        <meta name="viewport" content="width=device-width, initial-scale=1" />
        <link rel="icon" href="/favicon.ico" />
      </Head>
    </>
  )
}
