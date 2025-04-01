"use client";

import { ColumnDef } from "@tanstack/react-table";
import { format } from "date-fns";
import Link from "next/link";

export type Invoice = {
  id: string;
  invoiceNumber: number;
  issuedAt: Date;
  carRegistration: string;
  paymentOption: "CASH" | "BANK";
  amount: number;
};

export const columns: ColumnDef<Invoice>[] = [
  {
    accessorKey: "issuedAt",
    header: "Issued At",
    cell: ({ row }) => (
      <Link className="text-blue-500" href={`/invoices/${row.original.id}`}>
        {format(row.original.issuedAt, "MMM dd, yyyy")}
      </Link>
    ),
  },
  {
    accessorKey: "carRegistration",
    header: "Car registration",
  },
  {
    accessorKey: "paymentOption",
    header: "Payment option",
  },
  {
    accessorKey: "amount",
    header: "Amount",
    cell: ({ row }) => row.original.amount.toFixed(2),
  },
];
