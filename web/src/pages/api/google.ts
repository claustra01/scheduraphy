import type { NextApiRequest, NextApiResponse } from 'next'
import { google } from 'googleapis'
 
export default async function handler(req: NextApiRequest, res: NextApiResponse) {
 
  const oauth2Client = new google.auth.OAuth2({
    clientId: process.env.GOOGLE_OAUTH_CLIENT_ID,
    clientSecret: process.env.GOOGLE_OAUTH_CLIENT_SECRET,
    redirectUri: process.env.ORIGIN_URI + '/callback'
  })

  switch (req.method) {

    // OAuth用URLを生成して返す
    case 'GET':
      const authorizeUrl = oauth2Client.generateAuthUrl({
        access_type: 'offline',
        scope: 'https://www.googleapis.com/auth/calendar'
      })
      res.status(200).json({ authorizeUrl })
      break  
  
    // AuthCodeからTokenを生成して返す
    case 'POST':
      const { authorizationCode } = req.body
      const response = await oauth2Client.getToken(authorizationCode)
      res.status(200).json({ tokens: response.tokens })
      break

    // 405を返す
    default:
      res.status(405).end()
      break

  }

}