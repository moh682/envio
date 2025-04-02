import "./globals.css";

import { Button } from "@/components/ui/button";
import { Card, CardContent, CardFooter, CardHeader, CardTitle } from "@/components/ui/card";
import Link from "next/link";

export default function NotFound() {
  return (
    <div className="h-lvh flex items-center justify-center">
      <Card>
        <CardHeader>
          <CardTitle className="text-center">Not Found</CardTitle>
        </CardHeader>
        <CardContent>
          <p>Could not find requested resource</p>
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
