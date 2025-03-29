import { Account } from "@/components/account";
import { TryRefreshComponent } from "@/components/auth/try-refresh-client";
import { AppSidebar } from "@/components/side-bar";
import { SidebarInset, SidebarProvider, SidebarTrigger } from "@/components/ui/sidebar";
import { getSessionSSR } from "@/lib/auth";
import { redirect } from "next/navigation";
import { getOrganization } from "../domains/organization/service";
import { Dialog, DialogContent, DialogDescription, DialogHeader, DialogTitle } from "@/components/ui/dialog";
import { CreateOrganization } from "@/components/create.organization";

type Props = Readonly<{
  children: React.ReactNode;
}>;

export default async function layout({ children }: Props) {
  const { accessTokenPayload, hasToken, error } = await getSessionSSR();

  if (error) {
    return <div>Something went wrong while trying to get the session. Error - {error.message}</div>;
  }

  // `accessTokenPayload` will be undefined if it the session does not exist or has expired
  if (accessTokenPayload === undefined) {
    if (!hasToken) {
      /**
       * This means that the user is not logged in. If you want to display some other UI in this
       * case, you can do so here.
       */
      return redirect("/auth");
    }

    /**
     * This means that the session does not exist but we have session tokens for the user. In this case
     * the `TryRefreshComponent` will try to refresh the session.
     *
     * To learn about why the 'key' attribute is required refer to: https://github.com/supertokens/supertokens-node/issues/826#issuecomment-2092144048
     */
    return <TryRefreshComponent key={Date.now()} />;
  }

  const organizaiton = await getOrganization();

  console.log(organizaiton);

  if (!organizaiton) return <CreateOrganization />;

  const financeYears = [
    {
      id: "2025",
      year: "2025",
    },
    {
      id: "2024",
      year: "2024",
    },
    {
      id: "2023",
      year: "2023",
    },
  ];

  return (
    <SidebarProvider>
      <AppSidebar financeYearId="2025" financeYears={financeYears} />
      <SidebarInset>
        <header className="flex justify-between h-16 shrink-0 items-center gap-2 border-b px-4">
          <SidebarTrigger className="-ml-1" />
          <Account />
        </header>
        <div className="p-3 self-center w-full xl:max-w-[1280px]">{children}</div>
      </SidebarInset>
    </SidebarProvider>
  );
}
