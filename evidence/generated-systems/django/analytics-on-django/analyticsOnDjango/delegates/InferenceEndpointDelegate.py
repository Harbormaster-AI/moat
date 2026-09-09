from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from analyticsOnDjango.models.InferenceEndpoint import InferenceEndpoint
from analyticsOnDjango.models.ModelVersion import ModelVersion
from analyticsOnDjango.models.AnalyticsWorkspace import AnalyticsWorkspace
from analyticsOnDjango.models.Prediction import Prediction
from analyticsOnDjango.exceptions import Exceptions

 #======================================================================
# 
# Encapsulates data for model InferenceEndpoint
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class InferenceEndpointDelegate Declaration
#======================================================================
class InferenceEndpointDelegate :

#======================================================================
# Function Declarations
#======================================================================

	def get(self, inferenceEndpointId ):
		try:	
			inferenceEndpoint = InferenceEndpoint.objects.filter(id=inferenceEndpointId)
			return inferenceEndpoint.first();
		except InferenceEndpoint.DoesNotExist:
			raise ProcessingError("InferenceEndpoint with id " + str(inferenceEndpointId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 

	def createFromJson(self, inferenceEndpoint):
		for model in serializers.deserialize("json", inferenceEndpoint):
			model.save()
			return model;

	def create(self, inferenceEndpoint):
		inferenceEndpoint.save()
		return inferenceEndpoint;

	def saveFromJson(self, inferenceEndpoint):
		for model in serializers.deserialize("json", inferenceEndpoint):
			model.save()
			return inferenceEndpoint;
	
	def save(self, inferenceEndpoint):
		inferenceEndpoint.save()
		return inferenceEndpoint;
	
	def delete(self, inferenceEndpointId ):
		errMsg = "Failed to delete InferenceEndpoint from db using id " + str(inferenceEndpointId)
		
		try:
			inferenceEndpoint = InferenceEndpoint.objects.get(id=inferenceEndpointId)
			inferenceEndpoint.delete()
			return True
		except InferenceEndpoint.DoesNotExist:
			raise ProcessingError("InferenceEndpoint with id " + str(inferenceEndpointId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 
	
	def getAll(self):
		try:
			all = InferenceEndpoint.objects.all()
			return all;
		except utils.DatabaseError:
			raise StorageReadError("Failed to get all InferenceEndpoint from db")
		except Exception:
			return None;
		
	def assignModelVersion( self, inferenceEndpointId, modelVersionId ):
		# lazy importing avoids circular dependencies
		from analyticsOnDjango.delegates.ModelVersionDelegate import ModelVersionDelegate

		errMsg = "Failed to assign element " + str(modelVersionId) + " for ModelVersion on InferenceEndpoint"

		try:
			# get the InferenceEndpoint from db
			inferenceEndpoint = self.get( inferenceEndpointId ).first()	
			
			# get the ModelVersion from db
			modelVersion = ModelVersionDelegate().get(modelVersionId).first();
			
			# assign the ModelVersion		
			inferenceEndpoint.modelVersion = modelVersion
			
			#save it
			inferenceEndpoint.save()

			# reload and return the appropriate version					
			return self.get( inferenceEndpointId );
		except InferenceEndpoint.DoesNotExist:
			raise ProcessingError(errMsg + " : InferenceEndpoint with id " + str(inferenceEndpointId) + " does not exist.")
		except ModelVersion.DoesNotExist:
			raise ProcessingError(errMsg + " : ModelVersion with id " + str(modelVersionId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignModelVersion( self, inferenceEndpointId ):
		errMsg = "Failed to unassign element " + str(modelVersionId) + " for ModelVersion on InferenceEndpoint"

		try:
			# get the InferenceEndpoint from db
			inferenceEndpoint = self.get( inferenceEndpointId ).first()	
			
			# assign to None for unassignment
			inferenceEndpoint.modelVersion = None			

			#save it
			inferenceEndpoint.save()

			# reload and return the appropriate version					
			return self.get( inferenceEndpointId );
		except InferenceEndpoint.DoesNotExist:
			raise ProcessingError(errMsg + " : InferenceEndpoint with id " + str(inferenceEndpointId) + " does not exist.")
		except Exception:
			return None;
		
	def assignWorkspace( self, inferenceEndpointId, workspaceId ):
		# lazy importing avoids circular dependencies
		from analyticsOnDjango.delegates.AnalyticsWorkspaceDelegate import AnalyticsWorkspaceDelegate

		errMsg = "Failed to assign element " + str(workspaceId) + " for Workspace on InferenceEndpoint"

		try:
			# get the InferenceEndpoint from db
			inferenceEndpoint = self.get( inferenceEndpointId ).first()	
			
			# get the AnalyticsWorkspace from db
			analyticsWorkspace = AnalyticsWorkspaceDelegate().get(workspaceId).first();
			
			# assign the Workspace		
			inferenceEndpoint.workspace = analyticsWorkspace
			
			#save it
			inferenceEndpoint.save()

			# reload and return the appropriate version					
			return self.get( inferenceEndpointId );
		except InferenceEndpoint.DoesNotExist:
			raise ProcessingError(errMsg + " : InferenceEndpoint with id " + str(inferenceEndpointId) + " does not exist.")
		except AnalyticsWorkspace.DoesNotExist:
			raise ProcessingError(errMsg + " : AnalyticsWorkspace with id " + str(workspaceId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignWorkspace( self, inferenceEndpointId ):
		errMsg = "Failed to unassign element " + str(workspaceId) + " for Workspace on InferenceEndpoint"

		try:
			# get the InferenceEndpoint from db
			inferenceEndpoint = self.get( inferenceEndpointId ).first()	
			
			# assign to None for unassignment
			inferenceEndpoint.analyticsWorkspace = None			

			#save it
			inferenceEndpoint.save()

			# reload and return the appropriate version					
			return self.get( inferenceEndpointId );
		except InferenceEndpoint.DoesNotExist:
			raise ProcessingError(errMsg + " : InferenceEndpoint with id " + str(inferenceEndpointId) + " does not exist.")
		except Exception:
			return None;
		
	def addPredictions( self, inferenceEndpointId, predictionsIds ):
		# lazy importing avoids circular dependencies
		from analyticsOnDjango.delegates.PredictionDelegate import PredictionDelegate

		errMsg = "Failed to add elements " + str(predictionsIds) + " for Predictions on InferenceEndpoint"

		try:
			# get the InferenceEndpoint
			inferenceEndpoint = self.get( inferenceEndpointId ).first()
				
			# split on a comma with no spaces
			idList = predictionsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the Prediction		
				prediction = PredictionDelegate().get(id).first();	
				# add the Prediction
				inferenceEndpoint.predictions.add(prediction)
				
			# save it		
			inferenceEndpoint.save()
			
			# reload and return the appropriate version
			return self.get( inferenceEndpointId );
		except InferenceEndpoint.DoesNotExist:
			raise ProcessingError(errMsg + " : InferenceEndpoint with id " + str(inferenceEndpointId) + " does not exist.")
		except Prediction.DoesNotExist:
			raise ProcessingError(errMsg + " : Prediction does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removePredictions( self, inferenceEndpointId, predictionsIds ):
		# lazy importing avoids circular dependencies
		from analyticsOnDjango.delegates.PredictionDelegate import PredictionDelegate

		errMsg = "Failed to remove elements " + str(predictionsIds) + " for Predictions on InferenceEndpoint"

		try:
			# get the InferenceEndpoint
			inferenceEndpoint = self.get( inferenceEndpointId ).first()
				
			# split on a comma with no spaces
			idList = predictionsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the Prediction		
				prediction = PredictionDelegate().get(id).first();	
				# add the Prediction
				inferenceEndpoint.predictions.remove(prediction)
				
			# save it		
			inferenceEndpoint.save()
			
			# reload and return the appropriate version
			return self.get( inferenceEndpointId );
		except InferenceEndpoint.DoesNotExist:
			raise ProcessingError(errMsg + " : InferenceEndpoint with id " + str(inferenceEndpointId) + " does not exist.")
		except Prediction.DoesNotExist:
			raise ProcessingError(errMsg + " : Prediction does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
