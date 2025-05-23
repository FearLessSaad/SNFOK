import axios from "axios";
import { Login_Path } from "../routes";
import { z } from "zod";
import loginSchema from "@/features/auth/schema/login-schema";

export const Login_Request = (data: z.infer<typeof loginSchema>) => {
    return axios.post(Login_Path, data, {
        headers: {
            'Content-Type': 'application/json'
        },
        withCredentials: true
    });
}
