export default defineEventHandler(event =>  {
	const apiHost = process.env.API_HOST || 'http://localhost:8080';
	return $fetch(`${apiHost}/api/technology`)
})
