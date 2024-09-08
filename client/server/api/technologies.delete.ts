export default defineEventHandler((event) => {
	const id = getRouterParam(event, 'id');
	console.log(`to delete id: ${id}`);
})
