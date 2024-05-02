const URL = import.meta.env.VITE_URL;

export const postEvaluationData = async (evaluationText: string) => {
	let url = `${URL}/submiteval`;

	const response = fetch(url, {
		method: "POST",
		body: JSON.stringify({
			evaluationText: evaluationText,
		}),
	})
		.then((response) => {
			if (!response.ok) {
				throw new Error(`HTTP error! Status: ${response.status}`);
			}
			return response.ok;
		})
		.catch((error) => {
			console.log(error);
		});
	return response;
};
