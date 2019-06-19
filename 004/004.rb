def isP(n)
    s = n.to_s
    return s == s.reverse
end

r = 0
(100..999).each do |a|
    (a..999).each do |b|
        p = a * b
        r = isP(p) && p > r ? p : r
    end
end

puts r
