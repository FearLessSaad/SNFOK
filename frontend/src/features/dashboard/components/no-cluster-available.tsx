import { Button } from "@/components/ui/button";
import {
  Dialog,
  DialogTrigger,
  DialogContent,
  DialogHeader,
  DialogTitle,
  DialogDescription,
  DialogFooter,
} from "@/components/ui/dialog";
import { Textarea } from "@/components/ui/textarea"

import { Input } from "@/components/ui/input";
import React from "react";
import { zodResolver } from "@hookform/resolvers/zod";
import { z } from "zod";
import {
  Form,
  FormControl,
  FormDescription,
  FormField,
  FormItem,
  FormLabel,
  FormMessage,
} from "@/components/ui/form";
import { useForm } from "react-hook-form";
import RegisterClusterSchema from "../schema/register-cluster-schema";

function NoClusterAvailable() {
  const form = useForm<z.infer<typeof RegisterClusterSchema>>({
    resolver: zodResolver(RegisterClusterSchema),
    defaultValues: {
      master_ip: "",
      agent_port: 1024,
      description: "",
    },
  });

  const handleSubmit = (values: z.infer<typeof RegisterClusterSchema>) => {};

  return (
    <div className="flex justify-center items-center flex-col mt-24 w-full">
      <div className="flex flex-col items-center justify-center gap-6 bg-background rounded-lg p-8 shadow-md">
        <div className="bg-muted rounded-full p-4 mb-2">
          <svg
            xmlns="http://www.w3.org/2000/svg"
            className="h-12 w-12 text-primary"
            fill="none"
            viewBox="0 0 24 24"
            stroke="currentColor"
          >
            <path
              strokeLinecap="round"
              strokeLinejoin="round"
              strokeWidth={2}
              d="M12 4v16m8-8H4"
            />
          </svg>
        </div>
        <h2 className="text-2xl font-bold text-foreground mb-1">
          No Kubernetes Cluster Found
        </h2>
        <p className="text-muted-foreground text-center max-w-md">
          You have not registered any Kubernetes clusters yet. To get started,
          create a new cluster and begin managing your infrastructure securely
          and efficiently.
        </p>

        <Dialog>
          <DialogTrigger asChild>
            <Button size="lg" type="button">
              <svg
                xmlns="http://www.w3.org/2000/svg"
                className="h-5 w-5 mr-2"
                fill="none"
                viewBox="0 0 24 24"
                stroke="currentColor"
              >
                <path
                  strokeLinecap="round"
                  strokeLinejoin="round"
                  strokeWidth={2}
                  d="M12 4v16m8-8H4"
                />
              </svg>
              Create New Cluster
            </Button>
          </DialogTrigger>
          <DialogContent>
            <DialogHeader>
              <DialogTitle>Create a New Kubernetes Cluster</DialogTitle>
              <DialogDescription>
                Fill in the details below to create a new cluster.
              </DialogDescription>
            </DialogHeader>
            {/* Replace this with a form if needed */}
            <Form {...form}>
              <form onSubmit={form.handleSubmit(handleSubmit)}>
                <div className="space-y-4">
                  <FormField
                    control={form.control}
                    name="master_ip"
                    render={({ field }) => (
                      <FormItem>
                        <FormLabel>Master IP</FormLabel>
                        <FormControl>
                          <Input
                            type="text"
                            placeholder="Enter master IP address"
                            {...field}
                            className="w-full p-2 border rounded-md"
                          />
                        </FormControl>
                        <FormDescription>
                          Enter the master node IP address.
                        </FormDescription>
                        <FormMessage />
                      </FormItem>
                    )}
                  />

                  <FormField
                    control={form.control}
                    name="agent_port"
                    render={({ field }) => (
                      <FormItem>
                        <FormLabel>Agent Port</FormLabel>
                        <FormControl>
                          <Input
                            type="number"
                            placeholder="1024"
                            {...field}
                            className="w-full p-2 border rounded-md"
                          />
                        </FormControl>
                        <FormDescription>
                          Enter the agent port (1024-65535).
                        </FormDescription>
                        <FormMessage />
                      </FormItem>
                    )}
                  />

                  <FormField
                    control={form.control}
                    name="description"
                    render={({ field }) => (
                      <FormItem>
                        <FormLabel>Description</FormLabel>
                        <FormControl>
                          <Textarea
                            placeholder="Cluster description"
                            {...field}
                            className="w-full p-2 border rounded-md"
                            rows={4}
                          />
                        </FormControl>
                        <FormDescription>
                          Enter a description for your cluster (min 15 characters).
                        </FormDescription>
                        <FormMessage />
                      </FormItem>
                    )}
                  />
                </div>
                <DialogFooter className="pt-4">
                  <Button type="submit">Create Cluster</Button>
                </DialogFooter>
              </form>
            </Form>
          </DialogContent>
        </Dialog>
      </div>
    </div>
  );
}

export default NoClusterAvailable;
