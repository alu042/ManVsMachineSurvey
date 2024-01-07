export const convertAnswerData = (value: string): number => {
	let result = 0;

	switch (value) {
		case "Veldig dårlig":
			result = 1;
			break;
		case "Dårlig":
			result = 2;
			break;
		case "Nøytral":
			result = 3;
			break;
		case "Bra":
			result = 4;
			break;
		case "Veldig Bra":
			result = 5;
			break;
		default:
			console.log("Wrong input");
			result = 0;
			break;
	}

	return result;
};
