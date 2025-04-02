"use client";

import React from "react";
import { UseFormReturn, useWatch } from "react-hook-form";
import { z } from "zod";
import { expenseSchema } from "./schema";
import { FormField, FormItem, FormLabel, FormControl, FormMessage } from "@/components/ui/form";
import { Input } from "@/components/ui/input";
import { Select, SelectContent, SelectItem, SelectTrigger, SelectValue } from "@/components/ui/select";
import { Switch } from "@/components/ui/switch";
import { Label } from "@/components/ui/label";

const paymentOptions = [
  {
    value: "CASH",
    text: "Cash",
  },
  {
    value: "BANK",
    text: "Bank",
  },
];

const accounts = [
  {
    value: "SALG_YDELSER",
    text: "Salg og Ydelser",
    isVat: true,
  },
  {
    value: "KONTOR_ARTIKLER",
    text: "Kontor artikler",
    isVat: false,
  },
];

type Props = {
  form: UseFormReturn<z.infer<typeof expenseSchema>>;
};

export const PaymentInfo = ({ form }: Props) => {
  useWatch({
    control: form.control,
    name: ["isVat", "amount"],
    exact: true,
  });

  const calcTotal = () => {
    let amount: string | number = form.getValues().amount.toString();
    const isVat = form.getValues().isVat;

    if (!amount) return "";

    amount = parseFloat(amount);

    if (isNaN(amount)) return "";

    if (isVat) amount = amount * 1.25;

    return amount.toFixed(2);
  };
  return (
    <div className="col-span-6 md:col-span-3 space-y-4">
      <FormField
        control={form.control}
        name="paymentOption"
        render={({ field }) => (
          <FormItem>
            <FormLabel>Payment option</FormLabel>
            <Select onValueChange={field.onChange} defaultValue={field.value}>
              <FormControl>
                <SelectTrigger className="w-full">
                  <SelectValue placeholder="Select payment option" />
                </SelectTrigger>
              </FormControl>
              <SelectContent>
                {paymentOptions.map((option) => (
                  <SelectItem key={option.value} value={option.value}>
                    {option.text}
                  </SelectItem>
                ))}
              </SelectContent>
            </Select>
          </FormItem>
        )}
      />
      <FormField
        control={form.control}
        name="account"
        render={({ field }) => (
          <FormItem>
            <FormLabel>Account</FormLabel>
            <Select
              onValueChange={(value) => {
                const option = accounts.find((account) => account.value === value);
                if (option) form.setValue("isVat", option.isVat, { shouldDirty: true });
                field.onChange(value);
              }}
              defaultValue={field.value}
            >
              <FormControl>
                <SelectTrigger className="w-full">
                  <SelectValue placeholder="Select account" />
                </SelectTrigger>
              </FormControl>
              <SelectContent>
                {accounts.map((account) => (
                  <SelectItem key={account.value} value={account.value}>
                    {account.text}
                  </SelectItem>
                ))}
              </SelectContent>
            </Select>
          </FormItem>
        )}
      />
      <FormField
        control={form.control}
        name="amount"
        render={({ field }) => (
          <FormItem>
            <FormLabel>Amount</FormLabel>
            <FormControl>
              <Input placeholder="Amount" {...field} />
            </FormControl>
            <FormMessage />
          </FormItem>
        )}
      />
      <FormField
        control={form.control}
        name="isVat"
        render={({ field }) => (
          <FormItem className="flex flex-row items-center justify-between rounded-lg border p-3 shadow-xs">
            <FormLabel>Vat</FormLabel>
            <FormControl>
              <Switch checked={field.value} onCheckedChange={field.onChange} />
            </FormControl>
          </FormItem>
        )}
      />

      <div className="flex flex-col gap-2">
        <Label>Total</Label>
        <div className="flex-1 border rounded-md py-[6px] px-3 shadow-xs text-sm">
          {calcTotal() ? <span>{calcTotal()}</span> : <span className="text-gray-500">Total</span>}
        </div>
      </div>
    </div>
  );
};
