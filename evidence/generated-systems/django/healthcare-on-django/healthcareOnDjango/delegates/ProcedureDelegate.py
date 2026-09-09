from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from healthcareOnDjango.models.Procedure import Procedure
from healthcareOnDjango.models.Encounter import Encounter
from healthcareOnDjango.models.Clinician import Clinician
from healthcareOnDjango.models.ProcedureOrder import ProcedureOrder
from healthcareOnDjango.exceptions import Exceptions

 #======================================================================
# 
# Encapsulates data for model Procedure
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class ProcedureDelegate Declaration
#======================================================================
class ProcedureDelegate :

#======================================================================
# Function Declarations
#======================================================================

	def get(self, procedureId ):
		try:	
			procedure = Procedure.objects.filter(id=procedureId)
			return procedure.first();
		except Procedure.DoesNotExist:
			raise ProcessingError("Procedure with id " + str(procedureId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 

	def createFromJson(self, procedure):
		for model in serializers.deserialize("json", procedure):
			model.save()
			return model;

	def create(self, procedure):
		procedure.save()
		return procedure;

	def saveFromJson(self, procedure):
		for model in serializers.deserialize("json", procedure):
			model.save()
			return procedure;
	
	def save(self, procedure):
		procedure.save()
		return procedure;
	
	def delete(self, procedureId ):
		errMsg = "Failed to delete Procedure from db using id " + str(procedureId)
		
		try:
			procedure = Procedure.objects.get(id=procedureId)
			procedure.delete()
			return True
		except Procedure.DoesNotExist:
			raise ProcessingError("Procedure with id " + str(procedureId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 
	
	def getAll(self):
		try:
			all = Procedure.objects.all()
			return all;
		except utils.DatabaseError:
			raise StorageReadError("Failed to get all Procedure from db")
		except Exception:
			return None;
		
	def assignEncounter( self, procedureId, encounterId ):
		# lazy importing avoids circular dependencies
		from healthcareOnDjango.delegates.EncounterDelegate import EncounterDelegate

		errMsg = "Failed to assign element " + str(encounterId) + " for Encounter on Procedure"

		try:
			# get the Procedure from db
			procedure = self.get( procedureId ).first()	
			
			# get the Encounter from db
			encounter = EncounterDelegate().get(encounterId).first();
			
			# assign the Encounter		
			procedure.encounter = encounter
			
			#save it
			procedure.save()

			# reload and return the appropriate version					
			return self.get( procedureId );
		except Procedure.DoesNotExist:
			raise ProcessingError(errMsg + " : Procedure with id " + str(procedureId) + " does not exist.")
		except Encounter.DoesNotExist:
			raise ProcessingError(errMsg + " : Encounter with id " + str(encounterId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignEncounter( self, procedureId ):
		errMsg = "Failed to unassign element " + str(encounterId) + " for Encounter on Procedure"

		try:
			# get the Procedure from db
			procedure = self.get( procedureId ).first()	
			
			# assign to None for unassignment
			procedure.encounter = None			

			#save it
			procedure.save()

			# reload and return the appropriate version					
			return self.get( procedureId );
		except Procedure.DoesNotExist:
			raise ProcessingError(errMsg + " : Procedure with id " + str(procedureId) + " does not exist.")
		except Exception:
			return None;
		
	def assignPerformer( self, procedureId, performerId ):
		# lazy importing avoids circular dependencies
		from healthcareOnDjango.delegates.ClinicianDelegate import ClinicianDelegate

		errMsg = "Failed to assign element " + str(performerId) + " for Performer on Procedure"

		try:
			# get the Procedure from db
			procedure = self.get( procedureId ).first()	
			
			# get the Clinician from db
			clinician = ClinicianDelegate().get(performerId).first();
			
			# assign the Performer		
			procedure.performer = clinician
			
			#save it
			procedure.save()

			# reload and return the appropriate version					
			return self.get( procedureId );
		except Procedure.DoesNotExist:
			raise ProcessingError(errMsg + " : Procedure with id " + str(procedureId) + " does not exist.")
		except Clinician.DoesNotExist:
			raise ProcessingError(errMsg + " : Clinician with id " + str(performerId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignPerformer( self, procedureId ):
		errMsg = "Failed to unassign element " + str(performerId) + " for Performer on Procedure"

		try:
			# get the Procedure from db
			procedure = self.get( procedureId ).first()	
			
			# assign to None for unassignment
			procedure.clinician = None			

			#save it
			procedure.save()

			# reload and return the appropriate version					
			return self.get( procedureId );
		except Procedure.DoesNotExist:
			raise ProcessingError(errMsg + " : Procedure with id " + str(procedureId) + " does not exist.")
		except Exception:
			return None;
		
	def assignProcedureOrder( self, procedureId, procedureOrderId ):
		# lazy importing avoids circular dependencies
		from healthcareOnDjango.delegates.ProcedureOrderDelegate import ProcedureOrderDelegate

		errMsg = "Failed to assign element " + str(procedureOrderId) + " for ProcedureOrder on Procedure"

		try:
			# get the Procedure from db
			procedure = self.get( procedureId ).first()	
			
			# get the ProcedureOrder from db
			procedureOrder = ProcedureOrderDelegate().get(procedureOrderId).first();
			
			# assign the ProcedureOrder		
			procedure.procedureOrder = procedureOrder
			
			#save it
			procedure.save()

			# reload and return the appropriate version					
			return self.get( procedureId );
		except Procedure.DoesNotExist:
			raise ProcessingError(errMsg + " : Procedure with id " + str(procedureId) + " does not exist.")
		except ProcedureOrder.DoesNotExist:
			raise ProcessingError(errMsg + " : ProcedureOrder with id " + str(procedureOrderId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignProcedureOrder( self, procedureId ):
		errMsg = "Failed to unassign element " + str(procedureOrderId) + " for ProcedureOrder on Procedure"

		try:
			# get the Procedure from db
			procedure = self.get( procedureId ).first()	
			
			# assign to None for unassignment
			procedure.procedureOrder = None			

			#save it
			procedure.save()

			# reload and return the appropriate version					
			return self.get( procedureId );
		except Procedure.DoesNotExist:
			raise ProcessingError(errMsg + " : Procedure with id " + str(procedureId) + " does not exist.")
		except Exception:
			return None;
		
