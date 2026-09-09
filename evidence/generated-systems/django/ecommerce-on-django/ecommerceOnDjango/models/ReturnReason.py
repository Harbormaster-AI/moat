from django.db import models
 #======================================================================
# 
# Encapsulates data for model ReturnReason
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class ReturnReason Declaration (enumerated type)
#======================================================================
from enum import Enum 
class ReturnReason(Enum):   # A subclass of Enum
	Defective = 'Defective'
	Damaged = 'Damaged'
	NotAsDescribed = 'NotAsDescribed'
	WrongItem = 'WrongItem'
	NoLongerNeeded = 'NoLongerNeeded'
	SizeFitIssue = 'SizeFitIssue'
	Other = 'Other'
