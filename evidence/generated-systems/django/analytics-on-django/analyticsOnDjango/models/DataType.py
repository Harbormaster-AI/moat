from django.db import models
 #======================================================================
# 
# Encapsulates data for model DataType
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class DataType Declaration (enumerated type)
#======================================================================
from enum import Enum 
class DataType(Enum):   # A subclass of Enum
	String = 'String'
	Integer = 'Integer'
	Decimal = 'Decimal'
	Boolean = 'Boolean'
	Date = 'Date'
	DateTime = 'DateTime'
