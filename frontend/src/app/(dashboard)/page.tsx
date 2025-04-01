import { getOrganization } from "../../domains/organization/actions";
import { redirect } from "next/navigation";

export default async function () {
  const organization = await getOrganization();
  const financialYear = organization?.financialYears.sort((a, b) => a.year - b.year)[0].year;
  redirect(`/${financialYear}/`);
}
