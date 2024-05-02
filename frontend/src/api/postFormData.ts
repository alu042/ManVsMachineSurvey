const URL = import.meta.env.VITE_URL;

export const postFormData = async (
	respondentID: number,
	allFormAnswers: string
) => {
	let url = `${URL}/submitanswers`;

	const response = fetch(url, {
		method: "POST",
		body: JSON.stringify({
			respondentID: respondentID,
			allFormAnswers: allFormAnswers,
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
