#!/usr/bin/env ruby

require 'json'
require 'stringio'

# Complete the repeatedString function below.
def repeatedString(s, n)
  remainder_matches = 0

  length    = s.length
  puts 'length: %d' % [length]

  matches   = s.scan('a').count
  puts 'matches: %d' % [matches]

  quotient  = n / length
  puts 'quotient: %d' % [quotient]

  remainder = n % length
  puts 'remainder: %d' % [remainder]

  remainder.times.each\
  {|idx|
    remainder_matches += 1 if s[idx] == 'a'
  }
  puts 'remainder_matches: %d' % [remainder_matches]

  return matches*quotient + remainder_matches
end

fptr = File.open(ENV['OUTPUT_PATH'], 'w')

s = gets.to_s.rstrip

n = gets.to_i

result = repeatedString s, n

fptr.write result
fptr.write "\n"

fptr.close()
