from django.db import models
 #======================================================================
# 
# Encapsulates data for model InvestmentAccountType
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class InvestmentAccountType Declaration (enumerated type)
#======================================================================
from enum import Enum 
class InvestmentAccountType(Enum):   # A subclass of Enum
	Brokerage = 'Brokerage'
	Retirement = 'Retirement'
	Custody = 'Custody'
	Margin = 'Margin'
