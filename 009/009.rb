def triplet(a, b, c)
    return c**2 == b**2 + a**2
end

(1..1000).each do |c|
    (1..c-1).each do |b|
        (1..b-1).each do |a|
            if triplet(a, b, c) && a+b+c==1000
                puts "#{a}² + #{b}² = #{c}²\n#{a} * #{b} * #{c} = #{a*b*c}"
            end
        end
    end
end
