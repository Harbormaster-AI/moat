from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from governanceOnDjango.models.Procedure import Procedure
from governanceOnDjango.models.Policy import Policy
from governanceOnDjango.models.Control import Control
from governanceOnDjango.exceptions import Exceptions

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
		
	def assignPolicy( self, procedureId, policyId ):
		# lazy importing avoids circular dependencies
		from governanceOnDjango.delegates.PolicyDelegate import PolicyDelegate

		errMsg = "Failed to assign element " + str(policyId) + " for Policy on Procedure"

		try:
			# get the Procedure from db
			procedure = self.get( procedureId ).first()	
			
			# get the Policy from db
			policy = PolicyDelegate().get(policyId).first();
			
			# assign the Policy		
			procedure.policy = policy
			
			#save it
			procedure.save()

			# reload and return the appropriate version					
			return self.get( procedureId );
		except Procedure.DoesNotExist:
			raise ProcessingError(errMsg + " : Procedure with id " + str(procedureId) + " does not exist.")
		except Policy.DoesNotExist:
			raise ProcessingError(errMsg + " : Policy with id " + str(policyId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignPolicy( self, procedureId ):
		errMsg = "Failed to unassign element " + str(policyId) + " for Policy on Procedure"

		try:
			# get the Procedure from db
			procedure = self.get( procedureId ).first()	
			
			# assign to None for unassignment
			procedure.policy = None			

			#save it
			procedure.save()

			# reload and return the appropriate version					
			return self.get( procedureId );
		except Procedure.DoesNotExist:
			raise ProcessingError(errMsg + " : Procedure with id " + str(procedureId) + " does not exist.")
		except Exception:
			return None;
		
	def addControls( self, procedureId, controlsIds ):
		# lazy importing avoids circular dependencies
		from governanceOnDjango.delegates.ControlDelegate import ControlDelegate

		errMsg = "Failed to add elements " + str(controlsIds) + " for Controls on Procedure"

		try:
			# get the Procedure
			procedure = self.get( procedureId ).first()
				
			# split on a comma with no spaces
			idList = controlsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the Control		
				control = ControlDelegate().get(id).first();	
				# add the Control
				procedure.controls.add(control)
				
			# save it		
			procedure.save()
			
			# reload and return the appropriate version
			return self.get( procedureId );
		except Procedure.DoesNotExist:
			raise ProcessingError(errMsg + " : Procedure with id " + str(procedureId) + " does not exist.")
		except Control.DoesNotExist:
			raise ProcessingError(errMsg + " : Control does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeControls( self, procedureId, controlsIds ):
		# lazy importing avoids circular dependencies
		from governanceOnDjango.delegates.ControlDelegate import ControlDelegate

		errMsg = "Failed to remove elements " + str(controlsIds) + " for Controls on Procedure"

		try:
			# get the Procedure
			procedure = self.get( procedureId ).first()
				
			# split on a comma with no spaces
			idList = controlsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the Control		
				control = ControlDelegate().get(id).first();	
				# add the Control
				procedure.controls.remove(control)
				
			# save it		
			procedure.save()
			
			# reload and return the appropriate version
			return self.get( procedureId );
		except Procedure.DoesNotExist:
			raise ProcessingError(errMsg + " : Procedure with id " + str(procedureId) + " does not exist.")
		except Control.DoesNotExist:
			raise ProcessingError(errMsg + " : Control does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
