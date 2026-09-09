from django.db import models
 #======================================================================
# 
# Encapsulates data for model TaxClass
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class TaxClass Declaration (enumerated type)
#======================================================================
from enum import Enum 
class TaxClass(Enum):   # A subclass of Enum
	Standard = 'Standard'
	Reduced = 'Reduced'
	Zero = 'Zero'
	Exempt = 'Exempt'
	DigitalServices = 'DigitalServices'
	Food = 'Food'
	Clothing = 'Clothing'
