import statsmodels.api as sm
import datapackage
import matplotlib.pyplot as plt
import numpy as np
import pandas as pd
import seaborn as sns

data_url_example = 'https://datahub.io/core/natural-gas/datapackage.json'

# to load Data Package into storage
package = datapackage.Package(data_url_example)

# to load only tabular data
resources = package.resources
for resource in resources:
    if resource.tabular:
        data = pd.read_csv(resource.descriptor['path'])
        print (data)
