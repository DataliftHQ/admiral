import { createSlice, type PayloadAction } from '@reduxjs/toolkit'

import { type RootState } from '@/store'

export interface UserState {
  id: string
  email: string
  givenName: string
  familyName: string
}

const initialState: UserState = {
  id: '',
  email: '',
  givenName: '',
  familyName: '',
}

const userSlice = createSlice({
  name: 'user',
  initialState,
  reducers: {
    // eslint-disable-next-line @typescript-eslint/no-explicit-any
    setUser(state, action: PayloadAction<any>) {
      state.id = action.payload.user.id
      state.email = action.payload.user.email
      state.givenName = action.payload.user.givenName
      state.familyName = action.payload.user.familyName
    },
  },
})

export const { setUser } = userSlice.actions

export const selectUser = (state: RootState): UserState => state.user

export default userSlice.reducer
