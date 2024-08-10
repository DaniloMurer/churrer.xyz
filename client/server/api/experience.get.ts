export default defineEventHandler(async (event): Promise<void> => {
	const apiHost = process.env.API_HOST || 'http://localhost:8080';
	return $fetch(`${apiHost}/api/experience`)
})
