from django.db import models
 #======================================================================
# 
# Encapsulates data for model LineOfBusiness
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class LineOfBusiness Declaration (enumerated type)
#======================================================================
from enum import Enum 
class LineOfBusiness(Enum):   # A subclass of Enum
	PersonalAuto = 'PersonalAuto'
	Homeowners = 'Homeowners'
	Renters = 'Renters'
	TermLife = 'TermLife'
	WholeLife = 'WholeLife'
	CommercialProperty = 'CommercialProperty'
	GeneralLiability = 'GeneralLiability'
	WorkersCompensation = 'WorkersCompensation'
