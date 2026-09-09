from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from analyticsOnDjango.models.Experiment import Experiment
from analyticsOnDjango.models.AnalyticsWorkspace import AnalyticsWorkspace
from analyticsOnDjango.models.TrainingRun import TrainingRun
from analyticsOnDjango.models.Model import Model
from analyticsOnDjango.models.Notebook import Notebook
from analyticsOnDjango.exceptions import Exceptions

 #======================================================================
# 
# Encapsulates data for model Experiment
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class ExperimentDelegate Declaration
#======================================================================
class ExperimentDelegate :

#======================================================================
# Function Declarations
#======================================================================

	def get(self, experimentId ):
		try:	
			experiment = Experiment.objects.filter(id=experimentId)
			return experiment.first();
		except Experiment.DoesNotExist:
			raise ProcessingError("Experiment with id " + str(experimentId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 

	def createFromJson(self, experiment):
		for model in serializers.deserialize("json", experiment):
			model.save()
			return model;

	def create(self, experiment):
		experiment.save()
		return experiment;

	def saveFromJson(self, experiment):
		for model in serializers.deserialize("json", experiment):
			model.save()
			return experiment;
	
	def save(self, experiment):
		experiment.save()
		return experiment;
	
	def delete(self, experimentId ):
		errMsg = "Failed to delete Experiment from db using id " + str(experimentId)
		
		try:
			experiment = Experiment.objects.get(id=experimentId)
			experiment.delete()
			return True
		except Experiment.DoesNotExist:
			raise ProcessingError("Experiment with id " + str(experimentId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 
	
	def getAll(self):
		try:
			all = Experiment.objects.all()
			return all;
		except utils.DatabaseError:
			raise StorageReadError("Failed to get all Experiment from db")
		except Exception:
			return None;
		
	def assignWorkspace( self, experimentId, workspaceId ):
		# lazy importing avoids circular dependencies
		from analyticsOnDjango.delegates.AnalyticsWorkspaceDelegate import AnalyticsWorkspaceDelegate

		errMsg = "Failed to assign element " + str(workspaceId) + " for Workspace on Experiment"

		try:
			# get the Experiment from db
			experiment = self.get( experimentId ).first()	
			
			# get the AnalyticsWorkspace from db
			analyticsWorkspace = AnalyticsWorkspaceDelegate().get(workspaceId).first();
			
			# assign the Workspace		
			experiment.workspace = analyticsWorkspace
			
			#save it
			experiment.save()

			# reload and return the appropriate version					
			return self.get( experimentId );
		except Experiment.DoesNotExist:
			raise ProcessingError(errMsg + " : Experiment with id " + str(experimentId) + " does not exist.")
		except AnalyticsWorkspace.DoesNotExist:
			raise ProcessingError(errMsg + " : AnalyticsWorkspace with id " + str(workspaceId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignWorkspace( self, experimentId ):
		errMsg = "Failed to unassign element " + str(workspaceId) + " for Workspace on Experiment"

		try:
			# get the Experiment from db
			experiment = self.get( experimentId ).first()	
			
			# assign to None for unassignment
			experiment.analyticsWorkspace = None			

			#save it
			experiment.save()

			# reload and return the appropriate version					
			return self.get( experimentId );
		except Experiment.DoesNotExist:
			raise ProcessingError(errMsg + " : Experiment with id " + str(experimentId) + " does not exist.")
		except Exception:
			return None;
		
	def addTrainingRuns( self, experimentId, trainingRunsIds ):
		# lazy importing avoids circular dependencies
		from analyticsOnDjango.delegates.TrainingRunDelegate import TrainingRunDelegate

		errMsg = "Failed to add elements " + str(trainingRunsIds) + " for TrainingRuns on Experiment"

		try:
			# get the Experiment
			experiment = self.get( experimentId ).first()
				
			# split on a comma with no spaces
			idList = trainingRunsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the TrainingRun		
				trainingRun = TrainingRunDelegate().get(id).first();	
				# add the TrainingRun
				experiment.trainingRuns.add(trainingRun)
				
			# save it		
			experiment.save()
			
			# reload and return the appropriate version
			return self.get( experimentId );
		except Experiment.DoesNotExist:
			raise ProcessingError(errMsg + " : Experiment with id " + str(experimentId) + " does not exist.")
		except TrainingRun.DoesNotExist:
			raise ProcessingError(errMsg + " : TrainingRun does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeTrainingRuns( self, experimentId, trainingRunsIds ):
		# lazy importing avoids circular dependencies
		from analyticsOnDjango.delegates.TrainingRunDelegate import TrainingRunDelegate

		errMsg = "Failed to remove elements " + str(trainingRunsIds) + " for TrainingRuns on Experiment"

		try:
			# get the Experiment
			experiment = self.get( experimentId ).first()
				
			# split on a comma with no spaces
			idList = trainingRunsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the TrainingRun		
				trainingRun = TrainingRunDelegate().get(id).first();	
				# add the TrainingRun
				experiment.trainingRuns.remove(trainingRun)
				
			# save it		
			experiment.save()
			
			# reload and return the appropriate version
			return self.get( experimentId );
		except Experiment.DoesNotExist:
			raise ProcessingError(errMsg + " : Experiment with id " + str(experimentId) + " does not exist.")
		except TrainingRun.DoesNotExist:
			raise ProcessingError(errMsg + " : TrainingRun does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addModels( self, experimentId, modelsIds ):
		# lazy importing avoids circular dependencies
		from analyticsOnDjango.delegates.ModelDelegate import ModelDelegate

		errMsg = "Failed to add elements " + str(modelsIds) + " for Models on Experiment"

		try:
			# get the Experiment
			experiment = self.get( experimentId ).first()
				
			# split on a comma with no spaces
			idList = modelsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the Model		
				model = ModelDelegate().get(id).first();	
				# add the Model
				experiment.models.add(model)
				
			# save it		
			experiment.save()
			
			# reload and return the appropriate version
			return self.get( experimentId );
		except Experiment.DoesNotExist:
			raise ProcessingError(errMsg + " : Experiment with id " + str(experimentId) + " does not exist.")
		except Model.DoesNotExist:
			raise ProcessingError(errMsg + " : Model does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeModels( self, experimentId, modelsIds ):
		# lazy importing avoids circular dependencies
		from analyticsOnDjango.delegates.ModelDelegate import ModelDelegate

		errMsg = "Failed to remove elements " + str(modelsIds) + " for Models on Experiment"

		try:
			# get the Experiment
			experiment = self.get( experimentId ).first()
				
			# split on a comma with no spaces
			idList = modelsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the Model		
				model = ModelDelegate().get(id).first();	
				# add the Model
				experiment.models.remove(model)
				
			# save it		
			experiment.save()
			
			# reload and return the appropriate version
			return self.get( experimentId );
		except Experiment.DoesNotExist:
			raise ProcessingError(errMsg + " : Experiment with id " + str(experimentId) + " does not exist.")
		except Model.DoesNotExist:
			raise ProcessingError(errMsg + " : Model does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addNotebooks( self, experimentId, notebooksIds ):
		# lazy importing avoids circular dependencies
		from analyticsOnDjango.delegates.NotebookDelegate import NotebookDelegate

		errMsg = "Failed to add elements " + str(notebooksIds) + " for Notebooks on Experiment"

		try:
			# get the Experiment
			experiment = self.get( experimentId ).first()
				
			# split on a comma with no spaces
			idList = notebooksIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the Notebook		
				notebook = NotebookDelegate().get(id).first();	
				# add the Notebook
				experiment.notebooks.add(notebook)
				
			# save it		
			experiment.save()
			
			# reload and return the appropriate version
			return self.get( experimentId );
		except Experiment.DoesNotExist:
			raise ProcessingError(errMsg + " : Experiment with id " + str(experimentId) + " does not exist.")
		except Notebook.DoesNotExist:
			raise ProcessingError(errMsg + " : Notebook does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeNotebooks( self, experimentId, notebooksIds ):
		# lazy importing avoids circular dependencies
		from analyticsOnDjango.delegates.NotebookDelegate import NotebookDelegate

		errMsg = "Failed to remove elements " + str(notebooksIds) + " for Notebooks on Experiment"

		try:
			# get the Experiment
			experiment = self.get( experimentId ).first()
				
			# split on a comma with no spaces
			idList = notebooksIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the Notebook		
				notebook = NotebookDelegate().get(id).first();	
				# add the Notebook
				experiment.notebooks.remove(notebook)
				
			# save it		
			experiment.save()
			
			# reload and return the appropriate version
			return self.get( experimentId );
		except Experiment.DoesNotExist:
			raise ProcessingError(errMsg + " : Experiment with id " + str(experimentId) + " does not exist.")
		except Notebook.DoesNotExist:
			raise ProcessingError(errMsg + " : Notebook does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
