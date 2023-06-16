import type { NextApiRequest, NextApiResponse } from 'next'
import { User } from './db/models';
import { NOW, UUIDV4 } from 'sequelize';

export default async function handler(req: NextApiRequest, res: NextApiResponse) {

  switch (req.method) {

    case 'POST':
      try {
        const { lineId, refreshToken } = req.body
        const user = await User.create({
          line_id: lineId,
          refresh_token: refreshToken,
        });
        res.status(200).json(user.toJSON())
      } catch (error) {
        console.error(error)   
        res.status(503).end()     
      }
      break

    default:
      res.status(405).end()
      break

  }
}