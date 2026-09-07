import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { HttpClient } from '@angular/common/http';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { AircraftVariantService } from '../../../services/AircraftVariant.service';
import { AircraftVariant } from '../../../models/AircraftVariant';
import { SubBaseComponent } from '../../AircraftVariant/sub.base.component';

@Component({
    selector: 'app-create-aircraftVariant',
    standalone: false,
    templateUrl: './create.component.html',
    styleUrls: ['./create.component.css']
})
export class CreateAircraftVariantComponent extends SubBaseComponent implements OnInit {

    title = 'Add AircraftVariant';

    aircraftVariantForm: FormGroup;
    aircraftVariant: AircraftVariant;

    constructor( http: HttpClient,
        private aircraftVariantService: AircraftVariantService,
        private fb: FormBuilder,
        private router: Router
) {
        super(http);
        this.aircraftVariantForm = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  variantCode: ['', Validators.required],
      rangeNm: ['', Validators.required],
      maxTakeoffWeightKg: ['', Validators.required],
      Model_: ['', ],
      EngineType: ['', ],
      AvionicsSuite: ['', ],
      Apu: ['', ],
      LandingGear: ['', ],
      CabinLayouts: ['', ],
      Options: ['', ],
      Packages: ['', ]
        });
    }

    
    addAircraftVariant(variantCode, rangeNm, maxTakeoffWeightKg, Model_, EngineType, AvionicsSuite, Apu, LandingGear, CabinLayouts, Options, Packages): void {
        this.aircraftVariantService
        .addAircraftVariant(variantCode, rangeNm, maxTakeoffWeightKg, Model_, EngineType, AvionicsSuite, Apu, LandingGear, CabinLayouts, Options, Packages)
            .subscribe(() => {
                this.router.navigate(['/indexAircraftVariant']);
            });
    }

    ngOnInit(): void {
    }
}