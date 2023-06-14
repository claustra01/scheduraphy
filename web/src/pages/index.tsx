import Head from 'next/head'
import { Inter } from 'next/font/google'
import axios from 'axios'
import { useEffect } from 'react'
import { useRouter } from 'next/router'
import { useCookies } from 'react-cookie'

const inter = Inter({ subsets: ['latin'] })

export default function Home() {

  const router = useRouter()
  const [cookie, setCookie, removeCookie] = useCookies(['id'])
  
  useEffect(() => {

    const lineId = router.query.id
    if (!lineId) return
    setCookie('id', lineId)

    const loginWithGoogle = async () => {
      const response = await axios.get('/api/google')
      const { authorizeUrl } = response.data
      location.href = authorizeUrl
    }
    loginWithGoogle()

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
