from django.db import models
 #======================================================================
# 
# Encapsulates data for model ReportStatus
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class ReportStatus Declaration (enumerated type)
#======================================================================
from enum import Enum 
class ReportStatus(Enum):   # A subclass of Enum
	Draft = 'Draft'
	Published = 'Published'
	Archived = 'Archived'
