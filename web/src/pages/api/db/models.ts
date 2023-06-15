import { DataTypes } from 'sequelize'
import { Psql } from "./db";

export const User = Psql.define('user', {
  id: {
    type: DataTypes.UUIDV4,
    defaultValue: DataTypes.UUIDV4,
    primaryKey: true
  },
  line_id: {
    type: DataTypes.STRING
  },
  refresh_token: {
    type: DataTypes.STRING
  },
  created_at: {
    type: DataTypes.TIME,
    defaultValue: DataTypes.NOW
  },
  updated_at: {
    type: DataTypes.TIME,
    defaultValue: DataTypes.NOW
  }
}, {
  timestamps: false
})