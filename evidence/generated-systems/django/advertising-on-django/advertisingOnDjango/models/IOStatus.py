from django.db import models
 #======================================================================
# 
# Encapsulates data for model IOStatus
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class IOStatus Declaration (enumerated type)
#======================================================================
from enum import Enum 
class IOStatus(Enum):   # A subclass of Enum
	Draft = 'Draft'
	Sent = 'Sent'
	Executed = 'Executed'
	OnHold = 'OnHold'
	Closed = 'Closed'
	Cancelled = 'Cancelled'
