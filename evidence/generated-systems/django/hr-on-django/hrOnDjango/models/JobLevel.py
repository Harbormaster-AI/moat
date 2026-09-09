from django.db import models
 #======================================================================
# 
# Encapsulates data for model JobLevel
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class JobLevel Declaration (enumerated type)
#======================================================================
from enum import Enum 
class JobLevel(Enum):   # A subclass of Enum
	Entry = 'Entry'
	Intermediate = 'Intermediate'
	Senior = 'Senior'
	Lead = 'Lead'
	Manager = 'Manager'
	Director = 'Director'
	Executive = 'Executive'
