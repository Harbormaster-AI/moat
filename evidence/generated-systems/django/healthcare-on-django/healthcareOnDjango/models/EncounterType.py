from django.db import models
 #======================================================================
# 
# Encapsulates data for model EncounterType
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class EncounterType Declaration (enumerated type)
#======================================================================
from enum import Enum 
class EncounterType(Enum):   # A subclass of Enum
	Inpatient = 'Inpatient'
	Outpatient = 'Outpatient'
	Emergency = 'Emergency'
	Observation = 'Observation'
	Telemedicine = 'Telemedicine'
