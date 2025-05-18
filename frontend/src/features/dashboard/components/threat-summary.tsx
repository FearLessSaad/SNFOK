'use client';

import React from 'react';
import { Card, CardHeader, CardContent } from '@/components/ui/card';
import { Badge } from '@/components/ui/badge';

const threatSummary = {
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
};

const ThreatSummary = () => (
  <>
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
  </>
);

export default ThreatSummary;
