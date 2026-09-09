from django.db import models
 #======================================================================
# 
# Encapsulates data for model SkillLevel
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class SkillLevel Declaration (enumerated type)
#======================================================================
from enum import Enum 
class SkillLevel(Enum):   # A subclass of Enum
	Novice = 'Novice'
	Competent = 'Competent'
	Proficient = 'Proficient'
	Expert = 'Expert'
