i = 2
num = 600851475143

until num <= i
    i += 1
    num = (num % i == 0) ? (num / i) : num
end

puts i
