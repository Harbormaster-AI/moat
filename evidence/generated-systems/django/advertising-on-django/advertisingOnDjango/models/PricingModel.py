from django.db import models
 #======================================================================
# 
# Encapsulates data for model PricingModel
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class PricingModel Declaration (enumerated type)
#======================================================================
from enum import Enum 
class PricingModel(Enum):   # A subclass of Enum
	CPM = 'CPM'
	CPC = 'CPC'
	CPA = 'CPA'
	CPL = 'CPL'
	CPV = 'CPV'
	FlatFee = 'FlatFee'
