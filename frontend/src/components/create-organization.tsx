"use client";

import React, { useState, useTransition } from "react";
import { Dialog, DialogContent, DialogFooter, DialogHeader, DialogTitle } from "./ui/dialog";
import { Form, FormControl, FormField, FormItem, FormLabel, FormMessage } from "./ui/form";
import { useForm } from "react-hook-form";
import { zodResolver } from "@hookform/resolvers/zod";
import { createOrganizationSchema } from "@/app/domains/organization/schema";
import { z } from "zod";
import { Button } from "./ui/button";
import { Input } from "./ui/input";
import { Loader } from "./loader";
import { createOrganization } from "@/app/domains/organization/actions";
import { toast } from "sonner";

export const CreateOrganization = () => {
  const [isPending, startTransition] = useTransition();

  const form = useForm({
    resolver: zodResolver(createOrganizationSchema),
    defaultValues: {
      name: "",
      invoiceNumberStart: "",
    },
  });

  const onSubmit = (values: z.infer<typeof createOrganizationSchema>) => {
    if (isNaN(Number(values.invoiceNumberStart))) {
      form.setError("invoiceNumberStart", {
        message: "Must be a number",
      });
      return;
    }

    startTransition(async () => {
      const response = await createOrganization({ ...values, invoiceNumberStart: parseInt(values.invoiceNumberStart) });

      if (response?.error) {
        toast.error(response.error.message);
        return;
      }
    });
  };

  return (
    <Dialog open={true}>
      <Form {...form}>
        <DialogContent>
          <DialogHeader>
            <DialogTitle>Create organization</DialogTitle>
          </DialogHeader>

          <form onSubmit={form.handleSubmit(onSubmit)} className="space-y-3">
            <FormField
              control={form.control}
              name="name"
              render={({ field }) => (
                <FormItem>
                  <FormLabel>Name</FormLabel>
                  <FormControl>
                    <Input {...field} />
                  </FormControl>
                  <FormMessage />
                </FormItem>
              )}
            />
            <FormField
              control={form.control}
              name="invoiceNumberStart"
              render={({ field }) => (
                <FormItem>
                  <FormLabel>Invoice start number</FormLabel>
                  <FormControl>
                    <Input {...field} />
                  </FormControl>
                  <FormMessage />
                </FormItem>
              )}
            />
            <DialogFooter>
              <Button onClick={form.handleSubmit(onSubmit)} disabled={!form.formState.isDirty}>
                {isPending ? <Loader /> : "Create"}
              </Button>
            </DialogFooter>
          </form>
        </DialogContent>
      </Form>
    </Dialog>
  );
};
