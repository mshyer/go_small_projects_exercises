const isPrime = (n) => {
  const sqrt = Math.sqrt(n);
  for (let idx = 2; idx <= sqrt; idx++) {
    if (n % idx === 0) {
      return false
    }
  }
  return true;
};
console.log(isPrime(1))
console.log(isPrime(2))
console.log(isPrime(3))
console.log(isPrime(4))
console.log(isPrime(10))
console.log(isPrime(27))
console.log(isPrime(17))
console.log(isPrime(107))

