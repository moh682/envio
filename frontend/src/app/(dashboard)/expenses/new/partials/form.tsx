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
import {
  Dialog,
  DialogContent,
  DialogDescription,
  DialogHeader,
  DialogTitle,
  DialogTrigger,
} from "@/components/ui/dialog";
import { PlusCircleIcon } from "lucide-react";
import { useState } from "react";
import { Label } from "@/components/ui/label";
import { SearchCompany } from "./search-company";
import { CreateCompany } from "./create-company";

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
  paymentOption: z.string().min(1),
  account: z.string().min(1),
  isVat: z.boolean(),
  amount: z.string().min(1),
  company: z.string().min(1),
});

export const ExponseForm = () => {
  const [open, setOpen] = useState(false);
  const [stage, setStage] = useState<"searchCompany" | "createCompany">("searchCompany");
  const form = useForm<z.infer<typeof ExpenseSchema>>({
    resolver: zodResolver(ExpenseSchema),
    defaultValues: {
      issuedAt: new Date(),
      paymentOption: "",
      account: "",
      isVat: true,
      amount: "",
      company: "",
    },
  });

  useWatch({
    control: form.control,
    name: ["isVat", "amount"],
    exact: true,
  });

  const calcTotal = () => {
    let amount = form.getValues().amount.toString();
    const isVat = form.getValues().isVat;

    if (!amount) return 0;

    let parsedAmount = parseFloat(amount);

    if (isNaN(parsedAmount)) return 0;

    if (isVat) parsedAmount = parsedAmount * 1.25;

    return parsedAmount.toFixed(2);
  };

  const setCompany = (value: string) => {
    form.setValue("company", value, {
      shouldDirty: true,
    });

    setOpen(false);
  };

  const onSubmit = async (values: z.infer<typeof ExpenseSchema>) => {
    if (isNaN(Number(values.amount))) {
      form.setError("amount", {
        message: "Must be a number",
      });
      return;
    }

    console.log(values);
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
          <div className="col-span-6 md:col-span-3 space-y-4">
            <FormField
              control={form.control}
              name="issuedAt"
              render={({ field }) => (
                <FormItem>
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
              name="company"
              render={({ field }) => (
                <FormItem className="col-span-6 md:col-span-3">
                  <FormLabel>Company</FormLabel>
                  <Dialog
                    open={open}
                    onOpenChange={(value) => {
                      setOpen(value);
                      setStage("searchCompany");
                    }}
                  >
                    <FormControl>
                      <div className="w-full flex gap-3">
                        <div className="flex-1 border rounded-md py-1 px-3 shadow-xs">
                          {field.value ? (
                            <span className="text-sm">{field.value}</span>
                          ) : (
                            <span className="text-gray-500 text-sm">Select company</span>
                          )}
                        </div>
                        <DialogTrigger asChild>
                          <Button variant="outline" className="cursor-pointer">
                            <PlusCircleIcon className="h-6 w-6" />
                          </Button>
                        </DialogTrigger>
                      </div>
                    </FormControl>
                    <DialogContent>
                      {stage === "searchCompany" && (
                        <SearchCompany setCompany={setCompany} setStageToCreate={() => setStage("createCompany")} />
                      )}

                      {stage === "createCompany" && <CreateCompany setCompany={setCompany} />}
                    </DialogContent>
                  </Dialog>
                </FormItem>
              )}
            />
          </div>
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
              <div className="flex-1 border rounded-md py-1 px-3 shadow-xs">
                {calcTotal() || <span className="text-gray-500 text-sm">Total</span>}
              </div>
            </div>

            <Button className="w-full md:w-fit" type="submit" disabled={!form.formState.isDirty}>
              Create Expense
            </Button>
          </div>
        </form>
      </Form>
    </section>
  );
};
