from django.db import models
 #======================================================================
# 
# Encapsulates data for model LineageNodeType
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class LineageNodeType Declaration (enumerated type)
#======================================================================
from enum import Enum 
class LineageNodeType(Enum):   # A subclass of Enum
	Dataset = 'Dataset'
	Pipeline = 'Pipeline'
	Model = 'Model'
	Dashboard = 'Dashboard'
	Report = 'Report'
	FeatureSet = 'FeatureSet'
	Notebook = 'Notebook'
