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


def main():
    numbers = [float(x) for x in sys.argv[1:]]
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
    else:
        print(f'{rd_mean:.0f}±{rd_sem:.0f}s')


if __name__ == '__main__':
    main()
