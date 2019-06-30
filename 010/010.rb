require "prime"

primes = [*1..2e6].select { |n| Prime.prime?(n) }
puts primes.inject(:+)
