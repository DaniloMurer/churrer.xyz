export default defineEventHandler(async (event) => {
	const apiHost = process.env.API_HOST || 'http://localhost:8080';
	const requestBody = await readBody(event);
	const authenticationToken = event.headers.get('Authorization');
	if (!authenticationToken) {
		setResponseStatus(event, 401);
		return {"message": "Invalid credentials"};
	}
	return $fetch(`${apiHost}/api/technology`, {
		method: 'POST',
		headers: {
			'Content-Type': 'application/json',
			'Accept': 'application/json',
			'Authorization': `${authenticationToken}`
		},
		body: JSON.stringify(requestBody)
	})
})
