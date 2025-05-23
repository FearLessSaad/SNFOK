import { createSlice } from '@reduxjs/toolkit'
import type { PayloadAction } from '@reduxjs/toolkit'

interface Authorization {
  status: boolean;
  loading: boolean;
}
const initialState: Authorization = {
  status: false,
  loading: true,
}

export const authorizationSlice = createSlice({
  name: 'authorization',
  initialState,
  reducers: {
    setAuthorizationState: (state, action: PayloadAction<boolean>) => {
      state.status = action.payload
    },
    setLoadingState: (state, action: PayloadAction<boolean>) => {
      state.loading = action.payload
    },
  },
})

export const { setAuthorizationState, setLoadingState } = authorizationSlice.actions
export default authorizationSlice.reducer