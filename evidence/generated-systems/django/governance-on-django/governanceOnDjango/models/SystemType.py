from django.db import models
 #======================================================================
# 
# Encapsulates data for model SystemType
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class SystemType Declaration (enumerated type)
#======================================================================
from enum import Enum 
class SystemType(Enum):   # A subclass of Enum
	Application = 'Application'
	Database = 'Database'
	DataWarehouse = 'DataWarehouse'
	SaaS = 'SaaS'
	Infrastructure = 'Infrastructure'
	Endpoint = 'Endpoint'
