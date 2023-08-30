def binary_search(array, search_value)
  lower_bound = 0
  upper_bound = array.length - 1

  while lower_bound <= upper_bound do
    midpoint = (upper_bound + lower_bound) / 2
    value_at_midpoint = array[midpoint]

    puts lower_bound
    puts midpoint, "is a midpoint"
    puts upper_bound
      if (search_value == value_at_midpoint)
        return midpoint
      elsif search_value < value_at_midpoint
        upper_bound = midpoint - 1
      elsif search_value > value_at_midpoint
        lower_bound =  midpoint + 1
      end
  end
  return nil
end

#emu = binary_search([1, 17, 23, 34, 66, 77], 17)
#puts emu

nono = binary_search([1, 17, 23, 34, 66, 77], 44)
print nono
