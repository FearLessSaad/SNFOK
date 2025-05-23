import Link from "next/link"
import {
    Table,
    TableBody,
    TableCaption,
    TableCell,
    TableFooter,
    TableHead,
    TableHeader,
    TableRow,
  } from "@/components/ui/table"
  
  const invoices = [
    {
      id: "1",
      title: "Alpha Scan",
      description: "This is test scan...",
      ip: "172.20.30.40",
      status: "Running",
      time: "13/02/2025"
    },
    {
      id: "2",
      title: "Beta Scan",
      description: "This is test scan...",
      ip: "172.2.3.4",
      status: "Paused",
      time: "01/11/2023"
    },
    {
      id: "3",
      title: "Gemma Scan",
      description: "This is test scan...",
      ip: "172.120.230.140",
      status: "Compeleted",
      time: "10/10/2024"
    },
  ]
  
  export function ScansTable() {
    return (
      <Table className="w-full px-10">
        <TableCaption>A list of your all scans for Audits.</TableCaption>
        <TableHeader>
          <TableRow>
            <TableHead className="w-[100px]">#</TableHead>
            <TableHead><Link href="http://localhost:3000/scans/5">Title</Link></TableHead>
            <TableHead>Description</TableHead>
            <TableHead>IP Address</TableHead>
            <TableHead>Status</TableHead>
            <TableHead>Time</TableHead>
          </TableRow>
        </TableHeader>
        <TableBody>
          {invoices.map((invoice) => (
            <TableRow key={invoice.id}>
              <TableCell className="font-medium">{invoice.id}</TableCell>
              <TableCell>{invoice.title}</TableCell>
              <TableCell>{invoice.description}</TableCell>
              <TableCell>{invoice.ip}</TableCell>
              <TableCell>{invoice.status}</TableCell>
              <TableCell>{invoice.time}</TableCell>
            </TableRow>
          ))}
        </TableBody>
        <TableFooter>
          <TableRow>
            <TableCell colSpan={5}>Total Scans</TableCell>
            <TableCell className="text-right">3</TableCell>
          </TableRow>
        </TableFooter>
      </Table>
    )
  }
  