from django.db import models
 #======================================================================
# 
# Encapsulates data for model BidStrategyType
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class BidStrategyType Declaration (enumerated type)
#======================================================================
from enum import Enum 
class BidStrategyType(Enum):   # A subclass of Enum
	Manual = 'Manual'
	AutoMaximizeClicks = 'AutoMaximizeClicks'
	AutoTargetCPA = 'AutoTargetCPA'
	AutoTargetROAS = 'AutoTargetROAS'
