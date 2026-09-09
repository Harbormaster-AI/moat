from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from analyticsOnDjango.models.Notebook import Notebook
from analyticsOnDjango.models.AnalyticsWorkspace import AnalyticsWorkspace
from analyticsOnDjango.models.DataSet import DataSet
from analyticsOnDjango.models.Experiment import Experiment
from analyticsOnDjango.models.BIQuery import BIQuery
from analyticsOnDjango.exceptions import Exceptions

 #======================================================================
# 
# Encapsulates data for model Notebook
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class NotebookDelegate Declaration
#======================================================================
class NotebookDelegate :

#======================================================================
# Function Declarations
#======================================================================

	def get(self, notebookId ):
		try:	
			notebook = Notebook.objects.filter(id=notebookId)
			return notebook.first();
		except Notebook.DoesNotExist:
			raise ProcessingError("Notebook with id " + str(notebookId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 

	def createFromJson(self, notebook):
		for model in serializers.deserialize("json", notebook):
			model.save()
			return model;

	def create(self, notebook):
		notebook.save()
		return notebook;

	def saveFromJson(self, notebook):
		for model in serializers.deserialize("json", notebook):
			model.save()
			return notebook;
	
	def save(self, notebook):
		notebook.save()
		return notebook;
	
	def delete(self, notebookId ):
		errMsg = "Failed to delete Notebook from db using id " + str(notebookId)
		
		try:
			notebook = Notebook.objects.get(id=notebookId)
			notebook.delete()
			return True
		except Notebook.DoesNotExist:
			raise ProcessingError("Notebook with id " + str(notebookId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 
	
	def getAll(self):
		try:
			all = Notebook.objects.all()
			return all;
		except utils.DatabaseError:
			raise StorageReadError("Failed to get all Notebook from db")
		except Exception:
			return None;
		
	def assignWorkspace( self, notebookId, workspaceId ):
		# lazy importing avoids circular dependencies
		from analyticsOnDjango.delegates.AnalyticsWorkspaceDelegate import AnalyticsWorkspaceDelegate

		errMsg = "Failed to assign element " + str(workspaceId) + " for Workspace on Notebook"

		try:
			# get the Notebook from db
			notebook = self.get( notebookId ).first()	
			
			# get the AnalyticsWorkspace from db
			analyticsWorkspace = AnalyticsWorkspaceDelegate().get(workspaceId).first();
			
			# assign the Workspace		
			notebook.workspace = analyticsWorkspace
			
			#save it
			notebook.save()

			# reload and return the appropriate version					
			return self.get( notebookId );
		except Notebook.DoesNotExist:
			raise ProcessingError(errMsg + " : Notebook with id " + str(notebookId) + " does not exist.")
		except AnalyticsWorkspace.DoesNotExist:
			raise ProcessingError(errMsg + " : AnalyticsWorkspace with id " + str(workspaceId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignWorkspace( self, notebookId ):
		errMsg = "Failed to unassign element " + str(workspaceId) + " for Workspace on Notebook"

		try:
			# get the Notebook from db
			notebook = self.get( notebookId ).first()	
			
			# assign to None for unassignment
			notebook.analyticsWorkspace = None			

			#save it
			notebook.save()

			# reload and return the appropriate version					
			return self.get( notebookId );
		except Notebook.DoesNotExist:
			raise ProcessingError(errMsg + " : Notebook with id " + str(notebookId) + " does not exist.")
		except Exception:
			return None;
		
	def addDatasets( self, notebookId, datasetsIds ):
		# lazy importing avoids circular dependencies
		from analyticsOnDjango.delegates.DataSetDelegate import DataSetDelegate

		errMsg = "Failed to add elements " + str(datasetsIds) + " for Datasets on Notebook"

		try:
			# get the Notebook
			notebook = self.get( notebookId ).first()
				
			# split on a comma with no spaces
			idList = datasetsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the DataSet		
				dataSet = DataSetDelegate().get(id).first();	
				# add the DataSet
				notebook.datasets.add(dataSet)
				
			# save it		
			notebook.save()
			
			# reload and return the appropriate version
			return self.get( notebookId );
		except Notebook.DoesNotExist:
			raise ProcessingError(errMsg + " : Notebook with id " + str(notebookId) + " does not exist.")
		except DataSet.DoesNotExist:
			raise ProcessingError(errMsg + " : DataSet does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeDatasets( self, notebookId, datasetsIds ):
		# lazy importing avoids circular dependencies
		from analyticsOnDjango.delegates.DataSetDelegate import DataSetDelegate

		errMsg = "Failed to remove elements " + str(datasetsIds) + " for Datasets on Notebook"

		try:
			# get the Notebook
			notebook = self.get( notebookId ).first()
				
			# split on a comma with no spaces
			idList = datasetsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the DataSet		
				dataSet = DataSetDelegate().get(id).first();	
				# add the DataSet
				notebook.datasets.remove(dataSet)
				
			# save it		
			notebook.save()
			
			# reload and return the appropriate version
			return self.get( notebookId );
		except Notebook.DoesNotExist:
			raise ProcessingError(errMsg + " : Notebook with id " + str(notebookId) + " does not exist.")
		except DataSet.DoesNotExist:
			raise ProcessingError(errMsg + " : DataSet does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addExperiments( self, notebookId, experimentsIds ):
		# lazy importing avoids circular dependencies
		from analyticsOnDjango.delegates.ExperimentDelegate import ExperimentDelegate

		errMsg = "Failed to add elements " + str(experimentsIds) + " for Experiments on Notebook"

		try:
			# get the Notebook
			notebook = self.get( notebookId ).first()
				
			# split on a comma with no spaces
			idList = experimentsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the Experiment		
				experiment = ExperimentDelegate().get(id).first();	
				# add the Experiment
				notebook.experiments.add(experiment)
				
			# save it		
			notebook.save()
			
			# reload and return the appropriate version
			return self.get( notebookId );
		except Notebook.DoesNotExist:
			raise ProcessingError(errMsg + " : Notebook with id " + str(notebookId) + " does not exist.")
		except Experiment.DoesNotExist:
			raise ProcessingError(errMsg + " : Experiment does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeExperiments( self, notebookId, experimentsIds ):
		# lazy importing avoids circular dependencies
		from analyticsOnDjango.delegates.ExperimentDelegate import ExperimentDelegate

		errMsg = "Failed to remove elements " + str(experimentsIds) + " for Experiments on Notebook"

		try:
			# get the Notebook
			notebook = self.get( notebookId ).first()
				
			# split on a comma with no spaces
			idList = experimentsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the Experiment		
				experiment = ExperimentDelegate().get(id).first();	
				# add the Experiment
				notebook.experiments.remove(experiment)
				
			# save it		
			notebook.save()
			
			# reload and return the appropriate version
			return self.get( notebookId );
		except Notebook.DoesNotExist:
			raise ProcessingError(errMsg + " : Notebook with id " + str(notebookId) + " does not exist.")
		except Experiment.DoesNotExist:
			raise ProcessingError(errMsg + " : Experiment does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addQueries( self, notebookId, queriesIds ):
		# lazy importing avoids circular dependencies
		from analyticsOnDjango.delegates.BIQueryDelegate import BIQueryDelegate

		errMsg = "Failed to add elements " + str(queriesIds) + " for Queries on Notebook"

		try:
			# get the Notebook
			notebook = self.get( notebookId ).first()
				
			# split on a comma with no spaces
			idList = queriesIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the BIQuery		
				bIQuery = BIQueryDelegate().get(id).first();	
				# add the BIQuery
				notebook.queries.add(bIQuery)
				
			# save it		
			notebook.save()
			
			# reload and return the appropriate version
			return self.get( notebookId );
		except Notebook.DoesNotExist:
			raise ProcessingError(errMsg + " : Notebook with id " + str(notebookId) + " does not exist.")
		except BIQuery.DoesNotExist:
			raise ProcessingError(errMsg + " : BIQuery does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeQueries( self, notebookId, queriesIds ):
		# lazy importing avoids circular dependencies
		from analyticsOnDjango.delegates.BIQueryDelegate import BIQueryDelegate

		errMsg = "Failed to remove elements " + str(queriesIds) + " for Queries on Notebook"

		try:
			# get the Notebook
			notebook = self.get( notebookId ).first()
				
			# split on a comma with no spaces
			idList = queriesIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the BIQuery		
				bIQuery = BIQueryDelegate().get(id).first();	
				# add the BIQuery
				notebook.queries.remove(bIQuery)
				
			# save it		
			notebook.save()
			
			# reload and return the appropriate version
			return self.get( notebookId );
		except Notebook.DoesNotExist:
			raise ProcessingError(errMsg + " : Notebook with id " + str(notebookId) + " does not exist.")
		except BIQuery.DoesNotExist:
			raise ProcessingError(errMsg + " : BIQuery does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
