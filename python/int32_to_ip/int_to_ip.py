def int32_to_ip(int32):
    b = bin(int32)
    if len(b) < 34:
        b = "0b"+"0" * (34 - len(b)) + b[2:]
    b1 = (b[2:10])
    b2 = (b[10:18])
    b3 = (b[18:26])
    b4 = (b[26:34])
    return '{}.{}.{}.{}'.format(int(b1, 2), int(b2, 2), int(b3, 2), int(b4, 2))
