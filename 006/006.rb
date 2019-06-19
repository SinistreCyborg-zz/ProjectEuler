squares = (1..100).map { |n| n ** 2 }
sum = (1..100).inject(:+)

puts (sum ** 2) - squares.inject(:+)
