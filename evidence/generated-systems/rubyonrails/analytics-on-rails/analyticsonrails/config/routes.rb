Rails.application.routes.draw do
  root "application#health"
  resources :analyticsworkspaces do
    resources :datasets
    resources :datasources
    resources :pipelines
    resources :dashboards
    resources :reports
    resources :notebooks
    resources :models
    resources :featuresets
    resources :policies
    resources :lineagenodes
  end
  resources :datasources do
    resource :workspace
    resources :produceddatasets
    resources :pipelines
  end
  resources :datasets do
    resource :workspace
    resources :sources
    resources :pipelines
    resources :semanticmodels
    resources :dimensions
    resources :measures
    resources :metrics
    resources :qualityrules
    resource :lineagenode
    resources :tags
  end
  resources :datapipelines do
    resource :workspace
    resources :tasks
    resources :sources
    resources :outputs
    resource :lineagenode
  end
  resources :datatasks do
    resource :pipeline
    resources :inputdatasets
    resources :outputdatasets
  end
  resources :semanticmodels do
    resources :datasets
    resources :metrics
    resources :dimensions
    resources :measures
    resources :glossaryterms
  end
  resources :dimensions do
    resource :semanticmodel
    resources :datasets
    resources :glossaryterms
  end
  resources :measures do
    resource :semanticmodel
    resources :datasets
    resources :glossaryterms
  end
  resources :metrics do
    resource :semanticmodel
    resources :datasets
    resources :glossaryterms
    resources :alerts
    resources :visualizations
  end
  resources :reports do
    resource :workspace
    resources :visualizations
    resources :datasets
    resources :semanticmodels
    resources :queries
    resources :tags
  end
  resources :dashboards do
    resource :workspace
    resources :visualizations
    resources :reports
    resources :datasets
    resources :alerts
    resources :queries
    resources :tags
  end
  resources :visualizations do
    resource :dashboard
    resource :report
    resources :metrics
    resources :dimensions
    resources :datasets
  end
  resources :notebooks do
    resource :workspace
    resources :datasets
    resources :experiments
    resources :queries
  end
  resources :biquerys do
    resource :workspace
    resources :datasets
    resources :reports
    resources :dashboards
    resources :notebooks
  end
  resources :experiments do
    resource :workspace
    resources :trainingruns
    resources :models
    resources :notebooks
  end
  resources :trainingruns do
    resource :experiment
    resource :modelversion
    resources :inputdatasets
    resources :features
    resources :runmetrics
    resources :runparameters
  end
  resources :runmetrics do
    resource :trainingrun
    resource :metric
    resource :dataset
  end
  resources :runparameters do
    resource :trainingrun
  end
  resources :models do
    resource :workspace
    resources :versions
    resources :featuresets
    resources :experiments
    resources :tags
  end
  resources :modelversions do
    resource :model
    resource :trainingrun
    resources :evaluationmetrics
    resources :deployments
    resources :featuresets
    resources :datasets
  end
  resources :evaluationmetrics do
    resource :modelversion
    resource :metric
    resource :dataset
  end
  resources :featuresets do
    resource :workspace
    resources :features
    resources :datasets
    resources :models
    resources :modelversions
    resources :tags
  end
  resources :features do
    resource :featureset
    resources :sourcedatasets
    resources :models
    resources :trainingruns
  end
  resources :inferenceendpoints do
    resource :modelversion
    resource :workspace
    resources :predictions
  end
  resources :predictions do
    resource :endpoint
    resource :modelversion
    resource :dataset
  end
  resources :forecasts do
    resource :modelversion
    resource :timeseries
    resources :datasets
  end
  resources :timeseriess do
    resources :datasets
    resources :forecasts
    resources :anomalies
  end
  resources :anomalys do
    resource :timeseries
    resource :alert
    resource :dataset
  end
  resources :qualityrules do
    resource :dataset
    resources :checks
  end
  resources :qualitychecks do
    resource :rule
    resource :dataset
  end
  resources :lineagenodes do
    resource :workspace
    resources :inputs
    resources :outputs
    resources :datasets
    resources :models
    resources :pipelines
    resources :dashboards
    resources :reports
  end
  resources :tags do
    resources :datasets
    resources :models
    resources :modelversions
    resources :dashboards
    resources :reports
    resources :featuresets
    resources :metrics
  end
  resources :accesspolicys do
    resource :workspace
    resources :datasets
    resources :dashboards
    resources :reports
    resources :models
    resources :featuresets
  end
  resources :alerts do
    resource :metric
    resource :dashboard
    resource :dataset
    resource :rule
    resources :anomalies
    resources :subscribers
  end
  resources :subscribers do
    resources :alerts
  end
  resources :businessglossaryterms do
    resources :relatedterms
    resources :metrics
    resources :datasets
    resources :dimensions
    resources :measures
  end
  resources :recommendationscenarios do
    resources :models
    resources :datasets
    resources :experiments
    resources :alerts
  end
  resources :fraudscenarios do
    resources :models
    resources :datasets
    resources :alerts
    resources :signals
  end
  resources :fraudsignals do
    resource :scenario
    resource :dataset
    resource :modelversion
  end
end
