"use client";
import { Button } from "@/components/ui/button";
import { Card, CardContent, CardFooter, CardHeader, CardTitle } from "@/components/ui/card";
import Link from "next/link";

export default function GlobalError() {
  return (
    <div className="h-lvh flex items-center justify-center">
      <Card>
        <CardHeader>
          <CardTitle className="text-center">Something went wrong</CardTitle>
        </CardHeader>
        <CardContent>
          <p>Please contact the admin of the page</p>
        </CardContent>
        <CardFooter>
          <Button asChild>
            <Link className="w-full" href="/">
              Return Home
            </Link>
          </Button>
        </CardFooter>
      </Card>
    </div>
  );
}
