import { useEffect, useState } from 'react'
import axios from 'axios'
import { useRouter } from 'next/router'
import { useCookies } from 'react-cookie'

export default function Callback() {

  const router = useRouter()
  const [cookie, setCookie, removeCookie] = useCookies(['line_id'])
  const [lineId, setLineId] = useState<string>('')
  const [refreshToken, setRefreshToken] = useState<string>('')

  useEffect(() => {

    const getCookie = () => {
      if (cookie.line_id) {
        setLineId(cookie.line_id)
        removeCookie('line_id')
      }
    }
    
    const getToken = async () => {
      const authCode = router.query.code
      if (!authCode) return
      const res = await axios.post('/api/google', {
        authorizationCode: authCode,
      })
      const { tokens } = res.data
      setRefreshToken(tokens.refresh_token)    
    }

    getCookie()
    getToken().then()

  }, [router.query.code])

  useEffect(() => {

    const registerUser = async (id: string, token: string) => {
      if (refreshToken === '') return
      const res = await axios.post('/api/users', {
        lineId: id,
        refreshToken: token
      })
      if (res.status == 200) {
        router.replace('/result')
      } else {
        alert('ログインに失敗しました')
        router.replace('/')
      }
    }
    registerUser(lineId, refreshToken).then()

  }, [refreshToken])
 
  return
}