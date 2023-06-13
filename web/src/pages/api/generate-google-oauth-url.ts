import type { NextApiRequest, NextApiResponse } from 'next'
import { google } from 'googleapis'

type Response = {
  authorizeUrl: string
}

// Google認証画面のURLを生成しレスポンスに含める
export default function handler(req: NextApiRequest, res: NextApiResponse<Response>) {
  const oauth2Client = new google.auth.OAuth2({
    clientId: process.env.GOOGLE_OAUTH_CLIENT_ID,
    clientSecret: process.env.GOOGLE_OAUTH_CLIENT_SECRET,
    redirectUri: process.env.ORIGIN_URI + '/google-callback'
  })

  // Google認証画面のURLを生成する
  const authorizeUrl = oauth2Client.generateAuthUrl({
    access_type: 'offline',
    scope: 'https://www.googleapis.com/auth/calendar'
  })

  res.status(200).json({ authorizeUrl })
}