export type CreateOrganizationPostResult = {
  id: string;
  name: string;
  invoiceNumberStart: number;
  financialYears: Year[];
};

export type Year = {
  year: number;
};
