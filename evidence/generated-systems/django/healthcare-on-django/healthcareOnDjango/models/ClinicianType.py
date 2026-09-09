from django.db import models
 #======================================================================
# 
# Encapsulates data for model ClinicianType
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class ClinicianType Declaration (enumerated type)
#======================================================================
from enum import Enum 
class ClinicianType(Enum):   # A subclass of Enum
	Physician = 'Physician'
	NursePractitioner = 'NursePractitioner'
	PhysicianAssistant = 'PhysicianAssistant'
	RegisteredNurse = 'RegisteredNurse'
	Pharmacist = 'Pharmacist'
	Therapist = 'Therapist'
	Technician = 'Technician'
