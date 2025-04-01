import { z } from "zod";

export const createOrganizationSchema = z.object({
  name: z.string().min(1).max(255),
  invoiceNumberStart: z.string().min(1),
});
