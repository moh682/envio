import React from "react";
import { columns, Invoice } from "./partials/columns";
import { Link } from "@/components/link";
import { PlusIcon } from "lucide-react";
import { DataTable } from "@/components/data-table";
import { rawFetch } from "@/lib/fetch";

const getData = async (): Promise<Invoice[]> => [
  {
    id: "1",
    invoiceNumber: 1,
    carRegistration: "YY58002",
    issuedAt: new Date("01-05-2024"),
    paymentOption: "CASH",
    amount: 1000.0,
  },
  {
    id: "2",
    invoiceNumber: 2,
    carRegistration: "HV53502",
    issuedAt: new Date("01-02-2024"),
    paymentOption: "BANK",
    amount: 500.31,
  },
  {
    id: "3",
    invoiceNumber: 3,
    carRegistration: "YG58362",
    issuedAt: new Date("01-02-2024"),
    paymentOption: "BANK",
    amount: 2456,
  },
];

const getDataFromServer = async () => {
  const response = await rawFetch("/invoices", {
    method: "GET",
  });

  return await response.json();
};

async function Page() {
  const data22 = await getDataFromServer();
  console.log(data22);
  const data = await getData();
  return (
    <div className="space-y-3">
      <Link href="/invoices/new">
        Invoice
        <PlusIcon className="h-5 w-5" />
      </Link>
      <DataTable columns={columns} data={data} />
    </div>
  );
}

export default Page;
