def reverse_complement(strand):
    new_strand = ""
    for base in reversed(strand):
        if base.upper() == "A":
            new_strand += "T"
        elif base.upper() == "T":
            new_strand += "A"
        elif base.upper() == "G":
            new_strand += "C"
        elif base.upper() == "C":
            new_strand += "G"
        else:
            print("Error: Non-DNA character found")
            return
    return new_strand

strand = "atGC"
reverse_complement(strand)
