import { verifyToken } from '../utils/token'

const isAuthorized = async (resolve, parent, args, ctx, info) => {
  const auth = ctx.request.get('Authorization') || ''
  const token = auth.replace('Bearer ', '')

  try {
    verifyToken(token)
    return resolve()
  } catch (err) {
    throw new Error('Invalid Authorization.')
  }
}

const authMiddleware = {
  Mutation: {
  }
}


export default authMiddleware
export {
  isAuthorized
}