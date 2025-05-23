"use client";

import { useState } from "react";
import { useRouter } from "next/navigation";
import { Card, CardContent, CardHeader } from "@/components/ui/card";
import { Tabs, TabsList, TabsTrigger, TabsContent } from "@/components/ui/tabs";
import {
  Select,
  SelectTrigger,
  SelectValue,
  SelectContent,
  SelectItem,
} from "@/components/ui/select";

import { Button } from "@/components/ui/button";
import { Badge } from "@/components/ui/badge";

import {
  Table,
  TableHeader,
  TableRow,
  TableHead,
  TableBody,
  TableCell,
} from "@/components/ui/table";

import { format } from "date-fns";
import { Label } from "@/components/ui/label";

const mockThreatData = {
  detectionResults: [
    {
      id: "det-1",
      timestamp: "2025-04-05T13:45:10Z",
      eventType: "syscall",
      anomalyScore: 0.92,
      isAnomaly: true,
      confidence: 0.89,
      threatType: "privilege_escalation",
      threatSeverity: "high",
      podName: "nginx-pod-1",
      namespace: "default",
      nodeName: "worker-1",
    },
    {
      id: "det-2",
      timestamp: "2025-04-05T13:30:40Z",
      eventType: "network",
      anomalyScore: 0.88,
      isAnomaly: true,
      confidence: 0.85,
      threatType: "command_and_control",
      threatSeverity: "critical",
      podName: "redis-pod-2",
      namespace: "default",
      nodeName: "worker-2",
    },
    {
      id: "det-3",
      timestamp: "2025-04-05T12:15:20Z",
      eventType: "file",
      anomalyScore: 0.95,
      isAnomaly: true,
      confidence: 0.91,
      threatType: "credential_access",
      threatSeverity: "critical",
      podName: "etcd-backup-job-1",
      namespace: "kube-system",
      nodeName: "master-1",
    },
    {
      id: "det-4",
      timestamp: "2025-04-05T11:05:15Z",
      eventType: "process",
      anomalyScore: 0.78,
      isAnomaly: true,
      confidence: 0.75,
      threatType: "suspicious_shell_execution",
      threatSeverity: "medium",
      podName: "web-app-3",
      namespace: "app",
      nodeName: "worker-3",
    },
    {
      id: "det-5",
      timestamp: "2025-04-05T10:30:00Z",
      eventType: "syscall",
      anomalyScore: 0.94,
      isAnomaly: true,
      confidence: 0.9,
      threatType: "container_escape",
      threatSeverity: "critical",
      podName: "monitoring-agent-2",
      namespace: "monitoring",
      nodeName: "worker-1",
    },
  ],
  modelMetrics: [
    {
      modelId: "anomaly-detector-1",
      modelType: "anomaly_detector",
      accuracy: 0.92,
      precision: 0.89,
      recall: 0.85,
      f1Score: 0.87,
      falsePositiveRate: 0.08,
      falseNegativeRate: 0.15,
      lastUpdated: "2025-04-04T14:30:00Z",
      sampleCount: 10000,
      version: "1.0.0",
    },
    {
      modelId: "threat-classifier-1",
      modelType: "classifier",
      accuracy: 0.88,
      precision: 0.85,
      recall: 0.82,
      f1Score: 0.83,
      falsePositiveRate: 0.12,
      falseNegativeRate: 0.18,
      lastUpdated: "2025-04-03T10:15:00Z",
      sampleCount: 8000,
      version: "1.0.0",
    },
    {
      modelId: "sequence-detector-1",
      modelType: "sequence_detector",
      accuracy: 0.85,
      precision: 0.82,
      recall: 0.79,
      f1Score: 0.8,
      falsePositiveRate: 0.15,
      falseNegativeRate: 0.21,
      lastUpdated: "2025-04-02T09:45:00Z",
      sampleCount: 5000,
      version: "1.0.0",
    },
  ],
};

const getSeverityColor = (
  severity: string
): "default" | "destructive" | "outline" | "secondary" => {
  switch (severity) {
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

export default function ThreatsPage() {
  const router = useRouter();
  const [tab, setTab] = useState("detections");
  const [severityFilter, setSeverityFilter] = useState("all");
  const [typeFilter, setTypeFilter] = useState("all");

  const filteredDetections = mockThreatData.detectionResults.filter((d) => {
    const matchesSeverity =
      severityFilter === "all" || d.threatSeverity === severityFilter;
    const matchesType = typeFilter === "all" || d.threatType === typeFilter;
    return matchesSeverity && matchesType;
  });

  return (
    <div className="p-6 space-y-6">
      <h1 className="text-3xl font-bold">Threat Detection</h1>

      <Tabs value={tab} onValueChange={setTab} className="space-y-4">
        <TabsList>
          <TabsTrigger value="detections">Detection Results</TabsTrigger>
          <TabsTrigger value="models">Model Metrics</TabsTrigger>
        </TabsList>

        <TabsContent value="detections">
          <div className="flex gap-4 pb-4">
            <div>
              <Label className="mb-1 text-xs text-muted-foreground">
                Filter by severity
              </Label>
              <Select value={severityFilter} onValueChange={setSeverityFilter}>
                <SelectTrigger className="w-[150px]">
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
              <Label className="mb-1 text-xs text-muted-foreground">
                Filter by threat
              </Label>
              <Select value={typeFilter} onValueChange={setTypeFilter}>
                <SelectTrigger className="w-[250px]">
                  <SelectValue placeholder="Threat Type" />
                </SelectTrigger>
                <SelectContent>
                  <SelectItem value="all">All</SelectItem>
                  <SelectItem value="privilege_escalation">
                    Privilege Escalation
                  </SelectItem>
                  <SelectItem value="command_and_control">
                    Command & Control
                  </SelectItem>
                  <SelectItem value="credential_access">
                    Credential Access
                  </SelectItem>
                  <SelectItem value="suspicious_shell_execution">
                    Suspicious Shell Execution
                  </SelectItem>
                  <SelectItem value="container_escape">
                    Container Escape
                  </SelectItem>
                </SelectContent>
              </Select>
            </div>
          </div>

          <Card>
            <CardContent>
            <Table>
            <TableHeader>
              <TableRow>
                <TableHead>Severity</TableHead>
                <TableHead>Threat Type</TableHead>
                <TableHead>Confidence</TableHead>
                <TableHead>Event</TableHead>
                <TableHead>Pod</TableHead>
                <TableHead>Namespace</TableHead>
                <TableHead>Time</TableHead>
                <TableHead>Actions</TableHead>
              </TableRow>
            </TableHeader>
            <TableBody>
              {filteredDetections.map((d) => (
                <TableRow key={d.id}>
                  <TableCell>
                    <Badge variant={getSeverityColor(d.threatSeverity)}>
                      {d.threatSeverity}
                    </Badge>
                  </TableCell>
                  <TableCell>{d.threatType.replace(/_/g, " ")}</TableCell>
                  <TableCell>{(d.confidence * 100).toFixed(1)}%</TableCell>
                  <TableCell>{d.eventType}</TableCell>
                  <TableCell>{d.podName}</TableCell>
                  <TableCell>{d.namespace}</TableCell>
                  <TableCell>{format(new Date(d.timestamp), "PPpp")}</TableCell>
                  <TableCell>
                    <Button
                      size="sm"
                      onClick={() => router.push(`/threats/${d.id}`)}
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
        </TabsContent>

        <TabsContent value="models">
          <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-4">
            {mockThreatData.modelMetrics.map((model) => (
              <Card key={model.modelId}>
                <CardHeader>
                  <div className="flex justify-between">
                    <h4 className="font-semibold capitalize">
                      {model.modelType.replace(/_/g, " ")}
                    </h4>
                    <span className="text-sm text-muted-foreground">
                      v{model.version}
                    </span>
                  </div>
                  <p className="text-xs text-muted-foreground">
                    Last updated: {format(new Date(model.lastUpdated), "PPpp")}
                  </p>
                </CardHeader>
                <CardContent className="space-y-1">
                  <p className="text-sm text-muted-foreground">
                    Trained on {model.sampleCount.toLocaleString()} samples
                  </p>
                  {[
                    ["Accuracy", model.accuracy],
                    ["Precision", model.precision],
                    ["Recall", model.recall],
                    ["F1 Score", model.f1Score],
                    ["False Positive Rate", model.falsePositiveRate],
                    ["False Negative Rate", model.falseNegativeRate],
                  ].map(([label, val]) => (
                    <div
                      className="flex justify-between text-sm"
                      key={label as string}
                    >
                      <span>{label}:</span>
                      <span
                        className={
                          (label as string).includes("False")
                            ? "text-destructive font-semibold"
                            : "font-semibold"
                        }
                      >
                        {((val as number) * 100).toFixed(1)}%
                      </span>
                    </div>
                  ))}
                  <Button
                    className="w-full mt-4"
                    onClick={() => alert(`Updating ${model.modelType}`)}
                  >
                    Update Model
                  </Button>
                </CardContent>
              </Card>
            ))}
          </div>
        </TabsContent>
      </Tabs>
    </div>
  );
}
