# Set cars number equal 100
cars = 100
space_in_a_car = 4.0
drivers = 30
passegers = 90
cars_not_driven = 100 - 30
cars_driven = drivers
carpool_capacity = cars_driven * space_in_a_car
average_passergers_per_car = passegers / cars_driven

print "There are", cars, "cars available"
print "There are only", drivers, "drivers available"
print "There will be", cars_not_driven, "empty cars today"
print "We can transport", carpool_capacity, "people today"
print "We have", passegers, "to carpool today"
print "We need to put about", average_passergers_per_car, "in each car"