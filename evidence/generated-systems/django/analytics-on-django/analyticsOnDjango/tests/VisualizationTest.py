import datetime

from django.test import TestCase
from django.utils import timezone
from analyticsOnDjango.models.Visualization import Visualization
from analyticsOnDjango.delegates.VisualizationDelegate import VisualizationDelegate

 #======================================================================
# 
# Encapsulates data for model Visualization
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class VisualizationTest Declaration
#======================================================================
class VisualizationTest (TestCase) :
	def test_crud(self) :
		visualization = Visualization()
		visualization.title = "default title field value"
		visualization.options = "default options field value"
		visualization.chartType = "default chartType field value"
		
		delegate = VisualizationDelegate()
		responseObj = delegate.create(visualization)
		
		self.assertEqual(responseObj, delegate.get( responseObj.id ))
	
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 1 )		
		delegate.delete(responseObj.id)
		
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 0 )		


