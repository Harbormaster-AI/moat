from django.db import models
 #======================================================================
# 
# Encapsulates data for model AssignmentType
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class AssignmentType Declaration (enumerated type)
#======================================================================
from enum import Enum 
class AssignmentType(Enum):   # A subclass of Enum
	Primary = 'Primary'
	Secondary = 'Secondary'
	Temporary = 'Temporary'
