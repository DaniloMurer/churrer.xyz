import {readBody} from "#imports";
import {base64urlEncode} from "iron-webcrypto";

/**
 * Handle call to backend for telemetry from client (BFF principle)
 */
export default defineEventHandler(async (event) => {
	const apiHost = process.env.API_HOST || 'http://localhost:8080';
	const apiUsername = process.env.API_USERNAME || 'admin';
	const apiPassword = process.env.API_PASSWORD || 'testico';
	const body = await readBody(event);
	$fetch(`${apiHost}/api/telemetry`, {
		method: 'POST',
		headers: {
			'Content-Type': 'application/json',
			'Accept': 'application/json',
			'Authorization': `Basic ${base64urlEncode(apiUsername.concat(':').concat(apiPassword))}`,
		},
		body: JSON.stringify(body)
	}).then((telemetry) => {
		console.trace('successfully created telemetry entry', telemetry);
	}).catch((error) => {
		console.error(error, JSON.stringify(body));
	})
})
