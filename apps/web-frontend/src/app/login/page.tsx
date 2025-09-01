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
	async function onSubmit(values: z.infer<typeof formSchema>) {
		console.log("onSubmit: ", values);

		try {
			const response = await fetch('http://localhost:8080/api/auth/login', {
				method: 'POST',
				headers: {
					'Content-Type': 'application/json',
				},
				body: JSON.stringify({
					email: values.email,
					password: values.password,
				}),
				credentials: "include",
			});

			const data = await response.json();
			console.log('Login response:', data);

			if (!response.ok) {
				console.error('Login failed:', response.status, data);
			} else {
				console.log('Login successful:', data);
			}
		} catch (error) {
			console.error('Login error:', error);
		}
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
					<Button type="submit" className="md:self-end cursor-pointer">Submit</Button>
				</form>
			</Form>
		</div>
	);
}
