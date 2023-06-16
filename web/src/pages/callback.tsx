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

    const searchUser = async (id: string) => {
      const res = await axios.get('/api/users?lineId=' + id)
      if (res.status === 200 && res.data.line_id === id) {
        return true
      } else {
        return false
      }
    }

    const registerUser = async (id: string, token: string) => {
      const res = await axios.post('/api/users', {
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

    const updateUser = async (id: string, token: string) => {
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
      const userExist = await searchUser(lineId)
      console.log(tokens)
      if (userExist) {
        if (tokens.refresh_token === '') {
          router.replace('/result')
        } else {
          updateUser(lineId, tokens.refresh_token)
        }
      } else {
        registerUser(lineId, tokens.refresh_token)
      }
    }
    redirect().then()

  }, [tokens])
 
  return
}