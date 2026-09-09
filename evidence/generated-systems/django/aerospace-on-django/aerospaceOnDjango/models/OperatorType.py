from django.db import models
 #======================================================================
# 
# Encapsulates data for model OperatorType
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class OperatorType Declaration (enumerated type)
#======================================================================
from enum import Enum 
class OperatorType(Enum):   # A subclass of Enum
	Airline = 'Airline'
	Cargo = 'Cargo'
	Government = 'Government'
	Private = 'Private'
	Lessor = 'Lessor'
