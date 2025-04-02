import { cookies } from "next/headers";

export const rawFetch = async (url: string, options: RequestInit) => {
  const cookiesFromReq = await cookies();

  return await fetch(`${process.env.API_URL}${url}`, {
    ...options,
    headers: {
      cookie: cookiesFromReq
        .getAll()
        .map((value) => `${value.name}=${value.value}`)
        .join("; "),
    },
  });
};
