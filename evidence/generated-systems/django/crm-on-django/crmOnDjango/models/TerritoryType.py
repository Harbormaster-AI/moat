from django.db import models
 #======================================================================
# 
# Encapsulates data for model TerritoryType
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class TerritoryType Declaration (enumerated type)
#======================================================================
from enum import Enum 
class TerritoryType(Enum):   # A subclass of Enum
	Geographic = 'Geographic'
	Industry = 'Industry'
	NamedAccount = 'NamedAccount'
	Segment = 'Segment'
	Hybrid = 'Hybrid'
