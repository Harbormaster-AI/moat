from django.db import models
 #======================================================================
# 
# Encapsulates data for model OperationType
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class OperationType Declaration (enumerated type)
#======================================================================
from enum import Enum 
class OperationType(Enum):   # A subclass of Enum
	Setup = 'Setup'
	Run = 'Run'
	Teardown = 'Teardown'
	Inspection = 'Inspection'
	Transfer = 'Transfer'
