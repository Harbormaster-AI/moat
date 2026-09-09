from django.db import models
 #======================================================================
# 
# Encapsulates data for model EquityType
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class EquityType Declaration (enumerated type)
#======================================================================
from enum import Enum 
class EquityType(Enum):   # A subclass of Enum
	RSU = 'RSU'
	StockOption = 'StockOption'
	ESPP = 'ESPP'
