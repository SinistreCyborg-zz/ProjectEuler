def calc(s)
    return s.at(s.length - 1) + s.at(s.length - 2)
end

fib = Array[1, 2]
while fib.at(fib.length - 1) < 4e6
    fib << calc(fib)
end

fib.delete_if { |n| n % 2 != 0 }
sum = fib.reduce(:+)
puts sum
