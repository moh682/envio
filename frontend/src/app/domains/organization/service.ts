import { rawFetch } from "@/lib/fetch";

export const getOrganization = async () => {
  const response = await rawFetch("/organization", {
    method: "GET",
  });

  if (response.status === 404) return undefined;

  return await response.json();
};
