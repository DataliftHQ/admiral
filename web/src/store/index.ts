import { type Action, configureStore, type ThunkAction } from '@reduxjs/toolkit'
import { useDispatch as useAppDispatch, useSelector as useAppSelector, type TypedUseSelectorHook } from 'react-redux'
import logger from 'redux-logger'
import { persistStore } from 'redux-persist'

import { listenerMiddleware } from './middleware'

import rootReducer from './reducer'

const store = configureStore({
  reducer: rootReducer,
  middleware: (getDefaultMiddleware) =>
    getDefaultMiddleware({ serializableCheck: false, immutableCheck: false }).concat(
      logger,
      listenerMiddleware.middleware
    ),
  devTools: true,
})

const persistor = persistStore(store)

export type RootState = ReturnType<typeof rootReducer>
export type AppDispatch = typeof store.dispatch
export type AppThunk<ReturnType = void> = ThunkAction<ReturnType, RootState, unknown, Action<string>>

const { dispatch } = store
const useDispatch = (): AppDispatch => useAppDispatch<AppDispatch>()
const useSelector: TypedUseSelectorHook<RootState> = useAppSelector

export { store, persistor, dispatch, useSelector, useDispatch }
