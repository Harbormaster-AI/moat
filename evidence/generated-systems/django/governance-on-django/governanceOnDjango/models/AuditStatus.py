from django.db import models
 #======================================================================
# 
# Encapsulates data for model AuditStatus
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class AuditStatus Declaration (enumerated type)
#======================================================================
from enum import Enum 
class AuditStatus(Enum):   # A subclass of Enum
	Planned = 'Planned'
	Fieldwork = 'Fieldwork'
	Reporting = 'Reporting'
	Closed = 'Closed'
	OnHold = 'OnHold'
