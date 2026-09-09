from django.db import models
 #======================================================================
# 
# Encapsulates data for model TreatyType
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class TreatyType Declaration (enumerated type)
#======================================================================
from enum import Enum 
class TreatyType(Enum):   # A subclass of Enum
	QuotaShare = 'QuotaShare'
	Surplus = 'Surplus'
	ExcessOfLoss = 'ExcessOfLoss'
	StopLoss = 'StopLoss'
