from django.db import models
 #======================================================================
# 
# Encapsulates data for model NonconformanceType
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class NonconformanceType Declaration (enumerated type)
#======================================================================
from enum import Enum 
class NonconformanceType(Enum):   # A subclass of Enum
	Dimension = 'Dimension'
	Functional = 'Functional'
	Cosmetic = 'Cosmetic'
	Documentation = 'Documentation'
	Supplier = 'Supplier'
	Process = 'Process'
