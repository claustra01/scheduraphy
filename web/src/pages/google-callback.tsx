import { useEffect } from 'react'
import { useRouter } from 'next/router'
 
export default function Callback() {

  const router = useRouter()

  useEffect(() => {
    
    const fn = async () => {
      // getパラメータの認証コードを取得する
      const code = router.query.code
      if (!code) return
 
      console.log('code: ', code)
    }
    fn().then()

    router.replace('/')

  }, [router.query.code])
 
  return
}