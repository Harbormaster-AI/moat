package router

import (

    AerospaceManufacturerController "aerospace-on-golang/internal/controller"
    AircraftProgramController "aerospace-on-golang/internal/controller"
    AircraftFamilyController "aerospace-on-golang/internal/controller"
    AircraftModelController "aerospace-on-golang/internal/controller"
    EngineTypeController "aerospace-on-golang/internal/controller"
    AircraftVariantController "aerospace-on-golang/internal/controller"
    AvionicsSuiteController "aerospace-on-golang/internal/controller"
    APUController "aerospace-on-golang/internal/controller"
    LandingGearController "aerospace-on-golang/internal/controller"
    AircraftOptionController "aerospace-on-golang/internal/controller"
    AircraftPackageController "aerospace-on-golang/internal/controller"
    SupplierController "aerospace-on-golang/internal/controller"
    Component_Controller "aerospace-on-golang/internal/controller"
    PlantController "aerospace-on-golang/internal/controller"
    ProductionLineController "aerospace-on-golang/internal/controller"
    WorkCenterController "aerospace-on-golang/internal/controller"
    ProductionOrderController "aerospace-on-golang/internal/controller"
    BuildScheduleController "aerospace-on-golang/internal/controller"
    WarehouseController "aerospace-on-golang/internal/controller"
    InventoryItemController "aerospace-on-golang/internal/controller"
    OperatorController "aerospace-on-golang/internal/controller"
    AircraftOrderController "aerospace-on-golang/internal/controller"
    QuoteController "aerospace-on-golang/internal/controller"
    PurchaseAgreementController "aerospace-on-golang/internal/controller"
    AircraftController "aerospace-on-golang/internal/controller"
    RegistrationController "aerospace-on-golang/internal/controller"
    WarrantyController "aerospace-on-golang/internal/controller"
    CabinLayoutController "aerospace-on-golang/internal/controller"
    MROFacilityController "aerospace-on-golang/internal/controller"
    MaintenanceAppointmentController "aerospace-on-golang/internal/controller"
    MaintenanceWorkOrderController "aerospace-on-golang/internal/controller"
    AirworthinessDirectiveController "aerospace-on-golang/internal/controller"
    ServiceBulletinController "aerospace-on-golang/internal/controller"
    ConnectedAircraftController "aerospace-on-golang/internal/controller"
    FlightHealthEventController "aerospace-on-golang/internal/controller"
    SoftwareLoadController "aerospace-on-golang/internal/controller"
    TypeCertificateController "aerospace-on-golang/internal/controller"
    ProductionCertificateController "aerospace-on-golang/internal/controller"
    SalesRegionController "aerospace-on-golang/internal/controller"
    SalesCampaignController "aerospace-on-golang/internal/controller"
    jsonResponseFormatter "aerospace-on-golang/internal/response"
    "github.com/gorilla/mux"

    PulseIndicatorController__ "aerospace-on-golang/internal/controller"

)

// Router is exported and used in main.go
func Router() *mux.Router {

    router := mux.NewRouter()

    //----------------------------------------------------------------------------
    // default controllers for health and availability checking
    //----------------------------------------------------------------------------

    router.HandleFunc("/", jsonResponseFormatter.FormatToJSON(PulseIndicatorController__.Default__)).Methods("GET", "OPTIONS")
    router.HandleFunc("/health", jsonResponseFormatter.FormatToJSON(PulseIndicatorController__.Health__)).Methods("GET", "OPTIONS")


    //----------------------------------------------------------------------------
    // AerospaceManufacturer Routes to JSON response formatter first
    // then to the correct Controller function
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Standard Lifecycle Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AerospaceManufacturer/{id}", jsonResponseFormatter.FormatToJSON(AerospaceManufacturerController.GetAerospaceManufacturer)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/AerospaceManufacturer", jsonResponseFormatter.FormatToJSON(AerospaceManufacturerController.GetAllAerospaceManufacturer)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/NewAerospaceManufacturer", jsonResponseFormatter.FormatToJSON(AerospaceManufacturerController.CreateAerospaceManufacturer)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/AerospaceManufacturer/{id}", jsonResponseFormatter.FormatToJSON(AerospaceManufacturerController.UpdateAerospaceManufacturer)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/DeleteAerospaceManufacturer/{id}", jsonResponseFormatter.FormatToJSON(AerospaceManufacturerController.DeleteAerospaceManufacturer)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AddProgramsToAerospaceManufacturer/{parentId}/programsId", jsonResponseFormatter.FormatToJSON(AerospaceManufacturerController.AddProgramsToAerospaceManufacturer)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveProgramsFromAerospaceManufacturer/{parentId}/programsIds", jsonResponseFormatter.FormatToJSON(AerospaceManufacturerController.RemoveProgramsFromAerospaceManufacturer)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddPlantsToAerospaceManufacturer/{parentId}/plantsId", jsonResponseFormatter.FormatToJSON(AerospaceManufacturerController.AddPlantsToAerospaceManufacturer)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemovePlantsFromAerospaceManufacturer/{parentId}/plantsIds", jsonResponseFormatter.FormatToJSON(AerospaceManufacturerController.RemovePlantsFromAerospaceManufacturer)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddSuppliersToAerospaceManufacturer/{parentId}/suppliersId", jsonResponseFormatter.FormatToJSON(AerospaceManufacturerController.AddSuppliersToAerospaceManufacturer)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveSuppliersFromAerospaceManufacturer/{parentId}/suppliersIds", jsonResponseFormatter.FormatToJSON(AerospaceManufacturerController.RemoveSuppliersFromAerospaceManufacturer)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddProductionCertificatesToAerospaceManufacturer/{parentId}/productionCertificatesId", jsonResponseFormatter.FormatToJSON(AerospaceManufacturerController.AddProductionCertificatesToAerospaceManufacturer)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveProductionCertificatesFromAerospaceManufacturer/{parentId}/productionCertificatesIds", jsonResponseFormatter.FormatToJSON(AerospaceManufacturerController.RemoveProductionCertificatesFromAerospaceManufacturer)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // AircraftProgram Routes to JSON response formatter first
    // then to the correct Controller function
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Standard Lifecycle Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AircraftProgram/{id}", jsonResponseFormatter.FormatToJSON(AircraftProgramController.GetAircraftProgram)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/AircraftProgram", jsonResponseFormatter.FormatToJSON(AircraftProgramController.GetAllAircraftProgram)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/NewAircraftProgram", jsonResponseFormatter.FormatToJSON(AircraftProgramController.CreateAircraftProgram)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/AircraftProgram/{id}", jsonResponseFormatter.FormatToJSON(AircraftProgramController.UpdateAircraftProgram)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/DeleteAircraftProgram/{id}", jsonResponseFormatter.FormatToJSON(AircraftProgramController.DeleteAircraftProgram)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AssignManufacturerToAircraftProgram/{parentId}/manufacturerId", jsonResponseFormatter.FormatToJSON(AircraftProgramController.AssignManufacturerToAircraftProgram)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignManufacturerFromAircraftProgram/{parentId}", jsonResponseFormatter.FormatToJSON(AircraftProgramController.UnassignManufacturerFromAircraftProgram)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AssignTypeCertificateToAircraftProgram/{parentId}/typeCertificateId", jsonResponseFormatter.FormatToJSON(AircraftProgramController.AssignTypeCertificateToAircraftProgram)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignTypeCertificateFromAircraftProgram/{parentId}", jsonResponseFormatter.FormatToJSON(AircraftProgramController.UnassignTypeCertificateFromAircraftProgram)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AddAircraftFamiliesToAircraftProgram/{parentId}/aircraftFamiliesId", jsonResponseFormatter.FormatToJSON(AircraftProgramController.AddAircraftFamiliesToAircraftProgram)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveAircraftFamiliesFromAircraftProgram/{parentId}/aircraftFamiliesIds", jsonResponseFormatter.FormatToJSON(AircraftProgramController.RemoveAircraftFamiliesFromAircraftProgram)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddKeySuppliersToAircraftProgram/{parentId}/keySuppliersId", jsonResponseFormatter.FormatToJSON(AircraftProgramController.AddKeySuppliersToAircraftProgram)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveKeySuppliersFromAircraftProgram/{parentId}/keySuppliersIds", jsonResponseFormatter.FormatToJSON(AircraftProgramController.RemoveKeySuppliersFromAircraftProgram)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // AircraftFamily Routes to JSON response formatter first
    // then to the correct Controller function
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Standard Lifecycle Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AircraftFamily/{id}", jsonResponseFormatter.FormatToJSON(AircraftFamilyController.GetAircraftFamily)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/AircraftFamily", jsonResponseFormatter.FormatToJSON(AircraftFamilyController.GetAllAircraftFamily)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/NewAircraftFamily", jsonResponseFormatter.FormatToJSON(AircraftFamilyController.CreateAircraftFamily)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/AircraftFamily/{id}", jsonResponseFormatter.FormatToJSON(AircraftFamilyController.UpdateAircraftFamily)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/DeleteAircraftFamily/{id}", jsonResponseFormatter.FormatToJSON(AircraftFamilyController.DeleteAircraftFamily)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AssignProgramToAircraftFamily/{parentId}/programId", jsonResponseFormatter.FormatToJSON(AircraftFamilyController.AssignProgramToAircraftFamily)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignProgramFromAircraftFamily/{parentId}", jsonResponseFormatter.FormatToJSON(AircraftFamilyController.UnassignProgramFromAircraftFamily)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AddAircraftModelsToAircraftFamily/{parentId}/aircraftModelsId", jsonResponseFormatter.FormatToJSON(AircraftFamilyController.AddAircraftModelsToAircraftFamily)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveAircraftModelsFromAircraftFamily/{parentId}/aircraftModelsIds", jsonResponseFormatter.FormatToJSON(AircraftFamilyController.RemoveAircraftModelsFromAircraftFamily)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // AircraftModel Routes to JSON response formatter first
    // then to the correct Controller function
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Standard Lifecycle Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AircraftModel/{id}", jsonResponseFormatter.FormatToJSON(AircraftModelController.GetAircraftModel)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/AircraftModel", jsonResponseFormatter.FormatToJSON(AircraftModelController.GetAllAircraftModel)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/NewAircraftModel", jsonResponseFormatter.FormatToJSON(AircraftModelController.CreateAircraftModel)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/AircraftModel/{id}", jsonResponseFormatter.FormatToJSON(AircraftModelController.UpdateAircraftModel)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/DeleteAircraftModel/{id}", jsonResponseFormatter.FormatToJSON(AircraftModelController.DeleteAircraftModel)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AssignFamilyToAircraftModel/{parentId}/familyId", jsonResponseFormatter.FormatToJSON(AircraftModelController.AssignFamilyToAircraftModel)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignFamilyFromAircraftModel/{parentId}", jsonResponseFormatter.FormatToJSON(AircraftModelController.UnassignFamilyFromAircraftModel)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AddVariantsToAircraftModel/{parentId}/variantsId", jsonResponseFormatter.FormatToJSON(AircraftModelController.AddVariantsToAircraftModel)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveVariantsFromAircraftModel/{parentId}/variantsIds", jsonResponseFormatter.FormatToJSON(AircraftModelController.RemoveVariantsFromAircraftModel)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddEngineTypesToAircraftModel/{parentId}/engineTypesId", jsonResponseFormatter.FormatToJSON(AircraftModelController.AddEngineTypesToAircraftModel)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveEngineTypesFromAircraftModel/{parentId}/engineTypesIds", jsonResponseFormatter.FormatToJSON(AircraftModelController.RemoveEngineTypesFromAircraftModel)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // EngineType Routes to JSON response formatter first
    // then to the correct Controller function
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Standard Lifecycle Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/EngineType/{id}", jsonResponseFormatter.FormatToJSON(EngineTypeController.GetEngineType)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/EngineType", jsonResponseFormatter.FormatToJSON(EngineTypeController.GetAllEngineType)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/NewEngineType", jsonResponseFormatter.FormatToJSON(EngineTypeController.CreateEngineType)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/EngineType/{id}", jsonResponseFormatter.FormatToJSON(EngineTypeController.UpdateEngineType)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/DeleteEngineType/{id}", jsonResponseFormatter.FormatToJSON(EngineTypeController.DeleteEngineType)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AssignSupplierToEngineType/{parentId}/supplierId", jsonResponseFormatter.FormatToJSON(EngineTypeController.AssignSupplierToEngineType)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignSupplierFromEngineType/{parentId}", jsonResponseFormatter.FormatToJSON(EngineTypeController.UnassignSupplierFromEngineType)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AddCompatibleModelsToEngineType/{parentId}/compatibleModelsId", jsonResponseFormatter.FormatToJSON(EngineTypeController.AddCompatibleModelsToEngineType)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveCompatibleModelsFromEngineType/{parentId}/compatibleModelsIds", jsonResponseFormatter.FormatToJSON(EngineTypeController.RemoveCompatibleModelsFromEngineType)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // AircraftVariant Routes to JSON response formatter first
    // then to the correct Controller function
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Standard Lifecycle Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AircraftVariant/{id}", jsonResponseFormatter.FormatToJSON(AircraftVariantController.GetAircraftVariant)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/AircraftVariant", jsonResponseFormatter.FormatToJSON(AircraftVariantController.GetAllAircraftVariant)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/NewAircraftVariant", jsonResponseFormatter.FormatToJSON(AircraftVariantController.CreateAircraftVariant)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/AircraftVariant/{id}", jsonResponseFormatter.FormatToJSON(AircraftVariantController.UpdateAircraftVariant)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/DeleteAircraftVariant/{id}", jsonResponseFormatter.FormatToJSON(AircraftVariantController.DeleteAircraftVariant)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AssignModel_ToAircraftVariant/{parentId}/model_Id", jsonResponseFormatter.FormatToJSON(AircraftVariantController.AssignModel_ToAircraftVariant)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignModel_FromAircraftVariant/{parentId}", jsonResponseFormatter.FormatToJSON(AircraftVariantController.UnassignModel_FromAircraftVariant)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AssignEngineTypeToAircraftVariant/{parentId}/engineTypeId", jsonResponseFormatter.FormatToJSON(AircraftVariantController.AssignEngineTypeToAircraftVariant)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignEngineTypeFromAircraftVariant/{parentId}", jsonResponseFormatter.FormatToJSON(AircraftVariantController.UnassignEngineTypeFromAircraftVariant)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AssignAvionicsSuiteToAircraftVariant/{parentId}/avionicsSuiteId", jsonResponseFormatter.FormatToJSON(AircraftVariantController.AssignAvionicsSuiteToAircraftVariant)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignAvionicsSuiteFromAircraftVariant/{parentId}", jsonResponseFormatter.FormatToJSON(AircraftVariantController.UnassignAvionicsSuiteFromAircraftVariant)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AssignApuToAircraftVariant/{parentId}/apuId", jsonResponseFormatter.FormatToJSON(AircraftVariantController.AssignApuToAircraftVariant)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignApuFromAircraftVariant/{parentId}", jsonResponseFormatter.FormatToJSON(AircraftVariantController.UnassignApuFromAircraftVariant)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AssignLandingGearToAircraftVariant/{parentId}/landingGearId", jsonResponseFormatter.FormatToJSON(AircraftVariantController.AssignLandingGearToAircraftVariant)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignLandingGearFromAircraftVariant/{parentId}", jsonResponseFormatter.FormatToJSON(AircraftVariantController.UnassignLandingGearFromAircraftVariant)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AddCabinLayoutsToAircraftVariant/{parentId}/cabinLayoutsId", jsonResponseFormatter.FormatToJSON(AircraftVariantController.AddCabinLayoutsToAircraftVariant)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveCabinLayoutsFromAircraftVariant/{parentId}/cabinLayoutsIds", jsonResponseFormatter.FormatToJSON(AircraftVariantController.RemoveCabinLayoutsFromAircraftVariant)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddOptionsToAircraftVariant/{parentId}/optionsId", jsonResponseFormatter.FormatToJSON(AircraftVariantController.AddOptionsToAircraftVariant)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveOptionsFromAircraftVariant/{parentId}/optionsIds", jsonResponseFormatter.FormatToJSON(AircraftVariantController.RemoveOptionsFromAircraftVariant)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddPackagesToAircraftVariant/{parentId}/packagesId", jsonResponseFormatter.FormatToJSON(AircraftVariantController.AddPackagesToAircraftVariant)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemovePackagesFromAircraftVariant/{parentId}/packagesIds", jsonResponseFormatter.FormatToJSON(AircraftVariantController.RemovePackagesFromAircraftVariant)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // AvionicsSuite Routes to JSON response formatter first
    // then to the correct Controller function
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Standard Lifecycle Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AvionicsSuite/{id}", jsonResponseFormatter.FormatToJSON(AvionicsSuiteController.GetAvionicsSuite)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/AvionicsSuite", jsonResponseFormatter.FormatToJSON(AvionicsSuiteController.GetAllAvionicsSuite)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/NewAvionicsSuite", jsonResponseFormatter.FormatToJSON(AvionicsSuiteController.CreateAvionicsSuite)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/AvionicsSuite/{id}", jsonResponseFormatter.FormatToJSON(AvionicsSuiteController.UpdateAvionicsSuite)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/DeleteAvionicsSuite/{id}", jsonResponseFormatter.FormatToJSON(AvionicsSuiteController.DeleteAvionicsSuite)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AssignSupplierToAvionicsSuite/{parentId}/supplierId", jsonResponseFormatter.FormatToJSON(AvionicsSuiteController.AssignSupplierToAvionicsSuite)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignSupplierFromAvionicsSuite/{parentId}", jsonResponseFormatter.FormatToJSON(AvionicsSuiteController.UnassignSupplierFromAvionicsSuite)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AddVariantsToAvionicsSuite/{parentId}/variantsId", jsonResponseFormatter.FormatToJSON(AvionicsSuiteController.AddVariantsToAvionicsSuite)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveVariantsFromAvionicsSuite/{parentId}/variantsIds", jsonResponseFormatter.FormatToJSON(AvionicsSuiteController.RemoveVariantsFromAvionicsSuite)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddSoftwareLoadsToAvionicsSuite/{parentId}/softwareLoadsId", jsonResponseFormatter.FormatToJSON(AvionicsSuiteController.AddSoftwareLoadsToAvionicsSuite)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveSoftwareLoadsFromAvionicsSuite/{parentId}/softwareLoadsIds", jsonResponseFormatter.FormatToJSON(AvionicsSuiteController.RemoveSoftwareLoadsFromAvionicsSuite)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // APU Routes to JSON response formatter first
    // then to the correct Controller function
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Standard Lifecycle Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/APU/{id}", jsonResponseFormatter.FormatToJSON(APUController.GetAPU)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/APU", jsonResponseFormatter.FormatToJSON(APUController.GetAllAPU)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/NewAPU", jsonResponseFormatter.FormatToJSON(APUController.CreateAPU)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/APU/{id}", jsonResponseFormatter.FormatToJSON(APUController.UpdateAPU)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/DeleteAPU/{id}", jsonResponseFormatter.FormatToJSON(APUController.DeleteAPU)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AssignSupplierToAPU/{parentId}/supplierId", jsonResponseFormatter.FormatToJSON(APUController.AssignSupplierToAPU)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignSupplierFromAPU/{parentId}", jsonResponseFormatter.FormatToJSON(APUController.UnassignSupplierFromAPU)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AddVariantsToAPU/{parentId}/variantsId", jsonResponseFormatter.FormatToJSON(APUController.AddVariantsToAPU)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveVariantsFromAPU/{parentId}/variantsIds", jsonResponseFormatter.FormatToJSON(APUController.RemoveVariantsFromAPU)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // LandingGear Routes to JSON response formatter first
    // then to the correct Controller function
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Standard Lifecycle Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/LandingGear/{id}", jsonResponseFormatter.FormatToJSON(LandingGearController.GetLandingGear)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/LandingGear", jsonResponseFormatter.FormatToJSON(LandingGearController.GetAllLandingGear)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/NewLandingGear", jsonResponseFormatter.FormatToJSON(LandingGearController.CreateLandingGear)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/LandingGear/{id}", jsonResponseFormatter.FormatToJSON(LandingGearController.UpdateLandingGear)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/DeleteLandingGear/{id}", jsonResponseFormatter.FormatToJSON(LandingGearController.DeleteLandingGear)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AssignSupplierToLandingGear/{parentId}/supplierId", jsonResponseFormatter.FormatToJSON(LandingGearController.AssignSupplierToLandingGear)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignSupplierFromLandingGear/{parentId}", jsonResponseFormatter.FormatToJSON(LandingGearController.UnassignSupplierFromLandingGear)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AddVariantsToLandingGear/{parentId}/variantsId", jsonResponseFormatter.FormatToJSON(LandingGearController.AddVariantsToLandingGear)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveVariantsFromLandingGear/{parentId}/variantsIds", jsonResponseFormatter.FormatToJSON(LandingGearController.RemoveVariantsFromLandingGear)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // AircraftOption Routes to JSON response formatter first
    // then to the correct Controller function
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Standard Lifecycle Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AircraftOption/{id}", jsonResponseFormatter.FormatToJSON(AircraftOptionController.GetAircraftOption)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/AircraftOption", jsonResponseFormatter.FormatToJSON(AircraftOptionController.GetAllAircraftOption)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/NewAircraftOption", jsonResponseFormatter.FormatToJSON(AircraftOptionController.CreateAircraftOption)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/AircraftOption/{id}", jsonResponseFormatter.FormatToJSON(AircraftOptionController.UpdateAircraftOption)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/DeleteAircraftOption/{id}", jsonResponseFormatter.FormatToJSON(AircraftOptionController.DeleteAircraftOption)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AddVariantsToAircraftOption/{parentId}/variantsId", jsonResponseFormatter.FormatToJSON(AircraftOptionController.AddVariantsToAircraftOption)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveVariantsFromAircraftOption/{parentId}/variantsIds", jsonResponseFormatter.FormatToJSON(AircraftOptionController.RemoveVariantsFromAircraftOption)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddPackagesToAircraftOption/{parentId}/packagesId", jsonResponseFormatter.FormatToJSON(AircraftOptionController.AddPackagesToAircraftOption)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemovePackagesFromAircraftOption/{parentId}/packagesIds", jsonResponseFormatter.FormatToJSON(AircraftOptionController.RemovePackagesFromAircraftOption)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // AircraftPackage Routes to JSON response formatter first
    // then to the correct Controller function
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Standard Lifecycle Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AircraftPackage/{id}", jsonResponseFormatter.FormatToJSON(AircraftPackageController.GetAircraftPackage)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/AircraftPackage", jsonResponseFormatter.FormatToJSON(AircraftPackageController.GetAllAircraftPackage)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/NewAircraftPackage", jsonResponseFormatter.FormatToJSON(AircraftPackageController.CreateAircraftPackage)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/AircraftPackage/{id}", jsonResponseFormatter.FormatToJSON(AircraftPackageController.UpdateAircraftPackage)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/DeleteAircraftPackage/{id}", jsonResponseFormatter.FormatToJSON(AircraftPackageController.DeleteAircraftPackage)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AddOptionsToAircraftPackage/{parentId}/optionsId", jsonResponseFormatter.FormatToJSON(AircraftPackageController.AddOptionsToAircraftPackage)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveOptionsFromAircraftPackage/{parentId}/optionsIds", jsonResponseFormatter.FormatToJSON(AircraftPackageController.RemoveOptionsFromAircraftPackage)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddVariantsToAircraftPackage/{parentId}/variantsId", jsonResponseFormatter.FormatToJSON(AircraftPackageController.AddVariantsToAircraftPackage)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveVariantsFromAircraftPackage/{parentId}/variantsIds", jsonResponseFormatter.FormatToJSON(AircraftPackageController.RemoveVariantsFromAircraftPackage)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Supplier Routes to JSON response formatter first
    // then to the correct Controller function
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Standard Lifecycle Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/Supplier/{id}", jsonResponseFormatter.FormatToJSON(SupplierController.GetSupplier)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/Supplier", jsonResponseFormatter.FormatToJSON(SupplierController.GetAllSupplier)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/NewSupplier", jsonResponseFormatter.FormatToJSON(SupplierController.CreateSupplier)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/Supplier/{id}", jsonResponseFormatter.FormatToJSON(SupplierController.UpdateSupplier)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/DeleteSupplier/{id}", jsonResponseFormatter.FormatToJSON(SupplierController.DeleteSupplier)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AddManufacturersToSupplier/{parentId}/manufacturersId", jsonResponseFormatter.FormatToJSON(SupplierController.AddManufacturersToSupplier)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveManufacturersFromSupplier/{parentId}/manufacturersIds", jsonResponseFormatter.FormatToJSON(SupplierController.RemoveManufacturersFromSupplier)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddComponentsToSupplier/{parentId}/componentsId", jsonResponseFormatter.FormatToJSON(SupplierController.AddComponentsToSupplier)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveComponentsFromSupplier/{parentId}/componentsIds", jsonResponseFormatter.FormatToJSON(SupplierController.RemoveComponentsFromSupplier)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddEngineTypesToSupplier/{parentId}/engineTypesId", jsonResponseFormatter.FormatToJSON(SupplierController.AddEngineTypesToSupplier)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveEngineTypesFromSupplier/{parentId}/engineTypesIds", jsonResponseFormatter.FormatToJSON(SupplierController.RemoveEngineTypesFromSupplier)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddAvionicsSuitesToSupplier/{parentId}/avionicsSuitesId", jsonResponseFormatter.FormatToJSON(SupplierController.AddAvionicsSuitesToSupplier)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveAvionicsSuitesFromSupplier/{parentId}/avionicsSuitesIds", jsonResponseFormatter.FormatToJSON(SupplierController.RemoveAvionicsSuitesFromSupplier)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddApusToSupplier/{parentId}/apusId", jsonResponseFormatter.FormatToJSON(SupplierController.AddApusToSupplier)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveApusFromSupplier/{parentId}/apusIds", jsonResponseFormatter.FormatToJSON(SupplierController.RemoveApusFromSupplier)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddLandingGearsToSupplier/{parentId}/landingGearsId", jsonResponseFormatter.FormatToJSON(SupplierController.AddLandingGearsToSupplier)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveLandingGearsFromSupplier/{parentId}/landingGearsIds", jsonResponseFormatter.FormatToJSON(SupplierController.RemoveLandingGearsFromSupplier)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Component_ Routes to JSON response formatter first
    // then to the correct Controller function
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Standard Lifecycle Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/Component_/{id}", jsonResponseFormatter.FormatToJSON(Component_Controller.GetComponent_)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/Component_", jsonResponseFormatter.FormatToJSON(Component_Controller.GetAllComponent_)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/NewComponent_", jsonResponseFormatter.FormatToJSON(Component_Controller.CreateComponent_)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/Component_/{id}", jsonResponseFormatter.FormatToJSON(Component_Controller.UpdateComponent_)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/DeleteComponent_/{id}", jsonResponseFormatter.FormatToJSON(Component_Controller.DeleteComponent_)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AssignSupplierToComponent_/{parentId}/supplierId", jsonResponseFormatter.FormatToJSON(Component_Controller.AssignSupplierToComponent_)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignSupplierFromComponent_/{parentId}", jsonResponseFormatter.FormatToJSON(Component_Controller.UnassignSupplierFromComponent_)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Plant Routes to JSON response formatter first
    // then to the correct Controller function
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Standard Lifecycle Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/Plant/{id}", jsonResponseFormatter.FormatToJSON(PlantController.GetPlant)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/Plant", jsonResponseFormatter.FormatToJSON(PlantController.GetAllPlant)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/NewPlant", jsonResponseFormatter.FormatToJSON(PlantController.CreatePlant)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/Plant/{id}", jsonResponseFormatter.FormatToJSON(PlantController.UpdatePlant)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/DeletePlant/{id}", jsonResponseFormatter.FormatToJSON(PlantController.DeletePlant)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AssignManufacturerToPlant/{parentId}/manufacturerId", jsonResponseFormatter.FormatToJSON(PlantController.AssignManufacturerToPlant)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignManufacturerFromPlant/{parentId}", jsonResponseFormatter.FormatToJSON(PlantController.UnassignManufacturerFromPlant)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AddProductionLinesToPlant/{parentId}/productionLinesId", jsonResponseFormatter.FormatToJSON(PlantController.AddProductionLinesToPlant)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveProductionLinesFromPlant/{parentId}/productionLinesIds", jsonResponseFormatter.FormatToJSON(PlantController.RemoveProductionLinesFromPlant)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddWarehousesToPlant/{parentId}/warehousesId", jsonResponseFormatter.FormatToJSON(PlantController.AddWarehousesToPlant)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveWarehousesFromPlant/{parentId}/warehousesIds", jsonResponseFormatter.FormatToJSON(PlantController.RemoveWarehousesFromPlant)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // ProductionLine Routes to JSON response formatter first
    // then to the correct Controller function
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Standard Lifecycle Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/ProductionLine/{id}", jsonResponseFormatter.FormatToJSON(ProductionLineController.GetProductionLine)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/ProductionLine", jsonResponseFormatter.FormatToJSON(ProductionLineController.GetAllProductionLine)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/NewProductionLine", jsonResponseFormatter.FormatToJSON(ProductionLineController.CreateProductionLine)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/ProductionLine/{id}", jsonResponseFormatter.FormatToJSON(ProductionLineController.UpdateProductionLine)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/DeleteProductionLine/{id}", jsonResponseFormatter.FormatToJSON(ProductionLineController.DeleteProductionLine)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AssignPlantToProductionLine/{parentId}/plantId", jsonResponseFormatter.FormatToJSON(ProductionLineController.AssignPlantToProductionLine)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignPlantFromProductionLine/{parentId}", jsonResponseFormatter.FormatToJSON(ProductionLineController.UnassignPlantFromProductionLine)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AddWorkCentersToProductionLine/{parentId}/workCentersId", jsonResponseFormatter.FormatToJSON(ProductionLineController.AddWorkCentersToProductionLine)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveWorkCentersFromProductionLine/{parentId}/workCentersIds", jsonResponseFormatter.FormatToJSON(ProductionLineController.RemoveWorkCentersFromProductionLine)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // WorkCenter Routes to JSON response formatter first
    // then to the correct Controller function
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Standard Lifecycle Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/WorkCenter/{id}", jsonResponseFormatter.FormatToJSON(WorkCenterController.GetWorkCenter)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/WorkCenter", jsonResponseFormatter.FormatToJSON(WorkCenterController.GetAllWorkCenter)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/NewWorkCenter", jsonResponseFormatter.FormatToJSON(WorkCenterController.CreateWorkCenter)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/WorkCenter/{id}", jsonResponseFormatter.FormatToJSON(WorkCenterController.UpdateWorkCenter)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/DeleteWorkCenter/{id}", jsonResponseFormatter.FormatToJSON(WorkCenterController.DeleteWorkCenter)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AssignProductionLineToWorkCenter/{parentId}/productionLineId", jsonResponseFormatter.FormatToJSON(WorkCenterController.AssignProductionLineToWorkCenter)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignProductionLineFromWorkCenter/{parentId}", jsonResponseFormatter.FormatToJSON(WorkCenterController.UnassignProductionLineFromWorkCenter)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // ProductionOrder Routes to JSON response formatter first
    // then to the correct Controller function
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Standard Lifecycle Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/ProductionOrder/{id}", jsonResponseFormatter.FormatToJSON(ProductionOrderController.GetProductionOrder)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/ProductionOrder", jsonResponseFormatter.FormatToJSON(ProductionOrderController.GetAllProductionOrder)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/NewProductionOrder", jsonResponseFormatter.FormatToJSON(ProductionOrderController.CreateProductionOrder)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/ProductionOrder/{id}", jsonResponseFormatter.FormatToJSON(ProductionOrderController.UpdateProductionOrder)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/DeleteProductionOrder/{id}", jsonResponseFormatter.FormatToJSON(ProductionOrderController.DeleteProductionOrder)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AssignVariantToProductionOrder/{parentId}/variantId", jsonResponseFormatter.FormatToJSON(ProductionOrderController.AssignVariantToProductionOrder)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignVariantFromProductionOrder/{parentId}", jsonResponseFormatter.FormatToJSON(ProductionOrderController.UnassignVariantFromProductionOrder)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AssignPlantToProductionOrder/{parentId}/plantId", jsonResponseFormatter.FormatToJSON(ProductionOrderController.AssignPlantToProductionOrder)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignPlantFromProductionOrder/{parentId}", jsonResponseFormatter.FormatToJSON(ProductionOrderController.UnassignPlantFromProductionOrder)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AssignAircraftOrderToProductionOrder/{parentId}/aircraftOrderId", jsonResponseFormatter.FormatToJSON(ProductionOrderController.AssignAircraftOrderToProductionOrder)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignAircraftOrderFromProductionOrder/{parentId}", jsonResponseFormatter.FormatToJSON(ProductionOrderController.UnassignAircraftOrderFromProductionOrder)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // BuildSchedule Routes to JSON response formatter first
    // then to the correct Controller function
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Standard Lifecycle Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/BuildSchedule/{id}", jsonResponseFormatter.FormatToJSON(BuildScheduleController.GetBuildSchedule)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/BuildSchedule", jsonResponseFormatter.FormatToJSON(BuildScheduleController.GetAllBuildSchedule)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/NewBuildSchedule", jsonResponseFormatter.FormatToJSON(BuildScheduleController.CreateBuildSchedule)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/BuildSchedule/{id}", jsonResponseFormatter.FormatToJSON(BuildScheduleController.UpdateBuildSchedule)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/DeleteBuildSchedule/{id}", jsonResponseFormatter.FormatToJSON(BuildScheduleController.DeleteBuildSchedule)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AddProductionOrdersToBuildSchedule/{parentId}/productionOrdersId", jsonResponseFormatter.FormatToJSON(BuildScheduleController.AddProductionOrdersToBuildSchedule)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveProductionOrdersFromBuildSchedule/{parentId}/productionOrdersIds", jsonResponseFormatter.FormatToJSON(BuildScheduleController.RemoveProductionOrdersFromBuildSchedule)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Warehouse Routes to JSON response formatter first
    // then to the correct Controller function
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Standard Lifecycle Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/Warehouse/{id}", jsonResponseFormatter.FormatToJSON(WarehouseController.GetWarehouse)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/Warehouse", jsonResponseFormatter.FormatToJSON(WarehouseController.GetAllWarehouse)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/NewWarehouse", jsonResponseFormatter.FormatToJSON(WarehouseController.CreateWarehouse)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/Warehouse/{id}", jsonResponseFormatter.FormatToJSON(WarehouseController.UpdateWarehouse)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/DeleteWarehouse/{id}", jsonResponseFormatter.FormatToJSON(WarehouseController.DeleteWarehouse)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AddInventoryItemsToWarehouse/{parentId}/inventoryItemsId", jsonResponseFormatter.FormatToJSON(WarehouseController.AddInventoryItemsToWarehouse)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveInventoryItemsFromWarehouse/{parentId}/inventoryItemsIds", jsonResponseFormatter.FormatToJSON(WarehouseController.RemoveInventoryItemsFromWarehouse)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // InventoryItem Routes to JSON response formatter first
    // then to the correct Controller function
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Standard Lifecycle Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/InventoryItem/{id}", jsonResponseFormatter.FormatToJSON(InventoryItemController.GetInventoryItem)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/InventoryItem", jsonResponseFormatter.FormatToJSON(InventoryItemController.GetAllInventoryItem)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/NewInventoryItem", jsonResponseFormatter.FormatToJSON(InventoryItemController.CreateInventoryItem)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/InventoryItem/{id}", jsonResponseFormatter.FormatToJSON(InventoryItemController.UpdateInventoryItem)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/DeleteInventoryItem/{id}", jsonResponseFormatter.FormatToJSON(InventoryItemController.DeleteInventoryItem)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AssignComponentToInventoryItem/{parentId}/componentId", jsonResponseFormatter.FormatToJSON(InventoryItemController.AssignComponentToInventoryItem)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignComponentFromInventoryItem/{parentId}", jsonResponseFormatter.FormatToJSON(InventoryItemController.UnassignComponentFromInventoryItem)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AssignWarehouseToInventoryItem/{parentId}/warehouseId", jsonResponseFormatter.FormatToJSON(InventoryItemController.AssignWarehouseToInventoryItem)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignWarehouseFromInventoryItem/{parentId}", jsonResponseFormatter.FormatToJSON(InventoryItemController.UnassignWarehouseFromInventoryItem)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Operator Routes to JSON response formatter first
    // then to the correct Controller function
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Standard Lifecycle Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/Operator/{id}", jsonResponseFormatter.FormatToJSON(OperatorController.GetOperator)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/Operator", jsonResponseFormatter.FormatToJSON(OperatorController.GetAllOperator)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/NewOperator", jsonResponseFormatter.FormatToJSON(OperatorController.CreateOperator)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/Operator/{id}", jsonResponseFormatter.FormatToJSON(OperatorController.UpdateOperator)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/DeleteOperator/{id}", jsonResponseFormatter.FormatToJSON(OperatorController.DeleteOperator)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AssignSalesRegionToOperator/{parentId}/salesRegionId", jsonResponseFormatter.FormatToJSON(OperatorController.AssignSalesRegionToOperator)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignSalesRegionFromOperator/{parentId}", jsonResponseFormatter.FormatToJSON(OperatorController.UnassignSalesRegionFromOperator)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AddAircraftOrdersToOperator/{parentId}/aircraftOrdersId", jsonResponseFormatter.FormatToJSON(OperatorController.AddAircraftOrdersToOperator)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveAircraftOrdersFromOperator/{parentId}/aircraftOrdersIds", jsonResponseFormatter.FormatToJSON(OperatorController.RemoveAircraftOrdersFromOperator)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddOperatedAircraftToOperator/{parentId}/operatedAircraftId", jsonResponseFormatter.FormatToJSON(OperatorController.AddOperatedAircraftToOperator)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveOperatedAircraftFromOperator/{parentId}/operatedAircraftIds", jsonResponseFormatter.FormatToJSON(OperatorController.RemoveOperatedAircraftFromOperator)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // AircraftOrder Routes to JSON response formatter first
    // then to the correct Controller function
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Standard Lifecycle Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AircraftOrder/{id}", jsonResponseFormatter.FormatToJSON(AircraftOrderController.GetAircraftOrder)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/AircraftOrder", jsonResponseFormatter.FormatToJSON(AircraftOrderController.GetAllAircraftOrder)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/NewAircraftOrder", jsonResponseFormatter.FormatToJSON(AircraftOrderController.CreateAircraftOrder)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/AircraftOrder/{id}", jsonResponseFormatter.FormatToJSON(AircraftOrderController.UpdateAircraftOrder)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/DeleteAircraftOrder/{id}", jsonResponseFormatter.FormatToJSON(AircraftOrderController.DeleteAircraftOrder)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AssignOperatorToAircraftOrder/{parentId}/operatorId", jsonResponseFormatter.FormatToJSON(AircraftOrderController.AssignOperatorToAircraftOrder)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignOperatorFromAircraftOrder/{parentId}", jsonResponseFormatter.FormatToJSON(AircraftOrderController.UnassignOperatorFromAircraftOrder)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AssignVariantToAircraftOrder/{parentId}/variantId", jsonResponseFormatter.FormatToJSON(AircraftOrderController.AssignVariantToAircraftOrder)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignVariantFromAircraftOrder/{parentId}", jsonResponseFormatter.FormatToJSON(AircraftOrderController.UnassignVariantFromAircraftOrder)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AssignQuoteToAircraftOrder/{parentId}/quoteId", jsonResponseFormatter.FormatToJSON(AircraftOrderController.AssignQuoteToAircraftOrder)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignQuoteFromAircraftOrder/{parentId}", jsonResponseFormatter.FormatToJSON(AircraftOrderController.UnassignQuoteFromAircraftOrder)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AssignPurchaseAgreementToAircraftOrder/{parentId}/purchaseAgreementId", jsonResponseFormatter.FormatToJSON(AircraftOrderController.AssignPurchaseAgreementToAircraftOrder)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignPurchaseAgreementFromAircraftOrder/{parentId}", jsonResponseFormatter.FormatToJSON(AircraftOrderController.UnassignPurchaseAgreementFromAircraftOrder)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Quote Routes to JSON response formatter first
    // then to the correct Controller function
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Standard Lifecycle Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/Quote/{id}", jsonResponseFormatter.FormatToJSON(QuoteController.GetQuote)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/Quote", jsonResponseFormatter.FormatToJSON(QuoteController.GetAllQuote)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/NewQuote", jsonResponseFormatter.FormatToJSON(QuoteController.CreateQuote)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/Quote/{id}", jsonResponseFormatter.FormatToJSON(QuoteController.UpdateQuote)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/DeleteQuote/{id}", jsonResponseFormatter.FormatToJSON(QuoteController.DeleteQuote)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AssignAircraftOrderToQuote/{parentId}/aircraftOrderId", jsonResponseFormatter.FormatToJSON(QuoteController.AssignAircraftOrderToQuote)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignAircraftOrderFromQuote/{parentId}", jsonResponseFormatter.FormatToJSON(QuoteController.UnassignAircraftOrderFromQuote)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // PurchaseAgreement Routes to JSON response formatter first
    // then to the correct Controller function
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Standard Lifecycle Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/PurchaseAgreement/{id}", jsonResponseFormatter.FormatToJSON(PurchaseAgreementController.GetPurchaseAgreement)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/PurchaseAgreement", jsonResponseFormatter.FormatToJSON(PurchaseAgreementController.GetAllPurchaseAgreement)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/NewPurchaseAgreement", jsonResponseFormatter.FormatToJSON(PurchaseAgreementController.CreatePurchaseAgreement)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/PurchaseAgreement/{id}", jsonResponseFormatter.FormatToJSON(PurchaseAgreementController.UpdatePurchaseAgreement)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/DeletePurchaseAgreement/{id}", jsonResponseFormatter.FormatToJSON(PurchaseAgreementController.DeletePurchaseAgreement)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AssignAircraftOrderToPurchaseAgreement/{parentId}/aircraftOrderId", jsonResponseFormatter.FormatToJSON(PurchaseAgreementController.AssignAircraftOrderToPurchaseAgreement)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignAircraftOrderFromPurchaseAgreement/{parentId}", jsonResponseFormatter.FormatToJSON(PurchaseAgreementController.UnassignAircraftOrderFromPurchaseAgreement)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Aircraft Routes to JSON response formatter first
    // then to the correct Controller function
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Standard Lifecycle Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/Aircraft/{id}", jsonResponseFormatter.FormatToJSON(AircraftController.GetAircraft)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/Aircraft", jsonResponseFormatter.FormatToJSON(AircraftController.GetAllAircraft)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/NewAircraft", jsonResponseFormatter.FormatToJSON(AircraftController.CreateAircraft)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/Aircraft/{id}", jsonResponseFormatter.FormatToJSON(AircraftController.UpdateAircraft)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/DeleteAircraft/{id}", jsonResponseFormatter.FormatToJSON(AircraftController.DeleteAircraft)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AssignVariantToAircraft/{parentId}/variantId", jsonResponseFormatter.FormatToJSON(AircraftController.AssignVariantToAircraft)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignVariantFromAircraft/{parentId}", jsonResponseFormatter.FormatToJSON(AircraftController.UnassignVariantFromAircraft)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AssignOperatorToAircraft/{parentId}/operatorId", jsonResponseFormatter.FormatToJSON(AircraftController.AssignOperatorToAircraft)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignOperatorFromAircraft/{parentId}", jsonResponseFormatter.FormatToJSON(AircraftController.UnassignOperatorFromAircraft)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AssignRegistrationToAircraft/{parentId}/registrationId", jsonResponseFormatter.FormatToJSON(AircraftController.AssignRegistrationToAircraft)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignRegistrationFromAircraft/{parentId}", jsonResponseFormatter.FormatToJSON(AircraftController.UnassignRegistrationFromAircraft)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AssignWarrantyToAircraft/{parentId}/warrantyId", jsonResponseFormatter.FormatToJSON(AircraftController.AssignWarrantyToAircraft)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignWarrantyFromAircraft/{parentId}", jsonResponseFormatter.FormatToJSON(AircraftController.UnassignWarrantyFromAircraft)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AssignConnectedAircraftToAircraft/{parentId}/connectedAircraftId", jsonResponseFormatter.FormatToJSON(AircraftController.AssignConnectedAircraftToAircraft)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignConnectedAircraftFromAircraft/{parentId}", jsonResponseFormatter.FormatToJSON(AircraftController.UnassignConnectedAircraftFromAircraft)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AssignCabinLayoutToAircraft/{parentId}/cabinLayoutId", jsonResponseFormatter.FormatToJSON(AircraftController.AssignCabinLayoutToAircraft)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignCabinLayoutFromAircraft/{parentId}", jsonResponseFormatter.FormatToJSON(AircraftController.UnassignCabinLayoutFromAircraft)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AddMaintenanceRecordsToAircraft/{parentId}/maintenanceRecordsId", jsonResponseFormatter.FormatToJSON(AircraftController.AddMaintenanceRecordsToAircraft)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveMaintenanceRecordsFromAircraft/{parentId}/maintenanceRecordsIds", jsonResponseFormatter.FormatToJSON(AircraftController.RemoveMaintenanceRecordsFromAircraft)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Registration Routes to JSON response formatter first
    // then to the correct Controller function
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Standard Lifecycle Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/Registration/{id}", jsonResponseFormatter.FormatToJSON(RegistrationController.GetRegistration)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/Registration", jsonResponseFormatter.FormatToJSON(RegistrationController.GetAllRegistration)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/NewRegistration", jsonResponseFormatter.FormatToJSON(RegistrationController.CreateRegistration)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/Registration/{id}", jsonResponseFormatter.FormatToJSON(RegistrationController.UpdateRegistration)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/DeleteRegistration/{id}", jsonResponseFormatter.FormatToJSON(RegistrationController.DeleteRegistration)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AssignAircraftToRegistration/{parentId}/aircraftId", jsonResponseFormatter.FormatToJSON(RegistrationController.AssignAircraftToRegistration)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignAircraftFromRegistration/{parentId}", jsonResponseFormatter.FormatToJSON(RegistrationController.UnassignAircraftFromRegistration)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Warranty Routes to JSON response formatter first
    // then to the correct Controller function
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Standard Lifecycle Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/Warranty/{id}", jsonResponseFormatter.FormatToJSON(WarrantyController.GetWarranty)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/Warranty", jsonResponseFormatter.FormatToJSON(WarrantyController.GetAllWarranty)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/NewWarranty", jsonResponseFormatter.FormatToJSON(WarrantyController.CreateWarranty)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/Warranty/{id}", jsonResponseFormatter.FormatToJSON(WarrantyController.UpdateWarranty)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/DeleteWarranty/{id}", jsonResponseFormatter.FormatToJSON(WarrantyController.DeleteWarranty)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AssignAircraftToWarranty/{parentId}/aircraftId", jsonResponseFormatter.FormatToJSON(WarrantyController.AssignAircraftToWarranty)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignAircraftFromWarranty/{parentId}", jsonResponseFormatter.FormatToJSON(WarrantyController.UnassignAircraftFromWarranty)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // CabinLayout Routes to JSON response formatter first
    // then to the correct Controller function
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Standard Lifecycle Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/CabinLayout/{id}", jsonResponseFormatter.FormatToJSON(CabinLayoutController.GetCabinLayout)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/CabinLayout", jsonResponseFormatter.FormatToJSON(CabinLayoutController.GetAllCabinLayout)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/NewCabinLayout", jsonResponseFormatter.FormatToJSON(CabinLayoutController.CreateCabinLayout)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/CabinLayout/{id}", jsonResponseFormatter.FormatToJSON(CabinLayoutController.UpdateCabinLayout)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/DeleteCabinLayout/{id}", jsonResponseFormatter.FormatToJSON(CabinLayoutController.DeleteCabinLayout)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AssignVariantToCabinLayout/{parentId}/variantId", jsonResponseFormatter.FormatToJSON(CabinLayoutController.AssignVariantToCabinLayout)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignVariantFromCabinLayout/{parentId}", jsonResponseFormatter.FormatToJSON(CabinLayoutController.UnassignVariantFromCabinLayout)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AddAircraftToCabinLayout/{parentId}/aircraftId", jsonResponseFormatter.FormatToJSON(CabinLayoutController.AddAircraftToCabinLayout)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveAircraftFromCabinLayout/{parentId}/aircraftIds", jsonResponseFormatter.FormatToJSON(CabinLayoutController.RemoveAircraftFromCabinLayout)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddOptionsToCabinLayout/{parentId}/optionsId", jsonResponseFormatter.FormatToJSON(CabinLayoutController.AddOptionsToCabinLayout)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveOptionsFromCabinLayout/{parentId}/optionsIds", jsonResponseFormatter.FormatToJSON(CabinLayoutController.RemoveOptionsFromCabinLayout)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // MROFacility Routes to JSON response formatter first
    // then to the correct Controller function
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Standard Lifecycle Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/MROFacility/{id}", jsonResponseFormatter.FormatToJSON(MROFacilityController.GetMROFacility)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/MROFacility", jsonResponseFormatter.FormatToJSON(MROFacilityController.GetAllMROFacility)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/NewMROFacility", jsonResponseFormatter.FormatToJSON(MROFacilityController.CreateMROFacility)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/MROFacility/{id}", jsonResponseFormatter.FormatToJSON(MROFacilityController.UpdateMROFacility)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/DeleteMROFacility/{id}", jsonResponseFormatter.FormatToJSON(MROFacilityController.DeleteMROFacility)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AddAppointmentsToMROFacility/{parentId}/appointmentsId", jsonResponseFormatter.FormatToJSON(MROFacilityController.AddAppointmentsToMROFacility)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveAppointmentsFromMROFacility/{parentId}/appointmentsIds", jsonResponseFormatter.FormatToJSON(MROFacilityController.RemoveAppointmentsFromMROFacility)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddWorkOrdersToMROFacility/{parentId}/workOrdersId", jsonResponseFormatter.FormatToJSON(MROFacilityController.AddWorkOrdersToMROFacility)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveWorkOrdersFromMROFacility/{parentId}/workOrdersIds", jsonResponseFormatter.FormatToJSON(MROFacilityController.RemoveWorkOrdersFromMROFacility)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // MaintenanceAppointment Routes to JSON response formatter first
    // then to the correct Controller function
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Standard Lifecycle Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/MaintenanceAppointment/{id}", jsonResponseFormatter.FormatToJSON(MaintenanceAppointmentController.GetMaintenanceAppointment)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/MaintenanceAppointment", jsonResponseFormatter.FormatToJSON(MaintenanceAppointmentController.GetAllMaintenanceAppointment)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/NewMaintenanceAppointment", jsonResponseFormatter.FormatToJSON(MaintenanceAppointmentController.CreateMaintenanceAppointment)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/MaintenanceAppointment/{id}", jsonResponseFormatter.FormatToJSON(MaintenanceAppointmentController.UpdateMaintenanceAppointment)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/DeleteMaintenanceAppointment/{id}", jsonResponseFormatter.FormatToJSON(MaintenanceAppointmentController.DeleteMaintenanceAppointment)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AssignAircraftToMaintenanceAppointment/{parentId}/aircraftId", jsonResponseFormatter.FormatToJSON(MaintenanceAppointmentController.AssignAircraftToMaintenanceAppointment)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignAircraftFromMaintenanceAppointment/{parentId}", jsonResponseFormatter.FormatToJSON(MaintenanceAppointmentController.UnassignAircraftFromMaintenanceAppointment)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AssignMroFacilityToMaintenanceAppointment/{parentId}/mroFacilityId", jsonResponseFormatter.FormatToJSON(MaintenanceAppointmentController.AssignMroFacilityToMaintenanceAppointment)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignMroFacilityFromMaintenanceAppointment/{parentId}", jsonResponseFormatter.FormatToJSON(MaintenanceAppointmentController.UnassignMroFacilityFromMaintenanceAppointment)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AssignWorkOrderToMaintenanceAppointment/{parentId}/workOrderId", jsonResponseFormatter.FormatToJSON(MaintenanceAppointmentController.AssignWorkOrderToMaintenanceAppointment)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignWorkOrderFromMaintenanceAppointment/{parentId}", jsonResponseFormatter.FormatToJSON(MaintenanceAppointmentController.UnassignWorkOrderFromMaintenanceAppointment)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // MaintenanceWorkOrder Routes to JSON response formatter first
    // then to the correct Controller function
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Standard Lifecycle Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/MaintenanceWorkOrder/{id}", jsonResponseFormatter.FormatToJSON(MaintenanceWorkOrderController.GetMaintenanceWorkOrder)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/MaintenanceWorkOrder", jsonResponseFormatter.FormatToJSON(MaintenanceWorkOrderController.GetAllMaintenanceWorkOrder)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/NewMaintenanceWorkOrder", jsonResponseFormatter.FormatToJSON(MaintenanceWorkOrderController.CreateMaintenanceWorkOrder)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/MaintenanceWorkOrder/{id}", jsonResponseFormatter.FormatToJSON(MaintenanceWorkOrderController.UpdateMaintenanceWorkOrder)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/DeleteMaintenanceWorkOrder/{id}", jsonResponseFormatter.FormatToJSON(MaintenanceWorkOrderController.DeleteMaintenanceWorkOrder)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AssignAircraftToMaintenanceWorkOrder/{parentId}/aircraftId", jsonResponseFormatter.FormatToJSON(MaintenanceWorkOrderController.AssignAircraftToMaintenanceWorkOrder)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignAircraftFromMaintenanceWorkOrder/{parentId}", jsonResponseFormatter.FormatToJSON(MaintenanceWorkOrderController.UnassignAircraftFromMaintenanceWorkOrder)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AssignAirworthinessDirectiveToMaintenanceWorkOrder/{parentId}/airworthinessDirectiveId", jsonResponseFormatter.FormatToJSON(MaintenanceWorkOrderController.AssignAirworthinessDirectiveToMaintenanceWorkOrder)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignAirworthinessDirectiveFromMaintenanceWorkOrder/{parentId}", jsonResponseFormatter.FormatToJSON(MaintenanceWorkOrderController.UnassignAirworthinessDirectiveFromMaintenanceWorkOrder)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AssignServiceBulletinToMaintenanceWorkOrder/{parentId}/serviceBulletinId", jsonResponseFormatter.FormatToJSON(MaintenanceWorkOrderController.AssignServiceBulletinToMaintenanceWorkOrder)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignServiceBulletinFromMaintenanceWorkOrder/{parentId}", jsonResponseFormatter.FormatToJSON(MaintenanceWorkOrderController.UnassignServiceBulletinFromMaintenanceWorkOrder)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // AirworthinessDirective Routes to JSON response formatter first
    // then to the correct Controller function
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Standard Lifecycle Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AirworthinessDirective/{id}", jsonResponseFormatter.FormatToJSON(AirworthinessDirectiveController.GetAirworthinessDirective)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/AirworthinessDirective", jsonResponseFormatter.FormatToJSON(AirworthinessDirectiveController.GetAllAirworthinessDirective)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/NewAirworthinessDirective", jsonResponseFormatter.FormatToJSON(AirworthinessDirectiveController.CreateAirworthinessDirective)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/AirworthinessDirective/{id}", jsonResponseFormatter.FormatToJSON(AirworthinessDirectiveController.UpdateAirworthinessDirective)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/DeleteAirworthinessDirective/{id}", jsonResponseFormatter.FormatToJSON(AirworthinessDirectiveController.DeleteAirworthinessDirective)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AddWorkOrdersToAirworthinessDirective/{parentId}/workOrdersId", jsonResponseFormatter.FormatToJSON(AirworthinessDirectiveController.AddWorkOrdersToAirworthinessDirective)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveWorkOrdersFromAirworthinessDirective/{parentId}/workOrdersIds", jsonResponseFormatter.FormatToJSON(AirworthinessDirectiveController.RemoveWorkOrdersFromAirworthinessDirective)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // ServiceBulletin Routes to JSON response formatter first
    // then to the correct Controller function
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Standard Lifecycle Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/ServiceBulletin/{id}", jsonResponseFormatter.FormatToJSON(ServiceBulletinController.GetServiceBulletin)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/ServiceBulletin", jsonResponseFormatter.FormatToJSON(ServiceBulletinController.GetAllServiceBulletin)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/NewServiceBulletin", jsonResponseFormatter.FormatToJSON(ServiceBulletinController.CreateServiceBulletin)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/ServiceBulletin/{id}", jsonResponseFormatter.FormatToJSON(ServiceBulletinController.UpdateServiceBulletin)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/DeleteServiceBulletin/{id}", jsonResponseFormatter.FormatToJSON(ServiceBulletinController.DeleteServiceBulletin)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AddWorkOrdersToServiceBulletin/{parentId}/workOrdersId", jsonResponseFormatter.FormatToJSON(ServiceBulletinController.AddWorkOrdersToServiceBulletin)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveWorkOrdersFromServiceBulletin/{parentId}/workOrdersIds", jsonResponseFormatter.FormatToJSON(ServiceBulletinController.RemoveWorkOrdersFromServiceBulletin)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddVariantsToServiceBulletin/{parentId}/variantsId", jsonResponseFormatter.FormatToJSON(ServiceBulletinController.AddVariantsToServiceBulletin)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveVariantsFromServiceBulletin/{parentId}/variantsIds", jsonResponseFormatter.FormatToJSON(ServiceBulletinController.RemoveVariantsFromServiceBulletin)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // ConnectedAircraft Routes to JSON response formatter first
    // then to the correct Controller function
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Standard Lifecycle Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/ConnectedAircraft/{id}", jsonResponseFormatter.FormatToJSON(ConnectedAircraftController.GetConnectedAircraft)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/ConnectedAircraft", jsonResponseFormatter.FormatToJSON(ConnectedAircraftController.GetAllConnectedAircraft)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/NewConnectedAircraft", jsonResponseFormatter.FormatToJSON(ConnectedAircraftController.CreateConnectedAircraft)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/ConnectedAircraft/{id}", jsonResponseFormatter.FormatToJSON(ConnectedAircraftController.UpdateConnectedAircraft)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/DeleteConnectedAircraft/{id}", jsonResponseFormatter.FormatToJSON(ConnectedAircraftController.DeleteConnectedAircraft)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AssignAircraftToConnectedAircraft/{parentId}/aircraftId", jsonResponseFormatter.FormatToJSON(ConnectedAircraftController.AssignAircraftToConnectedAircraft)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignAircraftFromConnectedAircraft/{parentId}", jsonResponseFormatter.FormatToJSON(ConnectedAircraftController.UnassignAircraftFromConnectedAircraft)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AddFlightHealthEventsToConnectedAircraft/{parentId}/flightHealthEventsId", jsonResponseFormatter.FormatToJSON(ConnectedAircraftController.AddFlightHealthEventsToConnectedAircraft)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveFlightHealthEventsFromConnectedAircraft/{parentId}/flightHealthEventsIds", jsonResponseFormatter.FormatToJSON(ConnectedAircraftController.RemoveFlightHealthEventsFromConnectedAircraft)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddSoftwareLoadsToConnectedAircraft/{parentId}/softwareLoadsId", jsonResponseFormatter.FormatToJSON(ConnectedAircraftController.AddSoftwareLoadsToConnectedAircraft)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveSoftwareLoadsFromConnectedAircraft/{parentId}/softwareLoadsIds", jsonResponseFormatter.FormatToJSON(ConnectedAircraftController.RemoveSoftwareLoadsFromConnectedAircraft)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // FlightHealthEvent Routes to JSON response formatter first
    // then to the correct Controller function
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Standard Lifecycle Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/FlightHealthEvent/{id}", jsonResponseFormatter.FormatToJSON(FlightHealthEventController.GetFlightHealthEvent)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/FlightHealthEvent", jsonResponseFormatter.FormatToJSON(FlightHealthEventController.GetAllFlightHealthEvent)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/NewFlightHealthEvent", jsonResponseFormatter.FormatToJSON(FlightHealthEventController.CreateFlightHealthEvent)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/FlightHealthEvent/{id}", jsonResponseFormatter.FormatToJSON(FlightHealthEventController.UpdateFlightHealthEvent)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/DeleteFlightHealthEvent/{id}", jsonResponseFormatter.FormatToJSON(FlightHealthEventController.DeleteFlightHealthEvent)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AssignConnectedAircraftToFlightHealthEvent/{parentId}/connectedAircraftId", jsonResponseFormatter.FormatToJSON(FlightHealthEventController.AssignConnectedAircraftToFlightHealthEvent)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignConnectedAircraftFromFlightHealthEvent/{parentId}", jsonResponseFormatter.FormatToJSON(FlightHealthEventController.UnassignConnectedAircraftFromFlightHealthEvent)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // SoftwareLoad Routes to JSON response formatter first
    // then to the correct Controller function
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Standard Lifecycle Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/SoftwareLoad/{id}", jsonResponseFormatter.FormatToJSON(SoftwareLoadController.GetSoftwareLoad)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/SoftwareLoad", jsonResponseFormatter.FormatToJSON(SoftwareLoadController.GetAllSoftwareLoad)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/NewSoftwareLoad", jsonResponseFormatter.FormatToJSON(SoftwareLoadController.CreateSoftwareLoad)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/SoftwareLoad/{id}", jsonResponseFormatter.FormatToJSON(SoftwareLoadController.UpdateSoftwareLoad)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/DeleteSoftwareLoad/{id}", jsonResponseFormatter.FormatToJSON(SoftwareLoadController.DeleteSoftwareLoad)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AssignConnectedAircraftToSoftwareLoad/{parentId}/connectedAircraftId", jsonResponseFormatter.FormatToJSON(SoftwareLoadController.AssignConnectedAircraftToSoftwareLoad)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignConnectedAircraftFromSoftwareLoad/{parentId}", jsonResponseFormatter.FormatToJSON(SoftwareLoadController.UnassignConnectedAircraftFromSoftwareLoad)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AssignAvionicsSuiteToSoftwareLoad/{parentId}/avionicsSuiteId", jsonResponseFormatter.FormatToJSON(SoftwareLoadController.AssignAvionicsSuiteToSoftwareLoad)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignAvionicsSuiteFromSoftwareLoad/{parentId}", jsonResponseFormatter.FormatToJSON(SoftwareLoadController.UnassignAvionicsSuiteFromSoftwareLoad)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // TypeCertificate Routes to JSON response formatter first
    // then to the correct Controller function
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Standard Lifecycle Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/TypeCertificate/{id}", jsonResponseFormatter.FormatToJSON(TypeCertificateController.GetTypeCertificate)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/TypeCertificate", jsonResponseFormatter.FormatToJSON(TypeCertificateController.GetAllTypeCertificate)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/NewTypeCertificate", jsonResponseFormatter.FormatToJSON(TypeCertificateController.CreateTypeCertificate)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/TypeCertificate/{id}", jsonResponseFormatter.FormatToJSON(TypeCertificateController.UpdateTypeCertificate)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/DeleteTypeCertificate/{id}", jsonResponseFormatter.FormatToJSON(TypeCertificateController.DeleteTypeCertificate)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AssignProgramToTypeCertificate/{parentId}/programId", jsonResponseFormatter.FormatToJSON(TypeCertificateController.AssignProgramToTypeCertificate)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignProgramFromTypeCertificate/{parentId}", jsonResponseFormatter.FormatToJSON(TypeCertificateController.UnassignProgramFromTypeCertificate)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // ProductionCertificate Routes to JSON response formatter first
    // then to the correct Controller function
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Standard Lifecycle Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/ProductionCertificate/{id}", jsonResponseFormatter.FormatToJSON(ProductionCertificateController.GetProductionCertificate)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/ProductionCertificate", jsonResponseFormatter.FormatToJSON(ProductionCertificateController.GetAllProductionCertificate)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/NewProductionCertificate", jsonResponseFormatter.FormatToJSON(ProductionCertificateController.CreateProductionCertificate)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/ProductionCertificate/{id}", jsonResponseFormatter.FormatToJSON(ProductionCertificateController.UpdateProductionCertificate)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/DeleteProductionCertificate/{id}", jsonResponseFormatter.FormatToJSON(ProductionCertificateController.DeleteProductionCertificate)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AssignManufacturerToProductionCertificate/{parentId}/manufacturerId", jsonResponseFormatter.FormatToJSON(ProductionCertificateController.AssignManufacturerToProductionCertificate)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignManufacturerFromProductionCertificate/{parentId}", jsonResponseFormatter.FormatToJSON(ProductionCertificateController.UnassignManufacturerFromProductionCertificate)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // SalesRegion Routes to JSON response formatter first
    // then to the correct Controller function
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Standard Lifecycle Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/SalesRegion/{id}", jsonResponseFormatter.FormatToJSON(SalesRegionController.GetSalesRegion)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/SalesRegion", jsonResponseFormatter.FormatToJSON(SalesRegionController.GetAllSalesRegion)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/NewSalesRegion", jsonResponseFormatter.FormatToJSON(SalesRegionController.CreateSalesRegion)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/SalesRegion/{id}", jsonResponseFormatter.FormatToJSON(SalesRegionController.UpdateSalesRegion)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/DeleteSalesRegion/{id}", jsonResponseFormatter.FormatToJSON(SalesRegionController.DeleteSalesRegion)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AddOperatorsToSalesRegion/{parentId}/operatorsId", jsonResponseFormatter.FormatToJSON(SalesRegionController.AddOperatorsToSalesRegion)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveOperatorsFromSalesRegion/{parentId}/operatorsIds", jsonResponseFormatter.FormatToJSON(SalesRegionController.RemoveOperatorsFromSalesRegion)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddSalesCampaignsToSalesRegion/{parentId}/salesCampaignsId", jsonResponseFormatter.FormatToJSON(SalesRegionController.AddSalesCampaignsToSalesRegion)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveSalesCampaignsFromSalesRegion/{parentId}/salesCampaignsIds", jsonResponseFormatter.FormatToJSON(SalesRegionController.RemoveSalesCampaignsFromSalesRegion)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // SalesCampaign Routes to JSON response formatter first
    // then to the correct Controller function
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Standard Lifecycle Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/SalesCampaign/{id}", jsonResponseFormatter.FormatToJSON(SalesCampaignController.GetSalesCampaign)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/SalesCampaign", jsonResponseFormatter.FormatToJSON(SalesCampaignController.GetAllSalesCampaign)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/NewSalesCampaign", jsonResponseFormatter.FormatToJSON(SalesCampaignController.CreateSalesCampaign)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/SalesCampaign/{id}", jsonResponseFormatter.FormatToJSON(SalesCampaignController.UpdateSalesCampaign)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/DeleteSalesCampaign/{id}", jsonResponseFormatter.FormatToJSON(SalesCampaignController.DeleteSalesCampaign)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AssignRegionToSalesCampaign/{parentId}/regionId", jsonResponseFormatter.FormatToJSON(SalesCampaignController.AssignRegionToSalesCampaign)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignRegionFromSalesCampaign/{parentId}", jsonResponseFormatter.FormatToJSON(SalesCampaignController.UnassignRegionFromSalesCampaign)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AssignOperatorToSalesCampaign/{parentId}/operatorId", jsonResponseFormatter.FormatToJSON(SalesCampaignController.AssignOperatorToSalesCampaign)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignOperatorFromSalesCampaign/{parentId}", jsonResponseFormatter.FormatToJSON(SalesCampaignController.UnassignOperatorFromSalesCampaign)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AddQuotesToSalesCampaign/{parentId}/quotesId", jsonResponseFormatter.FormatToJSON(SalesCampaignController.AddQuotesToSalesCampaign)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveQuotesFromSalesCampaign/{parentId}/quotesIds", jsonResponseFormatter.FormatToJSON(SalesCampaignController.RemoveQuotesFromSalesCampaign)).Methods("DELETE", "OPTIONS")

    return router
}
