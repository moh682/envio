"use client";

import { Button } from "@/components/ui/button";
import { Form } from "@/components/ui/form";
import { Separator } from "@/components/ui/separator";
import { zodResolver } from "@hookform/resolvers/zod";
import React from "react";
import { useForm } from "react-hook-form";
import z from "zod";
import { InvoiceInfo } from "./invoice-info";
import { CustomerInfo } from "./customer-info";
import { Summary } from "./summary";
import { Products } from "./products";
import { InvoiceSchema } from "./schema";

export const InvoiceForm = () => {
  const form = useForm<z.infer<typeof InvoiceSchema>>({
    resolver: zodResolver(InvoiceSchema),
    defaultValues: {
      invoiceNumber: 1,
      issuedAt: new Date(),
      address: "",
      carRegistration: "",
      email: "",
      name: "",
      phone: "",
      products: [
        {
          id: "",
          description: "",
          price: 0,
          quantity: 0,
        },
      ],
    },
  });

  const onSubmit = async (values: z.infer<typeof InvoiceSchema>) => {
    console.log(values);
  };

  const getSum = () => {
    return form.watch("products").reduce((acc, item) => acc + item.quantity * item.price, 0);
  };

  return (
    <section>
      <Form {...form}>
        <form
          onSubmit={form.handleSubmit(onSubmit)}
          className="grid grid-cols-6 gap-3"
          onKeyDown={(e) => {
            if (e.key === "Enter") e.preventDefault();
          }}
        >
          <InvoiceInfo form={form} />
          <CustomerInfo form={form} />
          <Products form={form} />
          <Separator className="col-span-6" />
          <Summary getSum={getSum} />
          <Button className="col-span-6 md:col-span-3 md:col-start-4" type="submit">
            Create Invoice
          </Button>
        </form>
      </Form>
    </section>
  );
};
