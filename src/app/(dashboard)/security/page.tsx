"use client";
import Reacts, { useState } from "react";
import { Card, CardContent } from "@/components/ui/card";
import {
  Table,
  TableHeader,
  TableRow,
  TableHead,
  TableBody,
  TableCell,
} from "@/components/ui/table";
import { Label } from "@/components/ui/label";
import {
  Select,
  SelectTrigger,
  SelectValue,
  SelectContent,
  SelectItem,
} from "@/components/ui/select";

import { useRouter } from "next/navigation";
import { Button } from "@/components/ui/button";
import { randomUUID } from "crypto";
import { PlusCircleIcon, SearchCheckIcon, SearchSlashIcon } from "lucide-react";
import NewMonitor from "./_components/new-monitor";

const page = () => {
  const [namespaceFilter, setNamespaceFilter] = useState("all");
  const [PodFilter, setPodFilter] = useState("all");

  const router = useRouter();
  return (
    <>
      <div className="p-6 space-y-6">
        <h1 className="text-3xl font-semibold">Security Configuration</h1>
      

      <div className="flex items-end justify-between flex-row">
      <div className="flex gap-4 mb-4">
        <div>
          <Label className="mb-1 text-xs text-muted-foreground">
            Filter by namespace
          </Label>
          <Select onValueChange={setNamespaceFilter} defaultValue="all">
            <SelectTrigger className="w-[180px]">
              <SelectValue placeholder="Namespace" />
            </SelectTrigger>
            <SelectContent>
              <SelectItem value="default">default</SelectItem>
              <SelectItem value="kube-system">kube-systm</SelectItem>
              <SelectItem value="public">public</SelectItem>
            </SelectContent>
          </Select>
        </div>

        <div>
          <Label className="mb-1 text-xs text-muted-foreground">
            Filter by pods
          </Label>
          <Select onValueChange={setPodFilter} defaultValue="all">
            <SelectTrigger className="w-[180px]">
              <SelectValue placeholder="Status" />
            </SelectTrigger>
            <SelectContent>
              <SelectItem value="new">New</SelectItem>
              <SelectItem value="acknowledged">Acknowledged</SelectItem>
              <SelectItem value="resolved">Resolved</SelectItem>
              <SelectItem value="false_positive">False Positive</SelectItem>
            </SelectContent>
          </Select>
        </div>
      </div>
      <NewMonitor/>
      </div>
      <Card>
        <CardContent>
          <Table>
            <TableHeader>
              <TableRow>
                <TableHead>#</TableHead>
                <TableHead>Namespace</TableHead>
                <TableHead>Pod</TableHead>
                <TableHead>Ip Address</TableHead>
                <TableHead>Image</TableHead>
                <TableHead>Threats</TableHead>
                <TableHead>Actions</TableHead>
              </TableRow>
            </TableHeader>
            <TableBody>
              <TableRow>
                <TableCell>1</TableCell>
                <TableCell>hashx-web</TableCell>
                <TableCell>frontend-1</TableCell>
                <TableCell>10.12.13.60</TableCell>
                <TableCell>nignx:latest</TableCell>
                <TableCell>50</TableCell>
                <TableCell>
                  <Button
                    size="sm"
                    onClick={() => router.push(`/security/djashdsadjksahdjkh`)}
                  >
                    Details
                  </Button>
                </TableCell>
              </TableRow>
            </TableBody>
          </Table>
        </CardContent>
      </Card>
      </div>
    </>
  );
};

export default page;
