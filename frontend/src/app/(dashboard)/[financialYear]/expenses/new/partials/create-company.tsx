"use client";

import { Button } from "@/components/ui/button";
import { DialogFooter, DialogHeader, DialogTitle } from "@/components/ui/dialog";
import { Form, FormControl, FormField, FormItem, FormMessage } from "@/components/ui/form";
import { Input } from "@/components/ui/input";
import { Label } from "@/components/ui/label";
import { zodResolver } from "@hookform/resolvers/zod";
import React from "react";
import { useForm } from "react-hook-form";
import { z } from "zod";

type Props = {
  setCompany: (name: string) => void;
};

const formSchema = z.object({
  id: z.string().min(1),
  name: z.string().min(1),
});

export const CreateCompany = ({ setCompany }: Props) => {
  const form = useForm({
    resolver: zodResolver(formSchema),
    defaultValues: {
      id: "asdada",
      name: "",
    },
  });

  const onSubmit = (values: z.infer<typeof formSchema>) => {
    setCompany(values.name);
  };

  return (
    <>
      <DialogHeader>
        <DialogTitle>Create new company</DialogTitle>
      </DialogHeader>

      <Form {...form}>
        <form onSubmit={form.handleSubmit(onSubmit)}>
          <FormField
            name="name"
            control={form.control}
            render={({ field }) => (
              <FormItem>
                <Label>Name</Label>
                <FormControl>
                  <Input {...field} autoFocus />
                </FormControl>
                <FormMessage />
              </FormItem>
            )}
          />
        </form>
      </Form>

      <DialogFooter>
        <Button disabled={!form.formState.isDirty} onClick={form.handleSubmit(onSubmit)}>
          Create
        </Button>
      </DialogFooter>
    </>
  );
};
