import client from './client'

interface User {
  id: string
  email: string
  givenName: string
  familyName: string
}

export class UserService {
  public async get(): Promise<User> {
    return await client.get('/api/v1/users/me').then((res) => {
      return res.data as User
    })
  }
}
