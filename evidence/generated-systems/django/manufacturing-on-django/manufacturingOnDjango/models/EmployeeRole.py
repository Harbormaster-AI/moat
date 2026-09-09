from django.db import models
 #======================================================================
# 
# Encapsulates data for model EmployeeRole
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class EmployeeRole Declaration (enumerated type)
#======================================================================
from enum import Enum 
class EmployeeRole(Enum):   # A subclass of Enum
	Operator = 'Operator'
	Technician = 'Technician'
	Supervisor = 'Supervisor'
	Planner = 'Planner'
	QualityEngineer = 'QualityEngineer'
	Buyer = 'Buyer'
