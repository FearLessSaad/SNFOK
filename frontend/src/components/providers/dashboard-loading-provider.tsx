"use client"
import { useAppSelector } from "@/store/hooks";
import React from "react";
import Loading from "../loading";

function DashboardLoadingProvider({ children }: { children: React.ReactNode }) {
  const loading = useAppSelector((state) => state.authorization.loading);

  if (loading) {
    return <Loading />;
  }

  return <>{children}</>;
}

export default DashboardLoadingProvider;
