from django.db import models
 #======================================================================
# 
# Encapsulates data for model DataFormat
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class DataFormat Declaration (enumerated type)
#======================================================================
from enum import Enum 
class DataFormat(Enum):   # A subclass of Enum
	CSV = 'CSV'
	JSON = 'JSON'
	Parquet = 'Parquet'
	Avro = 'Avro'
	ORC = 'ORC'
	XML = 'XML'
