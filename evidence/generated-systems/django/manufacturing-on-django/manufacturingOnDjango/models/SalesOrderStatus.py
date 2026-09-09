from django.db import models
 #======================================================================
# 
# Encapsulates data for model SalesOrderStatus
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class SalesOrderStatus Declaration (enumerated type)
#======================================================================
from enum import Enum 
class SalesOrderStatus(Enum):   # A subclass of Enum
	Draft = 'Draft'
	Confirmed = 'Confirmed'
	Allocated = 'Allocated'
	InProduction = 'InProduction'
	Shipped = 'Shipped'
	Invoiced = 'Invoiced'
	Closed = 'Closed'
	Cancelled = 'Cancelled'
