"use client";
import { ScansTable } from "@/components/scans/table/table";
import { useState, ChangeEvent, FormEvent } from "react";
import { Button } from "@/components/ui/button";
import { FC } from "react";
import { FaPlus } from "react-icons/fa6";
import { Textarea } from "@/components/ui/textarea";
import { TfiReload } from "react-icons/tfi";
import {
  Select,
  SelectContent,
  SelectItem,
  SelectTrigger,
  SelectValue,
} from "@/components/ui/select";

import {
  Dialog,
  DialogContent,
  DialogDescription,
  DialogFooter,
  DialogHeader,
  DialogTitle,
  DialogTrigger,
} from "@/components/ui/dialog";
import { Input } from "@/components/ui/input";
import { Label } from "@/components/ui/label";

const Page: FC = () => {
  const [formData, setFormData] = useState({
    title: "",
    description: "",
    ip: "",
    ig: "",
  });

  const handleRefreshedIp = () => {
    console.log("Clicked on refresh button");
  };

  // Handle input change
  const handleChange = (e: ChangeEvent<HTMLInputElement | HTMLTextAreaElement>) => {
    setFormData({
      ...formData,
      [e.target.name]: e.target.value,
    });
  };

  // Handle Select change
  const handleSelectChange = (field: string, value: string) => {
    setFormData({
      ...formData,
      [field]: value,
    });
  };

  const handleSubmit = (e: FormEvent<HTMLButtonElement>) => {
    e.preventDefault();
    console.log("Form Data:", formData);
  };

  return (
    <>
      <div className="m-3 flex flex-col mx-5 w-full">
        <div className="flex justify-between align-center items-center">
          <h1 className="text-3xl font-bold p-2 mb-5">Scans</h1>
          <Dialog>
            <DialogTrigger asChild>
              <Button variant="outline">
                <FaPlus /> New Scan
              </Button>
            </DialogTrigger>
            <DialogContent className="sm:max-w-[425px]">
              <DialogHeader>
                <DialogTitle>Launch New Scan</DialogTitle>
                <DialogDescription>
                  Please fill all the fields to scan for misconfigurations in the
                  Kubernetes cluster using the CIS Benchmark.
                </DialogDescription>
              </DialogHeader>
              <div className="grid gap-4 py-4">
                <div className="grid items-center gap-4">
                  <Label htmlFor="title">Title</Label>
                  <Input
                    type="text"
                    id="title"
                    name="title"
                    placeholder="Enter title for scan"
                    value={formData.title}
                    onChange={handleChange}
                    className="col-span-3"
                  />
                </div>
                <div className="grid items-center gap-4">
                  <Label htmlFor="description">Description</Label>
                  <Textarea
                    id="description"
                    name="description"
                    placeholder="Enter description for new scan"
                    value={formData.description}
                    onChange={handleChange}
                  />
                </div>
                <div className="grid items-center gap-4">
                  <Label htmlFor="IP">IP Address</Label>
                  <div className="flex">
                    <Select onValueChange={(value) => handleSelectChange("ip", value)}>
                      <SelectTrigger className="w-full">
                        <SelectValue placeholder="Select IP Address" />
                      </SelectTrigger>
                      <SelectContent>
                        <SelectItem value="127.0.0.1">127.0.0.1</SelectItem>
                      </SelectContent>
                    </Select>
                    <Button onClick={handleRefreshedIp} className="ms-1" variant="outline">
                      <TfiReload />
                    </Button>
                  </div>
                </div>
                <div className="grid items-center gap-4">
                  <Label htmlFor="IG">Implementation Group</Label>
                  <Select onValueChange={(value) => handleSelectChange("ig", value)}>
                    <SelectTrigger className="w-full">
                      <SelectValue placeholder="Select IG" />
                    </SelectTrigger>
                    <SelectContent>
                      <SelectItem value="IG-1">IG-1</SelectItem>
                      <SelectItem value="IG-2">IG-2</SelectItem>
                      <SelectItem value="IG-3">IG-3</SelectItem>
                    </SelectContent>
                  </Select>
                </div>
              </div>
              <DialogFooter>
                <Button type="submit" onClick={handleSubmit}>
                  Launch Scan
                </Button>
              </DialogFooter>
            </DialogContent>
          </Dialog>
        </div>
        <ScansTable />
      </div>
    </>
  );
};

export default Page;
