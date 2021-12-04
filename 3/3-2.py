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

def rating_recur(report, bit_index, common):
    if len(report) == 1 or bit_index >= len(report[0]):
        return report

    s = sum([number[bit_index] for number in report])

    if s >= len(report) / 2:
        # 1 is most common or there's a tie
        new_report = filter_report(report, bit_index, 1 if common else 0)
    elif s < len(report) / 2:
        # 0 is most common
        new_report = filter_report(report, bit_index, 0 if common else 1)

    # If new_report is 0 it means that we had multiples of the same number, so return it
    # Wrap it in a list to look like the "normal" case return value
    if len(new_report) == 0:
        return [report[0]]

    return rating_recur(new_report, bit_index + 1, common)

def oxygen(report):
    return rating_recur(report, 0, True)

def co2(report):
    return rating_recur(report, 0, False)

def bin_to_dec(bin_list):
    dec = 0
    for i, bit in enumerate(bin_list):
         dec = dec + 2**(len(bin_list)-i-1)*bit
    return dec

# -------------------------------------------------------

with open("input.txt", "r") as f:
    lines = f.read().splitlines()

nums = []
for line in lines:
    l = list(line)
    nums.append([int(i) for i in l])

oxygen_bin = oxygen(nums)[0]
co2_bin = co2(nums)[0]

print(bin_to_dec(oxygen_bin) * bin_to_dec(co2_bin))
