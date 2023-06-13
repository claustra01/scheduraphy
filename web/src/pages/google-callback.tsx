import { useEffect } from 'react'
import { useRouter } from 'next/router'
import axios from 'axios'

export default function Callback() {

  const router = useRouter()

  useEffect(() => {
    
    const fn = async () => {
      // getパラメータの認証コードを取得する
      const code = router.query.code
      if (!code) return
 
      const response = await axios.post('/api/get-google-auth-token', {
        authorizationCode: code,
      })

      const { tokens } = response.data
      console.log('tokens: ', tokens)
      
    }
    fn().then()

    router.replace('/')

  }, [router.query.code])
 
  return
}