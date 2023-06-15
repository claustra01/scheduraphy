import { Sequelize } from 'sequelize'

export const Psql = new Sequelize(process.env.DB_CONNECTION_STRINGS as string)