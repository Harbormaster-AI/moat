from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from healthcareOnDjango.models.ProcedureOrder import ProcedureOrder
from healthcareOnDjango.models.ClinicalOrder import ClinicalOrder
from healthcareOnDjango.models.Facility import Facility
from healthcareOnDjango.models.Procedure import Procedure
from healthcareOnDjango.exceptions import Exceptions

 #======================================================================
# 
# Encapsulates data for model ProcedureOrder
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class ProcedureOrderDelegate Declaration
#======================================================================
class ProcedureOrderDelegate :

#======================================================================
# Function Declarations
#======================================================================

	def get(self, procedureOrderId ):
		try:	
			procedureOrder = ProcedureOrder.objects.filter(id=procedureOrderId)
			return procedureOrder.first();
		except ProcedureOrder.DoesNotExist:
			raise ProcessingError("ProcedureOrder with id " + str(procedureOrderId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 

	def createFromJson(self, procedureOrder):
		for model in serializers.deserialize("json", procedureOrder):
			model.save()
			return model;

	def create(self, procedureOrder):
		procedureOrder.save()
		return procedureOrder;

	def saveFromJson(self, procedureOrder):
		for model in serializers.deserialize("json", procedureOrder):
			model.save()
			return procedureOrder;
	
	def save(self, procedureOrder):
		procedureOrder.save()
		return procedureOrder;
	
	def delete(self, procedureOrderId ):
		errMsg = "Failed to delete ProcedureOrder from db using id " + str(procedureOrderId)
		
		try:
			procedureOrder = ProcedureOrder.objects.get(id=procedureOrderId)
			procedureOrder.delete()
			return True
		except ProcedureOrder.DoesNotExist:
			raise ProcessingError("ProcedureOrder with id " + str(procedureOrderId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 
	
	def getAll(self):
		try:
			all = ProcedureOrder.objects.all()
			return all;
		except utils.DatabaseError:
			raise StorageReadError("Failed to get all ProcedureOrder from db")
		except Exception:
			return None;
		
	def assignOrder( self, procedureOrderId, orderId ):
		# lazy importing avoids circular dependencies
		from healthcareOnDjango.delegates.ClinicalOrderDelegate import ClinicalOrderDelegate

		errMsg = "Failed to assign element " + str(orderId) + " for Order on ProcedureOrder"

		try:
			# get the ProcedureOrder from db
			procedureOrder = self.get( procedureOrderId ).first()	
			
			# get the ClinicalOrder from db
			clinicalOrder = ClinicalOrderDelegate().get(orderId).first();
			
			# assign the Order		
			procedureOrder.order = clinicalOrder
			
			#save it
			procedureOrder.save()

			# reload and return the appropriate version					
			return self.get( procedureOrderId );
		except ProcedureOrder.DoesNotExist:
			raise ProcessingError(errMsg + " : ProcedureOrder with id " + str(procedureOrderId) + " does not exist.")
		except ClinicalOrder.DoesNotExist:
			raise ProcessingError(errMsg + " : ClinicalOrder with id " + str(orderId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignOrder( self, procedureOrderId ):
		errMsg = "Failed to unassign element " + str(orderId) + " for Order on ProcedureOrder"

		try:
			# get the ProcedureOrder from db
			procedureOrder = self.get( procedureOrderId ).first()	
			
			# assign to None for unassignment
			procedureOrder.clinicalOrder = None			

			#save it
			procedureOrder.save()

			# reload and return the appropriate version					
			return self.get( procedureOrderId );
		except ProcedureOrder.DoesNotExist:
			raise ProcessingError(errMsg + " : ProcedureOrder with id " + str(procedureOrderId) + " does not exist.")
		except Exception:
			return None;
		
	def assignFacility( self, procedureOrderId, facilityId ):
		# lazy importing avoids circular dependencies
		from healthcareOnDjango.delegates.FacilityDelegate import FacilityDelegate

		errMsg = "Failed to assign element " + str(facilityId) + " for Facility on ProcedureOrder"

		try:
			# get the ProcedureOrder from db
			procedureOrder = self.get( procedureOrderId ).first()	
			
			# get the Facility from db
			facility = FacilityDelegate().get(facilityId).first();
			
			# assign the Facility		
			procedureOrder.facility = facility
			
			#save it
			procedureOrder.save()

			# reload and return the appropriate version					
			return self.get( procedureOrderId );
		except ProcedureOrder.DoesNotExist:
			raise ProcessingError(errMsg + " : ProcedureOrder with id " + str(procedureOrderId) + " does not exist.")
		except Facility.DoesNotExist:
			raise ProcessingError(errMsg + " : Facility with id " + str(facilityId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignFacility( self, procedureOrderId ):
		errMsg = "Failed to unassign element " + str(facilityId) + " for Facility on ProcedureOrder"

		try:
			# get the ProcedureOrder from db
			procedureOrder = self.get( procedureOrderId ).first()	
			
			# assign to None for unassignment
			procedureOrder.facility = None			

			#save it
			procedureOrder.save()

			# reload and return the appropriate version					
			return self.get( procedureOrderId );
		except ProcedureOrder.DoesNotExist:
			raise ProcessingError(errMsg + " : ProcedureOrder with id " + str(procedureOrderId) + " does not exist.")
		except Exception:
			return None;
		
	def assignProcedure( self, procedureOrderId, procedureId ):
		# lazy importing avoids circular dependencies
		from healthcareOnDjango.delegates.ProcedureDelegate import ProcedureDelegate

		errMsg = "Failed to assign element " + str(procedureId) + " for Procedure on ProcedureOrder"

		try:
			# get the ProcedureOrder from db
			procedureOrder = self.get( procedureOrderId ).first()	
			
			# get the Procedure from db
			procedure = ProcedureDelegate().get(procedureId).first();
			
			# assign the Procedure		
			procedureOrder.procedure = procedure
			
			#save it
			procedureOrder.save()

			# reload and return the appropriate version					
			return self.get( procedureOrderId );
		except ProcedureOrder.DoesNotExist:
			raise ProcessingError(errMsg + " : ProcedureOrder with id " + str(procedureOrderId) + " does not exist.")
		except Procedure.DoesNotExist:
			raise ProcessingError(errMsg + " : Procedure with id " + str(procedureId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignProcedure( self, procedureOrderId ):
		errMsg = "Failed to unassign element " + str(procedureId) + " for Procedure on ProcedureOrder"

		try:
			# get the ProcedureOrder from db
			procedureOrder = self.get( procedureOrderId ).first()	
			
			# assign to None for unassignment
			procedureOrder.procedure = None			

			#save it
			procedureOrder.save()

			# reload and return the appropriate version					
			return self.get( procedureOrderId );
		except ProcedureOrder.DoesNotExist:
			raise ProcessingError(errMsg + " : ProcedureOrder with id " + str(procedureOrderId) + " does not exist.")
		except Exception:
			return None;
		
