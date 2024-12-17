import argparse

def find_valid_x(y, r):
    results = []
    for b in range(8):  # b is a 3-bit number
        result = int(r, 2)
        yb = (int(y, 2) << 3) | b
        computed_result = b ^ 0b011 ^ (yb >> (b ^ 0b101))
        if computed_result == result:
            results.append(format(b, '03b'))
    return results


def main():
    parser = argparse.ArgumentParser(description="Find values of X that satisfy the equation.")
    parser.add_argument("-r", required=True, help="The 3-bit binary value for Z.")
    parser.add_argument("-y", required=True, help="The binary value for Y.")
    args = parser.parse_args()

    if len(args.r) != 3 or not set(args.r).issubset("01"):
        print("Error: -r must be a 3-bit binary number.")
        return

    if not set(args.y).issubset("01"):
        print("Error: -y must be a binary number.")
        return

    results = find_valid_x(args.y, args.r)
    if results:
        for b in results:
            print(f"b={b}")
    else:
        print("No valid X found.")


if __name__ == "__main__":
    main()
