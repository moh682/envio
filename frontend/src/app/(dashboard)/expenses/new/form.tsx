"use client";

import { useForm, useWatch } from "react-hook-form";
import { z } from "zod";
import { zodResolver } from "@hookform/resolvers/zod";
import { Form, FormControl, FormField, FormItem, FormLabel, FormMessage } from "@/components/ui/form";
import { Button } from "@/components/ui/button";
import { Select, SelectContent, SelectItem, SelectTrigger, SelectValue } from "@/components/ui/select";
import { Input } from "@/components/ui/input";
import { DatePicker } from "@/components/date-picker";
import { Switch } from "@/components/ui/switch";
import { Label } from "@radix-ui/react-label";

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
const ExpenseSchema = z.object({
  issuedAt: z.date(),
  paymentOption: z.string(),
  account: z.string(),
  isVat: z.boolean(),
  amount: z.string().min(1),
});

export const ExponseForm = () => {
  const form = useForm<z.infer<typeof ExpenseSchema>>({
    resolver: zodResolver(ExpenseSchema),
    defaultValues: {
      issuedAt: new Date(),
      paymentOption: "",
      account: "",
      isVat: true,
      amount: "",
    },
  });

  useWatch({
    control: form.control,
    name: ["isVat", "amount"],
    exact: true,
  });

  const calcTotal = () => {
    let amount = form.getValues().amount;
    const isVat = form.getValues().isVat;

    if (!amount) return 0;

    let parsedAmount = parseFloat(amount);

    if (isNaN(parsedAmount)) return 0;

    if (isVat) parsedAmount = parsedAmount * 1.25;

    return parsedAmount.toFixed(2);
  };

  const onSubmit = async (data: z.infer<typeof ExpenseSchema>) => {
    console.log(data);
  };

  const onError = (errors: any) => {
    console.log({ errors });
  };

  return (
    <section>
      <Form {...form}>
        <form
          onSubmit={form.handleSubmit(onSubmit, onError)}
          className="grid grid-cols-6 gap-4"
          onKeyDown={(e) => {
            if (e.key === "Enter") e.preventDefault();
          }}
        >
          <FormField
            control={form.control}
            name="issuedAt"
            render={({ field }) => (
              <FormItem className="col-span-6 md:col-span-3">
                <FormLabel>Issued at</FormLabel>
                <FormControl>
                  <DatePicker className="w-full" selected={field.value} onSelect={field.onChange} />
                </FormControl>
                <FormMessage />
              </FormItem>
            )}
          />
          <FormField
            control={form.control}
            name="paymentOption"
            render={({ field }) => (
              <FormItem className="col-span-6 md:col-span-3">
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
                <FormMessage />
              </FormItem>
            )}
          />
          <FormField
            control={form.control}
            name="account"
            render={({ field }) => (
              <FormItem className="col-span-6 md:col-span-3 md:col-start-4">
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
                <FormMessage />
              </FormItem>
            )}
          />
          <FormField
            control={form.control}
            name="amount"
            render={({ field }) => (
              <FormItem className="col-span-6 md:col-span-3 md:col-start-4">
                <FormLabel>Amount</FormLabel>
                <FormControl>
                  <Input placeholder="Amount" {...field} />
                </FormControl>
              </FormItem>
            )}
          />
          <FormField
            control={form.control}
            name="isVat"
            render={({ field }) => (
              <FormItem className="flex flex-row items-center justify-between rounded-lg border p-3 col-span-6 md:col-span-3 md:col-start-4">
                <FormLabel>Vat</FormLabel>
                <FormControl>
                  <Switch checked={field.value} onCheckedChange={field.onChange} />
                </FormControl>
              </FormItem>
            )}
          />

          <div className="flex flex-col col-span-6 md:col-span-3 md:col-start-4">
            <Label>Total</Label>
            <Input disabled value={calcTotal()} />
          </div>
          <Button className="col-span-6 md:col-span-3 md:col-start-4" type="submit">
            Create Expense
          </Button>
        </form>
      </Form>
    </section>
  );
};
