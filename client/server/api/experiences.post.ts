export default defineEventHandler(async (event) => {
	const requestBody = await readBody(event);
	const apiHost = process.env.API_HOST || 'http:/localhost:8080';
 	const authorizationHeader = event.headers.get('Authorization');
	if (!authorizationHeader) {
		setResponseStatus(event, 401);
		return {"message": "Invalid credentials"};
	}
	return $fetch(`${apiHost}/api/experience`, {
		method: 'POST',
		headers: {
			'Content-Type': 'application/json',
			'Accept': 'application/json',
			'Authorization': `${authorizationHeader}`,
		},
		body: JSON.stringify(requestBody)
	});
})
