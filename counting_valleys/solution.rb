#!/usr/bin/env ruby

require 'json'
require 'stringio'

def is_neg?(v)
  return false if v >= 0
  return true
end

# Complete the countingValleys function below.
def countingValleys(n, s)
  valley_count = 0
  here, last_step = 0, 0

	#puts "s is: %s" % [s]
  s.scan(/./).each \
  {|step|
    #puts "step is: %s" % [step]

    here += 1 if step == 'U'
    here -= 1 if step == 'D'

    #puts "here is: %d" % [here]

    is_valley = is_neg? here
    valley_count += 1 if is_valley and last_step == 0
    #puts "lastHere is: %d    here is: %d   is_neg is: %s\n" % [last_step, here, is_valley]

    last_step = here
  }

  valley_count
end

fptr = File.open(ENV['OUTPUT_PATH'], 'w')

n = gets.to_i

s = gets.to_s.rstrip

result = countingValleys n, s

fptr.write result
fptr.write "\n"

fptr.close()

