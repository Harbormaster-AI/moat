from django.db import models
 #======================================================================
# 
# Encapsulates data for model PolicyType
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class PolicyType Declaration (enumerated type)
#======================================================================
from enum import Enum 
class PolicyType(Enum):   # A subclass of Enum
	InformationSecurity = 'InformationSecurity'
	DataProtection = 'DataProtection'
	Ethics = 'Ethics'
	RecordsManagement = 'RecordsManagement'
	RiskManagement = 'RiskManagement'
	Compliance = 'Compliance'
	Privacy = 'Privacy'
	AcceptableUse = 'AcceptableUse'
