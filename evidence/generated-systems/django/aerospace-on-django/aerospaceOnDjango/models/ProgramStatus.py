from django.db import models
 #======================================================================
# 
# Encapsulates data for model ProgramStatus
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class ProgramStatus Declaration (enumerated type)
#======================================================================
from enum import Enum 
class ProgramStatus(Enum):   # A subclass of Enum
	Concept = 'Concept'
	Development = 'Development'
	Certification = 'Certification'
	Production = 'Production'
	InService = 'InService'
	Sunset = 'Sunset'
