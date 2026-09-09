from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from analyticsOnDjango.models.Prediction import Prediction
from analyticsOnDjango.models.InferenceEndpoint import InferenceEndpoint
from analyticsOnDjango.models.ModelVersion import ModelVersion
from analyticsOnDjango.models.DataSet import DataSet
from analyticsOnDjango.exceptions import Exceptions

 #======================================================================
# 
# Encapsulates data for model Prediction
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class PredictionDelegate Declaration
#======================================================================
class PredictionDelegate :

#======================================================================
# Function Declarations
#======================================================================

	def get(self, predictionId ):
		try:	
			prediction = Prediction.objects.filter(id=predictionId)
			return prediction.first();
		except Prediction.DoesNotExist:
			raise ProcessingError("Prediction with id " + str(predictionId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 

	def createFromJson(self, prediction):
		for model in serializers.deserialize("json", prediction):
			model.save()
			return model;

	def create(self, prediction):
		prediction.save()
		return prediction;

	def saveFromJson(self, prediction):
		for model in serializers.deserialize("json", prediction):
			model.save()
			return prediction;
	
	def save(self, prediction):
		prediction.save()
		return prediction;
	
	def delete(self, predictionId ):
		errMsg = "Failed to delete Prediction from db using id " + str(predictionId)
		
		try:
			prediction = Prediction.objects.get(id=predictionId)
			prediction.delete()
			return True
		except Prediction.DoesNotExist:
			raise ProcessingError("Prediction with id " + str(predictionId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 
	
	def getAll(self):
		try:
			all = Prediction.objects.all()
			return all;
		except utils.DatabaseError:
			raise StorageReadError("Failed to get all Prediction from db")
		except Exception:
			return None;
		
	def assignEndpoint( self, predictionId, endpointId ):
		# lazy importing avoids circular dependencies
		from analyticsOnDjango.delegates.InferenceEndpointDelegate import InferenceEndpointDelegate

		errMsg = "Failed to assign element " + str(endpointId) + " for Endpoint on Prediction"

		try:
			# get the Prediction from db
			prediction = self.get( predictionId ).first()	
			
			# get the InferenceEndpoint from db
			inferenceEndpoint = InferenceEndpointDelegate().get(endpointId).first();
			
			# assign the Endpoint		
			prediction.endpoint = inferenceEndpoint
			
			#save it
			prediction.save()

			# reload and return the appropriate version					
			return self.get( predictionId );
		except Prediction.DoesNotExist:
			raise ProcessingError(errMsg + " : Prediction with id " + str(predictionId) + " does not exist.")
		except InferenceEndpoint.DoesNotExist:
			raise ProcessingError(errMsg + " : InferenceEndpoint with id " + str(endpointId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignEndpoint( self, predictionId ):
		errMsg = "Failed to unassign element " + str(endpointId) + " for Endpoint on Prediction"

		try:
			# get the Prediction from db
			prediction = self.get( predictionId ).first()	
			
			# assign to None for unassignment
			prediction.inferenceEndpoint = None			

			#save it
			prediction.save()

			# reload and return the appropriate version					
			return self.get( predictionId );
		except Prediction.DoesNotExist:
			raise ProcessingError(errMsg + " : Prediction with id " + str(predictionId) + " does not exist.")
		except Exception:
			return None;
		
	def assignModelVersion( self, predictionId, modelVersionId ):
		# lazy importing avoids circular dependencies
		from analyticsOnDjango.delegates.ModelVersionDelegate import ModelVersionDelegate

		errMsg = "Failed to assign element " + str(modelVersionId) + " for ModelVersion on Prediction"

		try:
			# get the Prediction from db
			prediction = self.get( predictionId ).first()	
			
			# get the ModelVersion from db
			modelVersion = ModelVersionDelegate().get(modelVersionId).first();
			
			# assign the ModelVersion		
			prediction.modelVersion = modelVersion
			
			#save it
			prediction.save()

			# reload and return the appropriate version					
			return self.get( predictionId );
		except Prediction.DoesNotExist:
			raise ProcessingError(errMsg + " : Prediction with id " + str(predictionId) + " does not exist.")
		except ModelVersion.DoesNotExist:
			raise ProcessingError(errMsg + " : ModelVersion with id " + str(modelVersionId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignModelVersion( self, predictionId ):
		errMsg = "Failed to unassign element " + str(modelVersionId) + " for ModelVersion on Prediction"

		try:
			# get the Prediction from db
			prediction = self.get( predictionId ).first()	
			
			# assign to None for unassignment
			prediction.modelVersion = None			

			#save it
			prediction.save()

			# reload and return the appropriate version					
			return self.get( predictionId );
		except Prediction.DoesNotExist:
			raise ProcessingError(errMsg + " : Prediction with id " + str(predictionId) + " does not exist.")
		except Exception:
			return None;
		
	def assignDataset( self, predictionId, datasetId ):
		# lazy importing avoids circular dependencies
		from analyticsOnDjango.delegates.DataSetDelegate import DataSetDelegate

		errMsg = "Failed to assign element " + str(datasetId) + " for Dataset on Prediction"

		try:
			# get the Prediction from db
			prediction = self.get( predictionId ).first()	
			
			# get the DataSet from db
			dataSet = DataSetDelegate().get(datasetId).first();
			
			# assign the Dataset		
			prediction.dataset = dataSet
			
			#save it
			prediction.save()

			# reload and return the appropriate version					
			return self.get( predictionId );
		except Prediction.DoesNotExist:
			raise ProcessingError(errMsg + " : Prediction with id " + str(predictionId) + " does not exist.")
		except DataSet.DoesNotExist:
			raise ProcessingError(errMsg + " : DataSet with id " + str(datasetId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignDataset( self, predictionId ):
		errMsg = "Failed to unassign element " + str(datasetId) + " for Dataset on Prediction"

		try:
			# get the Prediction from db
			prediction = self.get( predictionId ).first()	
			
			# assign to None for unassignment
			prediction.dataSet = None			

			#save it
			prediction.save()

			# reload and return the appropriate version					
			return self.get( predictionId );
		except Prediction.DoesNotExist:
			raise ProcessingError(errMsg + " : Prediction with id " + str(predictionId) + " does not exist.")
		except Exception:
			return None;
		
