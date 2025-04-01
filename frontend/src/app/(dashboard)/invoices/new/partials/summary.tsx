"use client";

import { Input } from "@/components/ui/input";
import React from "react";

type Props = {
  getSum: () => number;
};

export const Summary = ({ getSum }: Props) => {
  return (
    <>
      <div className="col-span-6 md:col-span-3 md:col-start-4 text-sm flex gap-4 items-center">
        <p className="w-36">Total Excl. Vat</p>
        <Input disabled value={getSum()} />
      </div>
      <div className="col-span-6 md:col-span-3 md:col-start-4 text-sm flex gap-4 items-center">
        <p className="w-36">Vat</p>
        <Input disabled value={"25%"} />
      </div>
      <div className="col-span-6 md:col-span-3 md:col-start-4 text-sm flex gap-4 items-center">
        <p className="w-36">Vat Cost</p>
        <Input disabled value={getSum() * 1.25 - getSum()} />
      </div>
      <div className="col-span-6 md:col-span-3 md:col-start-4 text-sm flex gap-4 items-center">
        <p className="w-36">Total Incl. Vat</p>
        <Input disabled value={getSum() * 1.25} />
      </div>
    </>
  );
};
