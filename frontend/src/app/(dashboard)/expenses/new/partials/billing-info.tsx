"use client";

import React, { useState } from "react";
import { UseFormReturn } from "react-hook-form";
import { z } from "zod";
import { expenseSchema } from "./schema";
import { FormControl, FormField, FormItem, FormLabel, FormMessage } from "@/components/ui/form";
import { DatePicker } from "@/components/date-picker";
import { PlusCircleIcon } from "lucide-react";
import { CreateCompany } from "./create-company";
import { SearchCompany } from "./search-company";
import { Dialog, DialogContent, DialogTrigger } from "@/components/ui/dialog";
import { Button } from "@/components/ui/button";

type Props = {
  form: UseFormReturn<z.infer<typeof expenseSchema>>;
};

export const BillingInfo = ({ form }: Props) => {
  const [open, setOpen] = useState(false);
  const [stage, setStage] = useState<"searchCompany" | "createCompany">("searchCompany");
  const setCompany = (value: string) => {
    form.setValue("company", value, {
      shouldDirty: true,
    });

    setOpen(false);
  };
  return (
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
  );
};
