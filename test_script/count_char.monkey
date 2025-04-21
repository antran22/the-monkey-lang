fn count_char(str) {
  if (len(str) == 0) {
    return {}
  }

  let prev = count_char(tail(str)) 
  let currChar = first(str)

  if (prev[currChar] == null) {
    return prev + { currChar: 1 }
  } else {
    return prev + { currChar: prev[currChar] + 1 }
  }
}

print(count_char("hello_world"))
