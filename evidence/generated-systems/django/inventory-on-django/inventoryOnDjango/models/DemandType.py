from django.db import models
 #======================================================================
# 
# Encapsulates data for model DemandType
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class DemandType Declaration (enumerated type)
#======================================================================
from enum import Enum 
class DemandType(Enum):   # A subclass of Enum
	SalesOrder = 'SalesOrder'
	WorkOrder = 'WorkOrder'
	TransferOrder = 'TransferOrder'
	Forecast = 'Forecast'
	SampleRequest = 'SampleRequest'
