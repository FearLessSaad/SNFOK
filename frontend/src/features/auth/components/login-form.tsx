"use client"

import { cn } from "@/lib/utils"
import { Button } from "@/components/ui/button"
import {
    Card,
    CardContent,
    CardDescription,
    CardHeader,
    CardTitle,
} from "@/components/ui/card"
import { Input } from "@/components/ui/input"
import { zodResolver } from "@hookform/resolvers/zod"
import { z } from "zod"
import loginSchema from "../schema/login-schema"
import { Form, FormControl, FormDescription, FormField, FormItem, FormLabel, FormMessage } from "@/components/ui/form"
import { useForm } from "react-hook-form"

export function LoginForm({
    className,
    ...props
}: React.ComponentPropsWithoutRef<"div">) {

    const form = useForm<z.infer<typeof loginSchema>>({
        resolver: zodResolver(loginSchema),
        defaultValues: {
            email: "",
            token: "",
        },
    })




    return (
        <div className={cn("flex flex-col gap-6", className)} {...props}>
            <Card>
                <CardHeader>
                    <CardTitle className="text-2xl">Login</CardTitle>
                    <CardDescription>
                        Enter your email below to login to your account
                    </CardDescription>
                </CardHeader>
                <CardContent>
                    <Form {...form}>
                        <form>
                            <div className="flex flex-col gap-6">
                                <div className="grid gap-2">
                                    <FormField
                                        control={form.control}
                                        name="email"
                                        render={({ field }) => (
                                            <FormItem>
                                                <FormLabel>Email</FormLabel>
                                                <FormControl>
                                                <Input
                                                    placeholder="m@example.com"
                                                    required
                                                    {...field}
                                                />
                                                </FormControl>
                                                <FormDescription>
                                                    Emter your registred email address.
                                                </FormDescription>
                                                <FormMessage/>
                                            </FormItem>
                                        )}
                                    />

                                </div>
                                <div className="grid gap-2">
                                <FormField
                                        control={form.control}
                                        name="token"
                                        render={({ field }) => (
                                            <FormItem>
                                                <FormLabel>Token</FormLabel>
                                                <FormControl>
                                                <Input
                                                    placeholder="Enter your token here"
                                                    type="password"
                                                    required
                                                    {...field}
                                                />
                                                </FormControl>
                                                <FormDescription>
                                                    Enter your access token to login
                                                </FormDescription>
                                                <FormMessage/>
                                            </FormItem>
                                        )}/>
                                </div>
                                <Button type="submit" className="w-full">
                                    Login
                                </Button>
                            </div>
                        </form>
                    </Form>
                </CardContent>
            </Card>
        </div>
    )
}
