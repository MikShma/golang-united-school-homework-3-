package homework

func average(input [15]float32) (result float32) {
	//Place your code here
	var summ, count float32 = 0.0, 0.0
	for _, n := range input {
		if n == 0 {
			continue
		}
		summ += n
		count += 1
	}
	result = summ / count
	return
}
