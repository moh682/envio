import { Account } from "@/components/account";
import { AppSidebar } from "@/components/side-bar";
import { SidebarProvider, SidebarInset, SidebarTrigger } from "@/components/ui/sidebar";
import { getOrganization } from "@/domains/organization/actions";
import { redirect } from "next/navigation";
import React from "react";

type Props = Readonly<{
  params: Promise<{
    organizationId: string;
    financialYear: number;
  }>;
  children: React.ReactNode;
}>;

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

export default async function layout({ params, children }: Props) {
  const { organizationId, financialYear } = await params;
  return (
    <SidebarProvider>
      <AppSidebar
        financeYearId="2025"
        financeYears={financeYears}
        organizationId={organizationId}
        financialYear={financialYear}
      />
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
