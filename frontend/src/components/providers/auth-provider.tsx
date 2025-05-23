"use client";

import { ExcludedAuthRoutes, Validate_Path_Access_Request } from "@/lib/auth";
import { useMutation } from "@tanstack/react-query";
import { usePathname, useRouter } from "next/navigation";
import React, { useEffect, useState } from "react";
import Loading from "@/components/loading";
import { toast } from "sonner";
import { setLoadingState } from "@/store/features/authorization";
import { useAppDispatch } from "@/store/hooks";

const isExcluded = (path: string): boolean => {
  return ExcludedAuthRoutes.includes(path);
};

function AuthProvider({ children }: { children: React.ReactNode }) {
  const router = useRouter();
  const path = usePathname();

  const [loading, setLoading] = useState<boolean>(false);
  const [excluded, setExcluded] = useState<boolean>(true);
  const dispatch = useAppDispatch()

  const { isPending, mutate } = useMutation({
    mutationFn: Validate_Path_Access_Request,
    onSuccess: () => {
      setLoading(false);
      dispatch(setLoadingState(false));
    },
    onError: () => {
      toast.error("Session Expired", {
        description: "Please login again to access your account.",
        duration: 5000,
      });
      router.replace("/login");
    },
  });

  // useEffect(() => {
  //   if (isExcluded(path)) {
  //     setExcluded(true);
  //     setLoading(false);
  //     dispatch(setLoadingState(false));
  //   } else {
  //     setExcluded(false);
  //     dispatch(setLoadingState(true));
  //     setLoading(true);
  //     mutate();
  //   }
  // }, [path, dispatch, mutate]);

  dispatch(setLoadingState(false));
  if ((isPending || loading) && excluded) {
    return <Loading />;
  }

  // Check if the path is excluded first
  if (!excluded) {
    return <>{children}</>;
  }

  return <>{children}</>;
}

export default AuthProvider;