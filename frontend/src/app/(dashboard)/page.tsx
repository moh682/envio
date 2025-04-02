import { TryRefreshComponent } from "@/components/auth/try-refresh-client";
import { getSessionSSR } from "@/lib/auth";
import { redirect } from "next/navigation";
import { getOrganization } from "../../domains/organization/actions";
import { CreateOrganization } from "@/components/create-organization";

export default async function Page() {
  const { accessTokenPayload, hasToken, error } = await getSessionSSR();

  if (error) {
    return <div>Something went wrong while trying to get the session. Error - {error.message}</div>;
  }

  if (accessTokenPayload === undefined) {
    if (!hasToken) {
      return redirect("/auth");
    }

    return <TryRefreshComponent key={Date.now()} />;
  }

  const organization = await getOrganization();

  if (!organization) return <CreateOrganization />;

  redirect(`/${organization.id}/${organization.financialYears[0].year}`);
}
