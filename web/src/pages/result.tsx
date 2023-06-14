import { useRouter } from 'next/router'
import React, { useEffect } from 'react'

export default function Result() {

  const router = useRouter()

  useEffect(() => {
    router.replace('line://')
  })
 
  return (
    <>
      Login Successful!
    </>
  )
}