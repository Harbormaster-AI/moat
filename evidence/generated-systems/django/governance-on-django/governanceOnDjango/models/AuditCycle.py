from django.db import models
 #======================================================================
# 
# Encapsulates data for model AuditCycle
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class AuditCycle Declaration (enumerated type)
#======================================================================
from enum import Enum 
class AuditCycle(Enum):   # A subclass of Enum
	Annual = 'Annual'
	SemiAnnual = 'SemiAnnual'
	Quarterly = 'Quarterly'
	Continuous = 'Continuous'
	OneTime = 'OneTime'
