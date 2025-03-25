"use client";

import { DatePicker } from "@/components/date-picker";
import { Alert } from "@/components/ui/alert";
import { Button } from "@/components/ui/button";
import { Form, FormControl, FormDescription, FormField, FormItem, FormLabel, FormMessage } from "@/components/ui/form";
import { Input } from "@/components/ui/input";
import { Separator } from "@/components/ui/separator";
import { Table, TableBody, TableCaption, TableCell, TableHead, TableHeader, TableRow } from "@/components/ui/table";
import { zodResolver } from "@hookform/resolvers/zod";
import React from "react";
import { useFieldArray, useForm } from "react-hook-form";
import z from "zod";

const InvoiceSchema = z
  .object({
    invoiceNumber: z.number().readonly(),
    issuedAt: z.date().default(new Date()),
    customerName: z.string().optional(),
    customerEmail: z.string().email().optional(),
    customerCarRegistration: z.string().optional(),
    customerPhone: z.string().optional(),
    customerAddress: z.string().optional(),

    products: z.array(
      z.object({
        id: z.string().optional(),
        description: z.string(),
        quantity: z.number().positive().default(0),
        price: z.number().positive().default(0),
      })
    ),
  })
  .superRefine((invoice, ctx) => {
    if (
      !invoice.customerAddress &&
      !invoice.customerEmail &&
      !invoice.customerCarRegistration &&
      !invoice.customerPhone &&
      !invoice.customerName
    ) {
      const keys = Object.keys(invoice).filter((key) => key.includes("customer"));
      keys.forEach((key) => {
        ctx.addIssue({
          code: "custom",
          message: `This field is Required`,
          path: [key],
        });
      });
    }
  });

export const InvoiceForm = () => {
  const invoiceNumber = 1;

  const form = useForm<z.infer<typeof InvoiceSchema>>({
    resolver: zodResolver(InvoiceSchema),
    defaultValues: {
      invoiceNumber,
      issuedAt: new Date(),
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

  const products = useFieldArray({ control: form.control, name: "products" });

  const onSubmit = async (data: z.infer<typeof InvoiceSchema>) => {
    const totalExVat = data.products.reduce((acc, item) => acc + item.quantity * item.price, 0);
    const totalIncVat = totalExVat * 1.25;
  };

  const onError = (errors: any) => {
    console.log({ errors });
  };

  const getSum = () => {
    return form.watch("products").reduce((acc, item) => acc + item.quantity * item.price, 0);
  };

  return (
    <section className="w-full p-2 flex justify-center content-center items-center">
      <Form {...form}>
        <form onSubmit={form.handleSubmit(onSubmit, onError)} className="flex flex-col w-5/6 gap-8 py-12">
          {Object.values(form.formState.errors).length > 0 && (
            <Alert variant={"destructive"}>Please fix the errors below</Alert>
          )}
          <div className="flex justify-between">
            <div className="flex w-96 flex-col gap-4">
              <FormField
                control={form.control}
                name="invoiceNumber"
                render={({ field }) => (
                  <FormItem className="w-72">
                    <FormLabel>Invoice Number</FormLabel>
                    <FormControl>
                      <Input placeholder="shadcn" disabled {...field} />
                    </FormControl>
                    <FormDescription>
                      This invoice number is automatically generated and cannot be changed.
                    </FormDescription>
                  </FormItem>
                )}
              />
              <FormField
                control={form.control}
                name="issuedAt"
                render={({ field }) => (
                  <FormItem className="w-72">
                    <FormLabel>Issued At</FormLabel>
                    <FormControl>
                      <DatePicker onSelect={(selected) => field.onChange(selected)} selected={field.value} />
                    </FormControl>
                  </FormItem>
                )}
              />
            </div>
            <div className="flex w-72 flex-col gap-4">
              <h3 className="font-bold text-lg">Customer</h3>
              <FormField
                control={form.control}
                name="customerName"
                render={({ field, fieldState }) => (
                  <FormItem>
                    <FormLabel>Customer Name</FormLabel>
                    <FormControl>
                      <Input {...field} />
                    </FormControl>
                    {fieldState.error && <FormMessage>{fieldState.error.message}</FormMessage>}
                  </FormItem>
                )}
              />
              <FormField
                control={form.control}
                name="customerEmail"
                render={({ field, fieldState }) => (
                  <FormItem>
                    <FormLabel>Customer Email</FormLabel>
                    <FormControl>
                      <Input {...field} />
                    </FormControl>
                    {fieldState.error && <FormMessage>{fieldState.error.message}</FormMessage>}
                  </FormItem>
                )}
              />
              <FormField
                control={form.control}
                name="customerPhone"
                render={({ field, fieldState }) => (
                  <FormItem>
                    <FormLabel>Customer Phone</FormLabel>
                    <FormControl>
                      <Input {...field} />
                    </FormControl>
                    {fieldState.error && <FormMessage>{fieldState.error.message}</FormMessage>}
                  </FormItem>
                )}
              />
              <FormField
                control={form.control}
                name="customerAddress"
                render={({ field, fieldState }) => (
                  <FormItem>
                    <FormLabel>Customer Address</FormLabel>
                    <FormControl>
                      <Input {...field} />
                    </FormControl>
                    {fieldState.error && <FormMessage>{fieldState.error.message}</FormMessage>}
                  </FormItem>
                )}
              />
              <FormField
                control={form.control}
                name="customerCarRegistration"
                render={({ field, fieldState }) => (
                  <FormItem>
                    <FormLabel>Customer CarRegistration</FormLabel>
                    <FormControl>
                      <Input {...field} />
                    </FormControl>
                    {fieldState.error && <FormMessage>{fieldState.error.message}</FormMessage>}
                  </FormItem>
                )}
              />
            </div>
          </div>
          <div className="flex flex-col gap-4">
            <h3 className="font-bold text-lg">Products</h3>
            <Table>
              <TableHeader>
                <TableRow>
                  <TableHead className="w-28">Id</TableHead>
                  <TableHead>Description</TableHead>
                  <TableHead className="w-20">Quantity</TableHead>
                  <TableHead className="w-28">Price</TableHead>
                  <TableHead className="w-28 text-right">Total</TableHead>
                  <TableHead className="w-20" />
                </TableRow>
              </TableHeader>
              <TableBody>
                {products.fields.map((product, index) => (
                  <TableRow key={product.id}>
                    <TableCell>
                      <FormField
                        control={form.control}
                        name={`products.${index}.id`}
                        render={({ field, fieldState }) => (
                          <FormItem>
                            <FormControl>
                              <Input {...field} />
                            </FormControl>
                          </FormItem>
                        )}
                      />
                    </TableCell>
                    <TableCell>
                      <FormField
                        control={form.control}
                        name={`products.${index}.description`}
                        render={({ field, fieldState }) => (
                          <FormItem>
                            <FormControl>
                              <Input {...field} />
                            </FormControl>
                          </FormItem>
                        )}
                      />
                    </TableCell>
                    <TableCell>
                      <FormField
                        control={form.control}
                        name={`products.${index}.quantity`}
                        render={({ field, fieldState }) => (
                          <FormItem>
                            <FormControl>
                              <Input
                                type="number"
                                {...field}
                                onChange={(e) => field.onChange(parseInt(e.target.value))}
                              />
                            </FormControl>
                          </FormItem>
                        )}
                      />
                    </TableCell>
                    <TableCell>
                      <FormField
                        control={form.control}
                        name={`products.${index}.price`}
                        render={({ field, fieldState }) => (
                          <FormItem>
                            <FormControl>
                              <Input
                                type="string"
                                {...field}
                                onChange={(e) => field.onChange(parseInt(e.target.value))}
                              />
                            </FormControl>
                          </FormItem>
                        )}
                      />
                    </TableCell>
                    <TableCell>
                      <Input
                        disabled
                        value={form.watch(`products.${index}.quantity`) * form.watch(`products.${index}.price`)}
                      />
                    </TableCell>
                    <TableCell>
                      <Button variant={"secondary"} onClick={() => products.remove(index)}>
                        X
                      </Button>
                    </TableCell>
                  </TableRow>
                ))}
              </TableBody>
              <TableCaption>
                <Button
                  variant={"secondary"}
                  onClick={() => products.append({ description: "", price: 0, quantity: 0, id: "" })}
                >
                  +
                </Button>
              </TableCaption>
            </Table>
          </div>

          <Separator />

          <div className="flex flex-col gap-2 items-end">
            <div className="flex gap-4 items-center">
              <p className="w-36">Total Excl. Vat</p>
              <div className="w-36">
                <Input disabled value={getSum()} />
              </div>
            </div>
            <div className="flex gap-4 items-center">
              <p className="w-36">Vat</p>
              <div className="w-36">
                <Input disabled value={"25%"} />
              </div>
            </div>
            <div className="flex gap-4 items-center">
              <p className="w-36">Vat Cost</p>
              <div className="w-36">
                <Input disabled value={getSum() * 1.25 - getSum()} />
              </div>
            </div>
            <div className="flex gap-4 items-center">
              <p className="w-36">Total Incl. Vat</p>
              <div className="w-36">
                <Input disabled value={getSum() * 1.25} />
              </div>
            </div>
          </div>

          <Button className="mb-10" type="submit">
            Create Invoice
          </Button>
        </form>
      </Form>
    </section>
  );
};
