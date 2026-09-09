from django.db import models
 #======================================================================
# 
# Encapsulates data for model RiskCategory
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class RiskCategory Declaration (enumerated type)
#======================================================================
from enum import Enum 
class RiskCategory(Enum):   # A subclass of Enum
	Strategic = 'Strategic'
	Operational = 'Operational'
	Financial = 'Financial'
	Compliance = 'Compliance'
	Reputational = 'Reputational'
	Privacy = 'Privacy'
	Cybersecurity = 'Cybersecurity'
	ThirdParty = 'ThirdParty'
