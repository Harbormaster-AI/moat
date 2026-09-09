from django.db import models
 #======================================================================
# 
# Encapsulates data for model ReportType
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class ReportType Declaration (enumerated type)
#======================================================================
from enum import Enum 
class ReportType(Enum):   # A subclass of Enum
	Performance = 'Performance'
	Delivery = 'Delivery'
	Inventory = 'Inventory'
	Billing = 'Billing'
