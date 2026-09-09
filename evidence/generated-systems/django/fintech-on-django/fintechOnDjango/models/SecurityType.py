from django.db import models
 #======================================================================
# 
# Encapsulates data for model SecurityType
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class SecurityType Declaration (enumerated type)
#======================================================================
from enum import Enum 
class SecurityType(Enum):   # A subclass of Enum
	Equity = 'Equity'
	Bond = 'Bond'
	ETF = 'ETF'
	MutualFund = 'MutualFund'
	Derivative = 'Derivative'
	Crypto = 'Crypto'
