import sys
import math


def round_sem(sem):
    '''
    Round SEM to the first significant digit, rounding up
    '''
    dot = - math.floor(math.log10(sem))  # Position after period
    pw = 10 ** dot
    # Fix 0.098 -> 0.010 issue, should be 0.01
    if math.ceil(sem * pw) >= 10:
        dot -= 1
        pw /= 10
    ret = math.ceil(sem * pw) / pw
    # System error is 0.01, so shouldn't be less
    if ret < 0.01:
        ret = 0.01
        dot = 2
    return ret, dot


def round_mean(mean, dot):
    '''
    Round mean to the last significant digit, rounding half to even
    '''
    pw = 10 ** dot
    ret = round(mean * pw) / pw
    return ret


def compute(numbers):
    n = len(numbers)
    xsum = sum(numbers)
    xmean = xsum / n
    sqsum = sum([(x - xmean)**2 for x in numbers])
    sem = math.sqrt(sqsum / n / (n-1))
    rd_sem, dot = round_sem(sem)
    rd_mean = round_mean(xmean, dot)
    if dot > 0:
        fmtstr = f'{{rd_mean:.{dot}f}}±{{rd_sem:.{dot}f}}s'
        print(fmtstr.format(rd_mean=rd_mean, rd_sem=rd_sem))
        return xmean, sem
    else:
        print(f'{rd_mean:.0f}±{rd_sem:.0f}s')
        return xmean, sem


def linear_reg(m1, s1, m2, s2, m3, s3, dkb, dnc, kb1, nc1):
    kkb = (m2 - m1) / dkb
    skb = math.sqrt(s1*s1 + s2*s2) / dkb
    knc = (m3 - m1) / dnc
    snc = math.sqrt(s1*s1 + s3*s3) / dnc
    base = m1 - kkb * kb1 - knc * nc1
    sbase = math.sqrt(s1*s1 + kb1*kb1*skb*skb + nc1*nc1*snc*snc)
    # Print result
    rd_sbase, dot = round_sem(sbase)
    rd_base = round_mean(base, dot)
    if dot > 0:
        fmtstr = f'BAS={{rd_mean:.{dot}f}}±{{rd_sem:.{dot}f}}s'
        print(fmtstr.format(rd_mean=rd_base, rd_sem=rd_sbase))
    else:
        print(f'BAS={rd_base:.0f}±{rd_sbase:.0f}s')

    rd_skb, dot = round_sem(skb)
    rd_kkb = round_mean(kkb, dot)
    if dot > 0:
        fmtstr = f'KKB={{rd_mean:.{dot}f}}±{{rd_sem:.{dot}f}}s'
        print(fmtstr.format(rd_mean=rd_kkb, rd_sem=rd_skb))
    else:
        print(f'KKB={rd_kkb:.0f}±{rd_skb:.0f}s')

    rd_snc, dot = round_sem(snc)
    rd_knc = round_mean(knc, dot)
    if dot > 0:
        fmtstr = f'KNC={{rd_mean:.{dot}f}}±{{rd_sem:.{dot}f}}s'
        print(fmtstr.format(rd_mean=rd_knc, rd_sem=rd_snc))
    else:
        print(f'KNC={rd_knc:.0f}±{rd_snc:.0f}s')


def main():
    numbers = [float(x) for x in sys.argv[1:]]
    m1, s1 = compute(numbers[0:5])
    m2, s2 = compute(numbers[5:10])
    m3, s3 = compute(numbers[10:15])
    linear_reg(m1, s1, m2, s2, m3, s3, 3.904, 30.0, 0.189, 3)


if __name__ == '__main__':
    main()
