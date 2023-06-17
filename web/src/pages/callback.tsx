import { useEffect, useState } from 'react'
import axios from 'axios'
import { useRouter } from 'next/router'
import { useCookies } from 'react-cookie'

export default function Callback() {

  const router = useRouter()
  const [cookie, setCookie, removeCookie] = useCookies(['line_id'])
  const [lineId, setLineId] = useState<string>('')
  const [tokens, setTokens] = useState<any>({}) // Object

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
      setTokens(tokens)    
    }

    getCookie()
    getToken().then()

  }, [router.query.code])

  useEffect(() => {

    const updateToken = async (id: string, token: string) => {
      const res = await axios.patch('/api/users', {
        lineId: id,
        refreshToken: token
      })
      if (res.status === 200) {
        router.replace('/result')
      } else {
        alert('ログインに失敗しました')
        router.replace('/')
      }
    }

    const redirect = async () => {
      if (lineId === '' || tokens.toString() === '{}') return
      if (!tokens.refresh_token) {
        // トークンが既に発行済み
        router.replace('/result')
      } else {
        // トークン更新
        updateToken(lineId, tokens.refresh_token)
      }
    }
    redirect().then()

  }, [tokens])
 
  return
}