"use client";

import { Button } from "@/components/ui/button";
import { FormField, FormItem, FormControl } from "@/components/ui/form";
import { Input } from "@/components/ui/input";
import { TableHeader, TableRow, TableHead, TableBody, TableCell, TableCaption, Table } from "@/components/ui/table";
import React from "react";
import { useFieldArray, UseFormReturn } from "react-hook-form";
import { z } from "zod";
import { InvoiceSchema } from "./form";

type Props = {
  form: UseFormReturn<z.infer<typeof InvoiceSchema>>;
};

export const Products = ({ form }: Props) => {
  const products = useFieldArray({ control: form.control, name: "products" });
  return (
    <div className="col-span-6">
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
          {products.fields.map((product, index, arr) => (
            <TableRow key={product.id}>
              <TableCell>
                <FormField
                  control={form.control}
                  name={`products.${index}.id`}
                  render={({ field }) => (
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
                  render={({ field }) => (
                    <FormItem>
                      <FormControl>
                        <Input {...field} autoFocus />
                      </FormControl>
                    </FormItem>
                  )}
                />
              </TableCell>
              <TableCell>
                <FormField
                  control={form.control}
                  name={`products.${index}.quantity`}
                  render={({ field }) => (
                    <FormItem>
                      <FormControl>
                        <Input type="number" {...field} onChange={(e) => field.onChange(parseInt(e.target.value))} />
                      </FormControl>
                    </FormItem>
                  )}
                />
              </TableCell>
              <TableCell>
                <FormField
                  control={form.control}
                  name={`products.${index}.price`}
                  render={({ field }) => (
                    <FormItem>
                      <FormControl>
                        <Input type="string" {...field} onChange={(e) => field.onChange(parseInt(e.target.value))} />
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
                {arr.length > 1 && (
                  <Button variant={"secondary"} onClick={() => products.remove(index)}>
                    X
                  </Button>
                )}
              </TableCell>
            </TableRow>
          ))}
        </TableBody>
        <TableCaption>
          <Button
            variant={"secondary"}
            type="button"
            onClick={() =>
              products.append(
                { description: "", price: 0, quantity: 0, id: "" },
                {
                  focusName: `products.${products.fields.length}.description`,
                }
              )
            }
          >
            +
          </Button>
        </TableCaption>
      </Table>
    </div>
  );
};
