from django.db import models
 #======================================================================
# 
# Encapsulates data for model EventSeverity
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class EventSeverity Declaration (enumerated type)
#======================================================================
from enum import Enum 
class EventSeverity(Enum):   # A subclass of Enum
	Info = 'Info'
	Warning = 'Warning'
	Critical = 'Critical'
