import { UserService } from './user.service'

export interface Services {
  user: UserService
}

export const services: Services = {
  user: new UserService(),
}
