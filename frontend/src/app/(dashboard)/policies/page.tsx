"use client";

import React, { useState } from 'react';
import { useRouter } from 'next/navigation';
import {
  Card,
  CardContent,
  CardHeader,
} from "@/components/ui/card";
import {
  Table,
  TableBody,
  TableCell,
  TableHead,
  TableHeader,
  TableRow,
} from "@/components/ui/table";
import { Switch } from "@/components/ui/switch";
import { Label } from "@/components/ui/label";
import { Badge } from "@/components/ui/badge";
import { Button } from "@/components/ui/button";
import { Select, SelectContent, SelectItem, SelectTrigger, SelectValue } from "@/components/ui/select";
import { format } from 'date-fns';

const mockPolicies = [
  {
    id: "policy-privilege-escalation",
    name: "Privilege Escalation Detection",
    description: "Detects attempts to escalate privileges within containers",
    enabled: true,
    severity: "high",
    createdAt: "2025-03-15T10:30:00Z",
    updatedAt: "2025-04-01T14:45:00Z",
    rules: [
      {
        id: "rule-syscall-ptrace",
        type: "pattern",
        description: "Detect ptrace syscalls",
        conditions: {
          event_type: "syscall",
          syscall_name: "ptrace"
        }
      },
      {
        id: "rule-syscall-capabilities",
        type: "pattern",
        description: "Detect capability changes",
        conditions: {
          event_type: "syscall",
          syscall_name: "capset"
        }
      }
    ],
    actions: [
      {
        id: "action-alert-high",
        type: "alert",
        description: "Generate high severity alert",
        parameters: {
          severity: "high",
          title: "Privilege Escalation Attempt"
        },
        automatic: true
      },
      {
        id: "action-isolate-pod",
        type: "isolate",
        description: "Isolate the affected pod",
        parameters: {
          isolation_type: "network"
        },
        automatic: false
      }
    ]
  },
  {
    id: "policy-suspicious-network",
    name: "Suspicious Network Activity",
    description: "Detects suspicious network connections",
    enabled: true,
    severity: "medium",
    createdAt: "2025-03-20T09:15:00Z",
    updatedAt: "2025-04-02T11:30:00Z",
    rules: [
      {
        id: "rule-suspicious-ports",
        type: "pattern",
        description: "Detect connections to suspicious ports",
        conditions: {
          event_type: "network",
          destination_port: [4444, 8080, 6379]
        }
      },
      {
        id: "rule-high-volume",
        type: "threshold",
        description: "Detect high volume data transfers",
        conditions: {
          event_type: "network",
          bytes_sent: {
            threshold: 10000000,
            operator: ">"
          }
        }
      }
    ],
    actions: [
      {
        id: "action-alert-medium",
        type: "alert",
        description: "Generate medium severity alert",
        parameters: {
          severity: "medium",
          title: "Suspicious Network Activity"
        },
        automatic: true
      }
    ]
  },
  {
    id: "policy-sensitive-file-access",
    name: "Sensitive File Access",
    description: "Detects access to sensitive files",
    enabled: true,
    severity: "critical",
    createdAt: "2025-03-25T13:45:00Z",
    updatedAt: "2025-04-03T16:20:00Z",
    rules: [
      {
        id: "rule-credential-files",
        type: "pattern",
        description: "Detect access to credential files",
        conditions: {
          event_type: "file",
          path: ["/etc/passwd", "/etc/shadow", "/etc/kubernetes/pki"],
          operation: ["write", "append"]
        }
      }
    ],
    actions: [
      {
        id: "action-alert-critical",
        type: "alert",
        description: "Generate critical severity alert",
        parameters: {
          severity: "critical",
          title: "Sensitive File Access"
        },
        automatic: true
      },
      {
        id: "action-terminate-pod",
        type: "block",
        description: "Terminate the affected pod",
        parameters: {
          action: "terminate"
        },
        automatic: false
      }
    ]
  }
];

const getSeverityColor = (severity: string): "default" | "destructive" | "outline" | "secondary" => {
  switch (severity) {
    case 'critical':
    case 'high':
      return 'destructive';
    case 'medium':
      return 'secondary';
    case 'low':
      return 'outline';
    default:
      return 'default';
  }
};

const PoliciesPage = () => {
  const router = useRouter();
  const [severityFilter, setSeverityFilter] = useState('all');

  const filteredPolicies = mockPolicies.filter(policy => {
    return severityFilter === 'all' || policy.severity === severityFilter;
  });

  const handleTogglePolicy = (id: string, currentState: boolean) => {
    alert(`${currentState ? 'Disabling' : 'Enabling'} policy ${id}...`);
  };

  return (
    <div className="p-6 space-y-6">
      <div className="flex justify-between items-center">
        <h1 className="text-2xl font-semibold">Security Policies</h1>
        <Button onClick={() => router.push('/dashboard/policies/new')}>
          Create New Policy
        </Button>
      </div>

      <div>
        <Label className='mb-1 text-xs text-muted-foreground'>Filter by severity</Label>
        <Select value={severityFilter} onValueChange={setSeverityFilter}>
          <SelectTrigger className="w-[180px]">
            <SelectValue placeholder="Select severity" />
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

      <div className="grid gap-6">
        {filteredPolicies.map(policy => (
          <Card key={policy.id}>
            <CardHeader className="flex flex-col sm:flex-row sm:items-center sm:justify-between">
              <div>
                <h2 className="text-lg font-medium">{policy.name}</h2>
                <div className="text-sm text-muted-foreground flex gap-2 mt-1">
                  <Badge variant={getSeverityColor(policy.severity)}>{policy.severity}</Badge>
                  <span>Last updated: {format(new Date(policy.updatedAt), 'PPpp')}</span>
                </div>
              </div>
              <div className="flex items-center gap-2 mt-2 sm:mt-0">
                <Label>{policy.enabled ? "Enabled" : "Disabled"}</Label>
                <Switch
                  checked={policy.enabled}
                  onCheckedChange={() => handleTogglePolicy(policy.id, policy.enabled)}
                />
              </div>
            </CardHeader>

            <CardContent className="space-y-4">
              <p>{policy.description}</p>

              <div>
                <h3 className="text-sm font-semibold mb-2">Rules ({policy.rules.length})</h3>
                <Table>
                  <TableHeader>
                    <TableRow>
                      <TableHead>Type</TableHead>
                      <TableHead>Description</TableHead>
                    </TableRow>
                  </TableHeader>
                  <TableBody>
                    {policy.rules.map(rule => (
                      <TableRow key={rule.id}>
                        <TableCell>{rule.type}</TableCell>
                        <TableCell>{rule.description}</TableCell>
                      </TableRow>
                    ))}
                  </TableBody>
                </Table>
              </div>

              <div>
                <h3 className="text-sm font-semibold mb-2">Actions ({policy.actions.length})</h3>
                <Table>
                  <TableHeader>
                    <TableRow>
                      <TableHead>Type</TableHead>
                      <TableHead>Description</TableHead>
                      <TableHead>Automatic</TableHead>
                    </TableRow>
                  </TableHeader>
                  <TableBody>
                    {policy.actions.map(action => (
                      <TableRow key={action.id}>
                        <TableCell>{action.type}</TableCell>
                        <TableCell>{action.description}</TableCell>
                        <TableCell>
                          <Badge variant={action.automatic ? "default" : "secondary"}>
                            {action.automatic ? "Yes" : "No"}
                          </Badge>
                        </TableCell>
                      </TableRow>
                    ))}
                  </TableBody>
                </Table>
              </div>

              <div className="flex justify-end gap-2">
                <Button variant="outline" onClick={() => router.push(`/dashboard/policies/${policy.id}`)}>Edit</Button>
                <Button variant="destructive" onClick={() => alert(`Deleting policy ${policy.id}...`)}>Delete</Button>
              </div>
            </CardContent>
          </Card>
        ))}
      </div>
    </div>
  );
};

export default PoliciesPage;