export default defineEventHandler(async (event) => {
	const apiHost = process.env.API_HOST || 'http://localhost:8080';

	return $fetch(`${apiHost}/api/experience`, {
		method: 'GET',
		headers: {
			'Content-Type': 'application/json',
			'Accept': 'application/json'
		}
	});
})
