import axios from "axios";
import { Get_All_Clusters_Path } from "../routes";
import { Response } from "@/types/response";
import { ClusterData_T } from "@/types/clusters/cluster.type";

export const Get_All_Clusters_Request = async (): Promise<Response<any>> => {
    const { data }:{data:Response<ClusterData_T>} = await axios.get(Get_All_Clusters_Path, {
        withCredentials: true
    });
    return data;
}
