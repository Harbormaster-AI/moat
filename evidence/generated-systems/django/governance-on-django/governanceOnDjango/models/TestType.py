from django.db import models
 #======================================================================
# 
# Encapsulates data for model TestType
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class TestType Declaration (enumerated type)
#======================================================================
from enum import Enum 
class TestType(Enum):   # A subclass of Enum
	DesignEffectiveness = 'DesignEffectiveness'
	OperatingEffectiveness = 'OperatingEffectiveness'
	Walkthrough = 'Walkthrough'
	Reperformance = 'Reperformance'
	Inquiry = 'Inquiry'
	Observation = 'Observation'
	Inspection = 'Inspection'
	DataAnalysis = 'DataAnalysis'
