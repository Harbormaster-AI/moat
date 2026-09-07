import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { FormGroup, FormBuilder, Validators } from '@angular/forms';

import { AircraftVariantService } from '../../../services/AircraftVariant.service';
import { SubBaseComponent } from '../../AircraftVariant/sub.base.component';


@Component({
    selector: 'app-edit-aircraftVariant',
    standalone: false,
    templateUrl: './edit.component.html',
    styleUrls: ['./edit.component.css']
})
export class EditAircraftVariantComponent extends SubBaseComponent implements OnInit {

    title = 'Edit AircraftVariant';

    aircraftVariantForm: FormGroup;
    aircraftVariant: any;

    constructor( http: HttpClient,
        private route: ActivatedRoute,
        private router: Router,
        private service: AircraftVariantService,
        private fb: FormBuilder
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

    
    updateAircraftVariant(variantCode, rangeNm, maxTakeoffWeightKg, Model_, EngineType, AvionicsSuite, Apu, LandingGear, CabinLayouts, Options, Packages): void {
        this.route.params.subscribe((params) => {

                        this.service.updateAircraftVariant(variantCode, rangeNm, maxTakeoffWeightKg, Model_, EngineType, AvionicsSuite, Apu, LandingGear, CabinLayouts, Options, Packages, params['id'])
                            .subscribe(() => {
                    this.router.navigate(['/indexAircraftVariant']);
                });
        });
    }

    ngOnInit(): void {
        this.route.params.subscribe((params) => {
            this.service.getAircraftVariant(params['id']).subscribe(res => {
                this.aircraftVariant = res;
            });
        });
    }
}