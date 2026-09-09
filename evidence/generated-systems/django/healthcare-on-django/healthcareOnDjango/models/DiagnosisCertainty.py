from django.db import models
 #======================================================================
# 
# Encapsulates data for model DiagnosisCertainty
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class DiagnosisCertainty Declaration (enumerated type)
#======================================================================
from enum import Enum 
class DiagnosisCertainty(Enum):   # A subclass of Enum
	Suspected = 'Suspected'
	Presumptive = 'Presumptive'
	Confirmed = 'Confirmed'
	RuledOut = 'RuledOut'
