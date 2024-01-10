export const postEvaluationData = (evaluationText: string) => {
	let url = "http://localhost:8080/submiteval";

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
