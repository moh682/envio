import { DataTable } from "@/components/data-table";
import React from "react";
import { columns, Expense } from "./partials/columns";
import { PlusIcon } from "lucide-react";
import { Link } from "@/components/link";

const getData = async (): Promise<Expense[]> => [
  {
    id: "1",
    issuedAt: new Date("01-05-2024"),
    company: "FTZ",
    paymentOption: "CASH",
    account: "Kontor artikler",
    amount: 1000.0,
    isVat: true,
  },
  {
    id: "2",
    issuedAt: new Date("01-02-2024"),
    company: "AD Danmark",
    paymentOption: "BANK",
    account: "Kontor artikler",
    amount: 500.31,
    isVat: false,
  },
  {
    id: "3",
    issuedAt: new Date("01-02-2024"),
    company: "FTZ",
    paymentOption: "BANK",
    account: "Kontor artikler",
    amount: 2456,
    isVat: true,
  },
];

async function Page() {
  const data = await getData();
  return (
    <div className="space-y-3">
      <Link href="/expenses/new">
        Expense
        <PlusIcon className="h-5 w-5" />
      </Link>
      <DataTable columns={columns} data={data} />
    </div>
  );
}

export default Page;
