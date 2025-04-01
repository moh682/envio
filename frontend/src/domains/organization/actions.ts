"use server";

import { rawFetch } from "@/lib/fetch";
import { CreateOrganizationPost } from "./types";
import { redirect } from "next/navigation";
import { CreateOrganizationPostResult } from "./type";

export const getOrganization = async (): Promise<CreateOrganizationPostResult | undefined> => {
  const response = await rawFetch("/organization", {
    method: "GET",
  });

  if (response.status === 404) return undefined;

  const data = await response.json();
  console.log(data);

  return data;
};

export const createOrganization = async (values: CreateOrganizationPost) => {
  const response = await rawFetch("/organization", {
    method: "POST",
    body: JSON.stringify(values),
  });

  if (response.status !== 201) return { error: { message: "SOMETHING WENT WRONG" } };

  const body: CreateOrganizationPostResult = await response.json();

  const financialYear = body.financialYears.sort((a, b) => a.year - b.year)[0].year;

  redirect(`/${body.id}/${financialYear}`);
};
