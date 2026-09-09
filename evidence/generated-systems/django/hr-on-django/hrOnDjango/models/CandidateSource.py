from django.db import models
 #======================================================================
# 
# Encapsulates data for model CandidateSource
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class CandidateSource Declaration (enumerated type)
#======================================================================
from enum import Enum 
class CandidateSource(Enum):   # A subclass of Enum
	Referral = 'Referral'
	Agency = 'Agency'
	JobBoard = 'JobBoard'
	CareerSite = 'CareerSite'
	Campus = 'Campus'
	Social = 'Social'
	Internal = 'Internal'
