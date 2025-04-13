"use client";

import React, { useState } from 'react';
import {
  Tabs,
  TabsList,
  TabsTrigger,
  TabsContent,
} from "@/components/ui/tabs";
import {
  Select,
  SelectTrigger,
  SelectValue,
  SelectContent,
  SelectItem,
} from "@/components/ui/select";
import { Button } from "@/components/ui/button";
import { Badge } from "@/components/ui/badge";
import { Table, TableHeader, TableRow, TableHead, TableBody, TableCell } from "@/components/ui/table";
import { useRouter } from "next/navigation";

// Mock data for initial development
const mockKubernetesData = {
  pods: [
    {
      name: "nginx-pod-1",
      namespace: "default", 
      status: "Running",
      node: "worker-1",
      ip: "10.0.0.1",
      containers: [
        {
          name: "nginx",
          image: "nginx:latest",
          ready: "true"
        }
      ],
      created: "2025-04-03T14:30:00Z",
      securityStatus: "critical"
    },
    {
      name: "redis-pod-2",
      namespace: "default",
      status: "Running", 
      node: "worker-2",
      ip: "10.0.0.2",
      containers: [
        {
          name: "redis",
          image: "redis:alpine",
          ready: "true"
        }
      ],
      created: "2025-04-04T10:15:00Z",
      securityStatus: "high"
    },
    {
      name: "etcd-backup-job-1",
      namespace: "kube-system",
      status: "Completed",
      node: "master-1", 
      ip: "10.0.0.3",
      containers: [
        {
          name: "etcd-backup",
          image: "etcd-backup:v1",
          ready: "false"
        }
      ],
      created: "2025-04-05T08:45:00Z",
      securityStatus: "high"
    },
    {
      name: "web-app-3",
      namespace: "app",
      status: "Running",
      node: "worker-3",
      ip: "10.0.0.4", 
      containers: [
        {
          name: "web-app",
          image: "web-app:v2",
          ready: "true"
        },
        {
          name: "sidecar",
          image: "sidecar:v1",
          ready: "true"
        }
      ],
      created: "2025-04-04T16:20:00Z",
      securityStatus: "medium"
    },
    {
      name: "monitoring-agent-2",
      namespace: "monitoring",
      status: "Running",
      node: "worker-1",
      ip: "10.0.0.5",
      containers: [
        {
          name: "prometheus",
          image: "prometheus:v2.35.0",
          ready: "true"
        }
      ],
      created: "2025-04-03T09:10:00Z",
      securityStatus: "critical"
    }
  ],
  nodes: [
    {
      name: "master-1",
      status: "Ready",
      role: "master",
      ip: "192.168.1.10",
      pods: 10,
      cpu_usage: "25%",
      mem_usage: "30%",
      disk_usage: "20%",
      created: "2025-03-01T00:00:00Z",
      securityStatus: "medium"
    },
    {
      name: "worker-1",
      status: "Ready",
      role: "worker",
      ip: "192.168.1.11",
      pods: 15,
      cpu_usage: "40%",
      mem_usage: "45%",
      disk_usage: "30%",
      created: "2025-03-01T00:00:00Z",
      securityStatus: "critical"
    },
    {
      name: "worker-2",
      status: "Ready",
      role: "worker",
      ip: "192.168.1.12",
      pods: 12,
      cpu_usage: "35%",
      mem_usage: "40%",
      disk_usage: "25%",
      created: "2025-03-01T00:00:00Z",
      securityStatus: "high"
    },
    {
      name: "worker-3",
      status: "Ready",
      role: "worker",
      ip: "192.168.1.13",
      pods: 8,
      cpu_usage: "20%",
      mem_usage: "25%",
      disk_usage: "15%",
      created: "2025-03-15T00:00:00Z",
      securityStatus: "low"
    }
  ],
  namespaces: [
    {
      name: "default",
      status: "Active",
      pods: 25,
      created: "2025-03-01T00:00:00Z",
      securityStatus: "high"
    },
    {
      name: "kube-system",
      status: "Active",
      pods: 15,
      created: "2025-03-01T00:00:00Z",
      securityStatus: "medium"
    },
    {
      name: "monitoring",
      status: "Active",
      pods: 5,
      created: "2025-03-02T00:00:00Z",
      securityStatus: "critical"
    },
    {
      name: "app",
      status: "Active",
      pods: 8,
      created: "2025-03-05T00:00:00Z",
      securityStatus: "medium"
    },
    {
      name: "dev",
      status: "Active",
      pods: 3,
      created: "2025-03-10T00:00:00Z",
      securityStatus: "low"
    }
  ]
};

const KubernetesPage = () => {
  const router = useRouter();
  const [tabValue, setTabValue] = useState("pods");
  const [namespaceFilter, setNamespaceFilter] = useState("all");
  const [securityFilter, setSecurityFilter] = useState("all");

  const filteredPods = mockKubernetesData.pods.filter(pod => {
    const matchesNamespace = namespaceFilter === "all" || pod.namespace === namespaceFilter;
    const matchesSecurity = securityFilter === "all" || pod.securityStatus === securityFilter;
    return matchesNamespace && matchesSecurity;
  });

  const getSecurityColor = (status: string): "default" | "destructive" | "secondary" | "outline" => {
    switch (status) {
      case "critical":
      case "high":
        return "destructive";
      case "medium":
        return "secondary";
      case "low":
        return "outline";
      default:
        return "default";
    }
  };

  return (
    <div className="p-6">
      <h1 className="text-3xl font-bold mb-6">Kubernetes Resources</h1>
      
      <Tabs value={tabValue} onValueChange={setTabValue} className="mb-6">
        <TabsList>
          <TabsTrigger value="pods">Pods</TabsTrigger>
          <TabsTrigger value="nodes">Nodes</TabsTrigger>
          <TabsTrigger value="namespaces">Namespaces</TabsTrigger>
        </TabsList>

        <TabsContent value="pods">
          <div className="flex gap-4 mb-6">
            <Select value={namespaceFilter} onValueChange={setNamespaceFilter}>
              <SelectTrigger className="w-[200px]">
                <SelectValue placeholder="Namespace" />
              </SelectTrigger>
              <SelectContent>
                <SelectItem value="all">All</SelectItem>
                {mockKubernetesData.namespaces.map(ns => (
                  <SelectItem key={ns.name} value={ns.name}>{ns.name}</SelectItem>
                ))}
              </SelectContent>
            </Select>

            <Select value={securityFilter} onValueChange={setSecurityFilter}>
              <SelectTrigger className="w-[200px]">
                <SelectValue placeholder="Security Status" />
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

          <Table>
            <TableHeader>
              <TableRow>
                <TableHead>Name</TableHead>
                <TableHead>Namespace</TableHead>
                <TableHead>Status</TableHead>
                <TableHead>Node</TableHead>
                <TableHead>IP</TableHead>
                <TableHead>Containers</TableHead>
                <TableHead>Security</TableHead>
                <TableHead>Actions</TableHead>
              </TableRow>
            </TableHeader>
            <TableBody>
              {filteredPods.map((pod) => (
                <TableRow key={pod.name}>
                  <TableCell>{pod.name}</TableCell>
                  <TableCell>{pod.namespace}</TableCell>
                  <TableCell>{pod.status}</TableCell>
                  <TableCell>{pod.node}</TableCell>
                  <TableCell>{pod.ip}</TableCell>
                  <TableCell>{pod.containers.length}</TableCell>
                  <TableCell>
                    <Badge variant={getSecurityColor(pod.securityStatus)}>
                      {pod.securityStatus}
                    </Badge>
                  </TableCell>
                  <TableCell>
                    <div className="flex gap-2">
                      <Button 
                        variant="outline"
                        size="sm"
                        onClick={() => router.push(`/dashboard/kubernetes/pods/${pod.namespace}/${pod.name}`)}
                      >
                        Details
                      </Button>
                      <Button 
                        variant="outline"
                        size="sm"
                        onClick={() => alert(`Isolating pod ${pod.name}...`)}
                      >
                        Isolate
                      </Button>
                    </div>
                  </TableCell>
                </TableRow>
              ))}
            </TableBody>
          </Table>
        </TabsContent>

        <TabsContent value="nodes">
          <Table>
            <TableHeader>
              <TableRow>
                <TableHead>Name</TableHead>
                <TableHead>Role</TableHead>
                <TableHead>Status</TableHead>
                <TableHead>IP</TableHead>
                <TableHead>Pods</TableHead>
                <TableHead>CPU</TableHead>
                <TableHead>Memory</TableHead>
                <TableHead>Security</TableHead>
                <TableHead>Actions</TableHead>
              </TableRow>
            </TableHeader>
            <TableBody>
              {mockKubernetesData.nodes.map((node) => (
                <TableRow key={node.name}>
                  <TableCell>{node.name}</TableCell>
                  <TableCell>{node.role}</TableCell>
                  <TableCell>{node.status}</TableCell>
                  <TableCell>{node.ip}</TableCell>
                  <TableCell>{node.pods}</TableCell>
                  <TableCell>{node.cpu_usage}</TableCell>
                  <TableCell>{node.mem_usage}</TableCell>
                  <TableCell>
                    <Badge variant={getSecurityColor(node.securityStatus)}>
                      {node.securityStatus}
                    </Badge>
                  </TableCell>
                  <TableCell>
                    <Button 
                      variant="outline"
                      size="sm"
                      onClick={() => router.push(`/dashboard/kubernetes/nodes/${node.name}`)}
                    >
                      Details
                    </Button>
                  </TableCell>
                </TableRow>
              ))}
            </TableBody>
          </Table>
        </TabsContent>

        <TabsContent value="namespaces">
          <Table>
            <TableHeader>
              <TableRow>
                <TableHead>Name</TableHead>
                <TableHead>Status</TableHead>
                <TableHead>Pods</TableHead>
                <TableHead>Security</TableHead>
                <TableHead>Actions</TableHead>
              </TableRow>
            </TableHeader>
            <TableBody>
              {mockKubernetesData.namespaces.map((namespace) => (
                <TableRow key={namespace.name}>
                  <TableCell>{namespace.name}</TableCell>
                  <TableCell>{namespace.status}</TableCell>
                  <TableCell>{namespace.pods}</TableCell>
                  <TableCell>
                    <Badge variant={getSecurityColor(namespace.securityStatus)}>
                      {namespace.securityStatus}
                    </Badge>
                  </TableCell>
                  <TableCell>
                    <Button 
                      variant="outline"
                      size="sm"
                      onClick={() => router.push(`/dashboard/kubernetes/namespaces/${namespace.name}`)}
                    >
                      Details
                    </Button>
                  </TableCell>
                </TableRow>
              ))}
            </TableBody>
          </Table>
        </TabsContent>
      </Tabs>
    </div>
  );
};

export default KubernetesPage;