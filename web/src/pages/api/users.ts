import type { NextApiRequest, NextApiResponse } from 'next'
import { User } from './db/models';
import { NOW, UUIDV4 } from 'sequelize';
import { type } from 'os';

export default async function handler(req: NextApiRequest, res: NextApiResponse) {

  switch (req.method) {

    // lineIdからUserを取得
    case 'GET':
      try {
        const lineId = req.query.lineId
        const user = await User.findOne({
          where: {
            line_id: lineId
          }
        })
        if (user != null) {
          res.status(200).json(user.toJSON())
        } else {
          res.status(200).json({})
        }
      } catch(error) {
        console.error(error)   
        res.status(503).end()     
      }
      break

    // 新規ユーザー作成
    case 'POST':
      try {
        const { lineId, refreshToken } = req.body
        const user = await User.create({
          line_id: lineId,
          refresh_token: refreshToken,
        })
        if (user != null) {
          res.status(200).json(user.toJSON())
        }
      } catch(error) {
        console.error(error)   
        res.status(503).end()     
      }
      break
    
    // refreshToken更新
    case 'PATCH':
      try {
        const { lineId, refreshToken } = req.body
        const user = await User.findOne({
          where: {
            line_id: lineId
          }
        })
        if (user != null) {
          user.set({
            refresh_token: refreshToken,
            updated_at: new Date()
          })
          await user.save()
          res.status(200).json(user.toJSON())
        } else {
          res.status(503).end()     
        }
      } catch(error) {
        console.error(error)   
        res.status(503).end()     
      }

    // メソッドエラー
    default:
      res.status(405).end()
      break

  }
}