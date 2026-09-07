import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { HttpClient } from '@angular/common/http';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { ExposureService } from '../../../services/Exposure.service';
import { Exposure } from '../../../models/Exposure';
import { SubBaseComponent } from '../../Exposure/sub.base.component';

@Component({
    selector: 'app-create-exposure',
    standalone: false,
    templateUrl: './create.component.html',
    styleUrls: ['./create.component.css']
})
export class CreateExposureComponent extends SubBaseComponent implements OnInit {

    title = 'Add Exposure';

    exposureForm: FormGroup;
    exposure: Exposure;

    constructor( http: HttpClient,
        private exposureService: ExposureService,
        private fb: FormBuilder,
        private router: Router
) {
        super(http);
        this.exposureForm = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  Claim: ['', ],
      PolicyCoverage: ['', ],
      InsuredObject: ['', ],
      Reserves: ['', ],
      Payments: ['', ],
      ExposureType: ['', ],
      Status: ['', ]
        });
    }

    
    addExposure(Claim, PolicyCoverage, InsuredObject, Reserves, Payments, ExposureType, Status): void {
        this.exposureService
        .addExposure(Claim, PolicyCoverage, InsuredObject, Reserves, Payments, ExposureType, Status)
            .subscribe(() => {
                this.router.navigate(['/indexExposure']);
            });
    }

    ngOnInit(): void {
    }
}