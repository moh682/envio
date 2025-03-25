"use client";

import { useFieldArray, useForm } from "react-hook-form";
import { z } from "zod";
import { zodResolver } from "@hookform/resolvers/zod";
import { Form, FormControl, FormDescription, FormField, FormItem, FormLabel, FormMessage } from "@/components/ui/form";
import { Alert } from "@/components/ui/alert";
import { Popover, PopoverContent, PopoverTrigger } from "@/components/ui/popover";
import { Button } from "@/components/ui/button";
import { cn } from "@/lib/utils";
import { ChevronsUpDown, Command } from "lucide-react";
import { CommandEmpty, CommandGroup, CommandInput, CommandItem, CommandList } from "@/components/ui/command";
import { Select, SelectContent, SelectItem, SelectTrigger, SelectValue } from "@/components/ui/select";
import { Separator } from "@/components/ui/separator";
import { Input } from "@/components/ui/input";
import { Table, TableBody, TableCaption, TableCell, TableHead, TableHeader, TableRow } from "@/components/ui/table";

const VAT_RATE = 0.25;
const ExpenseSchema = z
  .object({
    company: z.string(),
    issuedAt: z.date(),
    paidAt: z.date(),
    paidWith: z.number(),
    vatRate: z.number().readonly().default(VAT_RATE),
    products: z.array(
      z
        .object({
          serial: z.string(),
          description: z.string(),
          quantity: z.number().positive().default(0),
          unitPrice: z.number().positive().default(0),
        })
        .transform((data) => {
          const totalExclVat = data.quantity * data.unitPrice;
          const totalInclVat = totalExclVat > 0 ? totalExclVat * (1 + VAT_RATE) : 0;
          return {
            ...data,
            total: data.quantity * data.unitPrice,
            totalInclVat,
            totalExclVat,
          };
        })
    ),
  })
  .transform((data) => {
    const totalExclVat = data.products.reduce((acc, item) => acc + item.quantity * item.unitPrice, 0);
    const totalInclVat = totalExclVat > 0 ? totalExclVat * (1 + VAT_RATE) : 0;
    const vatAmount = totalInclVat - totalExclVat;

    return {
      ...data,
      totalInclVat,
      totalExclVat,
      vatAmount,
    };
  });

export const ExponseForm = () => {
  const form = useForm<z.infer<typeof ExpenseSchema>>({
    resolver: zodResolver(ExpenseSchema),
    defaultValues: {
      issuedAt: new Date(),
      paidAt: new Date(),
      company: "",
      paidWith: 0,

      products: [
        {
          serial: "",
          description: "",
          quantity: 0,
          unitPrice: 0,
        },
      ],
    },
  });

  const products = useFieldArray({ control: form.control, name: "products" });

  const getSum = () => {
    return form.watch("products").reduce((acc, item) => acc + item.quantity * item.unitPrice, 0);
  };

  const onSubmit = async (data: z.infer<typeof ExpenseSchema>) => {
    console.log(data);
  };

  const onError = (errors: any) => {
    console.log({ errors });
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
                name="company"
                render={({ field }) => (
                  <FormItem className="relative w-72">
                    <FormLabel>Company</FormLabel>

                    <Popover>
                      <PopoverTrigger asChild>
                        <FormControl>
                          <Button
                            type="button"
                            variant="outline"
                            role="combobox"
                            className={cn("w-[200px] justify-between", !field.value && "text-muted-foreground")}
                          >
                            asdasd
                            <ChevronsUpDown className="ml-2 h-4 w-4 shrink-0 opacity-50" />
                          </Button>
                        </FormControl>
                      </PopoverTrigger>
                      <PopoverContent className="w-[200px]">
                        <Command>
                          <CommandInput placeholder="Search language..." />
                          <CommandList>
                            <CommandEmpty>No language found.</CommandEmpty>
                            <CommandGroup>
                              <CommandItem key={"add-button"}>
                                <Button>Add</Button>
                              </CommandItem>
                            </CommandGroup>
                          </CommandList>
                        </Command>
                      </PopoverContent>
                    </Popover>

                    <FormDescription>Company name the expense is from</FormDescription>
                  </FormItem>
                )}
              />
              <div className="flex gap-2">
                <FormField
                  control={form.control}
                  name="issuedAt"
                  render={({ field }) => (
                    <FormItem className="w-72">
                      <FormLabel>Issued At</FormLabel>
                      <FormControl>
                        {/* <DatePicker onSelect={(selected) => field.onChange(selected)} selected={field.value} /> */}
                      </FormControl>
                    </FormItem>
                  )}
                />
                <FormField
                  control={form.control}
                  name="paidAt"
                  render={({ field }) => (
                    <FormItem className="w-72">
                      <FormLabel>Paid At</FormLabel>
                      <FormControl>
                        {/* <DatePicker onSelect={(selected) => field.onChange(selected)} selected={field.value} /> */}
                      </FormControl>
                    </FormItem>
                  )}
                />
              </div>
            </div>
            <div className="flex w-72 flex-col gap-4">
              <h3 className="font-bold text-lg">Customer</h3>
              <FormField
                control={form.control}
                name="paidWith"
                render={({ field, fieldState }) => (
                  <FormItem>
                    <FormLabel>Paid With</FormLabel>
                    <FormControl>
                      <Select>
                        <SelectTrigger className="w-[180px]">
                          <SelectValue placeholder={field.value} />
                        </SelectTrigger>
                        <SelectContent>
                          <SelectItem value="set value">Set Value</SelectItem>
                        </SelectContent>
                      </Select>
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
                  <TableHead className="w-28">Serial</TableHead>
                  <TableHead>Description</TableHead>
                  <TableHead className="w-20">Quantity</TableHead>
                  <TableHead className="w-28">Price</TableHead>
                  <TableHead className="w-28 text-right">Total</TableHead>
                  <TableHead className="w-20" />
                </TableRow>
              </TableHeader>
              <TableBody>
                {products.fields.map((product, index) => (
                  <TableRow key={product.id + index}>
                    <TableCell>
                      <FormField
                        control={form.control}
                        name={`products.${index}.serial`}
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
                        name={`products.${index}.unitPrice`}
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
                        value={form.watch(`products.${index}.quantity`) * form.watch(`products.${index}.unitPrice`)}
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
                  onClick={() =>
                    products.append({
                      totalInclVat: 0,
                      totalExclVat: 0,
                      unitPrice: 0,
                      description: "",
                      total: 0,
                      quantity: 0,
                      serial: "",
                    })
                  }
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
            Create Expense
          </Button>
        </form>
      </Form>
    </section>
  );
};
