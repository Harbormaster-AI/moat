import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { HttpClient } from '@angular/common/http';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { LandingGearService } from '../../../services/LandingGear.service';
import { LandingGear } from '../../../models/LandingGear';
import { SubBaseComponent } from '../../LandingGear/sub.base.component';

@Component({
    selector: 'app-create-landingGear',
    standalone: false,
    templateUrl: './create.component.html',
    styleUrls: ['./create.component.css']
})
export class CreateLandingGearComponent extends SubBaseComponent implements OnInit {

    title = 'Add LandingGear';

    landingGearForm: FormGroup;
    landingGear: LandingGear;

    constructor( http: HttpClient,
        private landingGearService: LandingGearService,
        private fb: FormBuilder,
        private router: Router
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

    
    addLandingGear(supplierPartNumber, Supplier, Variants, GearType): void {
        this.landingGearService
        .addLandingGear(supplierPartNumber, Supplier, Variants, GearType)
            .subscribe(() => {
                this.router.navigate(['/indexLandingGear']);
            });
    }

    ngOnInit(): void {
    }
}