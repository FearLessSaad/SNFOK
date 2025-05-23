import React from "react";

function page({ params }: { params: { id: string } }) {
  const { id } = params;
  return (
    <>
      <div className="p-6 space-y-6">
        <h1 className="text-3xl">Pod</h1>
        <h3 className="text-xl">{id}</h3>
      </div>
    </>
  );
}

export default page;
