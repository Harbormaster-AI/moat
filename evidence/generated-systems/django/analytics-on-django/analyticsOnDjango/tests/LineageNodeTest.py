import datetime

from django.test import TestCase
from django.utils import timezone
from analyticsOnDjango.models.LineageNode import LineageNode
from analyticsOnDjango.delegates.LineageNodeDelegate import LineageNodeDelegate

 #======================================================================
# 
# Encapsulates data for model LineageNode
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class LineageNodeTest Declaration
#======================================================================
class LineageNodeTest (TestCase) :
	def test_crud(self) :
		lineageNode = LineageNode()
		lineageNode.name = "default name field value"
		lineageNode.qualifiedName = "default qualifiedName field value"
		lineageNode.nodeType = "default nodeType field value"
		
		delegate = LineageNodeDelegate()
		responseObj = delegate.create(lineageNode)
		
		self.assertEqual(responseObj, delegate.get( responseObj.id ))
	
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 1 )		
		delegate.delete(responseObj.id)
		
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 0 )		


