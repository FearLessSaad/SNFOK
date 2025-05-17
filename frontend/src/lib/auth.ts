import axios from "axios";
import { Validate_Access_Request_Path } from "./paths";

export const Validate_Path_Access_Request = ()=>{
    return axios.get(Validate_Access_Request_Path, {
        withCredentials: true,
    })
}

export const ExcludedAuthRoutes: string[] = [
    "/login",
];