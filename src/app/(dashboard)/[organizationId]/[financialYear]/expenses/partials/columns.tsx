"use client";

import { ColumnDef } from "@tanstack/react-table";
import { format } from "date-fns";
import Link from "next/link";

export type Expense = {
  id: string;
  issuedAt: Date;
  company: string;
  paymentOption: "CASH" | "BANK";
  account: string;
  amount: number;
  isVat: boolean;
};

export const columns: ColumnDef<Expense>[] = [
  {
    accessorKey: "issuedAt",
    header: "Issued At",
    cell: ({ row }) => (
      <Link className="text-blue-500" href={`/expenses/${row.original.id}`}>
        {format(row.original.issuedAt, "MMM dd, yyyy")}
      </Link>
    ),
  },
  {
    accessorKey: "company",
    header: "Company",
  },
  {
    accessorKey: "paymentOption",
    header: "Payment option",
  },
  {
    accessorKey: "amount",
    header: "Amount",
    cell: ({ row }) => (row.original.isVat ? (row.original.amount * 1.25).toFixed(2) : row.original.amount),
  },
];
