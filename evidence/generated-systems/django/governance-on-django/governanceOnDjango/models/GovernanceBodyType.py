from django.db import models
 #======================================================================
# 
# Encapsulates data for model GovernanceBodyType
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class GovernanceBodyType Declaration (enumerated type)
#======================================================================
from enum import Enum 
class GovernanceBodyType(Enum):   # A subclass of Enum
	Board = 'Board'
	Committee = 'Committee'
	Council = 'Council'
	WorkingGroup = 'WorkingGroup'
