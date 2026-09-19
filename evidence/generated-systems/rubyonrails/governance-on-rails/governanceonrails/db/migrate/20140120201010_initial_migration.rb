class InitialMigration < ActiveRecord::Migration[6.1]
  def change
    create_table :organizations do |t|
      t.string :name      
      t.string :legalName      
      t.string :jurisdiction      
      t.string :industrySector      
      t.timestamps
    end
    create_table :governanceBodys do |t|
      t.string :name      
      t.string :charterUrl      
      t.string :chair      
      t.integer :BodyType      
      t.timestamps
    end
    create_table :persons do |t|
      t.string :firstName      
      t.string :lastName      
      t.string :email      
      t.string :department      
      t.timestamps
    end
    create_table :roles do |t|
      t.string :name      
      t.string :responsibility      
      t.timestamps
    end
    create_table :roleAssignments do |t|
      t.date :effectiveFrom      
      t.date :effectiveTo      
      t.timestamps
    end
    create_table :policys do |t|
      t.string :title      
      t.string :versionLabel      
      t.date :approvalDate      
      t.date :nextReviewDate      
      t.string :documentUrl      
      t.integer :PolicyType      
      t.integer :Status      
      t.timestamps
    end
    create_table :procedures do |t|
      t.string :title      
      t.string :versionLabel      
      t.integer :Status      
      t.timestamps
    end
    create_table :regulations do |t|
      t.string :name      
      t.string :citation      
      t.string :jurisdiction      
      t.string :publicationUrl      
      t.timestamps
    end
    create_table :obligations do |t|
      t.string :referenceNumber      
      t.string :descriptionText      
      t.integer :ObligationType      
      t.integer :ReviewFrequency      
      t.timestamps
    end
    create_table :controls do |t|
      t.string :name      
      t.string :objective      
      t.string :ownerDepartment      
      t.integer :ControlType      
      t.integer :Frequency      
      t.integer :Status      
      t.timestamps
    end
    create_table :controlTest_s do |t|
      t.string :name      
      t.date :testPeriodStart      
      t.date :testPeriodEnd      
      t.integer :sampleSize      
      t.integer :TestType      
      t.integer :Effectiveness      
      t.integer :Status      
      t.timestamps
    end
    create_table :evidences do |t|
      t.string :title      
      t.string :locationUrl      
      t.date :receivedDate      
      t.integer :EvidenceType      
      t.timestamps
    end
    create_table :risks do |t|
      t.string :name      
      t.string :description      
      t.integer :inherentRiskScore      
      t.integer :residualRiskScore      
      t.integer :Category      
      t.integer :Impact      
      t.integer :Likelihood      
      t.integer :Status      
      t.timestamps
    end
    create_table :riskAssessments do |t|
      t.date :assessmentDate      
      t.string :assessor      
      t.string :summary      
      t.integer :AssessmentType      
      t.timestamps
    end
    create_table :compliancePrograms do |t|
      t.string :name      
      t.string :framework      
      t.integer :Status      
      t.timestamps
    end
    create_table :complianceRequirements do |t|
      t.string :name      
      t.string :source      
      t.string :citation      
      t.integer :Applicability      
      t.integer :Status      
      t.timestamps
    end
    create_table :attestations do |t|
      t.string :statement      
      t.string :attestor      
      t.date :dateSigned      
      t.integer :Result      
      t.timestamps
    end
    create_table :auditPrograms do |t|
      t.string :name      
      t.string :scope      
      t.integer :Cycle      
      t.integer :Status      
      t.timestamps
    end
    create_table :auditEngagements do |t|
      t.string :title      
      t.date :startDate      
      t.date :endDate      
      t.integer :Status      
      t.timestamps
    end
    create_table :auditWorkpapers do |t|
      t.string :workpaperRef      
      t.string :subject      
      t.string :workpaperUrl      
      t.timestamps
    end
    create_table :auditFindings do |t|
      t.string :title      
      t.string :description      
      t.date :dueDate      
      t.integer :Severity      
      t.integer :Status      
      t.timestamps
    end
    create_table :correctiveActions do |t|
      t.string :actionTitle      
      t.string :owner      
      t.date :targetDate      
      t.integer :Status      
      t.timestamps
    end
    create_table :issues do |t|
      t.string :title      
      t.date :openedDate      
      t.date :closedDate      
      t.integer :IssueType      
      t.integer :Priority      
      t.integer :Status      
      t.timestamps
    end
    create_table :businessUnits do |t|
      t.string :name      
      t.string :leader      
      t.timestamps
    end
    create_table :dataProcessingActivitys do |t|
      t.string :name      
      t.string :purpose      
      t.date :startDate      
      t.integer :LawfulBasis      
      t.timestamps
    end
    create_table :dataCategorys do |t|
      t.string :name      
      t.string :description      
      t.integer :Classification      
      t.timestamps
    end
    create_table :system_s do |t|
      t.string :name      
      t.string :ownerDepartment      
      t.integer :SystemType      
      t.timestamps
    end
    create_table :privacyNotices do |t|
      t.string :title      
      t.string :audience      
      t.string :versionLabel      
      t.date :publicationDate      
      t.string :publicationUrl      
      t.integer :Status      
      t.timestamps
    end
    create_table :dataSubjectRequests do |t|
      t.date :receivedDate      
      t.date :dueDate      
      t.string :requesterCountry      
      t.integer :RequestType      
      t.integer :Status      
      t.timestamps
    end
    create_table :recordsRepositorys do |t|
      t.string :name      
      t.string :location      
      t.string :ownerDepartment      
      t.integer :RepositoryType      
      t.timestamps
    end
    create_table :record_s do |t|
      t.string :title      
      t.date :creationDate      
      t.integer :RecordType      
      t.integer :Classification      
      t.integer :Status      
      t.timestamps
    end
    create_table :retentionSchedules do |t|
      t.string :name      
      t.integer :retentionPeriodMonths      
      t.integer :RetentionTrigger      
      t.integer :DispositionAction      
      t.integer :Status      
      t.timestamps
    end
    create_table :dispositionReviews do |t|
      t.date :reviewDate      
      t.string :reviewer      
      t.string :notes      
      t.integer :Outcome      
      t.timestamps
    end
    create_table :legalHolds do |t|
      t.string :name      
      t.string :reason      
      t.date :issuedDate      
      t.date :releaseDate      
      t.integer :HoldStatus      
      t.timestamps
    end
    create_table :matters do |t|
      t.string :matterName      
      t.string :leadCounsel      
      t.integer :MatterType      
      t.integer :Status      
      t.timestamps
    end
    create_table :thirdPartys do |t|
      t.string :name      
      t.string :country      
      t.string :contactEmail      
      t.integer :ThirdPartyType      
      t.integer :Criticality      
      t.timestamps
    end
    create_table :thirdPartyAssessments do |t|
      t.date :assessmentDate      
      t.string :assessor      
      t.integer :AssessmentType      
      t.integer :Result      
      t.timestamps
    end
    create_table :contracts do |t|
      t.string :title      
      t.date :effectiveDate      
      t.date :expiryDate      
      t.string :repositoryUrl      
      t.integer :Status      
      t.timestamps
    end
    create_table :exception_s do |t|
      t.string :title      
      t.string :justification      
      t.date :startDate      
      t.date :endDate      
      t.integer :ExceptionType      
      t.integer :Status      
      t.timestamps
    end
    create_table :consents do |t|
      t.string :subjectIdentifier      
      t.date :captureDate      
      t.date :expiryDate      
      t.integer :ConsentType      
      t.integer :Status      
      t.timestamps
    end
    create_table :dataBreachs do |t|
      t.date :incidentDate      
      t.string :description      
      t.integer :recordsAffected      
      t.boolean :notificationRequired      
      t.integer :Severity      
      t.integer :Status      
      t.timestamps
    end
  end
end
