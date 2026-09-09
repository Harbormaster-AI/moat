from django.db import models
 #======================================================================
# 
# Encapsulates data for model DataSourceType
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class DataSourceType Declaration (enumerated type)
#======================================================================
from enum import Enum 
class DataSourceType(Enum):   # A subclass of Enum
	Database = 'Database'
	File = 'File'
	Stream = 'Stream'
	API = 'API'
	DataWarehouse = 'DataWarehouse'
	DataLake = 'DataLake'
