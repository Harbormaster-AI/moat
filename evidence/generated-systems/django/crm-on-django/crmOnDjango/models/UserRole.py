from django.db import models
 #======================================================================
# 
# Encapsulates data for model UserRole
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class UserRole Declaration (enumerated type)
#======================================================================
from enum import Enum 
class UserRole(Enum):   # A subclass of Enum
	SalesRep = 'SalesRep'
	SalesManager = 'SalesManager'
	ServiceAgent = 'ServiceAgent'
	MarketingSpecialist = 'MarketingSpecialist'
	Administrator = 'Administrator'
	Executive = 'Executive'
