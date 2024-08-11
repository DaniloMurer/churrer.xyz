import {readBody} from "#imports";

export default defineEventHandler(async (event) : Promise<void> => {
	const apiHost = process.env.API_HOST || 'http://localhost:8080';
	const body = await readBody(event);
	return $fetch(`${apiHost}/api/authentication`, {
		method: 'POST',
		headers: {
			'Content-Type': 'application/json',
			'Accept': 'application/json',
		},
		body: JSON.stringify(body)
	})
})
