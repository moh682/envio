import { z } from "zod";

export const InvoiceSchema = z.object({
  invoiceNumber: z.number(),
  issuedAt: z.date(),
  name: z.string(),
  email: z.union([z.string(), z.string().email()]),
  carRegistration: z
    .string()
    .min(1)
    .toLowerCase()
    .transform((value) => value.replaceAll(" ", "")),
  phone: z.string(),
  address: z.string(),
  products: z.array(
    z.object({
      id: z.string(),
      description: z.string().min(1),
      quantity: z.number().positive(),
      price: z.number().positive(),
    })
  ),
});
