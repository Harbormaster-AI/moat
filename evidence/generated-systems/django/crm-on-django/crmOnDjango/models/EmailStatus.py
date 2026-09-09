from django.db import models
 #======================================================================
# 
# Encapsulates data for model EmailStatus
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class EmailStatus Declaration (enumerated type)
#======================================================================
from enum import Enum 
class EmailStatus(Enum):   # A subclass of Enum
	Draft = 'Draft'
	Sent = 'Sent'
	Delivered = 'Delivered'
	Opened = 'Opened'
	Bounced = 'Bounced'
	Failed = 'Failed'
	Replied = 'Replied'
