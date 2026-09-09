from django.db import models
 #======================================================================
# 
# Encapsulates data for model ComplianceStatus
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class ComplianceStatus Declaration (enumerated type)
#======================================================================
from enum import Enum 
class ComplianceStatus(Enum):   # A subclass of Enum
	NotStarted = 'NotStarted'
	InProgress = 'InProgress'
	Compliant = 'Compliant'
	NonCompliant = 'NonCompliant'
	Waived = 'Waived'
