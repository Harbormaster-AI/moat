from django.db import models
 #======================================================================
# 
# Encapsulates data for model EngineCategory
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class EngineCategory Declaration (enumerated type)
#======================================================================
from enum import Enum 
class EngineCategory(Enum):   # A subclass of Enum
	Turbofan = 'Turbofan'
	Turboprop = 'Turboprop'
	Turbojet = 'Turbojet'
	Piston = 'Piston'
	Electric = 'Electric'
	Rocket = 'Rocket'
