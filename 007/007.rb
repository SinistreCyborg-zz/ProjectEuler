def isP(n)
    (2..(n-1)).each do |a|
        if n % a == 0
            return false
        end
    end
    return true
end

primes = []
i = 1

until primes.length > 1e4
    i += 1
    if isP(i)
        primes << i
    end
end

puts primes.last
