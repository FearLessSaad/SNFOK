"use client"
 
import { z } from "zod"
 
const RegisterClusterSchema = z.object({
  master_ip: z.string().min(7, {
    message: "Minimum IP address should be 7 integers."
  }).max(15, {
    message: "Maximum IP address should be 15 integers"
  }),
  agent_port: z.number()
    .int({ message: "Port must be an integer." })
    .min(1024, { message: "Port must be at least 1024." })
    .max(65535, { message: "Port must be at most 65535." }),
  description: z.string().min(15, {
    message: "Description should be 15 characters long."
  }).max(500, {
    message:"Description cannot be 500 characters long."
  })
})

export default RegisterClusterSchema