from django.db import models
 #======================================================================
# 
# Encapsulates data for model ObjectiveType
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class ObjectiveType Declaration (enumerated type)
#======================================================================
from enum import Enum 
class ObjectiveType(Enum):   # A subclass of Enum
	Awareness = 'Awareness'
	Reach = 'Reach'
	Traffic = 'Traffic'
	Engagement = 'Engagement'
	Leads = 'Leads'
	Sales = 'Sales'
	AppInstalls = 'AppInstalls'
	VideoViews = 'VideoViews'
