#!/usr/bin/env ruby

require 'json'
require 'stringio'
require 'awesome_print'

# Complete the 'pickingNumbers' function below.
#
# The function is expected to return an INTEGER.
# The function accepts INTEGER_ARRAY a as parameter.
#
# 1 2 2 3 1 2

def pickingNumbers(a)
  m    = {}
  iter = 0
  last = 0
  ans  = nil

  a.each_with_index \
  {|ele, i|
    break if a[i + 1].nil?

    next_element = a[i + 1]

    #puts 'idx: %d   ele: %d   next_element: %d   subt(abs): %d' % [i, ele, next_element, (next_element - ele).abs]
    if (next_element - ele).abs <= 1
      ap iter
      (m[iter] ||= []) << ele
      next
    end

    iter += 1
  }

  ap m

  m.each_pair \
  {|k, v|
    size = v.count

    if last == 0
      last = size
      next
    end

    #puts 'last(%d) > size(%d)' % [last,size]
    if last > size
      ans = last
    end
  }

  ans
end

fptr = File.open(ENV['OUTPUT_PATH'], 'w')

gets.strip.to_i

a = gets.rstrip.split.map(&:to_i)

result = pickingNumbers a
ap result

fptr.write result
fptr.write "\n"

fptr.close()
