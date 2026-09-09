from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from analyticsOnDjango.models.FraudSignal import FraudSignal
from analyticsOnDjango.models.FraudScenario import FraudScenario
from analyticsOnDjango.models.DataSet import DataSet
from analyticsOnDjango.models.ModelVersion import ModelVersion
from analyticsOnDjango.exceptions import Exceptions

 #======================================================================
# 
# Encapsulates data for model FraudSignal
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class FraudSignalDelegate Declaration
#======================================================================
class FraudSignalDelegate :

#======================================================================
# Function Declarations
#======================================================================

	def get(self, fraudSignalId ):
		try:	
			fraudSignal = FraudSignal.objects.filter(id=fraudSignalId)
			return fraudSignal.first();
		except FraudSignal.DoesNotExist:
			raise ProcessingError("FraudSignal with id " + str(fraudSignalId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 

	def createFromJson(self, fraudSignal):
		for model in serializers.deserialize("json", fraudSignal):
			model.save()
			return model;

	def create(self, fraudSignal):
		fraudSignal.save()
		return fraudSignal;

	def saveFromJson(self, fraudSignal):
		for model in serializers.deserialize("json", fraudSignal):
			model.save()
			return fraudSignal;
	
	def save(self, fraudSignal):
		fraudSignal.save()
		return fraudSignal;
	
	def delete(self, fraudSignalId ):
		errMsg = "Failed to delete FraudSignal from db using id " + str(fraudSignalId)
		
		try:
			fraudSignal = FraudSignal.objects.get(id=fraudSignalId)
			fraudSignal.delete()
			return True
		except FraudSignal.DoesNotExist:
			raise ProcessingError("FraudSignal with id " + str(fraudSignalId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 
	
	def getAll(self):
		try:
			all = FraudSignal.objects.all()
			return all;
		except utils.DatabaseError:
			raise StorageReadError("Failed to get all FraudSignal from db")
		except Exception:
			return None;
		
	def assignScenario( self, fraudSignalId, scenarioId ):
		# lazy importing avoids circular dependencies
		from analyticsOnDjango.delegates.FraudScenarioDelegate import FraudScenarioDelegate

		errMsg = "Failed to assign element " + str(scenarioId) + " for Scenario on FraudSignal"

		try:
			# get the FraudSignal from db
			fraudSignal = self.get( fraudSignalId ).first()	
			
			# get the FraudScenario from db
			fraudScenario = FraudScenarioDelegate().get(scenarioId).first();
			
			# assign the Scenario		
			fraudSignal.scenario = fraudScenario
			
			#save it
			fraudSignal.save()

			# reload and return the appropriate version					
			return self.get( fraudSignalId );
		except FraudSignal.DoesNotExist:
			raise ProcessingError(errMsg + " : FraudSignal with id " + str(fraudSignalId) + " does not exist.")
		except FraudScenario.DoesNotExist:
			raise ProcessingError(errMsg + " : FraudScenario with id " + str(scenarioId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignScenario( self, fraudSignalId ):
		errMsg = "Failed to unassign element " + str(scenarioId) + " for Scenario on FraudSignal"

		try:
			# get the FraudSignal from db
			fraudSignal = self.get( fraudSignalId ).first()	
			
			# assign to None for unassignment
			fraudSignal.fraudScenario = None			

			#save it
			fraudSignal.save()

			# reload and return the appropriate version					
			return self.get( fraudSignalId );
		except FraudSignal.DoesNotExist:
			raise ProcessingError(errMsg + " : FraudSignal with id " + str(fraudSignalId) + " does not exist.")
		except Exception:
			return None;
		
	def assignDataset( self, fraudSignalId, datasetId ):
		# lazy importing avoids circular dependencies
		from analyticsOnDjango.delegates.DataSetDelegate import DataSetDelegate

		errMsg = "Failed to assign element " + str(datasetId) + " for Dataset on FraudSignal"

		try:
			# get the FraudSignal from db
			fraudSignal = self.get( fraudSignalId ).first()	
			
			# get the DataSet from db
			dataSet = DataSetDelegate().get(datasetId).first();
			
			# assign the Dataset		
			fraudSignal.dataset = dataSet
			
			#save it
			fraudSignal.save()

			# reload and return the appropriate version					
			return self.get( fraudSignalId );
		except FraudSignal.DoesNotExist:
			raise ProcessingError(errMsg + " : FraudSignal with id " + str(fraudSignalId) + " does not exist.")
		except DataSet.DoesNotExist:
			raise ProcessingError(errMsg + " : DataSet with id " + str(datasetId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignDataset( self, fraudSignalId ):
		errMsg = "Failed to unassign element " + str(datasetId) + " for Dataset on FraudSignal"

		try:
			# get the FraudSignal from db
			fraudSignal = self.get( fraudSignalId ).first()	
			
			# assign to None for unassignment
			fraudSignal.dataSet = None			

			#save it
			fraudSignal.save()

			# reload and return the appropriate version					
			return self.get( fraudSignalId );
		except FraudSignal.DoesNotExist:
			raise ProcessingError(errMsg + " : FraudSignal with id " + str(fraudSignalId) + " does not exist.")
		except Exception:
			return None;
		
	def assignModelVersion( self, fraudSignalId, modelVersionId ):
		# lazy importing avoids circular dependencies
		from analyticsOnDjango.delegates.ModelVersionDelegate import ModelVersionDelegate

		errMsg = "Failed to assign element " + str(modelVersionId) + " for ModelVersion on FraudSignal"

		try:
			# get the FraudSignal from db
			fraudSignal = self.get( fraudSignalId ).first()	
			
			# get the ModelVersion from db
			modelVersion = ModelVersionDelegate().get(modelVersionId).first();
			
			# assign the ModelVersion		
			fraudSignal.modelVersion = modelVersion
			
			#save it
			fraudSignal.save()

			# reload and return the appropriate version					
			return self.get( fraudSignalId );
		except FraudSignal.DoesNotExist:
			raise ProcessingError(errMsg + " : FraudSignal with id " + str(fraudSignalId) + " does not exist.")
		except ModelVersion.DoesNotExist:
			raise ProcessingError(errMsg + " : ModelVersion with id " + str(modelVersionId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignModelVersion( self, fraudSignalId ):
		errMsg = "Failed to unassign element " + str(modelVersionId) + " for ModelVersion on FraudSignal"

		try:
			# get the FraudSignal from db
			fraudSignal = self.get( fraudSignalId ).first()	
			
			# assign to None for unassignment
			fraudSignal.modelVersion = None			

			#save it
			fraudSignal.save()

			# reload and return the appropriate version					
			return self.get( fraudSignalId );
		except FraudSignal.DoesNotExist:
			raise ProcessingError(errMsg + " : FraudSignal with id " + str(fraudSignalId) + " does not exist.")
		except Exception:
			return None;
		
