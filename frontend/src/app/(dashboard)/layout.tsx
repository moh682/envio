import { AppSidebar } from "@/components/side-bar";
import { Separator } from "@/components/ui/separator";
import { SidebarInset, SidebarProvider, SidebarTrigger } from "@/components/ui/sidebar";

type Props = Readonly<{
  children: React.ReactNode;
}>;

export default async function layout({ children }: Props) {
  return (
    <SidebarProvider>
      <AppSidebar />
      <SidebarInset>
        <header className="flex h-16 shrink-0 items-center gap-2 border-b px-4">
          <SidebarTrigger className="-ml-1" />
        </header>
        <div className="p-3 self-center w-full xl:max-w-[1280px]">{children}</div>
      </SidebarInset>
    </SidebarProvider>
  );
}
