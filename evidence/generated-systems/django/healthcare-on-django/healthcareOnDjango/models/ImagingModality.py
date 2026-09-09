from django.db import models
 #======================================================================
# 
# Encapsulates data for model ImagingModality
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class ImagingModality Declaration (enumerated type)
#======================================================================
from enum import Enum 
class ImagingModality(Enum):   # A subclass of Enum
	XRay = 'XRay'
	CT = 'CT'
	MRI = 'MRI'
	Ultrasound = 'Ultrasound'
	PET = 'PET'
	Mammography = 'Mammography'
