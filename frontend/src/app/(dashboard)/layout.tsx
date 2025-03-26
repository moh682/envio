import { Account } from "@/components/account";
import { TryRefreshComponent } from "@/components/auth/try-refresh-client";
import { AppSidebar } from "@/components/side-bar";
import { SidebarInset, SidebarProvider, SidebarTrigger } from "@/components/ui/sidebar";
import { getSessionForSSR } from "@/lib/auth";
import { cookies } from "next/headers";
import { redirect } from "next/navigation";

type Props = Readonly<{
  children: React.ReactNode;
}>;

export default async function layout({ children }: Props) {
  const cookiesFromReq = await cookies();
  const cookiesArray: Array<{ name: string; value: string }> = Array.from(cookiesFromReq.getAll()).map(
    ({ name, value }) => ({
      name,
      value,
    })
  );
  const { accessTokenPayload, hasToken, error } = await getSessionForSSR(cookiesArray);

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
  return (
    <SidebarProvider>
      <AppSidebar />
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
