from django.db import models
 #======================================================================
# 
# Encapsulates data for model InterviewStage
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class InterviewStage Declaration (enumerated type)
#======================================================================
from enum import Enum 
class InterviewStage(Enum):   # A subclass of Enum
	PhoneScreen = 'PhoneScreen'
	Technical = 'Technical'
	Onsite = 'Onsite'
	Panel = 'Panel'
	HR = 'HR'
	Executive = 'Executive'
