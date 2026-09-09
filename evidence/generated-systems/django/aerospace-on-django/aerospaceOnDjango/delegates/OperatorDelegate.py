from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from aerospaceOnDjango.models.Operator import Operator
from aerospaceOnDjango.models.AircraftOrder import AircraftOrder
from aerospaceOnDjango.models.Aircraft import Aircraft
from aerospaceOnDjango.models.SalesRegion import SalesRegion
from aerospaceOnDjango.exceptions import Exceptions

 #======================================================================
# 
# Encapsulates data for model Operator
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class OperatorDelegate Declaration
#======================================================================
class OperatorDelegate :

#======================================================================
# Function Declarations
#======================================================================

	def get(self, operatorId ):
		try:	
			operator = Operator.objects.filter(id=operatorId)
			return operator.first();
		except Operator.DoesNotExist:
			raise ProcessingError("Operator with id " + str(operatorId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 

	def createFromJson(self, operator):
		for model in serializers.deserialize("json", operator):
			model.save()
			return model;

	def create(self, operator):
		operator.save()
		return operator;

	def saveFromJson(self, operator):
		for model in serializers.deserialize("json", operator):
			model.save()
			return operator;
	
	def save(self, operator):
		operator.save()
		return operator;
	
	def delete(self, operatorId ):
		errMsg = "Failed to delete Operator from db using id " + str(operatorId)
		
		try:
			operator = Operator.objects.get(id=operatorId)
			operator.delete()
			return True
		except Operator.DoesNotExist:
			raise ProcessingError("Operator with id " + str(operatorId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 
	
	def getAll(self):
		try:
			all = Operator.objects.all()
			return all;
		except utils.DatabaseError:
			raise StorageReadError("Failed to get all Operator from db")
		except Exception:
			return None;
		
	def assignSalesRegion( self, operatorId, salesRegionId ):
		# lazy importing avoids circular dependencies
		from aerospaceOnDjango.delegates.SalesRegionDelegate import SalesRegionDelegate

		errMsg = "Failed to assign element " + str(salesRegionId) + " for SalesRegion on Operator"

		try:
			# get the Operator from db
			operator = self.get( operatorId ).first()	
			
			# get the SalesRegion from db
			salesRegion = SalesRegionDelegate().get(salesRegionId).first();
			
			# assign the SalesRegion		
			operator.salesRegion = salesRegion
			
			#save it
			operator.save()

			# reload and return the appropriate version					
			return self.get( operatorId );
		except Operator.DoesNotExist:
			raise ProcessingError(errMsg + " : Operator with id " + str(operatorId) + " does not exist.")
		except SalesRegion.DoesNotExist:
			raise ProcessingError(errMsg + " : SalesRegion with id " + str(salesRegionId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignSalesRegion( self, operatorId ):
		errMsg = "Failed to unassign element " + str(salesRegionId) + " for SalesRegion on Operator"

		try:
			# get the Operator from db
			operator = self.get( operatorId ).first()	
			
			# assign to None for unassignment
			operator.salesRegion = None			

			#save it
			operator.save()

			# reload and return the appropriate version					
			return self.get( operatorId );
		except Operator.DoesNotExist:
			raise ProcessingError(errMsg + " : Operator with id " + str(operatorId) + " does not exist.")
		except Exception:
			return None;
		
	def addAircraftOrders( self, operatorId, aircraftOrdersIds ):
		# lazy importing avoids circular dependencies
		from aerospaceOnDjango.delegates.AircraftOrderDelegate import AircraftOrderDelegate

		errMsg = "Failed to add elements " + str(aircraftOrdersIds) + " for AircraftOrders on Operator"

		try:
			# get the Operator
			operator = self.get( operatorId ).first()
				
			# split on a comma with no spaces
			idList = aircraftOrdersIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the AircraftOrder		
				aircraftOrder = AircraftOrderDelegate().get(id).first();	
				# add the AircraftOrder
				operator.aircraftOrders.add(aircraftOrder)
				
			# save it		
			operator.save()
			
			# reload and return the appropriate version
			return self.get( operatorId );
		except Operator.DoesNotExist:
			raise ProcessingError(errMsg + " : Operator with id " + str(operatorId) + " does not exist.")
		except AircraftOrder.DoesNotExist:
			raise ProcessingError(errMsg + " : AircraftOrder does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeAircraftOrders( self, operatorId, aircraftOrdersIds ):
		# lazy importing avoids circular dependencies
		from aerospaceOnDjango.delegates.AircraftOrderDelegate import AircraftOrderDelegate

		errMsg = "Failed to remove elements " + str(aircraftOrdersIds) + " for AircraftOrders on Operator"

		try:
			# get the Operator
			operator = self.get( operatorId ).first()
				
			# split on a comma with no spaces
			idList = aircraftOrdersIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the AircraftOrder		
				aircraftOrder = AircraftOrderDelegate().get(id).first();	
				# add the AircraftOrder
				operator.aircraftOrders.remove(aircraftOrder)
				
			# save it		
			operator.save()
			
			# reload and return the appropriate version
			return self.get( operatorId );
		except Operator.DoesNotExist:
			raise ProcessingError(errMsg + " : Operator with id " + str(operatorId) + " does not exist.")
		except AircraftOrder.DoesNotExist:
			raise ProcessingError(errMsg + " : AircraftOrder does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addOperatedAircraft( self, operatorId, operatedAircraftIds ):
		# lazy importing avoids circular dependencies
		from aerospaceOnDjango.delegates.AircraftDelegate import AircraftDelegate

		errMsg = "Failed to add elements " + str(operatedAircraftIds) + " for OperatedAircraft on Operator"

		try:
			# get the Operator
			operator = self.get( operatorId ).first()
				
			# split on a comma with no spaces
			idList = operatedAircraftIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the Aircraft		
				aircraft = AircraftDelegate().get(id).first();	
				# add the Aircraft
				operator.operatedAircraft.add(aircraft)
				
			# save it		
			operator.save()
			
			# reload and return the appropriate version
			return self.get( operatorId );
		except Operator.DoesNotExist:
			raise ProcessingError(errMsg + " : Operator with id " + str(operatorId) + " does not exist.")
		except Aircraft.DoesNotExist:
			raise ProcessingError(errMsg + " : Aircraft does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeOperatedAircraft( self, operatorId, operatedAircraftIds ):
		# lazy importing avoids circular dependencies
		from aerospaceOnDjango.delegates.AircraftDelegate import AircraftDelegate

		errMsg = "Failed to remove elements " + str(operatedAircraftIds) + " for OperatedAircraft on Operator"

		try:
			# get the Operator
			operator = self.get( operatorId ).first()
				
			# split on a comma with no spaces
			idList = operatedAircraftIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the Aircraft		
				aircraft = AircraftDelegate().get(id).first();	
				# add the Aircraft
				operator.operatedAircraft.remove(aircraft)
				
			# save it		
			operator.save()
			
			# reload and return the appropriate version
			return self.get( operatorId );
		except Operator.DoesNotExist:
			raise ProcessingError(errMsg + " : Operator with id " + str(operatorId) + " does not exist.")
		except Aircraft.DoesNotExist:
			raise ProcessingError(errMsg + " : Aircraft does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
