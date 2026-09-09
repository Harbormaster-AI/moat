from django.db import models
 #======================================================================
# 
# Encapsulates data for model ClinicianSpecialty
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class ClinicianSpecialty Declaration (enumerated type)
#======================================================================
from enum import Enum 
class ClinicianSpecialty(Enum):   # A subclass of Enum
	InternalMedicine = 'InternalMedicine'
	FamilyMedicine = 'FamilyMedicine'
	Cardiology = 'Cardiology'
	Oncology = 'Oncology'
	Orthopedics = 'Orthopedics'
	Pediatrics = 'Pediatrics'
	Radiology = 'Radiology'
	Pathology = 'Pathology'
	Anesthesiology = 'Anesthesiology'
	Surgery = 'Surgery'
	Psychiatry = 'Psychiatry'
