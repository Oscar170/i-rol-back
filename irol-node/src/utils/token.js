import { verify } from 'jsonwebtoken'

const jwtSecret = "myfancysecret"

const verifyToken = (token) => verify(token, jwtSecret)

export {
  verifyToken
}