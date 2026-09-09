from django.db import models
 #======================================================================
# 
# Encapsulates data for model ObservationInterpretation
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class ObservationInterpretation Declaration (enumerated type)
#======================================================================
from enum import Enum 
class ObservationInterpretation(Enum):   # A subclass of Enum
	Normal = 'Normal'
	AbnormalLow = 'AbnormalLow'
	AbnormalHigh = 'AbnormalHigh'
	CriticalLow = 'CriticalLow'
	CriticalHigh = 'CriticalHigh'
	Reactive = 'Reactive'
	Nonreactive = 'Nonreactive'
	Positive = 'Positive'
	Negative = 'Negative'
