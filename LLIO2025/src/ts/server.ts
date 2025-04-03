const VITE_API_BASE_URL = import.meta.env.VITE_API_BASE_URL;

const getHeaders = (token = localStorage.getItem("token")) => ({
    "Content-Type": "application/json",
    Authorization: token || "",
})

const handleResponse = async <T>(response: Response, redirectToLoginOn401 = true): Promise<T | undefined> => {
    if (!response.ok) {
        const statusCodeRedirects = {
            500: "/500",
            401: "/",
        }

        if (redirectToLoginOn401 && statusCodeRedirects[response.status]) {
            window.location.href = statusCodeRedirects[response.status]
            ReadableStreamDefaultController
        }

        if (response.status === 404) return undefined as T

        throw new Error(`Error: ${response.status} - ${response.statusText}`)
    }
    return response.json() as Promise<T>
}

async function request<T>(
    method: string,
    url: string,
    body?: any,
    redirectToLoginOn401 = true
): Promise<T> {
    try {
        const options: RequestInit = {
            method,
            credentials: 'include',
            headers: getHeaders(),
        }

        if (body) {
            options.body = JSON.stringify(body)
        }

        const response = await fetch(`${VITE_API_BASE_URL}${url}`, options)
        const data = await handleResponse<T>(response, redirectToLoginOn401)
        return data as T
    } catch (error) {
        console.error(`Error ${method}ing:`, error)
        throw error
    }
}

export const GET = <T>(url: string, redirectToLoginOn401?: boolean): Promise<T> =>
    request<T>("GET", url, undefined, redirectToLoginOn401)

export const POST = <T, R>(url: string, body: T, redirectToLoginOn401?: boolean): Promise<R> =>
    request<R>("POST", url, body, redirectToLoginOn401)

export const DELETE = (url: string): Promise<void> =>
    request<void>("DELETE", url)

export const PUT = <T, R>(url: string, body: T, redirectToLoginOn401?: boolean): Promise<R> =>
    request<R>("PUT", url, body, redirectToLoginOn401)

export const PATCH = <T>(url: string, body: T): Promise<void> =>
    request<void>("PATCH", url, body)