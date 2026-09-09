from django.db import models
 #======================================================================
# 
# Encapsulates data for model CustomerType
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class CustomerType Declaration (enumerated type)
#======================================================================
from enum import Enum 
class CustomerType(Enum):   # A subclass of Enum
	Distributor = 'Distributor'
	OEM = 'OEM'
	Retailer = 'Retailer'
	Direct = 'Direct'
