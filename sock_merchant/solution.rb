#!/usr/bin/env ruby
#

require 'json'
require 'stringio'
require 'awesome_print'

# Complete the sockMerchant function below.
def sockMerchant(ar)
  matched = {}
  pairs   = 0

  ar.each \
  {|sock|
    if not matched.has_key? sock
      matched[sock] = 0
    end
    matched[sock] += 1
  }

  matched.each\
  {|k,v|
    pairs += v / 2       if v % 2 == 0
    pairs += (v - 1) / 2 if v - 1 != 0 and (v - 1) % 2 == 0
  }
  pairs
end

fptr = File.open(ENV['OUTPUT_PATH'], 'w')

ar = gets.rstrip.split(' ').map(&:to_i)

result = sockMerchant ar

fptr.write result
fptr.write "\n"

fptr.close()

