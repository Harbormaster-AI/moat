from django.db import models
 #======================================================================
# 
# Encapsulates data for model ChartType
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class ChartType Declaration (enumerated type)
#======================================================================
from enum import Enum 
class ChartType(Enum):   # A subclass of Enum
	Table = 'Table'
	Bar = 'Bar'
	Line = 'Line'
	Area = 'Area'
	Pie = 'Pie'
	Scatter = 'Scatter'
	Heatmap = 'Heatmap'
	KPI = 'KPI'
