import React from "react";
import Link from "next/link";
import { IoMenu } from "react-icons/io5";
import { FaFilePdf, FaHtml5, FaFileExcel } from "react-icons/fa";
import { SiKubernetes } from "react-icons/si";
import { FaGlobe } from "react-icons/fa";
import { Badge } from "@/components/ui/badge";
// import { badgeVariants } from "@/components/ui/badge";
import {
  Accordion,
  AccordionContent,
  AccordionItem,
  AccordionTrigger,
} from "@/components/ui/accordion";
import {
  DropdownMenu,
  DropdownMenuContent,
  DropdownMenuItem,
  DropdownMenuLabel,
  DropdownMenuSeparator,
  DropdownMenuTrigger,
} from "@/components/ui/dropdown-menu";

import {
  Table,
  TableBody,
  TableCaption,
  TableCell,
  TableFooter,
  TableHead,
  TableHeader,
  TableRow,
} from "@/components/ui/table";

const invoices = [
  {
    id: "1",
    check_name: "Alpha Scan",
    check_id: "Level 1 - Master Node",
    status: "Running",
    reference: "https://kubernetes.io/docs/admin",
    ig_1: "Running",
    ig_2: "Running",
    ig_3: "Running",
    action: "13/02/2025",
  },
  {
    id: "2",
    check_name: "Beta Scan",
    check_id: "Level 2 - Master Node",
    status: "Failed",
    reference: "https://kubernetes.io/docx/admin",
    ig_1: "Paused",
    ig_2: "Paused",
    ig_3: "Paused",
    action: "01/11/2023",
  },
  {
    id: "3",
    check_name: "Gemma Scan",
    check_id: "Level 3 - Master Node",
    status: "Pending",
    reference: "https://kubernetes.io/docx/admin",
    ig_1: "Compeleted",
    ig_2: "Compeleted",
    ig_3: "Compeleted",
    action: "10/10/2024",
  },
];

const Page = () => {
  return (
    <>
      <div className="m-3 flex flex-col mx-5 w-full">
        <div className="flex justify-between align-center items-center">
          <div>
            <h1 className="text-3xl font-bold px-2">Scan Results</h1>
            <p className="px-2">This is title of Scans</p>
          </div>
          <DropdownMenu>
            <DropdownMenuTrigger>
              <IoMenu />
            </DropdownMenuTrigger>
            <DropdownMenuContent align="end">
              <DropdownMenuLabel>Export As</DropdownMenuLabel>
              <DropdownMenuSeparator />
              <DropdownMenuItem>
                <FaFilePdf className="me-1" />
                PDF
              </DropdownMenuItem>
              <DropdownMenuItem>
                <FaHtml5 className="me-1" />
                HTML
              </DropdownMenuItem>
              <DropdownMenuItem>
                <FaFileExcel className="me-1" />
                EXCEL
              </DropdownMenuItem>
            </DropdownMenuContent>
          </DropdownMenu>
        </div>
        <div className="mt-8 ms-8 w-[30%]">
          <div className="flex justify-around items-center">
            <div>
              <SiKubernetes className="text-9xl" />
            </div>
            <div>
              <h1 className="font-bold text-4xl mb-3">Hostname</h1>
              <div className="flex justify-start gap-5 !p-0 !m-0">
                <h1 className="font-bold">IP Address</h1>
                <p>192.168.2.200</p>
              </div>
              <div className="flex justify-start gap-5 !p-0 !m-0">
                <h1 className="font-bold">State</h1>
                <p>Online</p>
              </div>
              <div className="flex justify-start gap-5 !p-0 !m-0">
                <h1 className="font-bold">Last Updated At</h1>
                <p>29/12/2024</p>
              </div>
            </div>
          </div>
        </div>
        <div className="flex justify-center">
          <div className="mt-14 w-[90%] ">
            {/* 1st Accordian*/}
            <Accordion type="single" collapsible>
              <AccordionItem value="item-1">
                <AccordionTrigger>Master Node Specifications</AccordionTrigger>
                <AccordionContent>
                  <div className="flex">
                    <div className="w-[30%]">
                      <h1>CPU</h1>
                      <h1>Memory</h1>
                      <h1>OS</h1>
                      <h1>OS Version</h1>
                      <h1>Kubernetes Client Version</h1>
                      <h1>Kubernetes Server Version</h1>
                    </div>
                    <div className="w-[70%]">
                      <p>SALMAN</p>
                      <p>SAAD</p>
                      <p>HUSNNAIN</p>
                      <p>AHMED</p>
                      <p>ZEESHAN</p>
                      <p>AHMER</p>
                    </div>
                  </div>
                </AccordionContent>
              </AccordionItem>
            </Accordion>
          </div>
        </div>
        {/* Table */}
        <h1 className="mt-14 text-3xl font-bold px-2">Audits Checklist</h1>
        <p className="px-2 mb-5">Continuously evaluates your Kubernetes cluster against CIS security benchmarks, identifying misconfigurations and vulnerabilities to enhance compliance and security posture.</p>
        <div className="ms-20 mt-6 w-[90%]">
          {" "}
          <Table className="w-full px-10">
            <TableCaption>Lorem ipsum dolor sit amet consectetur adipisicing elit. Nobis enim vitae sequi, eum id ad vero laborum tempore pariatur. Quis culpa autem consequatur corporis maiores officiis dolores eum assumenda! Unde ea veniam perspiciatis nam dicta odit. Officiis, vel rem error deserunt atque commodi facilis qui quam eius, laboriosam eaque perferendis.</TableCaption>
            <TableHeader>
              <TableRow>
                <TableHead className="w-[100px]">#</TableHead>
                <TableHead>Check Name</TableHead>
                <TableHead>Applicibility</TableHead>
                <TableHead className="text-center">Reference</TableHead>
                <TableHead className="text-center">IG-1</TableHead>
                <TableHead className="text-center">IG-2</TableHead>
                <TableHead className="text-center">IG-3</TableHead>
                <TableHead>Status</TableHead>
                <TableHead>Action</TableHead>
              </TableRow>
            </TableHeader>
            <TableBody>
              {invoices.map((invoice) => (
                <TableRow key={invoice.id}>
                  <TableCell className="font-medium">{invoice.id}</TableCell>
                  <TableCell>{invoice.check_name}</TableCell>
                  <TableCell>{invoice.check_id}</TableCell>
                  <TableCell><div className="flex items-center justify-center"><Link href={invoice.reference}><FaGlobe/></Link></div></TableCell>
                  <TableCell>
                    <div className="flex items-center justify-center"><Badge className="p-1  bg-green-500" variant="outline"></Badge></div>
                  </TableCell>
                  <TableCell>
                  <div className="flex items-center justify-center"><Badge className="p-1  bg-orange-500" variant="outline"></Badge></div>
                  </TableCell>
                  <TableCell>
                  <div className="flex items-center justify-center"><Badge className="p-1  bg-blue-500" variant="outline"></Badge></div>
                  </TableCell>
                  <TableCell>{invoice.status}</TableCell>
                  <TableCell>{invoice.action}</TableCell>
                </TableRow>
              ))}
            </TableBody>
            <TableFooter>
              <TableRow>
                <TableCell colSpan={8}>Total Scans</TableCell>
                <TableCell className="text-right">3</TableCell>
              </TableRow>
            </TableFooter>
          </Table>
        </div>
      </div>
    </>
  );
};

export default Page;
