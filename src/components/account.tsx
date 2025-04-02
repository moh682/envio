import React from "react";
import {
  DropdownMenu,
  DropdownMenuContent,
  DropdownMenuLabel,
  DropdownMenuSeparator,
  DropdownMenuTrigger,
} from "./ui/dropdown-menu";
import { SignOut } from "./sign-out";
import Image from "next/image";

export const Account = () => {
  return (
    <DropdownMenu>
      <DropdownMenuTrigger asChild>
        <Image className="cursor-pointer rounded-full" src="/avatar.jpg" alt="profile_image" height={40} width={40} />
      </DropdownMenuTrigger>
      <DropdownMenuContent className="w-56">
        <DropdownMenuLabel>My Account</DropdownMenuLabel>
        <DropdownMenuSeparator />
        <SignOut />
      </DropdownMenuContent>
    </DropdownMenu>
  );
};
