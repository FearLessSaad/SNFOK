'use client';

import React from 'react';
import { Card, CardContent, CardHeader } from '@/components/ui/card';
import { CheckCircle, AlertCircle } from 'lucide-react';
import { cn } from '@/lib/utils';

const clusterSummary = {
  totalNodes: 5,
  totalPods: 42,
  totalContainers: 78,
  activeAlerts: 3,
  criticalAlerts: 1,
  highAlerts: 1,
  mediumAlerts: 1,
  lowAlerts: 0,
  systemHealthy: true,
  collectorStatus: 'running',
  detectionEngineStatus: 'running',
  modelAccuracy: 0.92,
  falsePositiveRate: 0.08,
};

const ClusterOverview = () => (
  <div className="grid grid-cols-1 md:grid-cols-4 gap-4">
    <Card className={cn("text-white", clusterSummary.systemHealthy ? "bg-green-600" : "bg-red-600")}>
      <CardHeader>
        <h2 className="text-lg font-medium">System Status</h2>
      </CardHeader>
      <CardContent className="flex flex-col items-center gap-2">
        {clusterSummary.systemHealthy
          ? <CheckCircle className="w-12 h-12" />
          : <AlertCircle className="w-12 h-12" />}
        <p>{clusterSummary.systemHealthy ? 'Healthy' : 'Issues Detected'}</p>
      </CardContent>
    </Card>

    <Card>
      <CardHeader>
        <h2 className="text-lg font-medium">Cluster Resources</h2>
      </CardHeader>
      <CardContent>
        <div>Nodes: <strong>{clusterSummary.totalNodes}</strong></div>
        <div>Pods: <strong>{clusterSummary.totalPods}</strong></div>
        <div>Containers: <strong>{clusterSummary.totalContainers}</strong></div>
      </CardContent>
    </Card>

    <Card>
      <CardHeader>
        <h2 className="text-lg font-medium">Active Alerts</h2>
      </CardHeader>
      <CardContent className="space-y-1">
        <div className="text-red-600">Critical: <strong>{clusterSummary.criticalAlerts}</strong></div>
        <div className="text-orange-500">High: <strong>{clusterSummary.highAlerts}</strong></div>
        <div className="text-yellow-600">Medium: <strong>{clusterSummary.mediumAlerts}</strong></div>
        <div className="text-blue-500">Low: <strong>{clusterSummary.lowAlerts}</strong></div>
      </CardContent>
    </Card>

    <Card>
      <CardHeader>
        <h2 className="text-lg font-medium">AI Model</h2>
      </CardHeader>
      <CardContent className="space-y-1">
        <div>Accuracy: <strong>{(clusterSummary.modelAccuracy * 100).toFixed(1)}%</strong></div>
        <div>False Positives: <strong>{(clusterSummary.falsePositiveRate * 100).toFixed(1)}%</strong></div>
        <div>Status: <strong>{clusterSummary.detectionEngineStatus}</strong></div>
      </CardContent>
    </Card>
  </div>
);

export default ClusterOverview;
