from django.db import models
 #======================================================================
# 
# Encapsulates data for model LeaveCategory
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class LeaveCategory Declaration (enumerated type)
#======================================================================
from enum import Enum 
class LeaveCategory(Enum):   # A subclass of Enum
	Vacation = 'Vacation'
	Sick = 'Sick'
	Parental = 'Parental'
	Bereavement = 'Bereavement'
	Unpaid = 'Unpaid'
	JuryDuty = 'JuryDuty'
