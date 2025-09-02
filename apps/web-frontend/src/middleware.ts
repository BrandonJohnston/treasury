import { NextResponse } from 'next/server';
import type { NextRequest } from 'next/server';

export function middleware(request: NextRequest) {
	const sessionToken = request.cookies.get('session-name')?.value;

	console.log("sessionToken", sessionToken);

	if (!sessionToken) {
		return NextResponse.redirect(new URL('/login', request.url));
	}

	if (request.nextUrl.pathname === '/login' && sessionToken) {
		return NextResponse.redirect(new URL('/dashboard', request.url));
	}

	return NextResponse.next();
}

export const config = {
	matcher: ['/login', '/dashboard', '/dashboard/:path*'], // apply middleware only to these routes
};
