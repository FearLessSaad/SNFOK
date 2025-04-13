'use client';

import React from 'react';
import { useRouter } from 'next/navigation';
import {
  Card, CardContent, CardHeader,
} from '@/components/ui/card';
import { Button } from '@/components/ui/button';
import { Badge } from '@/components/ui/badge';
import {
  AlertCircle, CheckCircle, Info, ShieldAlert
} from 'lucide-react';
import { cn } from '@/lib/utils';

// Mock data (same as your original)
const mockDashboardData = {
  clusterSummary: {
    totalNodes: 5,
    totalPods: 42,
    totalContainers: 78,
    activeAlerts: 3,
    criticalAlerts: 1,
    highAlerts: 1,
    mediumAlerts: 1,
    lowAlerts: 0,
    systemHealthy: true,
    collectorStatus: "running",
    detectionEngineStatus: "running",
    modelAccuracy: 0.92,
    falsePositiveRate: 0.08
  },
  threatSummary: {
    totalThreats: 12,
    byType: [
      { type: "privilege_escalation", count: 3 },
      { type: "suspicious_network", count: 5 },
      { type: "suspicious_process", count: 2 },
      { type: "data_exfiltration", count: 1 },
      { type: "crypto_mining", count: 1 }
    ],
    byNamespace: [
      { namespace: "default", count: 5 },
      { namespace: "kube-system", count: 3 },
      { namespace: "monitoring", count: 2 },
      { namespace: "app", count: 2 }
    ]
  },
  recentActivity: [
    {
      id: "act-1",
      timestamp: "2025-04-05T13:45:12Z",
      type: "alert",
      severity: "critical",
      title: "Privilege Escalation Attempt",
      description: "Detected attempt to escalate privileges in pod nginx-pod-1"
    },
    {
      id: "act-2",
      timestamp: "2025-04-05T13:30:45Z",
      type: "detection",
      severity: "high",
      title: "Suspicious Network Activity",
      description: "Unusual outbound connection from pod redis-pod-2"
    },
    {
      id: "act-3",
      timestamp: "2025-04-05T13:15:22Z",
      type: "system",
      severity: "info",
      title: "Model Update",
      description: "AI detection model updated successfully"
    }
  ]
};

const getSeverityStyle = (severity: string) => {
  switch (severity) {
    case 'critical': return 'bg-red-600 text-white';
    case 'high': return 'bg-orange-600 text-white';
    case 'medium': return 'bg-yellow-400 text-black';
    case 'low': return 'bg-blue-400 text-white';
    case 'info': return 'bg-gray-200 text-black';
    default: return '';
  }
};

const getSeverityIcon = (severity: string) => {
  switch (severity) {
    case 'critical': return <AlertCircle className="text-red-600 w-5 h-5" />;
    case 'high':
    case 'medium': return <ShieldAlert className="text-orange-500 w-5 h-5" />;
    case 'low':
    case 'info':
    default: return <Info className="text-blue-500 w-5 h-5" />;
  }
};

const Dashboard = () => {
  const router = useRouter();
  const { clusterSummary, threatSummary, recentActivity } = mockDashboardData;

  return (
    <div className="p-6 space-y-6">
      <h1 className="text-3xl font-semibold">Security Dashboard</h1>

      {/* Cluster Overview */}
      <div className="grid grid-cols-1 md:grid-cols-4 gap-4">
        <Card className={cn("text-white", clusterSummary.systemHealthy ? "bg-green-600" : "bg-red-600")}>
          <CardHeader>
            <h2 className="text-lg font-medium">System Status</h2>
          </CardHeader>
          <CardContent className="flex flex-col items-center gap-2">
            {clusterSummary.systemHealthy
              ? <CheckCircle className="w-12 h-12" />
              : <AlertCircle className="w-12 h-12" />
            }
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

      {/* Threat Summary */}
      <h2 className="text-2xl font-semibold">Threat Summary</h2>
      <div className="grid grid-cols-1 md:grid-cols-2 gap-4">
        <Card>
          <CardHeader>
            <h3 className="text-lg font-medium">Threats by Type</h3>
          </CardHeader>
          <CardContent className="space-y-2">
            {threatSummary.byType.map((t) => (
              <div key={t.type} className="flex justify-between items-center">
                <span className="capitalize">{t.type.replace(/_/g, ' ')}</span>
                <Badge variant="outline">{t.count}</Badge>
              </div>
            ))}
          </CardContent>
        </Card>
        <Card>
          <CardHeader>
            <h3 className="text-lg font-medium">Threats by Namespace</h3>
          </CardHeader>
          <CardContent className="space-y-2">
            {threatSummary.byNamespace.map((t) => (
              <div key={t.namespace} className="flex justify-between items-center">
                <span>{t.namespace}</span>
                <Badge variant="outline">{t.count}</Badge>
              </div>
            ))}
          </CardContent>
        </Card>
      </div>

      {/* Recent Activity */}
      <div className="flex justify-between items-center">
        <h2 className="text-2xl font-semibold">Recent Activity</h2>
        <Button variant="outline" onClick={() => router.push('/alerts')}>
          View All Alerts
        </Button>
      </div>
      <div className="space-y-4">
        {recentActivity.map((item) => (
          <Card key={item.id}>
            <CardContent className="flex items-center space-x-4">
              <div>{getSeverityIcon(item.severity)}</div>
              <div className="flex-grow">
                <div className="flex items-center space-x-2 mb-1">
                  <span className="font-medium">{item.title}</span>
                  <Badge className={getSeverityStyle(item.severity)}>{item.severity}</Badge>
                  <span className="text-sm text-muted-foreground">{new Date(item.timestamp).toLocaleString()}</span>
                </div>
                <p className="text-sm">{item.description}</p>
              </div>
              <Button variant="link" onClick={() => router.push(`/dashboard/alerts/${item.id}`)}>Details</Button>
            </CardContent>
          </Card>
        ))}
      </div>
    </div>
  );
};

export default Dashboard;