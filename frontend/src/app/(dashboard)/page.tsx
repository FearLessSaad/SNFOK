'use client';

import React, { useEffect, useState } from 'react';
import { useRouter } from 'next/navigation';
import ClusterOverview from '@/features/dashboard/components/cluster-overview';
import ThreatSummary from '@/features/dashboard/components/threat-summary';
import RecentActivity from '@/features/dashboard/components/recent-activity';
import LoadingSkeleton from '@/features/dashboard/components/loading-skeleton';
import { useQuery } from '@tanstack/react-query';
import { Get_All_Clusters_Request } from '@/features/dashboard/actions/get-all-clusters.action';
import { Response } from '@/types/response';
import { NO_CLUSTER_AVAILABLE } from '@/lib/response-cods';
import NoClusterAvailable from '@/features/dashboard/components/no-cluster-available';
import BrokenError from '@/components/broken-error';


const Dashboard = () => {
  const router = useRouter();

  const { data, isError, isPending, isSuccess } = useQuery<Response<any>>({
    queryFn: Get_All_Clusters_Request,
    queryKey: ["get-all-clusters"]
  });

  const [cluster, setCluster] = useState<boolean>(false)

  if(isSuccess){
    if(data.meta.code != NO_CLUSTER_AVAILABLE){
      setCluster(true)
    }
  }

  return (
    <div className="p-6 space-y-6 min-h-[100%]">
      <h1 className="text-3xl font-semibold">Security Dashboard</h1>

      {isPending && <LoadingSkeleton/>}

      <div className="flex flex-col items-center justify-center w-full min-h-[60vh]">
        {isSuccess && cluster ? (
          <div className="flex flex-col items-center justify-center w-full">
            <ClusterOverview />
            <ThreatSummary />
            <RecentActivity router={router} />
          </div>
        ) : 
          (!isPending && !isError) ? (
            <div className="flex flex-col items-center justify-center w-full">
              <NoClusterAvailable />
            </div>
          ) : isError && (
            <div className="flex flex-col items-center justify-center w-full">
              <BrokenError />
            </div>
          )
        }
      </div>

    </div>
  );
};

export default Dashboard;