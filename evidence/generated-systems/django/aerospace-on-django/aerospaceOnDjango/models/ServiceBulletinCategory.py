from django.db import models
 #======================================================================
# 
# Encapsulates data for model ServiceBulletinCategory
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class ServiceBulletinCategory Declaration (enumerated type)
#======================================================================
from enum import Enum 
class ServiceBulletinCategory(Enum):   # A subclass of Enum
	Recommended = 'Recommended'
	Optional = 'Optional'
	Alert = 'Alert'
	Mandatory = 'Mandatory'
