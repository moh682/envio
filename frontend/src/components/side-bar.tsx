import { Check, ChevronDown, Home, Inbox, Plus, ReceiptText, Settings } from "lucide-react";

import {
  Sidebar,
  SidebarContent,
  SidebarGroup,
  SidebarGroupContent,
  SidebarGroupLabel,
  SidebarHeader,
  SidebarMenu,
  SidebarMenuButton,
  SidebarMenuItem,
} from "@/components/ui/sidebar";
import { Dialog, DialogContent, DialogTrigger } from "./ui/dialog";
import {
  DropdownMenu,
  DropdownMenuContent,
  DropdownMenuItem,
  DropdownMenuLabel,
  DropdownMenuSeparator,
  DropdownMenuTrigger,
} from "./ui/dropdown-menu";
import Link from "next/link";

// Menu items.
const items = [
  {
    title: "Home",
    url: "/",
    icon: Home,
  },
  {
    title: "Invoices",
    url: "/invoices",
    icon: Inbox,
  },
  {
    title: "Expenses",
    url: "/expenses",
    icon: ReceiptText,
  },
  {
    title: "Settings",
    url: "/settings",
    icon: Settings,
  },
];

type FinanceYear = {
  id: string;
  year: string;
};

type Props = {
  financeYears: FinanceYear[];
  financeYearId: string;
};

export function AppSidebar({ financeYears, financeYearId }: Props) {
  return (
    <Sidebar>
      <SidebarHeader>
        <SidebarMenu>
          <SidebarMenuItem>
            <Dialog>
              <DropdownMenu>
                <DropdownMenuTrigger asChild>
                  <SidebarMenuButton>
                    {financeYears.find((financeYear) => financeYear.id === financeYearId)?.year}
                    <ChevronDown className="ml-auto" />
                  </SidebarMenuButton>
                </DropdownMenuTrigger>
                <DropdownMenuContent className="w-[--radix-popper-anchor-width]">
                  <DropdownMenuLabel className="text-xs text-muted-foreground">Years</DropdownMenuLabel>
                  {financeYears.map((financeYear) => (
                    <Link href={`/${financeYear.id}/`} key={financeYear.id}>
                      <DropdownMenuItem>
                        {financeYear.year}
                        {financeYear.id === financeYearId && <Check className="ml-auto" />}
                      </DropdownMenuItem>
                    </Link>
                  ))}
                  <DropdownMenuSeparator />
                  <DialogTrigger asChild>
                    <DropdownMenuItem>
                      <div className="flex size-6 items-center justify-center rounded-sm border">
                        <Plus className="size-4" />
                      </div>
                      <span className="font-medium text-muted-foreground">Add finance year</span>
                    </DropdownMenuItem>
                  </DialogTrigger>
                </DropdownMenuContent>
              </DropdownMenu>
              <DialogContent>This is where you create a new financial year</DialogContent>
            </Dialog>
          </SidebarMenuItem>
        </SidebarMenu>
      </SidebarHeader>
      <SidebarContent>
        <SidebarGroup>
          <SidebarGroupContent>
            <SidebarMenu>
              {items.map((item) => (
                <SidebarMenuItem key={item.title}>
                  <SidebarMenuButton asChild>
                    <a href={item.url}>
                      <item.icon />
                      <span>{item.title}</span>
                    </a>
                  </SidebarMenuButton>
                </SidebarMenuItem>
              ))}
            </SidebarMenu>
          </SidebarGroupContent>
        </SidebarGroup>
      </SidebarContent>
    </Sidebar>
  );
}
