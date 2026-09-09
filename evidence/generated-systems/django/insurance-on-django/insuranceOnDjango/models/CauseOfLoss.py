from django.db import models
 #======================================================================
# 
# Encapsulates data for model CauseOfLoss
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class CauseOfLoss Declaration (enumerated type)
#======================================================================
from enum import Enum 
class CauseOfLoss(Enum):   # A subclass of Enum
	Collision = 'Collision'
	Weather = 'Weather'
	MechanicalFailure = 'MechanicalFailure'
	HumanError = 'HumanError'
	NaturalDisaster = 'NaturalDisaster'
	Theft = 'Theft'
	Vandalism = 'Vandalism'
	LiabilityClaim = 'LiabilityClaim'
	Illness = 'Illness'
