import { DataTypes } from 'sequelize'
import { Psql } from "./db";

export const User = Psql.define('User', {
  id: {
    type: DataTypes.UUIDV4,
    primaryKey: true
  },
  lineId: {
    type: DataTypes.STRING
  },
  refreshToken: {
    type: DataTypes.STRING
  },
  createdAt: {
    type: DataTypes.TIME
  },
  updatedAt: {
    type: DataTypes.TIME
  }
})