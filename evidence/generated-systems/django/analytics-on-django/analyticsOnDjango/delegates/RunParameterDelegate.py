from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from analyticsOnDjango.models.RunParameter import RunParameter
from analyticsOnDjango.models.TrainingRun import TrainingRun
from analyticsOnDjango.exceptions import Exceptions

 #======================================================================
# 
# Encapsulates data for model RunParameter
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class RunParameterDelegate Declaration
#======================================================================
class RunParameterDelegate :

#======================================================================
# Function Declarations
#======================================================================

	def get(self, runParameterId ):
		try:	
			runParameter = RunParameter.objects.filter(id=runParameterId)
			return runParameter.first();
		except RunParameter.DoesNotExist:
			raise ProcessingError("RunParameter with id " + str(runParameterId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 

	def createFromJson(self, runParameter):
		for model in serializers.deserialize("json", runParameter):
			model.save()
			return model;

	def create(self, runParameter):
		runParameter.save()
		return runParameter;

	def saveFromJson(self, runParameter):
		for model in serializers.deserialize("json", runParameter):
			model.save()
			return runParameter;
	
	def save(self, runParameter):
		runParameter.save()
		return runParameter;
	
	def delete(self, runParameterId ):
		errMsg = "Failed to delete RunParameter from db using id " + str(runParameterId)
		
		try:
			runParameter = RunParameter.objects.get(id=runParameterId)
			runParameter.delete()
			return True
		except RunParameter.DoesNotExist:
			raise ProcessingError("RunParameter with id " + str(runParameterId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 
	
	def getAll(self):
		try:
			all = RunParameter.objects.all()
			return all;
		except utils.DatabaseError:
			raise StorageReadError("Failed to get all RunParameter from db")
		except Exception:
			return None;
		
	def assignTrainingRun( self, runParameterId, trainingRunId ):
		# lazy importing avoids circular dependencies
		from analyticsOnDjango.delegates.TrainingRunDelegate import TrainingRunDelegate

		errMsg = "Failed to assign element " + str(trainingRunId) + " for TrainingRun on RunParameter"

		try:
			# get the RunParameter from db
			runParameter = self.get( runParameterId ).first()	
			
			# get the TrainingRun from db
			trainingRun = TrainingRunDelegate().get(trainingRunId).first();
			
			# assign the TrainingRun		
			runParameter.trainingRun = trainingRun
			
			#save it
			runParameter.save()

			# reload and return the appropriate version					
			return self.get( runParameterId );
		except RunParameter.DoesNotExist:
			raise ProcessingError(errMsg + " : RunParameter with id " + str(runParameterId) + " does not exist.")
		except TrainingRun.DoesNotExist:
			raise ProcessingError(errMsg + " : TrainingRun with id " + str(trainingRunId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignTrainingRun( self, runParameterId ):
		errMsg = "Failed to unassign element " + str(trainingRunId) + " for TrainingRun on RunParameter"

		try:
			# get the RunParameter from db
			runParameter = self.get( runParameterId ).first()	
			
			# assign to None for unassignment
			runParameter.trainingRun = None			

			#save it
			runParameter.save()

			# reload and return the appropriate version					
			return self.get( runParameterId );
		except RunParameter.DoesNotExist:
			raise ProcessingError(errMsg + " : RunParameter with id " + str(runParameterId) + " does not exist.")
		except Exception:
			return None;
		
