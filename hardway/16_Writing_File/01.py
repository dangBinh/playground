from sys import argv

script, filename = argv

print "We're going to erase %r" % filename

target = open(filename, 'w')
target.truncate()

line1 = raw_input("line1: ")
line2 = raw_input("line2: ")

target.write(line1)
target.write("\n")
target.write(line2)
target.write("\n")

target.close()