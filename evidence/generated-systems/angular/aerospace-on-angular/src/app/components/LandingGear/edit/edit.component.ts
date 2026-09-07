import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { FormGroup, FormBuilder, Validators } from '@angular/forms';

import { LandingGearService } from '../../../services/LandingGear.service';
import { SubBaseComponent } from '../../LandingGear/sub.base.component';


@Component({
    selector: 'app-edit-landingGear',
    standalone: false,
    templateUrl: './edit.component.html',
    styleUrls: ['./edit.component.css']
})
export class EditLandingGearComponent extends SubBaseComponent implements OnInit {

    title = 'Edit LandingGear';

    landingGearForm: FormGroup;
    landingGear: any;

    constructor( http: HttpClient,
        private route: ActivatedRoute,
        private router: Router,
        private service: LandingGearService,
        private fb: FormBuilder
) {
        super(http);
        this.landingGearForm = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  supplierPartNumber: ['', Validators.required],
      Supplier: ['', ],
      Variants: ['', ],
      GearType: ['', ]
        });
    }

    
    updateLandingGear(supplierPartNumber, Supplier, Variants, GearType): void {
        this.route.params.subscribe((params) => {

                        this.service.updateLandingGear(supplierPartNumber, Supplier, Variants, GearType, params['id'])
                            .subscribe(() => {
                    this.router.navigate(['/indexLandingGear']);
                });
        });
    }

    ngOnInit(): void {
        this.route.params.subscribe((params) => {
            this.service.getLandingGear(params['id']).subscribe(res => {
                this.landingGear = res;
            });
        });
    }
}