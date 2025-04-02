import { z } from "zod";

export const expenseSchema = z.object({
  issuedAt: z.date(),
  paymentOption: z.string().min(1),
  account: z.string().min(1),
  isVat: z.boolean(),
  amount: z.string().min(1),
  company: z.string().min(1),
});
