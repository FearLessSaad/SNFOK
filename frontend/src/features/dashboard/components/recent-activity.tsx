'use client';

import React from 'react';
import { Card, CardContent } from '@/components/ui/card';
import { Button } from '@/components/ui/button';
import { Badge } from '@/components/ui/badge';
import { AlertCircle, Info, ShieldAlert } from 'lucide-react';
import { cn } from '@/lib/utils';

const recentActivity = [
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
];

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

const RecentActivity = ({ router }: { router: any }) => (
  <>
    <div className="flex justify-between items-center">
      <h2 className="text-2xl font-semibold">Recent Activity</h2>
      <Button variant="outline" onClick={() => router.push('/alerts')}>View All Alerts</Button>
    </div>
    <div className="space-y-4">
      {recentActivity.map((item) => (
        <Card key={item.id}>
          <CardContent className="flex items-center space-x-4">
            <div>{getSeverityIcon(item.severity)}</div>
            <div className="flex-grow">
              <div className="flex items-center space-x-2 mb-1">
                <span className="font-medium">{item.title}</span>
                <Badge className={cn(getSeverityStyle(item.severity))}>{item.severity}</Badge>
                <span className="text-sm text-muted-foreground">
                  {new Date(item.timestamp).toLocaleString()}
                </span>
              </div>
              <p className="text-sm">{item.description}</p>
            </div>
            <Button variant="link" onClick={() => router.push(`/dashboard/alerts/${item.id}`)}>Details</Button>
          </CardContent>
        </Card>
      ))}
    </div>
  </>
);

export default RecentActivity;
