import { error } from "@sveltejs/kit";

/** @type {import('./$types').PageLoad} */
export async function load({ params }) {
	let questionNumber = params;

	if (questionNumber) return questionNumber;

	error(404, "Not found");
}
