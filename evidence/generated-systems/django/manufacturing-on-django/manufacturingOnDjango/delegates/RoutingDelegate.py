from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from manufacturingOnDjango.models.Routing import Routing
from manufacturingOnDjango.models.Item import Item
from manufacturingOnDjango.models.Operation import Operation
from manufacturingOnDjango.exceptions import Exceptions

 #======================================================================
# 
# Encapsulates data for model Routing
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class RoutingDelegate Declaration
#======================================================================
class RoutingDelegate :

#======================================================================
# Function Declarations
#======================================================================

	def get(self, routingId ):
		try:	
			routing = Routing.objects.filter(id=routingId)
			return routing.first();
		except Routing.DoesNotExist:
			raise ProcessingError("Routing with id " + str(routingId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 

	def createFromJson(self, routing):
		for model in serializers.deserialize("json", routing):
			model.save()
			return model;

	def create(self, routing):
		routing.save()
		return routing;

	def saveFromJson(self, routing):
		for model in serializers.deserialize("json", routing):
			model.save()
			return routing;
	
	def save(self, routing):
		routing.save()
		return routing;
	
	def delete(self, routingId ):
		errMsg = "Failed to delete Routing from db using id " + str(routingId)
		
		try:
			routing = Routing.objects.get(id=routingId)
			routing.delete()
			return True
		except Routing.DoesNotExist:
			raise ProcessingError("Routing with id " + str(routingId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 
	
	def getAll(self):
		try:
			all = Routing.objects.all()
			return all;
		except utils.DatabaseError:
			raise StorageReadError("Failed to get all Routing from db")
		except Exception:
			return None;
		
	def assignItem( self, routingId, itemId ):
		# lazy importing avoids circular dependencies
		from manufacturingOnDjango.delegates.ItemDelegate import ItemDelegate

		errMsg = "Failed to assign element " + str(itemId) + " for Item on Routing"

		try:
			# get the Routing from db
			routing = self.get( routingId ).first()	
			
			# get the Item from db
			item = ItemDelegate().get(itemId).first();
			
			# assign the Item		
			routing.item = item
			
			#save it
			routing.save()

			# reload and return the appropriate version					
			return self.get( routingId );
		except Routing.DoesNotExist:
			raise ProcessingError(errMsg + " : Routing with id " + str(routingId) + " does not exist.")
		except Item.DoesNotExist:
			raise ProcessingError(errMsg + " : Item with id " + str(itemId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignItem( self, routingId ):
		errMsg = "Failed to unassign element " + str(itemId) + " for Item on Routing"

		try:
			# get the Routing from db
			routing = self.get( routingId ).first()	
			
			# assign to None for unassignment
			routing.item = None			

			#save it
			routing.save()

			# reload and return the appropriate version					
			return self.get( routingId );
		except Routing.DoesNotExist:
			raise ProcessingError(errMsg + " : Routing with id " + str(routingId) + " does not exist.")
		except Exception:
			return None;
		
	def addOperations( self, routingId, operationsIds ):
		# lazy importing avoids circular dependencies
		from manufacturingOnDjango.delegates.OperationDelegate import OperationDelegate

		errMsg = "Failed to add elements " + str(operationsIds) + " for Operations on Routing"

		try:
			# get the Routing
			routing = self.get( routingId ).first()
				
			# split on a comma with no spaces
			idList = operationsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the Operation		
				operation = OperationDelegate().get(id).first();	
				# add the Operation
				routing.operations.add(operation)
				
			# save it		
			routing.save()
			
			# reload and return the appropriate version
			return self.get( routingId );
		except Routing.DoesNotExist:
			raise ProcessingError(errMsg + " : Routing with id " + str(routingId) + " does not exist.")
		except Operation.DoesNotExist:
			raise ProcessingError(errMsg + " : Operation does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeOperations( self, routingId, operationsIds ):
		# lazy importing avoids circular dependencies
		from manufacturingOnDjango.delegates.OperationDelegate import OperationDelegate

		errMsg = "Failed to remove elements " + str(operationsIds) + " for Operations on Routing"

		try:
			# get the Routing
			routing = self.get( routingId ).first()
				
			# split on a comma with no spaces
			idList = operationsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the Operation		
				operation = OperationDelegate().get(id).first();	
				# add the Operation
				routing.operations.remove(operation)
				
			# save it		
			routing.save()
			
			# reload and return the appropriate version
			return self.get( routingId );
		except Routing.DoesNotExist:
			raise ProcessingError(errMsg + " : Routing with id " + str(routingId) + " does not exist.")
		except Operation.DoesNotExist:
			raise ProcessingError(errMsg + " : Operation does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
