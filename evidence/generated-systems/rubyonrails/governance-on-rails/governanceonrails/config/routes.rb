Rails.application.routes.draw do
  root "application#health"
  resources :organizations do
    resources :governancebodies
    resources :policies
    resources :risks
    resources :thirdparties
    resources :recordsrepositories
    resources :dataprocessingactivities
    resources :complianceprograms
    resources :auditprograms
    resources :businessunits
    resources :matters
    resources :databreaches
  end
  resources :governancebodys do
    resource :organization
    resources :roleassignments
    resources :policies
  end
  resources :persons do
    resources :roleassignments
    resources :ownedpolicies
    resources :correctiveactions
  end
  resources :roles do
    resources :assignments
  end
  resources :roleassignments do
    resource :person
    resource :role
    resource :governancebody
    resource :organization
  end
  resources :policys do
    resource :organization
    resources :owners
    resources :relatedrequirements
    resources :controls
    resources :procedures
    resources :exceptions
    resources :attestations
  end
  resources :procedures do
    resource :policy
    resources :controls
  end
  resources :regulations do
    resources :obligations
    resources :complianceprograms
  end
  resources :obligations do
    resource :regulation
    resources :controls
    resources :policies
    resources :contracts
  end
  resources :controls do
    resource :policy
    resources :controltests
    resources :evidence
    resources :risks
    resources :obligations
    resources :procedures
    resources :issues
  end
  resources :controltest_s do
    resource :control
    resources :evidence
    resource :engagement
  end
  resources :evidences do
    resource :controltest
    resource :control
    resource :obligation
    resource :workpaper
  end
  resources :risks do
    resource :organization
    resources :controls
    resources :assessments
    resources :issues
    resources :findings
  end
  resources :riskassessments do
    resource :risk
  end
  resources :complianceprograms do
    resource :organization
    resources :requirements
    resources :controls
    resources :attestations
    resources :regulations
  end
  resources :compliancerequirements do
    resource :complianceprogram
    resources :policies
    resources :controls
    resources :obligations
  end
  resources :attestations do
    resource :control
    resource :policy
    resource :complianceprogram
  end
  resources :auditprograms do
    resource :organization
    resources :engagements
  end
  resources :auditengagements do
    resource :auditprogram
    resources :businessunits
    resources :controltests
    resources :workpapers
    resources :findings
  end
  resources :auditworkpapers do
    resource :engagement
    resources :evidence
    resources :findings
  end
  resources :auditfindings do
    resource :engagement
    resource :workpaper
    resources :correctiveactions
    resources :relatedrisks
    resources :relatedcontrols
    resources :issues
  end
  resources :correctiveactions do
    resource :finding
    resource :issue
  end
  resources :issues do
    resource :risk
    resource :finding
    resources :correctiveactions
    resource :control
  end
  resources :businessunits do
    resource :organization
    resources :audits
  end
  resources :dataprocessingactivitys do
    resource :organization
    resources :datacategories
    resources :systems
    resources :records
    resources :privacynotices
    resources :thirdparties
    resources :consents
    resources :databreaches
    resources :datasubjectrequests
  end
  resources :datacategorys do
    resources :processingactivities
    resources :records
    resources :databreaches
  end
  resources :system_s do
    resources :processingactivities
    resources :recordsrepositories
  end
  resources :privacynotices do
    resources :processingactivities
    resource :organization
    resources :consents
  end
  resources :datasubjectrequests do
    resource :organization
    resources :processingactivities
    resources :records
  end
  resources :recordsrepositorys do
    resource :organization
    resources :records
    resources :systems
    resources :retentionschedules
    resources :legalholds
  end
  resources :record_s do
    resource :repository
    resource :retentionschedule
    resources :processingactivities
    resources :datacategories
    resources :legalholds
    resources :datasubjectrequests
  end
  resources :retentionschedules do
    resources :repositories
    resources :records
    resources :exceptions
    resources :dispositionreviews
  end
  resources :dispositionreviews do
    resource :record
    resource :retentionschedule
  end
  resources :legalholds do
    resources :repositories
    resources :records
    resource :matter
  end
  resources :matters do
    resources :legalholds
    resource :organization
    resources :databreaches
    resources :contracts
  end
  resources :thirdpartys do
    resource :organization
    resources :processingactivities
    resources :assessments
    resources :contracts
    resources :obligations
    resources :databreaches
  end
  resources :thirdpartyassessments do
    resource :thirdparty
    resources :issues
  end
  resources :contracts do
    resource :thirdparty
    resources :obligations
    resources :dataprocessingactivities
    resource :matter
  end
  resources :exception_s do
    resource :retentionschedule
    resource :policy
    resource :control
    resource :risk
  end
  resources :consents do
    resources :processingactivities
    resource :privacynotice
  end
  resources :databreachs do
    resource :organization
    resources :processingactivities
    resources :datacategories
    resources :thirdparties
    resource :matter
  end
end
