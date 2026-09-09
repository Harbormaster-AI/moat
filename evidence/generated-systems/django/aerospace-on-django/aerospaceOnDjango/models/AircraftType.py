from django.db import models
 #======================================================================
# 
# Encapsulates data for model AircraftType
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class AircraftType Declaration (enumerated type)
#======================================================================
from enum import Enum 
class AircraftType(Enum):   # A subclass of Enum
	NarrowBody = 'NarrowBody'
	WideBody = 'WideBody'
	RegionalJet = 'RegionalJet'
	Turboprop = 'Turboprop'
	BusinessJet = 'BusinessJet'
	Helicopter = 'Helicopter'
	eVTOL = 'eVTOL'
	CargoPlane = 'CargoPlane'
