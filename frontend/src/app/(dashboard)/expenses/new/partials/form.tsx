"use client";

import { useForm } from "react-hook-form";
import { z } from "zod";
import { zodResolver } from "@hookform/resolvers/zod";
import { Form } from "@/components/ui/form";
import { Button } from "@/components/ui/button";
import { ExpenseSchema } from "./schema";
import { BillingInfo } from "./billing-info";
import { PaymentInfo } from "./payment-info";

export const ExponseForm = () => {
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

  const onSubmit = async (values: z.infer<typeof ExpenseSchema>) => {
    if (isNaN(Number(values.amount))) {
      form.setError("amount", {
        message: "Must be a number",
      });
      return;
    }
    console.log(values);
  };

  return (
    <section>
      <Form {...form}>
        <form
          onSubmit={form.handleSubmit(onSubmit)}
          className="grid grid-cols-6 gap-4"
          onKeyDown={(e) => {
            if (e.key === "Enter") e.preventDefault();
          }}
        >
          <BillingInfo form={form} />
          <PaymentInfo form={form} />

          <Button className="w-full md:col-span-1 md:col-start-4" type="submit" disabled={!form.formState.isDirty}>
            Create Expense
          </Button>
        </form>
      </Form>
    </section>
  );
};
