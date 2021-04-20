import time
from ndn.encoding import make_data, MetaInfo, Component
from ndn.security.signer import DigestSha256Signer


def one_run(i: int):
    name = f"/go-ndn/test/{i:010}"
    content = bytes(4000)
    make_data(name, MetaInfo(0, 4000, Component.from_str("10000")), content, DigestSha256Signer())


def main():
    tic = time.perf_counter()
    for i in range(1000000):
        one_run(i)
    toc = time.perf_counter()
    print(f"Total time: {toc - tic:0.4f} seconds")


if __name__ == '__main__':
    main()
