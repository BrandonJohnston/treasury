"use client";
import { z } from 'zod';
import { zodResolver } from "@hookform/resolvers/zod";
import { useForm } from "react-hook-form";
import { Button } from "@/components/ui/button";
import { Form, FormControl, FormField, FormItem, FormLabel, FormMessage } from "@/components/ui/form";
import { Input } from "@/components/ui/input";

const formSchema = z.object({
	email: z.email().max(50),
	password: z.string().min(5, "Password must be more than 5 characters.").max(50, "Password must be less than 50 characters."),
});

export default function Login() {
	// Define the form
	const form = useForm<z.infer<typeof formSchema>>({
		mode: "onBlur",
		resolver: zodResolver(formSchema),
		defaultValues: {
			email: "",
			password: "",
		},
	});

	// Define submit handler
	function onSubmit(values: z.infer<typeof formSchema>) {
		console.log("onSubmit: ", values);
	}

	return (
		<div className="login-form items-center justify-items-center w-full md:w-1/2 md:mx-auto p-10">
			<Form {...form}>
				<form onSubmit={form.handleSubmit(onSubmit)} className="flex flex-col w-full max-w-100">
					<FormField control={form.control}
							   name="email"
							   render={({ field }) => (
								   <FormItem className="mb-6">
									   <FormLabel>Email</FormLabel>
									   <FormControl>
										   <Input placeholder="you@example.com" {...field} />
									   </FormControl>
									   <FormMessage />
								   </FormItem>
							   )}
				    />
					<FormField control={form.control}
							   name="password"
							   render={({ field }) => (
								   <FormItem className="mb-6">
									   <FormLabel>Password</FormLabel>
									   <FormControl>
										   <Input type="password" placeholder="***" {...field} />
									   </FormControl>
									   <FormMessage />
								   </FormItem>
							   )}
					/>
					<Button type="submit" className="md:self-end cursor-pointer" disabled={!form.formState.isValid}>Submit</Button>
				</form>
			</Form>
		</div>
	);
}
