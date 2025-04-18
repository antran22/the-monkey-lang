fn fibonacci(i) {
  if (i < 0) {
    return -1
  } 
  if (i == 1) {
    return 0
  }
  if (i == 2) {
    return 1
  }

  return fibonacci(i - 1) + fibonacci(i - 2)
}

print(fibonacci(7))
