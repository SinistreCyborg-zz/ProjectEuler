sum = 0
i = 1000

until i == 0
    i -= 1
    sum += i % 3 == 0 || i % 5 == 0 ? i : 0
end

puts sum
