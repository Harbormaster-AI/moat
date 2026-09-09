from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from aerospaceOnDjango.models.Aircraft import Aircraft
from aerospaceOnDjango.models.AircraftVariant import AircraftVariant
from aerospaceOnDjango.models.Operator import Operator
from aerospaceOnDjango.models.Registration import Registration
from aerospaceOnDjango.models.Warranty import Warranty
from aerospaceOnDjango.models.MaintenanceWorkOrder import MaintenanceWorkOrder
from aerospaceOnDjango.models.ConnectedAircraft import ConnectedAircraft
from aerospaceOnDjango.models.CabinLayout import CabinLayout
from aerospaceOnDjango.exceptions import Exceptions

 #======================================================================
# 
# Encapsulates data for model Aircraft
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class AircraftDelegate Declaration
#======================================================================
class AircraftDelegate :

#======================================================================
# Function Declarations
#======================================================================

	def get(self, aircraftId ):
		try:	
			aircraft = Aircraft.objects.filter(id=aircraftId)
			return aircraft.first();
		except Aircraft.DoesNotExist:
			raise ProcessingError("Aircraft with id " + str(aircraftId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 

	def createFromJson(self, aircraft):
		for model in serializers.deserialize("json", aircraft):
			model.save()
			return model;

	def create(self, aircraft):
		aircraft.save()
		return aircraft;

	def saveFromJson(self, aircraft):
		for model in serializers.deserialize("json", aircraft):
			model.save()
			return aircraft;
	
	def save(self, aircraft):
		aircraft.save()
		return aircraft;
	
	def delete(self, aircraftId ):
		errMsg = "Failed to delete Aircraft from db using id " + str(aircraftId)
		
		try:
			aircraft = Aircraft.objects.get(id=aircraftId)
			aircraft.delete()
			return True
		except Aircraft.DoesNotExist:
			raise ProcessingError("Aircraft with id " + str(aircraftId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 
	
	def getAll(self):
		try:
			all = Aircraft.objects.all()
			return all;
		except utils.DatabaseError:
			raise StorageReadError("Failed to get all Aircraft from db")
		except Exception:
			return None;
		
	def assignVariant( self, aircraftId, variantId ):
		# lazy importing avoids circular dependencies
		from aerospaceOnDjango.delegates.AircraftVariantDelegate import AircraftVariantDelegate

		errMsg = "Failed to assign element " + str(variantId) + " for Variant on Aircraft"

		try:
			# get the Aircraft from db
			aircraft = self.get( aircraftId ).first()	
			
			# get the AircraftVariant from db
			aircraftVariant = AircraftVariantDelegate().get(variantId).first();
			
			# assign the Variant		
			aircraft.variant = aircraftVariant
			
			#save it
			aircraft.save()

			# reload and return the appropriate version					
			return self.get( aircraftId );
		except Aircraft.DoesNotExist:
			raise ProcessingError(errMsg + " : Aircraft with id " + str(aircraftId) + " does not exist.")
		except AircraftVariant.DoesNotExist:
			raise ProcessingError(errMsg + " : AircraftVariant with id " + str(variantId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignVariant( self, aircraftId ):
		errMsg = "Failed to unassign element " + str(variantId) + " for Variant on Aircraft"

		try:
			# get the Aircraft from db
			aircraft = self.get( aircraftId ).first()	
			
			# assign to None for unassignment
			aircraft.aircraftVariant = None			

			#save it
			aircraft.save()

			# reload and return the appropriate version					
			return self.get( aircraftId );
		except Aircraft.DoesNotExist:
			raise ProcessingError(errMsg + " : Aircraft with id " + str(aircraftId) + " does not exist.")
		except Exception:
			return None;
		
	def assignOperator( self, aircraftId, operatorId ):
		# lazy importing avoids circular dependencies
		from aerospaceOnDjango.delegates.OperatorDelegate import OperatorDelegate

		errMsg = "Failed to assign element " + str(operatorId) + " for Operator on Aircraft"

		try:
			# get the Aircraft from db
			aircraft = self.get( aircraftId ).first()	
			
			# get the Operator from db
			operator = OperatorDelegate().get(operatorId).first();
			
			# assign the Operator		
			aircraft.operator = operator
			
			#save it
			aircraft.save()

			# reload and return the appropriate version					
			return self.get( aircraftId );
		except Aircraft.DoesNotExist:
			raise ProcessingError(errMsg + " : Aircraft with id " + str(aircraftId) + " does not exist.")
		except Operator.DoesNotExist:
			raise ProcessingError(errMsg + " : Operator with id " + str(operatorId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignOperator( self, aircraftId ):
		errMsg = "Failed to unassign element " + str(operatorId) + " for Operator on Aircraft"

		try:
			# get the Aircraft from db
			aircraft = self.get( aircraftId ).first()	
			
			# assign to None for unassignment
			aircraft.operator = None			

			#save it
			aircraft.save()

			# reload and return the appropriate version					
			return self.get( aircraftId );
		except Aircraft.DoesNotExist:
			raise ProcessingError(errMsg + " : Aircraft with id " + str(aircraftId) + " does not exist.")
		except Exception:
			return None;
		
	def assignRegistration( self, aircraftId, registrationId ):
		# lazy importing avoids circular dependencies
		from aerospaceOnDjango.delegates.RegistrationDelegate import RegistrationDelegate

		errMsg = "Failed to assign element " + str(registrationId) + " for Registration on Aircraft"

		try:
			# get the Aircraft from db
			aircraft = self.get( aircraftId ).first()	
			
			# get the Registration from db
			registration = RegistrationDelegate().get(registrationId).first();
			
			# assign the Registration		
			aircraft.registration = registration
			
			#save it
			aircraft.save()

			# reload and return the appropriate version					
			return self.get( aircraftId );
		except Aircraft.DoesNotExist:
			raise ProcessingError(errMsg + " : Aircraft with id " + str(aircraftId) + " does not exist.")
		except Registration.DoesNotExist:
			raise ProcessingError(errMsg + " : Registration with id " + str(registrationId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignRegistration( self, aircraftId ):
		errMsg = "Failed to unassign element " + str(registrationId) + " for Registration on Aircraft"

		try:
			# get the Aircraft from db
			aircraft = self.get( aircraftId ).first()	
			
			# assign to None for unassignment
			aircraft.registration = None			

			#save it
			aircraft.save()

			# reload and return the appropriate version					
			return self.get( aircraftId );
		except Aircraft.DoesNotExist:
			raise ProcessingError(errMsg + " : Aircraft with id " + str(aircraftId) + " does not exist.")
		except Exception:
			return None;
		
	def assignWarranty( self, aircraftId, warrantyId ):
		# lazy importing avoids circular dependencies
		from aerospaceOnDjango.delegates.WarrantyDelegate import WarrantyDelegate

		errMsg = "Failed to assign element " + str(warrantyId) + " for Warranty on Aircraft"

		try:
			# get the Aircraft from db
			aircraft = self.get( aircraftId ).first()	
			
			# get the Warranty from db
			warranty = WarrantyDelegate().get(warrantyId).first();
			
			# assign the Warranty		
			aircraft.warranty = warranty
			
			#save it
			aircraft.save()

			# reload and return the appropriate version					
			return self.get( aircraftId );
		except Aircraft.DoesNotExist:
			raise ProcessingError(errMsg + " : Aircraft with id " + str(aircraftId) + " does not exist.")
		except Warranty.DoesNotExist:
			raise ProcessingError(errMsg + " : Warranty with id " + str(warrantyId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignWarranty( self, aircraftId ):
		errMsg = "Failed to unassign element " + str(warrantyId) + " for Warranty on Aircraft"

		try:
			# get the Aircraft from db
			aircraft = self.get( aircraftId ).first()	
			
			# assign to None for unassignment
			aircraft.warranty = None			

			#save it
			aircraft.save()

			# reload and return the appropriate version					
			return self.get( aircraftId );
		except Aircraft.DoesNotExist:
			raise ProcessingError(errMsg + " : Aircraft with id " + str(aircraftId) + " does not exist.")
		except Exception:
			return None;
		
	def assignConnectedAircraft( self, aircraftId, connectedAircraftId ):
		# lazy importing avoids circular dependencies
		from aerospaceOnDjango.delegates.ConnectedAircraftDelegate import ConnectedAircraftDelegate

		errMsg = "Failed to assign element " + str(connectedAircraftId) + " for ConnectedAircraft on Aircraft"

		try:
			# get the Aircraft from db
			aircraft = self.get( aircraftId ).first()	
			
			# get the ConnectedAircraft from db
			connectedAircraft = ConnectedAircraftDelegate().get(connectedAircraftId).first();
			
			# assign the ConnectedAircraft		
			aircraft.connectedAircraft = connectedAircraft
			
			#save it
			aircraft.save()

			# reload and return the appropriate version					
			return self.get( aircraftId );
		except Aircraft.DoesNotExist:
			raise ProcessingError(errMsg + " : Aircraft with id " + str(aircraftId) + " does not exist.")
		except ConnectedAircraft.DoesNotExist:
			raise ProcessingError(errMsg + " : ConnectedAircraft with id " + str(connectedAircraftId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignConnectedAircraft( self, aircraftId ):
		errMsg = "Failed to unassign element " + str(connectedAircraftId) + " for ConnectedAircraft on Aircraft"

		try:
			# get the Aircraft from db
			aircraft = self.get( aircraftId ).first()	
			
			# assign to None for unassignment
			aircraft.connectedAircraft = None			

			#save it
			aircraft.save()

			# reload and return the appropriate version					
			return self.get( aircraftId );
		except Aircraft.DoesNotExist:
			raise ProcessingError(errMsg + " : Aircraft with id " + str(aircraftId) + " does not exist.")
		except Exception:
			return None;
		
	def assignCabinLayout( self, aircraftId, cabinLayoutId ):
		# lazy importing avoids circular dependencies
		from aerospaceOnDjango.delegates.CabinLayoutDelegate import CabinLayoutDelegate

		errMsg = "Failed to assign element " + str(cabinLayoutId) + " for CabinLayout on Aircraft"

		try:
			# get the Aircraft from db
			aircraft = self.get( aircraftId ).first()	
			
			# get the CabinLayout from db
			cabinLayout = CabinLayoutDelegate().get(cabinLayoutId).first();
			
			# assign the CabinLayout		
			aircraft.cabinLayout = cabinLayout
			
			#save it
			aircraft.save()

			# reload and return the appropriate version					
			return self.get( aircraftId );
		except Aircraft.DoesNotExist:
			raise ProcessingError(errMsg + " : Aircraft with id " + str(aircraftId) + " does not exist.")
		except CabinLayout.DoesNotExist:
			raise ProcessingError(errMsg + " : CabinLayout with id " + str(cabinLayoutId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignCabinLayout( self, aircraftId ):
		errMsg = "Failed to unassign element " + str(cabinLayoutId) + " for CabinLayout on Aircraft"

		try:
			# get the Aircraft from db
			aircraft = self.get( aircraftId ).first()	
			
			# assign to None for unassignment
			aircraft.cabinLayout = None			

			#save it
			aircraft.save()

			# reload and return the appropriate version					
			return self.get( aircraftId );
		except Aircraft.DoesNotExist:
			raise ProcessingError(errMsg + " : Aircraft with id " + str(aircraftId) + " does not exist.")
		except Exception:
			return None;
		
	def addMaintenanceRecords( self, aircraftId, maintenanceRecordsIds ):
		# lazy importing avoids circular dependencies
		from aerospaceOnDjango.delegates.MaintenanceWorkOrderDelegate import MaintenanceWorkOrderDelegate

		errMsg = "Failed to add elements " + str(maintenanceRecordsIds) + " for MaintenanceRecords on Aircraft"

		try:
			# get the Aircraft
			aircraft = self.get( aircraftId ).first()
				
			# split on a comma with no spaces
			idList = maintenanceRecordsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the MaintenanceWorkOrder		
				maintenanceWorkOrder = MaintenanceWorkOrderDelegate().get(id).first();	
				# add the MaintenanceWorkOrder
				aircraft.maintenanceRecords.add(maintenanceWorkOrder)
				
			# save it		
			aircraft.save()
			
			# reload and return the appropriate version
			return self.get( aircraftId );
		except Aircraft.DoesNotExist:
			raise ProcessingError(errMsg + " : Aircraft with id " + str(aircraftId) + " does not exist.")
		except MaintenanceWorkOrder.DoesNotExist:
			raise ProcessingError(errMsg + " : MaintenanceWorkOrder does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeMaintenanceRecords( self, aircraftId, maintenanceRecordsIds ):
		# lazy importing avoids circular dependencies
		from aerospaceOnDjango.delegates.MaintenanceWorkOrderDelegate import MaintenanceWorkOrderDelegate

		errMsg = "Failed to remove elements " + str(maintenanceRecordsIds) + " for MaintenanceRecords on Aircraft"

		try:
			# get the Aircraft
			aircraft = self.get( aircraftId ).first()
				
			# split on a comma with no spaces
			idList = maintenanceRecordsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the MaintenanceWorkOrder		
				maintenanceWorkOrder = MaintenanceWorkOrderDelegate().get(id).first();	
				# add the MaintenanceWorkOrder
				aircraft.maintenanceRecords.remove(maintenanceWorkOrder)
				
			# save it		
			aircraft.save()
			
			# reload and return the appropriate version
			return self.get( aircraftId );
		except Aircraft.DoesNotExist:
			raise ProcessingError(errMsg + " : Aircraft with id " + str(aircraftId) + " does not exist.")
		except MaintenanceWorkOrder.DoesNotExist:
			raise ProcessingError(errMsg + " : MaintenanceWorkOrder does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
