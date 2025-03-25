import React from "react";

type Props = {
  params: Promise<{ id: string }>;
};

async function Page({ params }: Props) {
  const { id } = await params;
  return <div>Expense Id {id}</div>;
}

export default Page;
