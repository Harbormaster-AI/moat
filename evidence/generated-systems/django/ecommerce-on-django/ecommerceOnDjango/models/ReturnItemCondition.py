from django.db import models
 #======================================================================
# 
# Encapsulates data for model ReturnItemCondition
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class ReturnItemCondition Declaration (enumerated type)
#======================================================================
from enum import Enum 
class ReturnItemCondition(Enum):   # A subclass of Enum
	New = 'New'
	OpenBox = 'OpenBox'
	Used = 'Used'
	Damaged = 'Damaged'
	MissingParts = 'MissingParts'
