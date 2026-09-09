from django.db import models
 #======================================================================
# 
# Encapsulates data for model MatterType
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class MatterType Declaration (enumerated type)
#======================================================================
from enum import Enum 
class MatterType(Enum):   # A subclass of Enum
	Litigation = 'Litigation'
	Investigation = 'Investigation'
	RegulatoryInquiry = 'RegulatoryInquiry'
	Complaint = 'Complaint'
	Arbitration = 'Arbitration'
