"use client";

import React from "react";
import { UseFormReturn } from "react-hook-form";
import { DatePicker } from "@/components/date-picker";
import { FormField, FormItem, FormLabel, FormControl } from "@/components/ui/form";
import { Input } from "@/components/ui/input";
import { z } from "zod";
import { InvoiceSchema } from "./schema";

type Prop = {
  form: UseFormReturn<z.infer<typeof InvoiceSchema>>;
};
export const InvoiceInfo = ({ form }: Prop) => {
  return (
    <div className="col-span-6 md:col-span-3 space-y-3">
      <FormField
        control={form.control}
        name="invoiceNumber"
        render={({ field }) => (
          <FormItem>
            <FormLabel>Invoice Number</FormLabel>
            <FormControl>
              <Input {...field} disabled />
            </FormControl>
          </FormItem>
        )}
      />
      <FormField
        control={form.control}
        name="issuedAt"
        render={({ field }) => (
          <FormItem>
            <FormLabel>Issued At</FormLabel>
            <FormControl>
              <DatePicker onSelect={(selected) => field.onChange(selected)} selected={field.value} className="w-full" />
            </FormControl>
          </FormItem>
        )}
      />
    </div>
  );
};
