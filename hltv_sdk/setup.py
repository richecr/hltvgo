import setuptools

with open("README.md", "r") as fh:
    long_description = fh.read()

with open("LICENSE", "r") as fh:
    license = fh.read()


class BinaryDistribution(setuptools.Distribution):
    def has_ext_modules(_):
        return True


setuptools.setup(
    name="hltvsdk",
    version="0.1.0",
    author="Rich Ramalho",
    author_email="richelton14@gmail.com",
    description="The unofficial HLTV Python API",
    license=license,
    long_description=long_description,
    long_description_content_type="text/markdown",
    keywords=["sdk", "api", "hltv", "cs", "csgo", "cs2", "python", "go" "golang"],
    url="https://github.com/richecr/hltvgo",
    packages=setuptools.find_packages(),
    classifiers=[
        "Programming Language :: Python :: 3",
        "License :: OSI Approved :: BSD License",
        "Operating System :: OS Independent",
    ],
    include_package_data=True,
)
