from django.db import models
 #======================================================================
# 
# Encapsulates data for model ThirdPartyType
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class ThirdPartyType Declaration (enumerated type)
#======================================================================
from enum import Enum 
class ThirdPartyType(Enum):   # A subclass of Enum
	Vendor = 'Vendor'
	Processor = 'Processor'
	JointController = 'JointController'
	Subprocessor = 'Subprocessor'
	Partner = 'Partner'
	Consultant = 'Consultant'
