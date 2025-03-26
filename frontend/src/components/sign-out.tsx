"use client";

import React from "react";
import { DropdownMenuItem } from "./ui/dropdown-menu";
import Session from "supertokens-auth-react/recipe/session";
import SuperTokens from "supertokens-auth-react";

export const SignOut = () => {
  return (
    <DropdownMenuItem
      onClick={async () => {
        await Session.signOut();
        SuperTokens.redirectToAuth();
      }}
    >
      Log out
    </DropdownMenuItem>
  );
};
