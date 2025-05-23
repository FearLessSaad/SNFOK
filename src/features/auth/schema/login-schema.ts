"use client"
 
import { z } from "zod"
 
const loginSchema = z.object({
  email: z.string().email({
    message: "Please enter a correct email address."
  }),
  token: z.string().min(64, {
    message: "Token must be 64 characters long."
  }).max(64, {
    message: "Token must be 64 characters long."
  })
})

export default loginSchema