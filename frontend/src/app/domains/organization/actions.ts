"use server";

import { rawFetch } from "@/lib/fetch";
import { CreateOrganizationPost } from "./types";
import { redirect } from "next/navigation";

export const getOrganization = async () => {
  const response = await rawFetch("/organization", {
    method: "GET",
  });

  if (response.status === 404) return undefined;

  return await response.json();
};

export const createOrganization = async (values: CreateOrganizationPost) => {
  const response = await rawFetch("/organization", {
    method: "POST",
    body: JSON.stringify(values),
  });

  if (response.status !== 201) return { error: { message: "SOMETHING WENT WRONG" } };
  redirect("/");
};
