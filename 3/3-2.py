# Verify Life Support Rating
#   = Oxygen Generator Rating * CO2 Scrubber Rating

# Oxygen Generator Rating
#   * Most common value (0 or 1) in the current bit position, keep only numbers with that bit in the position (tiebreak is 1)

# CO2 Scrubber Rating
#   * Same as Oxygen but least common value with tiebreak 0

# Returns a filtered report
# report    = the report
# bit_index = the bit index to filter on
# bit       = the wanted bit for that index
def filter_report(report, bit_index, bit):
    filtered_report = []
    for number in report:
        if number[bit_index] == bit:
            filtered_report.append(number)
    return filtered_report

def oxygen_recur(report, bit_index):
    if len(report) == 1 or bit_index > len(report[0]):
        return report

    s = sum([number[bit_index] for number in report])

    if s >= len(report) / 2:
        # 1 is most common
        report = filter_report(report, bit_index, 1)
    else:
        # 0 is most common
        report = filter_report(report, bit_index, 0)

    return oxygen_recur(report, bit_index + 1)

def oxygen(report):
    return oxygen_recur(report, 0)

# -------------------------------------------------------

with open("test1.txt", "r") as f:
    lines = f.read().splitlines()

nums = []
for line in lines:
    l = list(line)
    nums.append([int(i) for i in l])

# print(nums)
# filter(nums, 1, 1)
print(oxygen(nums))
