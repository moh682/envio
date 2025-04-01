"use client";

import { Button } from "@/components/ui/button";
import { DialogDescription, DialogHeader, DialogTitle } from "@/components/ui/dialog";
import { Input } from "@/components/ui/input";
import { Label } from "@/components/ui/label";
import React, { useState } from "react";

type Props = {
  setCompany: (name: string) => void;
  setStageToCreate: () => void;
};

const defaultCompanies = [
  {
    id: "1",
    name: "FTZ",
  },
  {
    id: "2",
    name: "AD Danmark",
  },
  {
    id: "3",
    name: "Viggo Larson",
  },
];

export const SearchCompany = ({ setCompany, setStageToCreate }: Props) => {
  const [companies, setCompanies] = useState<typeof defaultCompanies>([]);

  const searchCompanyByName = (name: string) => {
    setCompanies(
      defaultCompanies.filter((company) => company.name.toLowerCase().trim().startsWith(name.toLowerCase().trim()))
    );
  };

  return (
    <>
      <DialogHeader>
        <DialogTitle>Search for company</DialogTitle>
        <DialogDescription>
          Select the company the invoice is from. If there is not company in the search result, please create the
          company
        </DialogDescription>
      </DialogHeader>
      <div className="flex flex-col gap-3">
        <Label htmlFor="searchCompanyName">Company name</Label>
        <Input id="searchCompanyName" onChange={(e) => searchCompanyByName(e.target.value)} autoFocus />
      </div>
      {companies.map((company) => (
        <div key={company.id} className="flex border py-1 px-3 rounded-md items-center">
          <span className="text-sm flex-1">{company.name}</span>
          <Button variant="ghost" onClick={() => setCompany(company.name)}>
            Select
          </Button>
        </div>
      ))}

      {companies.length === 0 && (
        <Button type="button" onClick={setStageToCreate}>
          Create company
        </Button>
      )}
    </>
  );
};
